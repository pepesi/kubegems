// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: kubegems/services/v1/tenant_services.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_kubegems_services_v1_tenant_services_proto protoreflect.FileDescriptor

var file_kubegems_services_v1_tenant_services_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6b, 0x75,
	0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xcd, 0x05, 0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x76, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x26, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x7a, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x7a, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d,
	0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x1a,
	0x11, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x6d, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x12, 0x24, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d,
	0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0xdc, 0x01, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x3a, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d,
	0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x01, 0x2a, 0x22, 0x34, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x7d, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x42, 0x37, 0x5a, 0x35, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x69, 0x6f, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_kubegems_services_v1_tenant_services_proto_goTypes = []interface{}{
	(*CreateTenantRequest)(nil),                      // 0: kubegems.datas.v1.CreateTenantRequest
	(*DeleteTenantRequest)(nil),                      // 1: kubegems.datas.v1.DeleteTenantRequest
	(*ModifyTenantRequest)(nil),                      // 2: kubegems.datas.v1.ModifyTenantRequest
	(*ListTenantRequest)(nil),                        // 3: kubegems.datas.v1.ListTenantRequest
	(*CreateTenantClusterResourceQuotaRequest)(nil),  // 4: kubegems.datas.v1.CreateTenantClusterResourceQuotaRequest
	(*CreateTenantResponse)(nil),                     // 5: kubegems.datas.v1.CreateTenantResponse
	(*DeleteTenantResponse)(nil),                     // 6: kubegems.datas.v1.DeleteTenantResponse
	(*ModifyTenantResponse)(nil),                     // 7: kubegems.datas.v1.ModifyTenantResponse
	(*ListTenantResponse)(nil),                       // 8: kubegems.datas.v1.ListTenantResponse
	(*CreateTenantClusterResourceQuotaResponse)(nil), // 9: kubegems.datas.v1.CreateTenantClusterResourceQuotaResponse
}
var file_kubegems_services_v1_tenant_services_proto_depIdxs = []int32{
	0, // 0: kubegems.services.v1.TenantService.CreateTenant:input_type -> kubegems.datas.v1.CreateTenantRequest
	1, // 1: kubegems.services.v1.TenantService.DeleteTenant:input_type -> kubegems.datas.v1.DeleteTenantRequest
	2, // 2: kubegems.services.v1.TenantService.ModifyTenant:input_type -> kubegems.datas.v1.ModifyTenantRequest
	3, // 3: kubegems.services.v1.TenantService.ListTenant:input_type -> kubegems.datas.v1.ListTenantRequest
	4, // 4: kubegems.services.v1.TenantService.CreateTenantClusterResourceQuota:input_type -> kubegems.datas.v1.CreateTenantClusterResourceQuotaRequest
	5, // 5: kubegems.services.v1.TenantService.CreateTenant:output_type -> kubegems.datas.v1.CreateTenantResponse
	6, // 6: kubegems.services.v1.TenantService.DeleteTenant:output_type -> kubegems.datas.v1.DeleteTenantResponse
	7, // 7: kubegems.services.v1.TenantService.ModifyTenant:output_type -> kubegems.datas.v1.ModifyTenantResponse
	8, // 8: kubegems.services.v1.TenantService.ListTenant:output_type -> kubegems.datas.v1.ListTenantResponse
	9, // 9: kubegems.services.v1.TenantService.CreateTenantClusterResourceQuota:output_type -> kubegems.datas.v1.CreateTenantClusterResourceQuotaResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kubegems_services_v1_tenant_services_proto_init() }
func file_kubegems_services_v1_tenant_services_proto_init() {
	if File_kubegems_services_v1_tenant_services_proto != nil {
		return
	}
	file_kubegems_datas_v1_tenant_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kubegems_services_v1_tenant_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kubegems_services_v1_tenant_services_proto_goTypes,
		DependencyIndexes: file_kubegems_services_v1_tenant_services_proto_depIdxs,
	}.Build()
	File_kubegems_services_v1_tenant_services_proto = out.File
	file_kubegems_services_v1_tenant_services_proto_rawDesc = nil
	file_kubegems_services_v1_tenant_services_proto_goTypes = nil
	file_kubegems_services_v1_tenant_services_proto_depIdxs = nil
}
