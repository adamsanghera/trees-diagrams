// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models.proto

/*
Package treespb is a generated protocol buffer package.

It is generated from these files:
	models.proto

It has these top-level messages:
	Tree
	CondensedTree
	Location
	Filter
	SearchRequest
	SearchResponse
*/
package treespb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FilterKey int32

const (
	FilterKey_tree_id        FilterKey = 0
	FilterKey_created_at     FilterKey = 1
	FilterKey_tree_diameter  FilterKey = 2
	FilterKey_stump_diameter FilterKey = 3
	FilterKey_status         FilterKey = 4
	FilterKey_health         FilterKey = 5
	FilterKey_spc_latin      FilterKey = 6
	FilterKey_spc_common     FilterKey = 7
	FilterKey_steward        FilterKey = 8
	FilterKey_curb_location  FilterKey = 9
	FilterKey_guards         FilterKey = 10
	FilterKey_sidewalk       FilterKey = 11
	FilterKey_user_type      FilterKey = 12
	FilterKey_problems       FilterKey = 13
	FilterKey_root_stone     FilterKey = 14
	FilterKey_root_grate     FilterKey = 15
	FilterKey_root_other     FilterKey = 16
	FilterKey_trunk_other    FilterKey = 17
	FilterKey_trunk_wire     FilterKey = 18
	FilterKey_trunk_light    FilterKey = 29
	FilterKey_branch_light   FilterKey = 20
	FilterKey_branch_shoe    FilterKey = 21
	FilterKey_branch_other   FilterKey = 22
	FilterKey_address        FilterKey = 23
	FilterKey_zipcode        FilterKey = 24
	FilterKey_zip_city       FilterKey = 25
	FilterKey_borough_name   FilterKey = 26
)

var FilterKey_name = map[int32]string{
	0:  "tree_id",
	1:  "created_at",
	2:  "tree_diameter",
	3:  "stump_diameter",
	4:  "status",
	5:  "health",
	6:  "spc_latin",
	7:  "spc_common",
	8:  "steward",
	9:  "curb_location",
	10: "guards",
	11: "sidewalk",
	12: "user_type",
	13: "problems",
	14: "root_stone",
	15: "root_grate",
	16: "root_other",
	17: "trunk_other",
	18: "trunk_wire",
	29: "trunk_light",
	20: "branch_light",
	21: "branch_shoe",
	22: "branch_other",
	23: "address",
	24: "zipcode",
	25: "zip_city",
	26: "borough_name",
}
var FilterKey_value = map[string]int32{
	"tree_id":        0,
	"created_at":     1,
	"tree_diameter":  2,
	"stump_diameter": 3,
	"status":         4,
	"health":         5,
	"spc_latin":      6,
	"spc_common":     7,
	"steward":        8,
	"curb_location":  9,
	"guards":         10,
	"sidewalk":       11,
	"user_type":      12,
	"problems":       13,
	"root_stone":     14,
	"root_grate":     15,
	"root_other":     16,
	"trunk_other":    17,
	"trunk_wire":     18,
	"trunk_light":    29,
	"branch_light":   20,
	"branch_shoe":    21,
	"branch_other":   22,
	"address":        23,
	"zipcode":        24,
	"zip_city":       25,
	"borough_name":   26,
}

func (x FilterKey) String() string {
	return proto.EnumName(FilterKey_name, int32(x))
}
func (FilterKey) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Tree struct {
	TreeId        int64     `protobuf:"varint,1,opt,name=tree_id,json=treeId" json:"tree_id,omitempty"`
	CreatedAt     int64     `protobuf:"varint,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	TreeDiameter  int32     `protobuf:"varint,3,opt,name=tree_diameter,json=treeDiameter" json:"tree_diameter,omitempty"`
	StumpDiameter int32     `protobuf:"varint,4,opt,name=stump_diameter,json=stumpDiameter" json:"stump_diameter,omitempty"`
	Status        string    `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	Health        string    `protobuf:"bytes,6,opt,name=health" json:"health,omitempty"`
	SpcLatin      string    `protobuf:"bytes,7,opt,name=spc_latin,json=spcLatin" json:"spc_latin,omitempty"`
	SpcCommon     string    `protobuf:"bytes,8,opt,name=spc_common,json=spcCommon" json:"spc_common,omitempty"`
	Steward       string    `protobuf:"bytes,9,opt,name=steward" json:"steward,omitempty"`
	CurbLocation  string    `protobuf:"bytes,10,opt,name=curb_location,json=curbLocation" json:"curb_location,omitempty"`
	Guards        string    `protobuf:"bytes,11,opt,name=guards" json:"guards,omitempty"`
	Sidewalk      string    `protobuf:"bytes,12,opt,name=sidewalk" json:"sidewalk,omitempty"`
	UserType      string    `protobuf:"bytes,13,opt,name=user_type,json=userType" json:"user_type,omitempty"`
	Problems      string    `protobuf:"bytes,14,opt,name=problems" json:"problems,omitempty"`
	RootStone     string    `protobuf:"bytes,15,opt,name=root_stone,json=rootStone" json:"root_stone,omitempty"`
	RootGrate     string    `protobuf:"bytes,16,opt,name=root_grate,json=rootGrate" json:"root_grate,omitempty"`
	RootOther     string    `protobuf:"bytes,17,opt,name=root_other,json=rootOther" json:"root_other,omitempty"`
	TrunkOther    string    `protobuf:"bytes,18,opt,name=trunk_other,json=trunkOther" json:"trunk_other,omitempty"`
	TrunkWire     string    `protobuf:"bytes,19,opt,name=trunk_wire,json=trunkWire" json:"trunk_wire,omitempty"`
	TrunkLight    string    `protobuf:"bytes,20,opt,name=trunk_light,json=trunkLight" json:"trunk_light,omitempty"`
	BranchLight   string    `protobuf:"bytes,21,opt,name=branch_light,json=branchLight" json:"branch_light,omitempty"`
	BranchShoe    string    `protobuf:"bytes,22,opt,name=branch_shoe,json=branchShoe" json:"branch_shoe,omitempty"`
	BranchOther   string    `protobuf:"bytes,23,opt,name=branch_other,json=branchOther" json:"branch_other,omitempty"`
	Address       string    `protobuf:"bytes,24,opt,name=address" json:"address,omitempty"`
	Zipcode       string    `protobuf:"bytes,25,opt,name=zipcode" json:"zipcode,omitempty"`
	ZipCity       string    `protobuf:"bytes,26,opt,name=zip_city,json=zipCity" json:"zip_city,omitempty"`
	BoroughName   string    `protobuf:"bytes,27,opt,name=borough_name,json=boroughName" json:"borough_name,omitempty"`
	Location      *Location `protobuf:"bytes,28,opt,name=location" json:"location,omitempty"`
}

func (m *Tree) Reset()                    { *m = Tree{} }
func (m *Tree) String() string            { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()               {}
func (*Tree) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Tree) GetTreeId() int64 {
	if m != nil {
		return m.TreeId
	}
	return 0
}

func (m *Tree) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Tree) GetTreeDiameter() int32 {
	if m != nil {
		return m.TreeDiameter
	}
	return 0
}

func (m *Tree) GetStumpDiameter() int32 {
	if m != nil {
		return m.StumpDiameter
	}
	return 0
}

func (m *Tree) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Tree) GetHealth() string {
	if m != nil {
		return m.Health
	}
	return ""
}

func (m *Tree) GetSpcLatin() string {
	if m != nil {
		return m.SpcLatin
	}
	return ""
}

func (m *Tree) GetSpcCommon() string {
	if m != nil {
		return m.SpcCommon
	}
	return ""
}

func (m *Tree) GetSteward() string {
	if m != nil {
		return m.Steward
	}
	return ""
}

func (m *Tree) GetCurbLocation() string {
	if m != nil {
		return m.CurbLocation
	}
	return ""
}

func (m *Tree) GetGuards() string {
	if m != nil {
		return m.Guards
	}
	return ""
}

func (m *Tree) GetSidewalk() string {
	if m != nil {
		return m.Sidewalk
	}
	return ""
}

func (m *Tree) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *Tree) GetProblems() string {
	if m != nil {
		return m.Problems
	}
	return ""
}

func (m *Tree) GetRootStone() string {
	if m != nil {
		return m.RootStone
	}
	return ""
}

func (m *Tree) GetRootGrate() string {
	if m != nil {
		return m.RootGrate
	}
	return ""
}

func (m *Tree) GetRootOther() string {
	if m != nil {
		return m.RootOther
	}
	return ""
}

func (m *Tree) GetTrunkOther() string {
	if m != nil {
		return m.TrunkOther
	}
	return ""
}

func (m *Tree) GetTrunkWire() string {
	if m != nil {
		return m.TrunkWire
	}
	return ""
}

func (m *Tree) GetTrunkLight() string {
	if m != nil {
		return m.TrunkLight
	}
	return ""
}

func (m *Tree) GetBranchLight() string {
	if m != nil {
		return m.BranchLight
	}
	return ""
}

func (m *Tree) GetBranchShoe() string {
	if m != nil {
		return m.BranchShoe
	}
	return ""
}

func (m *Tree) GetBranchOther() string {
	if m != nil {
		return m.BranchOther
	}
	return ""
}

func (m *Tree) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Tree) GetZipcode() string {
	if m != nil {
		return m.Zipcode
	}
	return ""
}

func (m *Tree) GetZipCity() string {
	if m != nil {
		return m.ZipCity
	}
	return ""
}

func (m *Tree) GetBoroughName() string {
	if m != nil {
		return m.BoroughName
	}
	return ""
}

func (m *Tree) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type CondensedTree struct {
	TreeId    int64     `protobuf:"varint,1,opt,name=tree_id,json=treeId" json:"tree_id,omitempty"`
	SpcCommon string    `protobuf:"bytes,2,opt,name=spc_common,json=spcCommon" json:"spc_common,omitempty"`
	SpcLatin  string    `protobuf:"bytes,3,opt,name=spc_latin,json=spcLatin" json:"spc_latin,omitempty"`
	Location  *Location `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
}

func (m *CondensedTree) Reset()                    { *m = CondensedTree{} }
func (m *CondensedTree) String() string            { return proto.CompactTextString(m) }
func (*CondensedTree) ProtoMessage()               {}
func (*CondensedTree) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CondensedTree) GetTreeId() int64 {
	if m != nil {
		return m.TreeId
	}
	return 0
}

func (m *CondensedTree) GetSpcCommon() string {
	if m != nil {
		return m.SpcCommon
	}
	return ""
}

func (m *CondensedTree) GetSpcLatin() string {
	if m != nil {
		return m.SpcLatin
	}
	return ""
}

func (m *CondensedTree) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Location struct {
	Lat float32 `protobuf:"fixed32,1,opt,name=lat" json:"lat,omitempty"`
	Lon float32 `protobuf:"fixed32,2,opt,name=lon" json:"lon,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Location) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Location) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Filter struct {
	Key   FilterKey `protobuf:"varint,1,opt,name=key,enum=FilterKey" json:"key,omitempty"`
	Value string    `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Filter) GetKey() FilterKey {
	if m != nil {
		return m.Key
	}
	return FilterKey_tree_id
}

func (m *Filter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SearchRequest struct {
	Origin  *Location `protobuf:"bytes,1,opt,name=origin" json:"origin,omitempty"`
	Radius  float32   `protobuf:"fixed32,2,opt,name=radius" json:"radius,omitempty"`
	Filters []*Filter `protobuf:"bytes,3,rep,name=filters" json:"filters,omitempty"`
	Limit   int32     `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SearchRequest) GetOrigin() *Location {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *SearchRequest) GetRadius() float32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *SearchRequest) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *SearchRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type SearchResponse struct {
	Trees []*CondensedTree `protobuf:"bytes,1,rep,name=trees" json:"trees,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SearchResponse) GetTrees() []*CondensedTree {
	if m != nil {
		return m.Trees
	}
	return nil
}

func init() {
	proto.RegisterType((*Tree)(nil), "Tree")
	proto.RegisterType((*CondensedTree)(nil), "CondensedTree")
	proto.RegisterType((*Location)(nil), "Location")
	proto.RegisterType((*Filter)(nil), "Filter")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterEnum("FilterKey", FilterKey_name, FilterKey_value)
}

func init() { proto.RegisterFile("models.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 844 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xcd, 0x6e, 0xe3, 0x36,
	0x17, 0xfd, 0x1c, 0xc7, 0x7f, 0xf4, 0x4f, 0x18, 0x7e, 0x99, 0x84, 0xc9, 0x4c, 0x30, 0x1e, 0xb7,
	0x03, 0x18, 0x5d, 0x78, 0x91, 0x02, 0x5d, 0x75, 0xd3, 0xa6, 0x68, 0x51, 0x34, 0x68, 0x01, 0x65,
	0x80, 0x02, 0xdd, 0x08, 0xb4, 0x74, 0x6b, 0x11, 0x91, 0x44, 0x95, 0xa4, 0x1a, 0x78, 0xb6, 0x5d,
	0xf6, 0x5d, 0xfa, 0x4c, 0x7d, 0x94, 0xe2, 0x92, 0xb4, 0x24, 0x07, 0x28, 0xba, 0xf3, 0x39, 0xe7,
	0xf2, 0xf2, 0x28, 0xbc, 0xf7, 0x84, 0xcc, 0x0a, 0x95, 0x42, 0x6e, 0x36, 0x95, 0x56, 0x56, 0xad,
	0xfe, 0x1e, 0x92, 0xd3, 0x0f, 0x1a, 0x80, 0x5d, 0x91, 0x91, 0xd5, 0x00, 0xb1, 0x4c, 0x79, 0x6f,
	0xd9, 0x5b, 0xf7, 0xa3, 0x21, 0xc2, 0xef, 0x53, 0x76, 0x4b, 0x48, 0xa2, 0x41, 0x58, 0x48, 0x63,
	0x61, 0xf9, 0x89, 0xd3, 0x26, 0x81, 0xf9, 0xca, 0xb2, 0x4f, 0xc8, 0xdc, 0x9d, 0x4b, 0xa5, 0x28,
	0xc0, 0x82, 0xe6, 0xfd, 0x65, 0x6f, 0x3d, 0x88, 0x66, 0x48, 0x7e, 0x13, 0x38, 0xf6, 0x9e, 0x2c,
	0x8c, 0xad, 0x8b, 0xaa, 0xad, 0x3a, 0x75, 0x55, 0x73, 0xc7, 0x36, 0x65, 0x97, 0x64, 0x68, 0xac,
	0xb0, 0xb5, 0xe1, 0x83, 0x65, 0x6f, 0x3d, 0x89, 0x02, 0x42, 0x3e, 0x03, 0x91, 0xdb, 0x8c, 0x0f,
	0x3d, 0xef, 0x11, 0x7b, 0x4d, 0x26, 0xa6, 0x4a, 0xe2, 0x5c, 0x58, 0x59, 0xf2, 0x91, 0x93, 0xc6,
	0xa6, 0x4a, 0x1e, 0x10, 0xa3, 0x6f, 0x14, 0x13, 0x55, 0x14, 0xaa, 0xe4, 0x63, 0xa7, 0x62, 0xf9,
	0xbd, 0x23, 0x18, 0x27, 0x23, 0x63, 0xe1, 0x59, 0xe8, 0x94, 0x4f, 0x9c, 0x76, 0x80, 0xf8, 0x45,
	0x49, 0xad, 0xb7, 0x71, 0xae, 0x12, 0x61, 0xa5, 0x2a, 0x39, 0x71, 0xfa, 0x0c, 0xc9, 0x87, 0xc0,
	0xa1, 0xa5, 0x5d, 0x2d, 0x74, 0x6a, 0xf8, 0xd4, 0x5b, 0xf2, 0x88, 0xdd, 0x90, 0xb1, 0x91, 0x29,
	0x3c, 0x8b, 0xfc, 0x89, 0xcf, 0x82, 0xa3, 0x80, 0xd1, 0x6e, 0x6d, 0x40, 0xc7, 0x76, 0x5f, 0x01,
	0x9f, 0x7b, 0x11, 0x89, 0x0f, 0xfb, 0x0a, 0xf0, 0x60, 0xa5, 0xd5, 0x36, 0x87, 0xc2, 0xf0, 0x85,
	0xd7, 0x0e, 0x18, 0x3f, 0x45, 0x2b, 0x65, 0x63, 0x63, 0x55, 0x09, 0xfc, 0xcc, 0x7f, 0x0a, 0x32,
	0x8f, 0x48, 0x34, 0xf2, 0x4e, 0x0b, 0x0b, 0x9c, 0xb6, 0xf2, 0x77, 0x48, 0x34, 0xb2, 0xb2, 0x19,
	0x68, 0x7e, 0xde, 0xca, 0x3f, 0x21, 0xc1, 0xde, 0x92, 0xa9, 0xd5, 0x75, 0xf9, 0x14, 0x74, 0xe6,
	0x74, 0xe2, 0x28, 0x5f, 0x70, 0x4b, 0x3c, 0x8a, 0x9f, 0xa5, 0x06, 0xfe, 0x7f, 0x7f, 0xde, 0x31,
	0x3f, 0x4b, 0x0d, 0xed, 0xf9, 0x5c, 0xee, 0x32, 0xcb, 0x2f, 0x3a, 0xe7, 0x1f, 0x90, 0x61, 0xef,
	0xc8, 0x6c, 0xab, 0x45, 0x99, 0x64, 0xa1, 0xe2, 0x95, 0xab, 0x98, 0x7a, 0xce, 0x97, 0xbc, 0x25,
	0x01, 0xc6, 0x26, 0x53, 0xc0, 0x2f, 0x7d, 0x0f, 0x4f, 0x3d, 0x66, 0x0a, 0x3a, 0x3d, 0xbc, 0xcb,
	0xab, 0x6e, 0x0f, 0x6f, 0x93, 0x93, 0x91, 0x48, 0x53, 0x0d, 0xc6, 0x70, 0xee, 0x1f, 0x34, 0x40,
	0x54, 0x3e, 0xca, 0x2a, 0x51, 0x29, 0xf0, 0x6b, 0xaf, 0x04, 0xc8, 0xae, 0xc9, 0xf8, 0xa3, 0xac,
	0xe2, 0x44, 0xda, 0x3d, 0xbf, 0x69, 0xa4, 0x7b, 0x69, 0xf7, 0xee, 0x46, 0xa5, 0x55, 0xbd, 0xcb,
	0xe2, 0x52, 0x14, 0xc0, 0x5f, 0x87, 0x1b, 0x3d, 0xf7, 0xa3, 0x28, 0x80, 0xbd, 0x27, 0xe3, 0x66,
	0x46, 0xde, 0x2c, 0x7b, 0xeb, 0xe9, 0xdd, 0x64, 0x73, 0x18, 0x90, 0xa8, 0x91, 0x56, 0x7f, 0xf6,
	0xc8, 0xfc, 0x5e, 0x95, 0x29, 0x94, 0x06, 0xd2, 0xff, 0xdc, 0xb5, 0xce, 0xcc, 0x9e, 0xbc, 0x9c,
	0xd9, 0xa3, 0x79, 0xef, 0xbf, 0x98, 0xf7, 0xae, 0x9b, 0xd3, 0x7f, 0x77, 0xb3, 0x21, 0xe3, 0x66,
	0x88, 0x29, 0xe9, 0xe7, 0xc2, 0x3a, 0x0f, 0x27, 0x11, 0xfe, 0x74, 0x4c, 0xb8, 0x19, 0x19, 0x55,
	0xae, 0xbe, 0x24, 0xc3, 0x6f, 0x65, 0x8e, 0xdb, 0xf9, 0x86, 0xf4, 0x9f, 0x60, 0xef, 0xaa, 0x17,
	0x77, 0x64, 0xe3, 0xd9, 0x1f, 0x60, 0x1f, 0x21, 0xcd, 0x2e, 0xc8, 0xe0, 0x77, 0x91, 0xd7, 0x10,
	0x5c, 0x7b, 0xb0, 0xfa, 0xa3, 0x47, 0xe6, 0x8f, 0x20, 0x74, 0x92, 0x45, 0xf0, 0x5b, 0x0d, 0x06,
	0xa7, 0x61, 0xa8, 0xb4, 0xdc, 0xc9, 0xd2, 0x35, 0x3a, 0x32, 0x19, 0x04, 0xdc, 0x2d, 0x2d, 0x52,
	0x59, 0x9b, 0xe0, 0x23, 0x20, 0xf6, 0x8e, 0x8c, 0x7e, 0x75, 0x97, 0x1a, 0xde, 0x5f, 0xf6, 0xd7,
	0xd3, 0xbb, 0x51, 0x30, 0x11, 0x1d, 0x78, 0x74, 0x91, 0xcb, 0x42, 0xda, 0x90, 0x2f, 0x1e, 0xac,
	0xbe, 0x20, 0x8b, 0x83, 0x09, 0x53, 0xa9, 0xd2, 0x00, 0xfb, 0x94, 0x0c, 0xf0, 0x4f, 0x6e, 0x78,
	0xcf, 0x35, 0x5a, 0x6c, 0x8e, 0x1e, 0x28, 0xf2, 0xe2, 0x67, 0x7f, 0xf5, 0xc9, 0xa4, 0xf9, 0x4c,
	0x36, 0x6d, 0x5e, 0x8d, 0xfe, 0x8f, 0x2d, 0xba, 0xa9, 0x48, 0x7b, 0xec, 0xfc, 0x45, 0x0c, 0xd2,
	0x13, 0xc6, 0x5e, 0x86, 0x1e, 0xed, 0x33, 0x72, 0x48, 0x38, 0x7a, 0x8a, 0xbf, 0x7d, 0x8e, 0xd1,
	0x01, 0x9b, 0x77, 0x5e, 0x96, 0x0e, 0xb1, 0x7b, 0x3b, 0x07, 0x74, 0x84, 0x57, 0x87, 0x74, 0xa2,
	0x63, 0xbc, 0xea, 0x28, 0x9f, 0xe8, 0x04, 0x5b, 0xf9, 0xfc, 0xa1, 0x84, 0xcd, 0xda, 0x04, 0xa2,
	0x53, 0x6c, 0xdc, 0x64, 0x0e, 0x9d, 0xa1, 0x78, 0x48, 0x15, 0x3a, 0xc7, 0x6b, 0xda, 0x5c, 0xa1,
	0x8b, 0x06, 0xbb, 0x20, 0xa1, 0x67, 0x0d, 0x76, 0x3b, 0x47, 0x29, 0x3b, 0x3b, 0x8a, 0x0a, 0x7a,
	0x8e, 0x05, 0x6d, 0x34, 0x50, 0xd6, 0x16, 0xb8, 0x4d, 0xa7, 0xb7, 0x8c, 0x1e, 0xef, 0x3e, 0xbd,
	0xc0, 0x92, 0xce, 0xaa, 0xd3, 0x57, 0x9d, 0x12, 0xdf, 0xf5, 0x12, 0xbf, 0x36, 0xac, 0x2e, 0xbd,
	0x42, 0x10, 0xb6, 0x95, 0x72, 0xb4, 0x7f, 0xd8, 0x57, 0x7a, 0xed, 0x4e, 0x76, 0x56, 0x94, 0xde,
	0x7c, 0x3d, 0xf9, 0xc5, 0x3d, 0x91, 0xa9, 0xb6, 0xdb, 0xa1, 0xfb, 0xff, 0xf6, 0xf9, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd4, 0x56, 0x85, 0xf6, 0xef, 0x06, 0x00, 0x00,
}