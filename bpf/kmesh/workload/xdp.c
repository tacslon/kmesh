// SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause)
/* Copyright Authors of Kmesh */

#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>
#include <linux/in.h>
#include <linux/ip.h>
#include <linux/ipv6.h>
#include <linux/tcp.h>
#include <linux/if_ether.h>
#include "config.h"
#include "bpf_log.h"
#include "workload.h"
#include "authz.h"
#include "xdp.h"

static inline void parser_tuple(struct xdp_info *info, struct bpf_sock_tuple *tuple_info)
{
    if (info->iph->version == 4) {
        tuple_info->ipv4.saddr = info->iph->saddr;
        tuple_info->ipv4.daddr = info->iph->daddr;
        tuple_info->ipv4.sport = info->tcph->source;
        tuple_info->ipv4.dport = info->tcph->dest;
    } else {
        bpf_memcpy((__u8 *)tuple_info->ipv6.saddr, info->ip6h->saddr.in6_u.u6_addr8, IPV6_ADDR_LEN);
        bpf_memcpy((__u8 *)tuple_info->ipv6.daddr, info->ip6h->daddr.in6_u.u6_addr8, IPV6_ADDR_LEN);
        tuple_info->ipv6.sport = info->tcph->source;
        tuple_info->ipv6.dport = info->tcph->dest;
    }
}

static inline void shutdown_tuple(struct xdp_info *info)
{
    info->tcph->fin = 0;
    info->tcph->syn = 0;
    info->tcph->rst = 1;
    info->tcph->psh = 0;
    info->tcph->ack = 0;
}

static inline int should_shutdown(struct xdp_info *info, struct bpf_sock_tuple *tuple_info)
{
    __u32 *value = bpf_map_lookup_elem(&map_of_auth, tuple_info);
    if (value) {
        if (info->iph->version == 4)
            BPF_LOG(
                INFO,
                XDP,
                "auth denied, src ip: %s, port: %u\n",
                ip2str(&tuple_info->ipv4.saddr, true),
                bpf_ntohs(tuple_info->ipv4.sport));
        else
            BPF_LOG(
                INFO,
                XDP,
                "auth denied, src ip: %s, port: %u\n",
                ip2str(&tuple_info->ipv6.saddr[0], false),
                bpf_ntohs(tuple_info->ipv6.sport));
        bpf_map_delete_elem(&map_of_auth, tuple_info);
        return AUTH_FORBID;
    }
    return AUTH_PASS;
}

static inline int parser_xdp_info(struct xdp_md *ctx, struct xdp_info *info)
{
    void *begin = (void *)(long)(ctx->data);
    void *end = (void *)(long)(ctx->data_end);

    // eth header
    info->ethh = (struct ethhdr *)begin;
    if ((void *)(info->ethh + 1) > end)
        return PARSER_FAILED;

    // ip4|ip6 header
    begin = info->ethh + 1;
    if ((begin + 1) > end)
        return PARSER_FAILED;
    if (((struct iphdr *)begin)->version == 4) {
        info->iph = (struct iphdr *)begin;
        if ((void *)(info->iph + 1) > end || (info->iph->protocol != IPPROTO_TCP))
            return PARSER_FAILED;
        begin = (info->iph + 1);
    } else if (((struct iphdr *)begin)->version == 6) {
        info->ip6h = (struct ipv6hdr *)begin;
        if ((void *)(info->ip6h + 1) > end || (info->ip6h->nexthdr != IPPROTO_TCP))
            return PARSER_FAILED;
        begin = (info->ip6h + 1);
    } else
        return PARSER_FAILED;

    info->tcph = (struct tcphdr *)begin;
    if ((void *)(info->tcph + 1) > end)
        return PARSER_FAILED;
    return PARSER_SUCC;
}

static inline int xdp_deny_packet(struct xdp_info *info, struct bpf_sock_tuple *tuple_info)
{
    if (info->iph != NULL && info->iph->version == 4) {
        BPF_LOG(
            INFO,
            XDP,
            "auth denied, src ip: %s, dst ip %s, dst port: %u\n",
            ip2str(&tuple_info->ipv4.saddr, true),
            ip2str(&tuple_info->ipv4.daddr, true),
            bpf_ntohs(tuple_info->ipv4.dport));
    } else {
        BPF_LOG(
            INFO,
            XDP,
            "auth denied, src ip: %s, dst ip %s, dst port: %u\n",
            ip2str(&tuple_info->ipv6.saddr[0], false),
            ip2str(&tuple_info->ipv6.daddr[0], false),
            bpf_ntohs(tuple_info->ipv6.dport));
    }
    return XDP_DROP;
}

static inline wl_policies_v *get_workload_policies(struct xdp_info *info, struct bpf_sock_tuple *tuple_info)
{
    frontend_key frontend_k = {};
    frontend_value *frontend_v = NULL;
    __u32 workload_uid = 0;

    if (info->iph->version == 4) {
        frontend_k.addr.ip4 = tuple_info->ipv4.daddr;
    } else if (is_ipv4_mapped_addr(tuple_info->ipv6.daddr)) {
        frontend_k.addr.ip4 = tuple_info->ipv6.daddr[3];
    } else {
        bpf_memcpy(frontend_k.addr.ip6, tuple_info->ipv6.daddr, IPV6_ADDR_LEN);
    }
    frontend_v = kmesh_map_lookup_elem(&map_of_frontend, &frontend_k);
    if (!frontend_v) {
        BPF_LOG(INFO, XDP, "failed to get frontend in xdp");
        return AUTH_ALLOW;
    }
    workload_uid = frontend_v->upstream_id;
    return get_workload_policies_by_uid(workload_uid);
}

static inline int match_workload_policy(struct xdp_info *info, struct bpf_sock_tuple *tuple_info)
{
    int ret = 0;
    wl_policies_v *policies;
    __u32 policyId;
    Istio__Security__Authorization *policy;

    policies = get_workload_policies(info, tuple_info);
    if (!policies) {
        return AUTH_ALLOW;
    }

    for (int i = 0; i < MAX_MEMBER_NUM_PER_POLICY; i++) {
        policyId = policies->policyIds[i];
        if (policyId != 0) {
            policy = map_lookup_authz(policyId);
            if (!policy) {
                continue;
            }
            if (do_auth(policy, info, tuple_info) == AUTH_DENY) {
                BPF_LOG(ERR, AUTH, "policy %u manage result deny\n", policyId);
                return AUTH_DENY;
            }
        }
    }
    return AUTH_ALLOW;
}

static inline int xdp_rbac_manage(struct xdp_info *info, struct bpf_sock_tuple *tuple_info)
{
    return match_workload_policy(info, tuple_info);
}

SEC("xdp_auth")
int xdp_shutdown(struct xdp_md *ctx)
{
    struct xdp_info info = {0};
    struct bpf_sock_tuple tuple_info = {0};

    if (parser_xdp_info(ctx, &info) == PARSER_FAILED)
        return XDP_PASS;
    if (info.iph->version != 4 && info.iph->version != 6)
        return XDP_PASS;

    // never failed
    parser_tuple(&info, &tuple_info);
    // Before the authentication types supported by eBPF XDP are fully implemented,
    // this section only processes AUTH_DENY. If get AUTH_ALLOW,
    // it will still depend on the user-space authentication process to match other rule types.
    if (xdp_rbac_manage(&info, &tuple_info) == AUTH_DENY) {
        return xdp_deny_packet(&info, &tuple_info);
    }

    if (should_shutdown(&info, &tuple_info) == AUTH_FORBID)
        shutdown_tuple(&info);

    // If auth denied, it still returns XDP_PASS here, so next time when a client package is
    // sent to server, it will be shutdown since server's RST has been set
    return XDP_PASS;
}

char _license[] SEC("license") = "Dual BSD/GPL";
int _version SEC("version") = 1;