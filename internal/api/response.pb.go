// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.0
// source: response.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pos *FloatPoint `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetPos() *FloatPoint {
	if x != nil {
		return x.Pos
	}
	return nil
}

type PlayerConnected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *PlayerConnected) Reset() {
	*x = PlayerConnected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerConnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerConnected) ProtoMessage() {}

func (x *PlayerConnected) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerConnected.ProtoReflect.Descriptor instead.
func (*PlayerConnected) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerConnected) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type PlayerDisconnected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PlayerDisconnected) Reset() {
	*x = PlayerDisconnected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerDisconnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerDisconnected) ProtoMessage() {}

func (x *PlayerDisconnected) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerDisconnected.ProtoReflect.Descriptor instead.
func (*PlayerDisconnected) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerDisconnected) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EntityUpdates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *EntityUpdates) Reset() {
	*x = EntityUpdates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityUpdates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityUpdates) ProtoMessage() {}

func (x *EntityUpdates) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityUpdates.ProtoReflect.Descriptor instead.
func (*EntityUpdates) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{3}
}

func (x *EntityUpdates) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type NetUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks   []*Block       `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Entities *EntityUpdates `protobuf:"bytes,2,opt,name=entities,proto3" json:"entities,omitempty"`
	Tick     uint64         `protobuf:"varint,3,opt,name=tick,proto3" json:"tick,omitempty"`
}

func (x *NetUpdate) Reset() {
	*x = NetUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetUpdate) ProtoMessage() {}

func (x *NetUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetUpdate.ProtoReflect.Descriptor instead.
func (*NetUpdate) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{4}
}

func (x *NetUpdate) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *NetUpdate) GetEntities() *EntityUpdates {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *NetUpdate) GetTick() uint64 {
	if x != nil {
		return x.Tick
	}
	return 0
}

var File_response_proto protoreflect.FileDescriptor

var file_response_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x06, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x03, 0x70, 0x6f, 0x73, 0x22, 0x32, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32,
	0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x21, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x22, 0x6b, 0x0a, 0x09, 0x4e, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1e, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12,
	0x2a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_response_proto_rawDescOnce sync.Once
	file_response_proto_rawDescData = file_response_proto_rawDesc
)

func file_response_proto_rawDescGZIP() []byte {
	file_response_proto_rawDescOnce.Do(func() {
		file_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_proto_rawDescData)
	})
	return file_response_proto_rawDescData
}

var file_response_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_response_proto_goTypes = []interface{}{
	(*Player)(nil),             // 0: Player
	(*PlayerConnected)(nil),    // 1: PlayerConnected
	(*PlayerDisconnected)(nil), // 2: PlayerDisconnected
	(*EntityUpdates)(nil),      // 3: EntityUpdates
	(*NetUpdate)(nil),          // 4: NetUpdate
	(*FloatPoint)(nil),         // 5: FloatPoint
	(*Block)(nil),              // 6: Block
}
var file_response_proto_depIdxs = []int32{
	5, // 0: Player.pos:type_name -> FloatPoint
	0, // 1: PlayerConnected.player:type_name -> Player
	0, // 2: EntityUpdates.players:type_name -> Player
	6, // 3: NetUpdate.blocks:type_name -> Block
	3, // 4: NetUpdate.entities:type_name -> EntityUpdates
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_response_proto_init() }
func file_response_proto_init() {
	if File_response_proto != nil {
		return
	}
	file_app_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerConnected); i {
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
		file_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerDisconnected); i {
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
		file_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityUpdates); i {
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
		file_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetUpdate); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_proto_goTypes,
		DependencyIndexes: file_response_proto_depIdxs,
		MessageInfos:      file_response_proto_msgTypes,
	}.Build()
	File_response_proto = out.File
	file_response_proto_rawDesc = nil
	file_response_proto_goTypes = nil
	file_response_proto_depIdxs = nil
}
