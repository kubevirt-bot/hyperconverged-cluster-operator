// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/asset/v1/asset_service.proto

package asset

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	_ "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Asset content type.
type ContentType int32

const (
	// Unspecified content type.
	ContentType_CONTENT_TYPE_UNSPECIFIED ContentType = 0
	// Resource metadata.
	ContentType_RESOURCE ContentType = 1
	// The actual IAM policy set on a resource.
	ContentType_IAM_POLICY ContentType = 2
	// The Cloud Organization Policy set on an asset.
	ContentType_ORG_POLICY ContentType = 4
	// The Cloud Access context mananger Policy set on an asset.
	ContentType_ACCESS_POLICY ContentType = 5
)

var ContentType_name = map[int32]string{
	0: "CONTENT_TYPE_UNSPECIFIED",
	1: "RESOURCE",
	2: "IAM_POLICY",
	4: "ORG_POLICY",
	5: "ACCESS_POLICY",
}

var ContentType_value = map[string]int32{
	"CONTENT_TYPE_UNSPECIFIED": 0,
	"RESOURCE":                 1,
	"IAM_POLICY":               2,
	"ORG_POLICY":               4,
	"ACCESS_POLICY":            5,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}

func (ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5104159f18b2092a, []int{0}
}

// Export asset request.
type ExportAssetsRequest struct {
	// Required. The relative name of the root asset. This can only be an
	// organization number (such as "organizations/123"), a project ID (such as
	// "projects/my-project-id"), or a project number (such as "projects/12345"),
	// or a folder number (such as "folders/123").
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Timestamp to take an asset snapshot. This can only be set to a timestamp
	// between 2018-10-02 UTC (inclusive) and the current time. If not specified,
	// the current time will be used. Due to delays in resource data collection
	// and indexing, there is a volatile window during which running the same
	// query may get different results.
	ReadTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	// A list of asset types of which to take a snapshot for. For example:
	// "compute.googleapis.com/Disk". If specified, only matching assets will be
	// returned. See [Introduction to Cloud Asset
	// Inventory](https://cloud.google.com/asset-inventory/docs/overview)
	// for all supported asset types.
	AssetTypes []string `protobuf:"bytes,3,rep,name=asset_types,json=assetTypes,proto3" json:"asset_types,omitempty"`
	// Asset content type. If not specified, no content but the asset name will be
	// returned.
	ContentType ContentType `protobuf:"varint,4,opt,name=content_type,json=contentType,proto3,enum=google.cloud.asset.v1.ContentType" json:"content_type,omitempty"`
	// Required. Output configuration indicating where the results will be output
	// to. All results will be in newline delimited JSON format.
	OutputConfig         *OutputConfig `protobuf:"bytes,5,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportAssetsRequest) Reset()         { *m = ExportAssetsRequest{} }
func (m *ExportAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*ExportAssetsRequest) ProtoMessage()    {}
func (*ExportAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5104159f18b2092a, []int{0}
}

func (m *ExportAssetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportAssetsRequest.Unmarshal(m, b)
}
func (m *ExportAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportAssetsRequest.Marshal(b, m, deterministic)
}
func (m *ExportAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportAssetsRequest.Merge(m, src)
}
func (m *ExportAssetsRequest) XXX_Size() int {
	return xxx_messageInfo_ExportAssetsRequest.Size(m)
}
func (m *ExportAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportAssetsRequest proto.InternalMessageInfo

func (m *ExportAssetsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ExportAssetsRequest) GetReadTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

func (m *ExportAssetsRequest) GetAssetTypes() []string {
	if m != nil {
		return m.AssetTypes
	}
	return nil
}

func (m *ExportAssetsRequest) GetContentType() ContentType {
	if m != nil {
		return m.ContentType
	}
	return ContentType_CONTENT_TYPE_UNSPECIFIED
}

func (m *ExportAssetsRequest) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

// The export asset response. This message is returned by the
// [google.longrunning.Operations.GetOperation][google.longrunning.Operations.GetOperation] method in the returned
// [google.longrunning.Operation.response][google.longrunning.Operation.response] field.
type ExportAssetsResponse struct {
	// Time the snapshot was taken.
	ReadTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	// Output configuration indicating where the results were output to.
	// All results are in JSON format.
	OutputConfig         *OutputConfig `protobuf:"bytes,2,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportAssetsResponse) Reset()         { *m = ExportAssetsResponse{} }
func (m *ExportAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*ExportAssetsResponse) ProtoMessage()    {}
func (*ExportAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5104159f18b2092a, []int{1}
}

func (m *ExportAssetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportAssetsResponse.Unmarshal(m, b)
}
func (m *ExportAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportAssetsResponse.Marshal(b, m, deterministic)
}
func (m *ExportAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportAssetsResponse.Merge(m, src)
}
func (m *ExportAssetsResponse) XXX_Size() int {
	return xxx_messageInfo_ExportAssetsResponse.Size(m)
}
func (m *ExportAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportAssetsResponse proto.InternalMessageInfo

func (m *ExportAssetsResponse) GetReadTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

func (m *ExportAssetsResponse) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

// Batch get assets history request.
type BatchGetAssetsHistoryRequest struct {
	// Required. The relative name of the root asset. It can only be an
	// organization number (such as "organizations/123"), a project ID (such as
	// "projects/my-project-id")", or a project number (such as "projects/12345").
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// A list of the full names of the assets. For example:
	// `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`.
	// See [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// and [Resource Name
	// Format](https://cloud.google.com/asset-inventory/docs/resource-name-format)
	// for more info.
	//
	// The request becomes a no-op if the asset name list is empty, and the max
	// size of the asset name list is 100 in one request.
	AssetNames []string `protobuf:"bytes,2,rep,name=asset_names,json=assetNames,proto3" json:"asset_names,omitempty"`
	// Optional. The content type.
	ContentType ContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=google.cloud.asset.v1.ContentType" json:"content_type,omitempty"`
	// Optional. The time window for the asset history. Both start_time and
	// end_time are optional and if set, it must be after 2018-10-02 UTC. If
	// end_time is not set, it is default to current timestamp. If start_time is
	// not set, the snapshot of the assets at end_time will be returned. The
	// returned results contain all temporal assets whose time window overlap with
	// read_time_window.
	ReadTimeWindow       *TimeWindow `protobuf:"bytes,4,opt,name=read_time_window,json=readTimeWindow,proto3" json:"read_time_window,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BatchGetAssetsHistoryRequest) Reset()         { *m = BatchGetAssetsHistoryRequest{} }
func (m *BatchGetAssetsHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*BatchGetAssetsHistoryRequest) ProtoMessage()    {}
func (*BatchGetAssetsHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5104159f18b2092a, []int{2}
}

func (m *BatchGetAssetsHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Unmarshal(m, b)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Marshal(b, m, deterministic)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetAssetsHistoryRequest.Merge(m, src)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Size(m)
}
func (m *BatchGetAssetsHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetAssetsHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetAssetsHistoryRequest proto.InternalMessageInfo

func (m *BatchGetAssetsHistoryRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *BatchGetAssetsHistoryRequest) GetAssetNames() []string {
	if m != nil {
		return m.AssetNames
	}
	return nil
}

func (m *BatchGetAssetsHistoryRequest) GetContentType() ContentType {
	if m != nil {
		return m.ContentType
	}
	return ContentType_CONTENT_TYPE_UNSPECIFIED
}

func (m *BatchGetAssetsHistoryRequest) GetReadTimeWindow() *TimeWindow {
	if m != nil {
		return m.ReadTimeWindow
	}
	return nil
}

// Batch get assets history response.
type BatchGetAssetsHistoryResponse struct {
	// A list of assets with valid time windows.
	Assets               []*TemporalAsset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BatchGetAssetsHistoryResponse) Reset()         { *m = BatchGetAssetsHistoryResponse{} }
func (m *BatchGetAssetsHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*BatchGetAssetsHistoryResponse) ProtoMessage()    {}
func (*BatchGetAssetsHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5104159f18b2092a, []int{3}
}

func (m *BatchGetAssetsHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Unmarshal(m, b)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Marshal(b, m, deterministic)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetAssetsHistoryResponse.Merge(m, src)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Size(m)
}
func (m *BatchGetAssetsHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetAssetsHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetAssetsHistoryResponse proto.InternalMessageInfo

func (m *BatchGetAssetsHistoryResponse) GetAssets() []*TemporalAsset {
	if m != nil {
		return m.Assets
	}
	return nil
}

// Output configuration for export assets destination.
type OutputConfig struct {
	// Asset export destination.
	//
	// Types that are valid to be assigned to Destination:
	//	*OutputConfig_GcsDestination
	//	*OutputConfig_BigqueryDestination
	Destination          isOutputConfig_Destination `protobuf_oneof:"destination"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *OutputConfig) Reset()         { *m = OutputConfig{} }
func (m *OutputConfig) String() string { return proto.CompactTextString(m) }
func (*OutputConfig) ProtoMessage()    {}
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5104159f18b2092a, []int{4}
}

func (m *OutputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputConfig.Unmarshal(m, b)
}
func (m *OutputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputConfig.Marshal(b, m, deterministic)
}
func (m *OutputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputConfig.Merge(m, src)
}
func (m *OutputConfig) XXX_Size() int {
	return xxx_messageInfo_OutputConfig.Size(m)
}
func (m *OutputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputConfig proto.InternalMessageInfo

type isOutputConfig_Destination interface {
	isOutputConfig_Destination()
}

type OutputConfig_GcsDestination struct {
	GcsDestination *GcsDestination `protobuf:"bytes,1,opt,name=gcs_destination,json=gcsDestination,proto3,oneof"`
}

type OutputConfig_BigqueryDestination struct {
	BigqueryDestination *BigQueryDestination `protobuf:"bytes,2,opt,name=bigquery_destination,json=bigqueryDestination,proto3,oneof"`
}

func (*OutputConfig_GcsDestination) isOutputConfig_Destination() {}

func (*OutputConfig_BigqueryDestination) isOutputConfig_Destination() {}

func (m *OutputConfig) GetDestination() isOutputConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *OutputConfig) GetGcsDestination() *GcsDestination {
	if x, ok := m.GetDestination().(*OutputConfig_GcsDestination); ok {
		return x.GcsDestination
	}
	return nil
}

func (m *OutputConfig) GetBigqueryDestination() *BigQueryDestination {
	if x, ok := m.GetDestination().(*OutputConfig_BigqueryDestination); ok {
		return x.BigqueryDestination
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutputConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutputConfig_GcsDestination)(nil),
		(*OutputConfig_BigqueryDestination)(nil),
	}
}

// A Cloud Storage location.
type GcsDestination struct {
	// Required.
	//
	// Types that are valid to be assigned to ObjectUri:
	//	*GcsDestination_Uri
	//	*GcsDestination_UriPrefix
	ObjectUri            isGcsDestination_ObjectUri `protobuf_oneof:"object_uri"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GcsDestination) Reset()         { *m = GcsDestination{} }
func (m *GcsDestination) String() string { return proto.CompactTextString(m) }
func (*GcsDestination) ProtoMessage()    {}
func (*GcsDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_5104159f18b2092a, []int{5}
}

func (m *GcsDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsDestination.Unmarshal(m, b)
}
func (m *GcsDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsDestination.Marshal(b, m, deterministic)
}
func (m *GcsDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsDestination.Merge(m, src)
}
func (m *GcsDestination) XXX_Size() int {
	return xxx_messageInfo_GcsDestination.Size(m)
}
func (m *GcsDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsDestination.DiscardUnknown(m)
}

var xxx_messageInfo_GcsDestination proto.InternalMessageInfo

type isGcsDestination_ObjectUri interface {
	isGcsDestination_ObjectUri()
}

type GcsDestination_Uri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3,oneof"`
}

type GcsDestination_UriPrefix struct {
	UriPrefix string `protobuf:"bytes,2,opt,name=uri_prefix,json=uriPrefix,proto3,oneof"`
}

func (*GcsDestination_Uri) isGcsDestination_ObjectUri() {}

func (*GcsDestination_UriPrefix) isGcsDestination_ObjectUri() {}

func (m *GcsDestination) GetObjectUri() isGcsDestination_ObjectUri {
	if m != nil {
		return m.ObjectUri
	}
	return nil
}

func (m *GcsDestination) GetUri() string {
	if x, ok := m.GetObjectUri().(*GcsDestination_Uri); ok {
		return x.Uri
	}
	return ""
}

func (m *GcsDestination) GetUriPrefix() string {
	if x, ok := m.GetObjectUri().(*GcsDestination_UriPrefix); ok {
		return x.UriPrefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GcsDestination) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GcsDestination_Uri)(nil),
		(*GcsDestination_UriPrefix)(nil),
	}
}

// A BigQuery destination.
type BigQueryDestination struct {
	// Required. The BigQuery dataset in format
	// "projects/projectId/datasets/datasetId", to which the snapshot result
	// should be exported. If this dataset does not exist, the export call returns
	// an error.
	Dataset string `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// Required. The BigQuery table to which the snapshot result should be
	// written. If this table does not exist, a new table with the given name
	// will be created.
	Table string `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	// If the destination table already exists and this flag is `TRUE`, the
	// table will be overwritten by the contents of assets snapshot. If the flag
	// is not set and the destination table already exists, the export call
	// returns an error.
	Force                bool     `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BigQueryDestination) Reset()         { *m = BigQueryDestination{} }
func (m *BigQueryDestination) String() string { return proto.CompactTextString(m) }
func (*BigQueryDestination) ProtoMessage()    {}
func (*BigQueryDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_5104159f18b2092a, []int{6}
}

func (m *BigQueryDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigQueryDestination.Unmarshal(m, b)
}
func (m *BigQueryDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigQueryDestination.Marshal(b, m, deterministic)
}
func (m *BigQueryDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigQueryDestination.Merge(m, src)
}
func (m *BigQueryDestination) XXX_Size() int {
	return xxx_messageInfo_BigQueryDestination.Size(m)
}
func (m *BigQueryDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_BigQueryDestination.DiscardUnknown(m)
}

var xxx_messageInfo_BigQueryDestination proto.InternalMessageInfo

func (m *BigQueryDestination) GetDataset() string {
	if m != nil {
		return m.Dataset
	}
	return ""
}

func (m *BigQueryDestination) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *BigQueryDestination) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func init() {
	proto.RegisterEnum("google.cloud.asset.v1.ContentType", ContentType_name, ContentType_value)
	proto.RegisterType((*ExportAssetsRequest)(nil), "google.cloud.asset.v1.ExportAssetsRequest")
	proto.RegisterType((*ExportAssetsResponse)(nil), "google.cloud.asset.v1.ExportAssetsResponse")
	proto.RegisterType((*BatchGetAssetsHistoryRequest)(nil), "google.cloud.asset.v1.BatchGetAssetsHistoryRequest")
	proto.RegisterType((*BatchGetAssetsHistoryResponse)(nil), "google.cloud.asset.v1.BatchGetAssetsHistoryResponse")
	proto.RegisterType((*OutputConfig)(nil), "google.cloud.asset.v1.OutputConfig")
	proto.RegisterType((*GcsDestination)(nil), "google.cloud.asset.v1.GcsDestination")
	proto.RegisterType((*BigQueryDestination)(nil), "google.cloud.asset.v1.BigQueryDestination")
}

func init() {
	proto.RegisterFile("google/cloud/asset/v1/asset_service.proto", fileDescriptor_5104159f18b2092a)
}

var fileDescriptor_5104159f18b2092a = []byte{
	// 1003 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0xaf, 0x9d, 0x6d, 0x69, 0x5f, 0xd2, 0xd2, 0x9d, 0xb6, 0x22, 0x0d, 0xad, 0x9a, 0xf5, 0xf2,
	0x27, 0x1b, 0x09, 0x5b, 0xed, 0x22, 0x21, 0x15, 0x10, 0x4a, 0xb2, 0xa1, 0x2d, 0xa2, 0x4d, 0x70,
	0xb2, 0xbb, 0x5a, 0x54, 0x64, 0x39, 0xce, 0xc4, 0x1d, 0x70, 0x3c, 0xde, 0xf1, 0xb8, 0xd9, 0x0a,
	0x71, 0xd9, 0xaf, 0xc0, 0x09, 0x09, 0xf1, 0x01, 0x38, 0xf2, 0x21, 0x38, 0xf4, 0x84, 0xe0, 0x8e,
	0xf6, 0xc0, 0x37, 0xe0, 0xc6, 0x09, 0x79, 0xc6, 0x6e, 0x9d, 0x90, 0xac, 0xba, 0x12, 0xb7, 0xcc,
	0x6f, 0x7e, 0xef, 0x37, 0x3f, 0xbf, 0xf7, 0x66, 0x5e, 0xe0, 0x9e, 0x4b, 0xa9, 0xeb, 0x61, 0xc3,
	0xf1, 0x68, 0xd4, 0x37, 0xec, 0x30, 0xc4, 0xdc, 0x38, 0xdf, 0x95, 0x3f, 0xac, 0x10, 0xb3, 0x73,
	0xe2, 0x60, 0x3d, 0x60, 0x94, 0x53, 0xb4, 0x21, 0xa9, 0xba, 0xa0, 0xea, 0x82, 0xa1, 0x9f, 0xef,
	0x96, 0xb6, 0x12, 0x05, 0x3b, 0x20, 0x86, 0xed, 0xfb, 0x94, 0xdb, 0x9c, 0x50, 0x3f, 0x94, 0x41,
	0xa5, 0x37, 0x32, 0xbb, 0x8e, 0x47, 0xb0, 0xcf, 0x93, 0x8d, 0x9d, 0xcc, 0xc6, 0x80, 0x60, 0xaf,
	0x6f, 0xf5, 0xf0, 0x99, 0x7d, 0x4e, 0x28, 0x4b, 0x08, 0x9b, 0x19, 0x02, 0xc3, 0x21, 0x8d, 0x58,
	0xea, 0xa4, 0xa4, 0xbd, 0xc4, 0x74, 0x7a, 0xf0, 0xdd, 0x84, 0xe3, 0x51, 0xdf, 0x65, 0x91, 0xef,
	0x13, 0xdf, 0x35, 0x68, 0x80, 0xd9, 0x98, 0xbb, 0x37, 0x13, 0x92, 0x58, 0xf5, 0xa2, 0x81, 0x81,
	0x87, 0x01, 0xbf, 0x48, 0x36, 0xcb, 0x93, 0x9b, 0xd2, 0xe6, 0xd0, 0x0e, 0xbf, 0x99, 0xf8, 0x86,
	0x2b, 0x06, 0x27, 0x43, 0x1c, 0x72, 0x7b, 0x18, 0x48, 0x82, 0xf6, 0xab, 0x0a, 0x6b, 0xcd, 0x67,
	0x01, 0x65, 0xbc, 0x26, 0xbc, 0x99, 0xf8, 0x69, 0x84, 0x43, 0x8e, 0x3e, 0x81, 0x85, 0xc0, 0x66,
	0xd8, 0xe7, 0x45, 0xa5, 0xac, 0x54, 0x96, 0xea, 0xef, 0xbe, 0xa8, 0xa9, 0xff, 0xd4, 0xee, 0xa0,
	0x1d, 0xf1, 0x45, 0x32, 0xb5, 0x52, 0xdb, 0x0e, 0x48, 0xa8, 0x3b, 0x74, 0x68, 0x08, 0x01, 0x33,
	0x09, 0x43, 0x1f, 0xc0, 0x12, 0xc3, 0x76, 0xdf, 0x8a, 0x0f, 0x2c, 0xaa, 0x65, 0xa5, 0x92, 0xdf,
	0x2b, 0x25, 0x11, 0x7a, 0xea, 0x46, 0xef, 0xa6, 0x6e, 0xcc, 0xc5, 0x98, 0x1c, 0x2f, 0xd1, 0x0e,
	0xe4, 0x65, 0x6d, 0xf9, 0x45, 0x80, 0xc3, 0x62, 0xae, 0x9c, 0xab, 0x2c, 0x99, 0x20, 0xa0, 0x6e,
	0x8c, 0xa0, 0x26, 0x14, 0x1c, 0xea, 0x73, 0xec, 0x4b, 0x4a, 0xf1, 0x56, 0x59, 0xa9, 0xac, 0xec,
	0x69, 0xfa, 0xd4, 0xe2, 0xeb, 0x0d, 0x49, 0x8d, 0x43, 0xcd, 0xbc, 0x73, 0xbd, 0x40, 0xc7, 0xb0,
	0x4c, 0x23, 0x1e, 0x44, 0xdc, 0x72, 0xa8, 0x3f, 0x20, 0x6e, 0x71, 0x5e, 0x98, 0xbc, 0x3b, 0x43,
	0xa7, 0x25, 0xb8, 0x0d, 0x41, 0xad, 0xe7, 0x5e, 0xd4, 0x54, 0xb3, 0x40, 0x33, 0x90, 0xf6, 0x83,
	0x02, 0xeb, 0xe3, 0x89, 0x0c, 0x03, 0xea, 0x87, 0x78, 0x3c, 0x11, 0xca, 0x2b, 0x24, 0xe2, 0x70,
	0xd2, 0xa0, 0x7a, 0x63, 0x83, 0x13, 0xde, 0x7e, 0x52, 0x61, 0xab, 0x6e, 0x73, 0xe7, 0xec, 0x00,
	0x27, 0xee, 0x0e, 0x49, 0xc8, 0x29, 0xbb, 0xf8, 0xdf, 0xaa, 0x7d, 0x55, 0x34, 0xdf, 0x1e, 0xe2,
	0xb0, 0xa8, 0x66, 0x8a, 0x76, 0x12, 0x23, 0xe8, 0xb3, 0x89, 0xa2, 0xe5, 0x6e, 0x5a, 0xb4, 0x38,
	0xd7, 0xca, 0x78, 0xe5, 0xda, 0xb0, 0x7a, 0x95, 0x51, 0x6b, 0x44, 0xfc, 0x3e, 0x1d, 0x89, 0x26,
	0xc8, 0xef, 0xdd, 0x99, 0xa1, 0x17, 0xe7, 0xf3, 0xb1, 0x20, 0x4a, 0xb9, 0x95, 0x34, 0xc9, 0x12,
	0xd4, 0xbe, 0x82, 0xed, 0x19, 0xf9, 0x49, 0x8a, 0xf8, 0x11, 0x2c, 0xc8, 0xbb, 0x5b, 0x54, 0xca,
	0xb9, 0x4a, 0x7e, 0xef, 0xad, 0x59, 0x07, 0xe1, 0x61, 0x40, 0x99, 0xed, 0x25, 0xd9, 0x91, 0x31,
	0xda, 0x6f, 0x0a, 0x14, 0xb2, 0xe5, 0x41, 0x6d, 0x78, 0xdd, 0x75, 0x42, 0xab, 0x8f, 0x43, 0x4e,
	0x7c, 0x71, 0xdf, 0x93, 0xce, 0x78, 0x7b, 0x86, 0xee, 0x81, 0x13, 0x3e, 0xb8, 0x26, 0x1f, 0xce,
	0x99, 0x2b, 0xee, 0x18, 0x82, 0x2c, 0x58, 0xef, 0x11, 0xf7, 0x69, 0x84, 0xd9, 0xc5, 0x98, 0xac,
	0xec, 0x99, 0xea, 0x0c, 0xd9, 0x3a, 0x71, 0xbf, 0x88, 0x43, 0xc6, 0xb5, 0xd7, 0x52, 0xa5, 0x0c,
	0x5c, 0x5f, 0x86, 0x7c, 0x46, 0x57, 0xeb, 0xc0, 0xca, 0xb8, 0x27, 0x84, 0x20, 0x17, 0x31, 0x22,
	0x1b, 0xe8, 0x70, 0xce, 0x8c, 0x17, 0x68, 0x07, 0x20, 0x62, 0xc4, 0x0a, 0x18, 0x1e, 0x90, 0x67,
	0xc2, 0x4b, 0xbc, 0xb5, 0x14, 0x31, 0xd2, 0x16, 0x50, 0xbd, 0x00, 0x40, 0x7b, 0x5f, 0x63, 0x87,
	0x5b, 0x11, 0x23, 0x1a, 0x86, 0xb5, 0x29, 0x8e, 0xd0, 0x36, 0xbc, 0xd6, 0xb7, 0xb9, 0x1d, 0xe2,
	0xb4, 0x3d, 0xc5, 0xf5, 0x4b, 0x31, 0xb4, 0x09, 0xf3, 0xdc, 0xee, 0x79, 0xf2, 0x95, 0x49, 0x36,
	0x25, 0x82, 0xd6, 0x61, 0x7e, 0x40, 0x99, 0x23, 0xdb, 0x6d, 0xd1, 0x94, 0x8b, 0xaa, 0x07, 0xf9,
	0x4c, 0x83, 0xa1, 0x2d, 0x28, 0x36, 0x5a, 0x27, 0xdd, 0xe6, 0x49, 0xd7, 0xea, 0x3e, 0x69, 0x37,
	0xad, 0x87, 0x27, 0x9d, 0x76, 0xb3, 0x71, 0xf4, 0xe9, 0x51, 0xf3, 0xc1, 0xea, 0x1c, 0x2a, 0xc0,
	0xa2, 0xd9, 0xec, 0xb4, 0x1e, 0x9a, 0x8d, 0xe6, 0xaa, 0x82, 0x56, 0x00, 0x8e, 0x6a, 0xc7, 0x56,
	0xbb, 0xf5, 0xf9, 0x51, 0xe3, 0xc9, 0xaa, 0x1a, 0xaf, 0x5b, 0xe6, 0x41, 0xba, 0xbe, 0x85, 0x6e,
	0xc3, 0x72, 0xad, 0xd1, 0x68, 0x76, 0x3a, 0x29, 0x34, 0xbf, 0xf7, 0x77, 0x0e, 0x0a, 0xa2, 0x1d,
	0x3a, 0x72, 0x56, 0xa1, 0x3f, 0x15, 0x28, 0x64, 0x5f, 0x0a, 0x34, 0xab, 0x3a, 0x53, 0xde, 0xe5,
	0xd2, 0x76, 0xca, 0xcd, 0x4c, 0x0d, 0xbd, 0x95, 0x4e, 0x0d, 0xed, 0xb9, 0x72, 0x59, 0x7b, 0x0c,
	0xd5, 0x9b, 0xe8, 0x25, 0x9d, 0x7d, 0xef, 0xc6, 0x67, 0x3f, 0xff, 0xe3, 0xaf, 0xef, 0x55, 0x4d,
	0xdb, 0x8e, 0x07, 0xd9, 0xb7, 0xf2, 0xde, 0x7f, 0x5c, 0x35, 0xaa, 0xdf, 0xed, 0xe3, 0x0c, 0x77,
	0x5f, 0xa9, 0xa2, 0x5f, 0x14, 0xd8, 0x98, 0x7a, 0x9d, 0xd0, 0xfd, 0x59, 0x7d, 0xf8, 0x92, 0xc7,
	0xa9, 0xf4, 0xfe, 0xab, 0x05, 0xc9, 0xef, 0xd2, 0x74, 0x61, 0xb7, 0x82, 0xde, 0xf9, 0x8f, 0xdd,
	0xde, 0xb4, 0xb8, 0xd2, 0xf1, 0x65, 0x6d, 0x73, 0xe6, 0x6b, 0xf7, 0x7b, 0x4d, 0x3f, 0xe3, 0x3c,
	0x08, 0xf7, 0x0d, 0x63, 0x34, 0x1a, 0x4d, 0x3e, 0x85, 0x76, 0xc4, 0xcf, 0xe4, 0xa4, 0x7f, 0x2f,
	0xf0, 0x6c, 0x3e, 0xa0, 0x6c, 0x58, 0xff, 0x51, 0x81, 0x4d, 0x87, 0x0e, 0xa7, 0x5b, 0xaf, 0xdf,
	0xce, 0x36, 0x44, 0x3b, 0x1e, 0x02, 0x6d, 0xe5, 0xcb, 0xfd, 0x84, 0xeb, 0x52, 0xcf, 0xf6, 0x5d,
	0x9d, 0x32, 0xd7, 0x70, 0xb1, 0x2f, 0x46, 0x84, 0x71, 0x7d, 0xe2, 0xc4, 0x5f, 0x8a, 0x0f, 0xc5,
	0x8f, 0x9f, 0xd5, 0x8d, 0x03, 0x19, 0xdc, 0x10, 0x07, 0x09, 0x79, 0xfd, 0xd1, 0xee, 0x65, 0x8a,
	0x9f, 0x0a, 0xfc, 0x54, 0xe0, 0xa7, 0x8f, 0x76, 0x7b, 0x0b, 0x42, 0xf6, 0xfe, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x30, 0xdc, 0x92, 0xd5, 0x5d, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetServiceClient interface {
	// Exports assets with time and resource types to a given Cloud Storage
	// location. The output format is newline-delimited JSON.
	// This API implements the [google.longrunning.Operation][google.longrunning.Operation] API allowing you
	// to keep track of the export.
	ExportAssets(ctx context.Context, in *ExportAssetsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Batch gets the update history of assets that overlap a time window.
	// For RESOURCE content, this API outputs history with asset in both
	// non-delete or deleted status.
	// For IAM_POLICY content, this API outputs history when the asset and its
	// attached IAM POLICY both exist. This can create gaps in the output history.
	// If a specified asset does not exist, this API returns an INVALID_ARGUMENT
	// error.
	BatchGetAssetsHistory(ctx context.Context, in *BatchGetAssetsHistoryRequest, opts ...grpc.CallOption) (*BatchGetAssetsHistoryResponse, error)
}

type assetServiceClient struct {
	cc *grpc.ClientConn
}

func NewAssetServiceClient(cc *grpc.ClientConn) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) ExportAssets(ctx context.Context, in *ExportAssetsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1.AssetService/ExportAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) BatchGetAssetsHistory(ctx context.Context, in *BatchGetAssetsHistoryRequest, opts ...grpc.CallOption) (*BatchGetAssetsHistoryResponse, error) {
	out := new(BatchGetAssetsHistoryResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1.AssetService/BatchGetAssetsHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
type AssetServiceServer interface {
	// Exports assets with time and resource types to a given Cloud Storage
	// location. The output format is newline-delimited JSON.
	// This API implements the [google.longrunning.Operation][google.longrunning.Operation] API allowing you
	// to keep track of the export.
	ExportAssets(context.Context, *ExportAssetsRequest) (*longrunning.Operation, error)
	// Batch gets the update history of assets that overlap a time window.
	// For RESOURCE content, this API outputs history with asset in both
	// non-delete or deleted status.
	// For IAM_POLICY content, this API outputs history when the asset and its
	// attached IAM POLICY both exist. This can create gaps in the output history.
	// If a specified asset does not exist, this API returns an INVALID_ARGUMENT
	// error.
	BatchGetAssetsHistory(context.Context, *BatchGetAssetsHistoryRequest) (*BatchGetAssetsHistoryResponse, error)
}

// UnimplementedAssetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (*UnimplementedAssetServiceServer) ExportAssets(ctx context.Context, req *ExportAssetsRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportAssets not implemented")
}
func (*UnimplementedAssetServiceServer) BatchGetAssetsHistory(ctx context.Context, req *BatchGetAssetsHistoryRequest) (*BatchGetAssetsHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetAssetsHistory not implemented")
}

func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	s.RegisterService(&_AssetService_serviceDesc, srv)
}

func _AssetService_ExportAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).ExportAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1.AssetService/ExportAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).ExportAssets(ctx, req.(*ExportAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_BatchGetAssetsHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetAssetsHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).BatchGetAssetsHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1.AssetService/BatchGetAssetsHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).BatchGetAssetsHistory(ctx, req.(*BatchGetAssetsHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.asset.v1.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportAssets",
			Handler:    _AssetService_ExportAssets_Handler,
		},
		{
			MethodName: "BatchGetAssetsHistory",
			Handler:    _AssetService_BatchGetAssetsHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/asset/v1/asset_service.proto",
}