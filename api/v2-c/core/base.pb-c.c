/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: api/core/base.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "core/base.pb-c.h"
static const ProtobufCEnumValue core__api_status__enum_values_by_number[6] =
{
  { "NONE", "CORE__API_STATUS__NONE", 0 },
  { "DELETE", "CORE__API_STATUS__DELETE", 1 },
  { "UPDATE", "CORE__API_STATUS__UPDATE", 2 },
  { "UNCHANGED", "CORE__API_STATUS__UNCHANGED", 3 },
  { "ALL", "CORE__API_STATUS__ALL", 4 },
  { "WAITING", "CORE__API_STATUS__WAITING", 5 },
};
static const ProtobufCIntRange core__api_status__value_ranges[] = {
{0, 0},{0, 6}
};
static const ProtobufCEnumValueIndex core__api_status__enum_values_by_name[6] =
{
  { "ALL", 4 },
  { "DELETE", 1 },
  { "NONE", 0 },
  { "UNCHANGED", 3 },
  { "UPDATE", 2 },
  { "WAITING", 5 },
};
const ProtobufCEnumDescriptor core__api_status__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "core.ApiStatus",
  "ApiStatus",
  "Core__ApiStatus",
  "core",
  6,
  core__api_status__enum_values_by_number,
  6,
  core__api_status__enum_values_by_name,
  1,
  core__api_status__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCEnumValue core__routing_priority__enum_values_by_number[2] =
{
  { "DEFAULT", "CORE__ROUTING_PRIORITY__DEFAULT", 0 },
  { "HIGH", "CORE__ROUTING_PRIORITY__HIGH", 1 },
};
static const ProtobufCIntRange core__routing_priority__value_ranges[] = {
{0, 0},{0, 2}
};
static const ProtobufCEnumValueIndex core__routing_priority__enum_values_by_name[2] =
{
  { "DEFAULT", 0 },
  { "HIGH", 1 },
};
const ProtobufCEnumDescriptor core__routing_priority__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "core.RoutingPriority",
  "RoutingPriority",
  "Core__RoutingPriority",
  "core",
  2,
  core__routing_priority__enum_values_by_number,
  2,
  core__routing_priority__enum_values_by_name,
  1,
  core__routing_priority__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
