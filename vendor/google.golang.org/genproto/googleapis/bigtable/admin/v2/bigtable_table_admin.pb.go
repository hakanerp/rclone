// Code generated by protoc-gen-go.
// source: google/bigtable/admin/v2/bigtable_table_admin.proto
// DO NOT EDIT!

package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import _ "google.golang.org/genproto/googleapis/longrunning"
import _ "github.com/golang/protobuf/ptypes/duration"
import google_protobuf3 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.CreateTable][google.bigtable.admin.v2.BigtableTableAdmin.CreateTable]
type CreateTableRequest struct {
	// The unique name of the instance in which to create the table.
	// Values are of the form `projects/<project>/instances/<instance>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The name by which the new table should be referred to within the parent
	// instance, e.g., `foobar` rather than `<parent>/tables/foobar`.
	TableId string `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	// The Table to create.
	Table *Table `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	// The optional list of row keys that will be used to initially split the
	// table into several tablets (tablets are similar to HBase regions).
	// Given two split keys, `s1` and `s2`, three tablets will be created,
	// spanning the key ranges: `[, s1), [s1, s2), [s2, )`.
	//
	// Example:
	//
	// * Row keys := `["a", "apple", "custom", "customer_1", "customer_2",`
	//                `"other", "zz"]`
	// * initial_split_keys := `["apple", "customer_1", "customer_2", "other"]`
	// * Key assignment:
	//     - Tablet 1 `[, apple)                => {"a"}.`
	//     - Tablet 2 `[apple, customer_1)      => {"apple", "custom"}.`
	//     - Tablet 3 `[customer_1, customer_2) => {"customer_1"}.`
	//     - Tablet 4 `[customer_2, other)      => {"customer_2"}.`
	//     - Tablet 5 `[other, )                => {"other", "zz"}.`
	InitialSplits []*CreateTableRequest_Split `protobuf:"bytes,4,rep,name=initial_splits,json=initialSplits" json:"initial_splits,omitempty"`
}

func (m *CreateTableRequest) Reset()                    { *m = CreateTableRequest{} }
func (m *CreateTableRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTableRequest) ProtoMessage()               {}
func (*CreateTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateTableRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateTableRequest) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

func (m *CreateTableRequest) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *CreateTableRequest) GetInitialSplits() []*CreateTableRequest_Split {
	if m != nil {
		return m.InitialSplits
	}
	return nil
}

// An initial split point for a newly created table.
type CreateTableRequest_Split struct {
	// Row key to use as an initial tablet boundary.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *CreateTableRequest_Split) Reset()                    { *m = CreateTableRequest_Split{} }
func (m *CreateTableRequest_Split) String() string            { return proto.CompactTextString(m) }
func (*CreateTableRequest_Split) ProtoMessage()               {}
func (*CreateTableRequest_Split) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *CreateTableRequest_Split) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.DropRowRange][google.bigtable.admin.v2.BigtableTableAdmin.DropRowRange]
type DropRowRangeRequest struct {
	// The unique name of the table on which to drop a range of rows.
	// Values are of the form
	// `projects/<project>/instances/<instance>/tables/<table>`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Delete all rows or by prefix.
	//
	// Types that are valid to be assigned to Target:
	//	*DropRowRangeRequest_RowKeyPrefix
	//	*DropRowRangeRequest_DeleteAllDataFromTable
	Target isDropRowRangeRequest_Target `protobuf_oneof:"target"`
}

func (m *DropRowRangeRequest) Reset()                    { *m = DropRowRangeRequest{} }
func (m *DropRowRangeRequest) String() string            { return proto.CompactTextString(m) }
func (*DropRowRangeRequest) ProtoMessage()               {}
func (*DropRowRangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type isDropRowRangeRequest_Target interface {
	isDropRowRangeRequest_Target()
}

type DropRowRangeRequest_RowKeyPrefix struct {
	RowKeyPrefix []byte `protobuf:"bytes,2,opt,name=row_key_prefix,json=rowKeyPrefix,proto3,oneof"`
}
type DropRowRangeRequest_DeleteAllDataFromTable struct {
	DeleteAllDataFromTable bool `protobuf:"varint,3,opt,name=delete_all_data_from_table,json=deleteAllDataFromTable,oneof"`
}

func (*DropRowRangeRequest_RowKeyPrefix) isDropRowRangeRequest_Target()           {}
func (*DropRowRangeRequest_DeleteAllDataFromTable) isDropRowRangeRequest_Target() {}

func (m *DropRowRangeRequest) GetTarget() isDropRowRangeRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *DropRowRangeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DropRowRangeRequest) GetRowKeyPrefix() []byte {
	if x, ok := m.GetTarget().(*DropRowRangeRequest_RowKeyPrefix); ok {
		return x.RowKeyPrefix
	}
	return nil
}

func (m *DropRowRangeRequest) GetDeleteAllDataFromTable() bool {
	if x, ok := m.GetTarget().(*DropRowRangeRequest_DeleteAllDataFromTable); ok {
		return x.DeleteAllDataFromTable
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DropRowRangeRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DropRowRangeRequest_OneofMarshaler, _DropRowRangeRequest_OneofUnmarshaler, _DropRowRangeRequest_OneofSizer, []interface{}{
		(*DropRowRangeRequest_RowKeyPrefix)(nil),
		(*DropRowRangeRequest_DeleteAllDataFromTable)(nil),
	}
}

func _DropRowRangeRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DropRowRangeRequest)
	// target
	switch x := m.Target.(type) {
	case *DropRowRangeRequest_RowKeyPrefix:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RowKeyPrefix)
	case *DropRowRangeRequest_DeleteAllDataFromTable:
		t := uint64(0)
		if x.DeleteAllDataFromTable {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("DropRowRangeRequest.Target has unexpected type %T", x)
	}
	return nil
}

func _DropRowRangeRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DropRowRangeRequest)
	switch tag {
	case 2: // target.row_key_prefix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Target = &DropRowRangeRequest_RowKeyPrefix{x}
		return true, err
	case 3: // target.delete_all_data_from_table
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Target = &DropRowRangeRequest_DeleteAllDataFromTable{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _DropRowRangeRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DropRowRangeRequest)
	// target
	switch x := m.Target.(type) {
	case *DropRowRangeRequest_RowKeyPrefix:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RowKeyPrefix)))
		n += len(x.RowKeyPrefix)
	case *DropRowRangeRequest_DeleteAllDataFromTable:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.ListTables][google.bigtable.admin.v2.BigtableTableAdmin.ListTables]
type ListTablesRequest struct {
	// The unique name of the instance for which tables should be listed.
	// Values are of the form `projects/<project>/instances/<instance>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The view to be applied to the returned tables' fields.
	// Defaults to `NAME_ONLY` if unspecified; no others are currently supported.
	View Table_View `protobuf:"varint,2,opt,name=view,enum=google.bigtable.admin.v2.Table_View" json:"view,omitempty"`
	// The value of `next_page_token` returned by a previous call.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListTablesRequest) Reset()                    { *m = ListTablesRequest{} }
func (m *ListTablesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTablesRequest) ProtoMessage()               {}
func (*ListTablesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ListTablesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListTablesRequest) GetView() Table_View {
	if m != nil {
		return m.View
	}
	return Table_VIEW_UNSPECIFIED
}

func (m *ListTablesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for
// [google.bigtable.admin.v2.BigtableTableAdmin.ListTables][google.bigtable.admin.v2.BigtableTableAdmin.ListTables]
type ListTablesResponse struct {
	// The tables present in the requested instance.
	Tables []*Table `protobuf:"bytes,1,rep,name=tables" json:"tables,omitempty"`
	// Set if not all tables could be returned in a single response.
	// Pass this value to `page_token` in another request to get the next
	// page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListTablesResponse) Reset()                    { *m = ListTablesResponse{} }
func (m *ListTablesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTablesResponse) ProtoMessage()               {}
func (*ListTablesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ListTablesResponse) GetTables() []*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

func (m *ListTablesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.GetTable][google.bigtable.admin.v2.BigtableTableAdmin.GetTable]
type GetTableRequest struct {
	// The unique name of the requested table.
	// Values are of the form
	// `projects/<project>/instances/<instance>/tables/<table>`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The view to be applied to the returned table's fields.
	// Defaults to `SCHEMA_ONLY` if unspecified.
	View Table_View `protobuf:"varint,2,opt,name=view,enum=google.bigtable.admin.v2.Table_View" json:"view,omitempty"`
}

func (m *GetTableRequest) Reset()                    { *m = GetTableRequest{} }
func (m *GetTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTableRequest) ProtoMessage()               {}
func (*GetTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *GetTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetTableRequest) GetView() Table_View {
	if m != nil {
		return m.View
	}
	return Table_VIEW_UNSPECIFIED
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.DeleteTable][google.bigtable.admin.v2.BigtableTableAdmin.DeleteTable]
type DeleteTableRequest struct {
	// The unique name of the table to be deleted.
	// Values are of the form
	// `projects/<project>/instances/<instance>/tables/<table>`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteTableRequest) Reset()                    { *m = DeleteTableRequest{} }
func (m *DeleteTableRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTableRequest) ProtoMessage()               {}
func (*DeleteTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *DeleteTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.ModifyColumnFamilies][google.bigtable.admin.v2.BigtableTableAdmin.ModifyColumnFamilies]
type ModifyColumnFamiliesRequest struct {
	// The unique name of the table whose families should be modified.
	// Values are of the form
	// `projects/<project>/instances/<instance>/tables/<table>`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Modifications to be atomically applied to the specified table's families.
	// Entries are applied in order, meaning that earlier modifications can be
	// masked by later ones (in the case of repeated updates to the same family,
	// for example).
	Modifications []*ModifyColumnFamiliesRequest_Modification `protobuf:"bytes,2,rep,name=modifications" json:"modifications,omitempty"`
}

func (m *ModifyColumnFamiliesRequest) Reset()                    { *m = ModifyColumnFamiliesRequest{} }
func (m *ModifyColumnFamiliesRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyColumnFamiliesRequest) ProtoMessage()               {}
func (*ModifyColumnFamiliesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ModifyColumnFamiliesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModifyColumnFamiliesRequest) GetModifications() []*ModifyColumnFamiliesRequest_Modification {
	if m != nil {
		return m.Modifications
	}
	return nil
}

// A create, update, or delete of a particular column family.
type ModifyColumnFamiliesRequest_Modification struct {
	// The ID of the column family to be modified.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Column familiy modifications.
	//
	// Types that are valid to be assigned to Mod:
	//	*ModifyColumnFamiliesRequest_Modification_Create
	//	*ModifyColumnFamiliesRequest_Modification_Update
	//	*ModifyColumnFamiliesRequest_Modification_Drop
	Mod isModifyColumnFamiliesRequest_Modification_Mod `protobuf_oneof:"mod"`
}

func (m *ModifyColumnFamiliesRequest_Modification) Reset() {
	*m = ModifyColumnFamiliesRequest_Modification{}
}
func (m *ModifyColumnFamiliesRequest_Modification) String() string { return proto.CompactTextString(m) }
func (*ModifyColumnFamiliesRequest_Modification) ProtoMessage()    {}
func (*ModifyColumnFamiliesRequest_Modification) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{6, 0}
}

type isModifyColumnFamiliesRequest_Modification_Mod interface {
	isModifyColumnFamiliesRequest_Modification_Mod()
}

type ModifyColumnFamiliesRequest_Modification_Create struct {
	Create *ColumnFamily `protobuf:"bytes,2,opt,name=create,oneof"`
}
type ModifyColumnFamiliesRequest_Modification_Update struct {
	Update *ColumnFamily `protobuf:"bytes,3,opt,name=update,oneof"`
}
type ModifyColumnFamiliesRequest_Modification_Drop struct {
	Drop bool `protobuf:"varint,4,opt,name=drop,oneof"`
}

func (*ModifyColumnFamiliesRequest_Modification_Create) isModifyColumnFamiliesRequest_Modification_Mod() {
}
func (*ModifyColumnFamiliesRequest_Modification_Update) isModifyColumnFamiliesRequest_Modification_Mod() {
}
func (*ModifyColumnFamiliesRequest_Modification_Drop) isModifyColumnFamiliesRequest_Modification_Mod() {
}

func (m *ModifyColumnFamiliesRequest_Modification) GetMod() isModifyColumnFamiliesRequest_Modification_Mod {
	if m != nil {
		return m.Mod
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ModifyColumnFamiliesRequest_Modification) GetCreate() *ColumnFamily {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Create); ok {
		return x.Create
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetUpdate() *ColumnFamily {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Update); ok {
		return x.Update
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetDrop() bool {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Drop); ok {
		return x.Drop
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ModifyColumnFamiliesRequest_Modification) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ModifyColumnFamiliesRequest_Modification_OneofMarshaler, _ModifyColumnFamiliesRequest_Modification_OneofUnmarshaler, _ModifyColumnFamiliesRequest_Modification_OneofSizer, []interface{}{
		(*ModifyColumnFamiliesRequest_Modification_Create)(nil),
		(*ModifyColumnFamiliesRequest_Modification_Update)(nil),
		(*ModifyColumnFamiliesRequest_Modification_Drop)(nil),
	}
}

func _ModifyColumnFamiliesRequest_Modification_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	// mod
	switch x := m.Mod.(type) {
	case *ModifyColumnFamiliesRequest_Modification_Create:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *ModifyColumnFamiliesRequest_Modification_Update:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *ModifyColumnFamiliesRequest_Modification_Drop:
		t := uint64(0)
		if x.Drop {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ModifyColumnFamiliesRequest_Modification.Mod has unexpected type %T", x)
	}
	return nil
}

func _ModifyColumnFamiliesRequest_Modification_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	switch tag {
	case 2: // mod.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ColumnFamily)
		err := b.DecodeMessage(msg)
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Create{msg}
		return true, err
	case 3: // mod.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ColumnFamily)
		err := b.DecodeMessage(msg)
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Update{msg}
		return true, err
	case 4: // mod.drop
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Drop{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ModifyColumnFamiliesRequest_Modification_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	// mod
	switch x := m.Mod.(type) {
	case *ModifyColumnFamiliesRequest_Modification_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModifyColumnFamiliesRequest_Modification_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModifyColumnFamiliesRequest_Modification_Drop:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CreateTableRequest)(nil), "google.bigtable.admin.v2.CreateTableRequest")
	proto.RegisterType((*CreateTableRequest_Split)(nil), "google.bigtable.admin.v2.CreateTableRequest.Split")
	proto.RegisterType((*DropRowRangeRequest)(nil), "google.bigtable.admin.v2.DropRowRangeRequest")
	proto.RegisterType((*ListTablesRequest)(nil), "google.bigtable.admin.v2.ListTablesRequest")
	proto.RegisterType((*ListTablesResponse)(nil), "google.bigtable.admin.v2.ListTablesResponse")
	proto.RegisterType((*GetTableRequest)(nil), "google.bigtable.admin.v2.GetTableRequest")
	proto.RegisterType((*DeleteTableRequest)(nil), "google.bigtable.admin.v2.DeleteTableRequest")
	proto.RegisterType((*ModifyColumnFamiliesRequest)(nil), "google.bigtable.admin.v2.ModifyColumnFamiliesRequest")
	proto.RegisterType((*ModifyColumnFamiliesRequest_Modification)(nil), "google.bigtable.admin.v2.ModifyColumnFamiliesRequest.Modification")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BigtableTableAdmin service

type BigtableTableAdminClient interface {
	// Creates a new table in the specified instance.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*Table, error)
	// Lists all tables served from a specified instance.
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
	// Gets metadata information about the specified table.
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error)
	// Atomically performs a series of column family modifications
	// on the specified table.
	ModifyColumnFamilies(ctx context.Context, in *ModifyColumnFamiliesRequest, opts ...grpc.CallOption) (*Table, error)
	// Permanently drop/delete a row range from a specified table. The request can
	// specify whether to delete all rows in a table, or only those that match a
	// particular prefix.
	DropRowRange(ctx context.Context, in *DropRowRangeRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error)
}

type bigtableTableAdminClient struct {
	cc *grpc.ClientConn
}

func NewBigtableTableAdminClient(cc *grpc.ClientConn) BigtableTableAdminClient {
	return &bigtableTableAdminClient{cc}
}

func (c *bigtableTableAdminClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*Table, error) {
	out := new(Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/CreateTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/ListTables", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*Table, error) {
	out := new(Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/GetTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error) {
	out := new(google_protobuf3.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/DeleteTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) ModifyColumnFamilies(ctx context.Context, in *ModifyColumnFamiliesRequest, opts ...grpc.CallOption) (*Table, error) {
	out := new(Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/ModifyColumnFamilies", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableAdminClient) DropRowRange(ctx context.Context, in *DropRowRangeRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error) {
	out := new(google_protobuf3.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableTableAdmin/DropRowRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableTableAdmin service

type BigtableTableAdminServer interface {
	// Creates a new table in the specified instance.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(context.Context, *CreateTableRequest) (*Table, error)
	// Lists all tables served from a specified instance.
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
	// Gets metadata information about the specified table.
	GetTable(context.Context, *GetTableRequest) (*Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(context.Context, *DeleteTableRequest) (*google_protobuf3.Empty, error)
	// Atomically performs a series of column family modifications
	// on the specified table.
	ModifyColumnFamilies(context.Context, *ModifyColumnFamiliesRequest) (*Table, error)
	// Permanently drop/delete a row range from a specified table. The request can
	// specify whether to delete all rows in a table, or only those that match a
	// particular prefix.
	DropRowRange(context.Context, *DropRowRangeRequest) (*google_protobuf3.Empty, error)
}

func RegisterBigtableTableAdminServer(s *grpc.Server, srv BigtableTableAdminServer) {
	s.RegisterService(&_BigtableTableAdmin_serviceDesc, srv)
}

func _BigtableTableAdmin_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).ListTables(ctx, req.(*ListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).GetTable(ctx, req.(*GetTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_DeleteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).DeleteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/DeleteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).DeleteTable(ctx, req.(*DeleteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_ModifyColumnFamilies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyColumnFamiliesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).ModifyColumnFamilies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/ModifyColumnFamilies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).ModifyColumnFamilies(ctx, req.(*ModifyColumnFamiliesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableAdmin_DropRowRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropRowRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableAdminServer).DropRowRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableTableAdmin/DropRowRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableAdminServer).DropRowRange(ctx, req.(*DropRowRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigtableTableAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.v2.BigtableTableAdmin",
	HandlerType: (*BigtableTableAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _BigtableTableAdmin_CreateTable_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _BigtableTableAdmin_ListTables_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _BigtableTableAdmin_GetTable_Handler,
		},
		{
			MethodName: "DeleteTable",
			Handler:    _BigtableTableAdmin_DeleteTable_Handler,
		},
		{
			MethodName: "ModifyColumnFamilies",
			Handler:    _BigtableTableAdmin_ModifyColumnFamilies_Handler,
		},
		{
			MethodName: "DropRowRange",
			Handler:    _BigtableTableAdmin_DropRowRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/bigtable/admin/v2/bigtable_table_admin.proto",
}

func init() {
	proto.RegisterFile("google/bigtable/admin/v2/bigtable_table_admin.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xce, 0xd8, 0x4e, 0x9a, 0x1c, 0x3b, 0x29, 0x0c, 0x25, 0xb8, 0x6e, 0xa1, 0xd1, 0x52, 0x45,
	0xc1, 0x84, 0x5d, 0x69, 0xab, 0xa8, 0x28, 0xb4, 0x82, 0xba, 0xa1, 0x84, 0x3f, 0x29, 0x5a, 0x2a,
	0x24, 0xb8, 0x59, 0x4d, 0xbc, 0x93, 0xed, 0x90, 0xdd, 0x99, 0x65, 0x77, 0x1c, 0xc7, 0x42, 0xbd,
	0x41, 0x48, 0x48, 0xdc, 0xf6, 0xaa, 0xe2, 0x45, 0x10, 0x37, 0x3c, 0x04, 0xd7, 0xdc, 0xf1, 0x08,
	0x3c, 0x00, 0x9a, 0x1f, 0x27, 0x9b, 0x38, 0x1b, 0xd7, 0xb9, 0xb1, 0x66, 0xce, 0xf9, 0xce, 0x39,
	0xdf, 0x9c, 0x33, 0xdf, 0x8e, 0xe1, 0x5e, 0x2c, 0x44, 0x9c, 0x50, 0x6f, 0x9f, 0xc5, 0x92, 0xec,
	0x27, 0xd4, 0x23, 0x51, 0xca, 0xb8, 0x77, 0xe4, 0x9f, 0x58, 0x42, 0xf3, 0xab, 0xed, 0x6e, 0x96,
	0x0b, 0x29, 0x70, 0xdb, 0x04, 0xb9, 0x63, 0x88, 0x6b, 0x9c, 0x47, 0x7e, 0xe7, 0xb6, 0x4d, 0x47,
	0x32, 0xe6, 0x11, 0xce, 0x85, 0x24, 0x92, 0x09, 0x5e, 0x98, 0xb8, 0xce, 0x9b, 0x65, 0xef, 0x40,
	0x3e, 0xb3, 0xe6, 0xbb, 0x95, 0x1c, 0x4c, 0x76, 0x83, 0x7a, 0xd7, 0xa2, 0x12, 0xc1, 0xe3, 0x7c,
	0xc0, 0x39, 0xe3, 0xb1, 0x27, 0x32, 0x9a, 0x9f, 0xa9, 0xf0, 0x8e, 0x05, 0xe9, 0xdd, 0xfe, 0xe0,
	0xc0, 0x8b, 0x06, 0x06, 0x60, 0xfd, 0xb7, 0xce, 0xfb, 0x69, 0x9a, 0xc9, 0x91, 0x75, 0xde, 0x39,
	0xef, 0x94, 0x2c, 0xa5, 0x85, 0x24, 0x69, 0x66, 0x00, 0xce, 0x7f, 0x08, 0xf0, 0xe3, 0x9c, 0x12,
	0x49, 0x9f, 0x2a, 0x62, 0x01, 0xfd, 0x71, 0x40, 0x0b, 0x89, 0x57, 0x61, 0x21, 0x23, 0x39, 0xe5,
	0xb2, 0x8d, 0xd6, 0xd0, 0xc6, 0x52, 0x60, 0x77, 0xf8, 0x26, 0x2c, 0x9a, 0xde, 0xb1, 0xa8, 0x5d,
	0xd3, 0x9e, 0x6b, 0x7a, 0xff, 0x79, 0x84, 0xb7, 0x60, 0x5e, 0x2f, 0xdb, 0xf5, 0x35, 0xb4, 0xd1,
	0xf4, 0xef, 0xb8, 0x55, 0x1d, 0x75, 0x4d, 0x25, 0x83, 0xc6, 0xdf, 0xc1, 0x0a, 0xe3, 0x4c, 0x32,
	0x92, 0x84, 0x45, 0x96, 0x30, 0x59, 0xb4, 0x1b, 0x6b, 0xf5, 0x8d, 0xa6, 0xef, 0x57, 0xc7, 0x4f,
	0xf2, 0x75, 0xbf, 0x51, 0xa1, 0xc1, 0xb2, 0xcd, 0xa4, 0x77, 0x45, 0xe7, 0x26, 0xcc, 0xeb, 0x15,
	0x7e, 0x0d, 0xea, 0x87, 0x74, 0xa4, 0x8f, 0xd2, 0x0a, 0xd4, 0xd2, 0x79, 0x89, 0xe0, 0x8d, 0x9d,
	0x5c, 0x64, 0x81, 0x18, 0x06, 0x84, 0xc7, 0x27, 0xe7, 0xc6, 0xd0, 0xe0, 0x24, 0xa5, 0xf6, 0xd4,
	0x7a, 0x8d, 0xd7, 0x61, 0x25, 0x17, 0xc3, 0xf0, 0x90, 0x8e, 0xc2, 0x2c, 0xa7, 0x07, 0xec, 0x58,
	0x9f, 0xbc, 0xb5, 0x3b, 0x17, 0xb4, 0x72, 0x31, 0xfc, 0x92, 0x8e, 0xf6, 0xb4, 0x15, 0x3f, 0x80,
	0x4e, 0x44, 0x13, 0x2a, 0x69, 0x48, 0x92, 0x24, 0x8c, 0x88, 0x24, 0xe1, 0x41, 0x2e, 0xd2, 0xf0,
	0xb4, 0x2b, 0x8b, 0xbb, 0x73, 0xc1, 0xaa, 0xc1, 0x3c, 0x4a, 0x92, 0x1d, 0x22, 0xc9, 0x93, 0x5c,
	0xa4, 0xfa, 0x20, 0xbd, 0x45, 0x58, 0x90, 0x24, 0x8f, 0xa9, 0x74, 0x7e, 0x41, 0xf0, 0xfa, 0x57,
	0xac, 0x90, 0xda, 0x5e, 0x4c, 0x9b, 0xc8, 0x87, 0xd0, 0x38, 0x62, 0x74, 0xa8, 0x39, 0xad, 0xf8,
	0x77, 0xa7, 0x74, 0xdd, 0xfd, 0x96, 0xd1, 0x61, 0xa0, 0x23, 0xf0, 0xdb, 0x00, 0x19, 0x89, 0x69,
	0x28, 0xc5, 0x21, 0xe5, 0x9a, 0xdf, 0x52, 0xb0, 0xa4, 0x2c, 0x4f, 0x95, 0xc1, 0x19, 0x00, 0x2e,
	0xb3, 0x28, 0x32, 0xc1, 0x0b, 0x8a, 0xef, 0x2b, 0x9a, 0xca, 0xd2, 0x46, 0x7a, 0x4c, 0x53, 0xc7,
	0x6c, 0xe1, 0x78, 0x1d, 0xae, 0x73, 0x7a, 0x2c, 0xc3, 0x52, 0x49, 0x73, 0x81, 0x96, 0x95, 0x79,
	0xef, 0xa4, 0x6c, 0x08, 0xd7, 0x3f, 0xa3, 0xf2, 0xcc, 0x65, 0xbc, 0x68, 0x28, 0x57, 0x3e, 0xb6,
	0xb3, 0x01, 0x78, 0x47, 0x8f, 0x60, 0x5a, 0x0d, 0xe7, 0x9f, 0x1a, 0xdc, 0xfa, 0x5a, 0x44, 0xec,
	0x60, 0xf4, 0x58, 0x24, 0x83, 0x94, 0x3f, 0x21, 0x29, 0x4b, 0xd8, 0xe9, 0x48, 0x2e, 0xe2, 0xf5,
	0x0c, 0x96, 0x53, 0x15, 0xc2, 0xfa, 0x46, 0xc4, 0xed, 0x9a, 0x6e, 0x53, 0xaf, 0x9a, 0xe0, 0x25,
	0x15, 0x8c, 0xcf, 0xa6, 0x0a, 0xce, 0x26, 0xee, 0xfc, 0x85, 0xa0, 0x55, 0xf6, 0xe3, 0x15, 0xa8,
	0xb1, 0xc8, 0x92, 0xa9, 0xb1, 0x08, 0x7f, 0x02, 0x0b, 0x7d, 0xad, 0x14, 0xdd, 0xa4, 0xa6, 0xbf,
	0x7e, 0x89, 0xa2, 0x4e, 0xab, 0x8f, 0x76, 0xe7, 0x02, 0x1b, 0xa7, 0x32, 0x0c, 0xb2, 0x48, 0x65,
	0xa8, 0xcf, 0x9a, 0xc1, 0xc4, 0xe1, 0x1b, 0xd0, 0x88, 0x72, 0x91, 0xb5, 0x1b, 0xf6, 0xf6, 0xeb,
	0x5d, 0x6f, 0x1e, 0xea, 0xa9, 0x88, 0xfc, 0x3f, 0xae, 0x01, 0xee, 0xd9, 0x4c, 0x7a, 0x18, 0x8f,
	0x54, 0x36, 0xfc, 0x02, 0x41, 0xb3, 0x24, 0x71, 0xbc, 0x39, 0xcb, 0x97, 0xa0, 0x33, 0xed, 0x42,
	0x3a, 0x5b, 0x3f, 0xff, 0xfd, 0xef, 0x8b, 0x9a, 0xe7, 0x74, 0xd5, 0xd7, 0xf8, 0x27, 0xa3, 0xa2,
	0x87, 0x59, 0x2e, 0x7e, 0xa0, 0x7d, 0x59, 0x78, 0x5d, 0x8f, 0xf1, 0x42, 0x12, 0xde, 0xa7, 0x85,
	0xd7, 0x7d, 0x6e, 0xbe, 0xd6, 0xc5, 0x36, 0xea, 0xe2, 0xdf, 0x11, 0xc0, 0xa9, 0x1e, 0xf0, 0xfb,
	0xd5, 0x65, 0x26, 0xb4, 0xdb, 0xd9, 0x7c, 0x35, 0xb0, 0x91, 0x98, 0xe3, 0x6b, 0x82, 0x9b, 0x78,
	0x06, 0x82, 0xf8, 0x37, 0x04, 0x8b, 0x63, 0xd9, 0xe0, 0xf7, 0xaa, 0xcb, 0x9d, 0x93, 0xd6, 0xf4,
	0x6e, 0x9d, 0x25, 0xa3, 0xae, 0x78, 0x05, 0x15, 0xcb, 0xc4, 0xeb, 0x3e, 0xc7, 0xbf, 0x22, 0x68,
	0x96, 0x24, 0x76, 0xd9, 0x00, 0x27, 0x95, 0xd8, 0x59, 0x1d, 0xa3, 0xc7, 0x6f, 0x96, 0xfb, 0xa9,
	0x7a, 0xd0, 0xc6, 0x4c, 0xba, 0xb3, 0x30, 0xf9, 0x13, 0xc1, 0x8d, 0x8b, 0xf4, 0x85, 0xb7, 0xae,
	0xa4, 0xc7, 0xe9, 0xed, 0xfa, 0x42, 0x93, 0xdc, 0x71, 0x3e, 0x7e, 0x75, 0x92, 0xdb, 0xe9, 0x05,
	0x05, 0xd5, 0x8d, 0x7b, 0x89, 0xa0, 0x55, 0x7e, 0xa3, 0xf0, 0x07, 0x97, 0xf4, 0x71, 0xf2, 0x2d,
	0xab, 0x6c, 0x64, 0x4f, 0x73, 0x7c, 0xe0, 0xdc, 0x9f, 0x81, 0x63, 0x54, 0xca, 0xbf, 0x8d, 0xba,
	0xbd, 0x63, 0xb8, 0xdd, 0x17, 0x69, 0x25, 0x9f, 0xde, 0x5b, 0x93, 0xba, 0xde, 0x53, 0x2c, 0xf6,
	0xd0, 0xf7, 0x0f, 0x6d, 0x50, 0x2c, 0x12, 0xc2, 0x63, 0x57, 0xe4, 0xb1, 0x17, 0x53, 0xae, 0x39,
	0x7a, 0xc6, 0x45, 0x32, 0x56, 0x4c, 0xfe, 0x73, 0xfa, 0x48, 0x2f, 0xf6, 0x17, 0x34, 0xf2, 0xde,
	0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x5d, 0x18, 0x44, 0xe6, 0x09, 0x00, 0x00,
}
