// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.17.3
// source: api/route/route_components.proto

package route

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VirtualHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Domains []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	Routes  []*Route `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *VirtualHost) Reset() {
	*x = VirtualHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_route_route_components_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualHost) ProtoMessage() {}

func (x *VirtualHost) ProtoReflect() protoreflect.Message {
	mi := &file_api_route_route_components_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualHost.ProtoReflect.Descriptor instead.
func (*VirtualHost) Descriptor() ([]byte, []int) {
	return file_api_route_route_components_proto_rawDescGZIP(), []int{0}
}

func (x *VirtualHost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VirtualHost) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *VirtualHost) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string       `protobuf:"bytes,14,opt,name=name,proto3" json:"name,omitempty"`
	Match *RouteMatch  `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Route *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_route_route_components_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_api_route_route_components_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_api_route_route_components_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Route) GetMatch() *RouteMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *Route) GetRoute() *RouteAction {
	if x != nil {
		return x.Route
	}
	return nil
}

type RouteMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix        string           `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	CaseSensitive bool             `protobuf:"varint,4,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"`
	Headers       []*HeaderMatcher `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *RouteMatch) Reset() {
	*x = RouteMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_route_route_components_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatch) ProtoMessage() {}

func (x *RouteMatch) ProtoReflect() protoreflect.Message {
	mi := &file_api_route_route_components_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatch.ProtoReflect.Descriptor instead.
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return file_api_route_route_components_proto_rawDescGZIP(), []int{2}
}

func (x *RouteMatch) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *RouteMatch) GetCaseSensitive() bool {
	if x != nil {
		return x.CaseSensitive
	}
	return false
}

func (x *RouteMatch) GetHeaders() []*HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

type RouteAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// the matched prefix (or path) should be swapped with this value.
	PrefixRewrite string       `protobuf:"bytes,5,opt,name=prefix_rewrite,json=prefixRewrite,proto3" json:"prefix_rewrite,omitempty"`
	Timeout       uint32       `protobuf:"varint,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	RetryPolicy   *RetryPolicy `protobuf:"bytes,9,opt,name=retry_policy,json=retryPolicy,proto3" json:"retry_policy,omitempty"`
}

func (x *RouteAction) Reset() {
	*x = RouteAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_route_route_components_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteAction) ProtoMessage() {}

func (x *RouteAction) ProtoReflect() protoreflect.Message {
	mi := &file_api_route_route_components_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteAction.ProtoReflect.Descriptor instead.
func (*RouteAction) Descriptor() ([]byte, []int) {
	return file_api_route_route_components_proto_rawDescGZIP(), []int{3}
}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (x *RouteAction) GetCluster() string {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (x *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (x *RouteAction) GetPrefixRewrite() string {
	if x != nil {
		return x.PrefixRewrite
	}
	return ""
}

func (x *RouteAction) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *RouteAction) GetRetryPolicy() *RetryPolicy {
	if x != nil {
		return x.RetryPolicy
	}
	return nil
}

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	// Indicates the upstream cluster to which the request should be routed to.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	// Multiple upstream clusters can be specified for a given route. The
	// request is routed to one of the upstream clusters based on weights
	// assigned to each cluster.
	WeightedClusters *WeightedCluster `protobuf:"bytes,3,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

type RetryPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumRetries uint32 `protobuf:"varint,2,opt,name=num_retries,json=numRetries,proto3" json:"num_retries,omitempty"` //RetryPriority retry_priority = 4;
}

func (x *RetryPolicy) Reset() {
	*x = RetryPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_route_route_components_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryPolicy) ProtoMessage() {}

func (x *RetryPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_api_route_route_components_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryPolicy.ProtoReflect.Descriptor instead.
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return file_api_route_route_components_proto_rawDescGZIP(), []int{4}
}

func (x *RetryPolicy) GetNumRetries() uint32 {
	if x != nil {
		return x.NumRetries
	}
	return 0
}

type WeightedCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *WeightedCluster) Reset() {
	*x = WeightedCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_route_route_components_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightedCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightedCluster) ProtoMessage() {}

func (x *WeightedCluster) ProtoReflect() protoreflect.Message {
	mi := &file_api_route_route_components_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightedCluster.ProtoReflect.Descriptor instead.
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return file_api_route_route_components_proto_rawDescGZIP(), []int{5}
}

func (x *WeightedCluster) GetClusters() []*ClusterWeight {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type ClusterWeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *ClusterWeight) Reset() {
	*x = ClusterWeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_route_route_components_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterWeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterWeight) ProtoMessage() {}

func (x *ClusterWeight) ProtoReflect() protoreflect.Message {
	mi := &file_api_route_route_components_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterWeight.ProtoReflect.Descriptor instead.
func (*ClusterWeight) Descriptor() ([]byte, []int) {
	return file_api_route_route_components_proto_rawDescGZIP(), []int{6}
}

func (x *ClusterWeight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterWeight) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type HeaderMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the name of the header in the request.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to HeaderMatchSpecifier:
	//	*HeaderMatcher_ExactMatch
	//	*HeaderMatcher_PrefixMatch
	HeaderMatchSpecifier isHeaderMatcher_HeaderMatchSpecifier `protobuf_oneof:"header_match_specifier"`
}

func (x *HeaderMatcher) Reset() {
	*x = HeaderMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_route_route_components_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderMatcher) ProtoMessage() {}

func (x *HeaderMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_api_route_route_components_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderMatcher.ProtoReflect.Descriptor instead.
func (*HeaderMatcher) Descriptor() ([]byte, []int) {
	return file_api_route_route_components_proto_rawDescGZIP(), []int{7}
}

func (x *HeaderMatcher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *HeaderMatcher) GetHeaderMatchSpecifier() isHeaderMatcher_HeaderMatchSpecifier {
	if m != nil {
		return m.HeaderMatchSpecifier
	}
	return nil
}

func (x *HeaderMatcher) GetExactMatch() string {
	if x, ok := x.GetHeaderMatchSpecifier().(*HeaderMatcher_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (x *HeaderMatcher) GetPrefixMatch() string {
	if x, ok := x.GetHeaderMatchSpecifier().(*HeaderMatcher_PrefixMatch); ok {
		return x.PrefixMatch
	}
	return ""
}

type isHeaderMatcher_HeaderMatchSpecifier interface {
	isHeaderMatcher_HeaderMatchSpecifier()
}

type HeaderMatcher_ExactMatch struct {
	// If specified, header match will be performed based on the value of the header.
	ExactMatch string `protobuf:"bytes,4,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type HeaderMatcher_PrefixMatch struct {
	// If specified, header match will be performed based on the prefix of the header value.
	PrefixMatch string `protobuf:"bytes,9,opt,name=prefix_match,json=prefixMatch,proto3,oneof"`
}

func (*HeaderMatcher_ExactMatch) isHeaderMatcher_HeaderMatchSpecifier() {}

func (*HeaderMatcher_PrefixMatch) isHeaderMatcher_HeaderMatchSpecifier() {}

var File_api_route_route_components_proto protoreflect.FileDescriptor

var file_api_route_route_components_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x61, 0x0a, 0x0b, 0x56, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x05,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x7b, 0x0a, 0x0a,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0xfd, 0x01, 0x0a, 0x0b, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x11, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65,
	0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65,
	0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x10, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x35, 0x0a,
	0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x42, 0x13, 0x0a, 0x11, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x0b, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e,
	0x75, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x0f, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3b,
	0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x0d,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x61, 0x63, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x18, 0x0a, 0x16, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x42, 0x21, 0x5a, 0x1f, 0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x65, 0x74,
	0x2f, 0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_route_route_components_proto_rawDescOnce sync.Once
	file_api_route_route_components_proto_rawDescData = file_api_route_route_components_proto_rawDesc
)

func file_api_route_route_components_proto_rawDescGZIP() []byte {
	file_api_route_route_components_proto_rawDescOnce.Do(func() {
		file_api_route_route_components_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_route_route_components_proto_rawDescData)
	})
	return file_api_route_route_components_proto_rawDescData
}

var file_api_route_route_components_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_route_route_components_proto_goTypes = []interface{}{
	(*VirtualHost)(nil),     // 0: route.VirtualHost
	(*Route)(nil),           // 1: route.Route
	(*RouteMatch)(nil),      // 2: route.RouteMatch
	(*RouteAction)(nil),     // 3: route.RouteAction
	(*RetryPolicy)(nil),     // 4: route.RetryPolicy
	(*WeightedCluster)(nil), // 5: route.WeightedCluster
	(*ClusterWeight)(nil),   // 6: route.ClusterWeight
	(*HeaderMatcher)(nil),   // 7: route.HeaderMatcher
}
var file_api_route_route_components_proto_depIdxs = []int32{
	1, // 0: route.VirtualHost.routes:type_name -> route.Route
	2, // 1: route.Route.match:type_name -> route.RouteMatch
	3, // 2: route.Route.route:type_name -> route.RouteAction
	7, // 3: route.RouteMatch.headers:type_name -> route.HeaderMatcher
	5, // 4: route.RouteAction.weighted_clusters:type_name -> route.WeightedCluster
	4, // 5: route.RouteAction.retry_policy:type_name -> route.RetryPolicy
	6, // 6: route.WeightedCluster.clusters:type_name -> route.ClusterWeight
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_route_route_components_proto_init() }
func file_api_route_route_components_proto_init() {
	if File_api_route_route_components_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_route_route_components_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualHost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_route_route_components_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_route_route_components_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_route_route_components_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_route_route_components_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_route_route_components_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightedCluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_route_route_components_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterWeight); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_route_route_components_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_route_route_components_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
	file_api_route_route_components_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*HeaderMatcher_ExactMatch)(nil),
		(*HeaderMatcher_PrefixMatch)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_route_route_components_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_route_route_components_proto_goTypes,
		DependencyIndexes: file_api_route_route_components_proto_depIdxs,
		MessageInfos:      file_api_route_route_components_proto_msgTypes,
	}.Build()
	File_api_route_route_components_proto = out.File
	file_api_route_route_components_proto_rawDesc = nil
	file_api_route_route_components_proto_goTypes = nil
	file_api_route_route_components_proto_depIdxs = nil
}
