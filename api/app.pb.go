// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.0
// source: api/app.proto

package pb

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

type BlockType int32

const (
	BlockType_Debug BlockType = 0
	BlockType_Air   BlockType = 1
	BlockType_Root  BlockType = 2
)

// Enum value maps for BlockType.
var (
	BlockType_name = map[int32]string{
		0: "Debug",
		1: "Air",
		2: "Root",
	}
	BlockType_value = map[string]int32{
		"Debug": 0,
		"Air":   1,
		"Root":  2,
	}
)

func (x BlockType) Enum() *BlockType {
	p := new(BlockType)
	*p = x
	return p
}

func (x BlockType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_app_proto_enumTypes[0].Descriptor()
}

func (BlockType) Type() protoreflect.EnumType {
	return &file_api_app_proto_enumTypes[0]
}

func (x BlockType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockType.Descriptor instead.
func (BlockType) EnumDescriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick    uint64   `protobuf:"varint,1,opt,name=tick,proto3" json:"tick,omitempty"`
	Command *Command `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetTick() uint64 {
	if x != nil {
		return x.Tick
	}
	return 0
}

func (x *Request) GetCommand() *Command {
	if x != nil {
		return x.Command
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Command_Move
	//	*Command_Shoot
	//	*Command_Changes
	Type isCommand_Type `protobuf_oneof:"type"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{1}
}

func (m *Command) GetType() isCommand_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Command) GetMove() *Move {
	if x, ok := x.GetType().(*Command_Move); ok {
		return x.Move
	}
	return nil
}

func (x *Command) GetShoot() *Shoot {
	if x, ok := x.GetType().(*Command_Shoot); ok {
		return x.Shoot
	}
	return nil
}

func (x *Command) GetChanges() *GetChanges {
	if x, ok := x.GetType().(*Command_Changes); ok {
		return x.Changes
	}
	return nil
}

type isCommand_Type interface {
	isCommand_Type()
}

type Command_Move struct {
	Move *Move `protobuf:"bytes,1,opt,name=move,proto3,oneof"`
}

type Command_Shoot struct {
	Shoot *Shoot `protobuf:"bytes,2,opt,name=shoot,proto3,oneof"`
}

type Command_Changes struct {
	Changes *GetChanges `protobuf:"bytes,3,opt,name=changes,proto3,oneof"`
}

func (*Command_Move) isCommand_Type() {}

func (*Command_Shoot) isCommand_Type() {}

func (*Command_Changes) isCommand_Type() {}

type GetChanges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTick uint64 `protobuf:"varint,1,opt,name=startTick,proto3" json:"startTick,omitempty"`
	EndTick   uint64 `protobuf:"varint,2,opt,name=endTick,proto3" json:"endTick,omitempty"`
}

func (x *GetChanges) Reset() {
	*x = GetChanges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChanges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChanges) ProtoMessage() {}

func (x *GetChanges) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChanges.ProtoReflect.Descriptor instead.
func (*GetChanges) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{2}
}

func (x *GetChanges) GetStartTick() uint64 {
	if x != nil {
		return x.StartTick
	}
	return 0
}

func (x *GetChanges) GetEndTick() uint64 {
	if x != nil {
		return x.EndTick
	}
	return 0
}

type Shoot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point *WorldPoint `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Shoot) Reset() {
	*x = Shoot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shoot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shoot) ProtoMessage() {}

func (x *Shoot) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shoot.ProtoReflect.Descriptor instead.
func (*Shoot) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{3}
}

func (x *Shoot) GetPoint() *WorldPoint {
	if x != nil {
		return x.Point
	}
	return nil
}

type Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point *WorldPoint `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Move) Reset() {
	*x = Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Move) ProtoMessage() {}

func (x *Move) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Move.ProtoReflect.Descriptor instead.
func (*Move) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{4}
}

func (x *Move) GetPoint() *WorldPoint {
	if x != nil {
		return x.Point
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick    uint64   `protobuf:"varint,1,opt,name=tick,proto3" json:"tick,omitempty"`
	Payload *Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetTick() uint64 {
	if x != nil {
		return x.Tick
	}
	return 0
}

func (x *Response) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Payload_Players
	//	*Payload_PlayerConnect
	//	*Payload_PlayerDisconnect
	//	*Payload_Changes
	Type isPayload_Type `protobuf_oneof:"type"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{6}
}

func (m *Payload) GetType() isPayload_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Payload) GetPlayers() *AllPlayers {
	if x, ok := x.GetType().(*Payload_Players); ok {
		return x.Players
	}
	return nil
}

func (x *Payload) GetPlayerConnect() *Player {
	if x, ok := x.GetType().(*Payload_PlayerConnect); ok {
		return x.PlayerConnect
	}
	return nil
}

func (x *Payload) GetPlayerDisconnect() *Player {
	if x, ok := x.GetType().(*Payload_PlayerDisconnect); ok {
		return x.PlayerDisconnect
	}
	return nil
}

func (x *Payload) GetChanges() *Changes {
	if x, ok := x.GetType().(*Payload_Changes); ok {
		return x.Changes
	}
	return nil
}

type isPayload_Type interface {
	isPayload_Type()
}

type Payload_Players struct {
	Players *AllPlayers `protobuf:"bytes,1,opt,name=players,proto3,oneof"`
}

type Payload_PlayerConnect struct {
	PlayerConnect *Player `protobuf:"bytes,2,opt,name=playerConnect,proto3,oneof"`
}

type Payload_PlayerDisconnect struct {
	PlayerDisconnect *Player `protobuf:"bytes,3,opt,name=playerDisconnect,proto3,oneof"`
}

type Payload_Changes struct {
	Changes *Changes `protobuf:"bytes,4,opt,name=changes,proto3,oneof"`
}

func (*Payload_Players) isPayload_Type() {}

func (*Payload_PlayerConnect) isPayload_Type() {}

func (*Payload_PlayerDisconnect) isPayload_Type() {}

func (*Payload_Changes) isPayload_Type() {}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Point *FloatPoint `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[7]
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
	return file_api_app_proto_rawDescGZIP(), []int{7}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetPoint() *FloatPoint {
	if x != nil {
		return x.Point
	}
	return nil
}

type AllPlayers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player []*Player `protobuf:"bytes,1,rep,name=player,proto3" json:"player,omitempty"`
}

func (x *AllPlayers) Reset() {
	*x = AllPlayers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllPlayers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllPlayers) ProtoMessage() {}

func (x *AllPlayers) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllPlayers.ProtoReflect.Descriptor instead.
func (*AllPlayers) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{8}
}

func (x *AllPlayers) GetPlayer() []*Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type FloatPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *FloatPoint) Reset() {
	*x = FloatPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatPoint) ProtoMessage() {}

func (x *FloatPoint) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatPoint.ProtoReflect.Descriptor instead.
func (*FloatPoint) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{9}
}

func (x *FloatPoint) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *FloatPoint) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *FloatPoint) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type WorldPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *WorldPoint) Reset() {
	*x = WorldPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldPoint) ProtoMessage() {}

func (x *WorldPoint) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldPoint.ProtoReflect.Descriptor instead.
func (*WorldPoint) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{10}
}

func (x *WorldPoint) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *WorldPoint) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *WorldPoint) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Changes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *Changes) Reset() {
	*x = Changes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Changes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Changes) ProtoMessage() {}

func (x *Changes) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Changes.ProtoReflect.Descriptor instead.
func (*Changes) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{11}
}

func (x *Changes) GetChanges() []*Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point     []*WorldPoint `protobuf:"bytes,1,rep,name=point,proto3" json:"point,omitempty"`
	BlockType BlockType     `protobuf:"varint,2,opt,name=blockType,proto3,enum=pb.BlockType" json:"blockType,omitempty"`
}

func (x *Change) Reset() {
	*x = Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Change) ProtoMessage() {}

func (x *Change) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Change.ProtoReflect.Descriptor instead.
func (*Change) Descriptor() ([]byte, []int) {
	return file_api_app_proto_rawDescGZIP(), []int{12}
}

func (x *Change) GetPoint() []*WorldPoint {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *Change) GetBlockType() BlockType {
	if x != nil {
		return x.BlockType
	}
	return BlockType_Debug
}

var File_api_app_proto protoreflect.FileDescriptor

var file_api_app_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x44, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69,
	0x63, 0x6b, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52,
	0x04, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x6f, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x48, 0x00, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x63, 0x6b, 0x22, 0x2d, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x6f, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x22, 0x2c, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22,
	0x45, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12,
	0x25, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xd4, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x32,
	0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x38, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52, 0x10, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x48, 0x00, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a,
	0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x30, 0x0a,
	0x0a, 0x41, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22,
	0x36, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x22, 0x36, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x22,
	0x2f, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x22, 0x5b, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x29, 0x0a,
	0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x69, 0x72, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x74, 0x10, 0x02, 0x42, 0x05, 0x5a, 0x03, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_app_proto_rawDescOnce sync.Once
	file_api_app_proto_rawDescData = file_api_app_proto_rawDesc
)

func file_api_app_proto_rawDescGZIP() []byte {
	file_api_app_proto_rawDescOnce.Do(func() {
		file_api_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_app_proto_rawDescData)
	})
	return file_api_app_proto_rawDescData
}

var file_api_app_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_app_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_app_proto_goTypes = []interface{}{
	(BlockType)(0),     // 0: pb.BlockType
	(*Request)(nil),    // 1: pb.Request
	(*Command)(nil),    // 2: pb.Command
	(*GetChanges)(nil), // 3: pb.GetChanges
	(*Shoot)(nil),      // 4: pb.Shoot
	(*Move)(nil),       // 5: pb.Move
	(*Response)(nil),   // 6: pb.Response
	(*Payload)(nil),    // 7: pb.Payload
	(*Player)(nil),     // 8: pb.Player
	(*AllPlayers)(nil), // 9: pb.AllPlayers
	(*FloatPoint)(nil), // 10: pb.FloatPoint
	(*WorldPoint)(nil), // 11: pb.WorldPoint
	(*Changes)(nil),    // 12: pb.Changes
	(*Change)(nil),     // 13: pb.Change
}
var file_api_app_proto_depIdxs = []int32{
	2,  // 0: pb.Request.command:type_name -> pb.Command
	5,  // 1: pb.Command.move:type_name -> pb.Move
	4,  // 2: pb.Command.shoot:type_name -> pb.Shoot
	3,  // 3: pb.Command.changes:type_name -> pb.GetChanges
	11, // 4: pb.Shoot.point:type_name -> pb.WorldPoint
	11, // 5: pb.Move.point:type_name -> pb.WorldPoint
	7,  // 6: pb.Response.payload:type_name -> pb.Payload
	9,  // 7: pb.Payload.players:type_name -> pb.AllPlayers
	8,  // 8: pb.Payload.playerConnect:type_name -> pb.Player
	8,  // 9: pb.Payload.playerDisconnect:type_name -> pb.Player
	12, // 10: pb.Payload.changes:type_name -> pb.Changes
	10, // 11: pb.Player.point:type_name -> pb.FloatPoint
	8,  // 12: pb.AllPlayers.player:type_name -> pb.Player
	13, // 13: pb.Changes.changes:type_name -> pb.Change
	11, // 14: pb.Change.point:type_name -> pb.WorldPoint
	0,  // 15: pb.Change.blockType:type_name -> pb.BlockType
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_api_app_proto_init() }
func file_api_app_proto_init() {
	if File_api_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_api_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_api_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChanges); i {
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
		file_api_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shoot); i {
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
		file_api_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Move); i {
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
		file_api_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_api_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_api_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_app_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllPlayers); i {
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
		file_api_app_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloatPoint); i {
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
		file_api_app_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldPoint); i {
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
		file_api_app_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Changes); i {
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
		file_api_app_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Change); i {
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
	file_api_app_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Command_Move)(nil),
		(*Command_Shoot)(nil),
		(*Command_Changes)(nil),
	}
	file_api_app_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Payload_Players)(nil),
		(*Payload_PlayerConnect)(nil),
		(*Payload_PlayerDisconnect)(nil),
		(*Payload_Changes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_app_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_app_proto_goTypes,
		DependencyIndexes: file_api_app_proto_depIdxs,
		EnumInfos:         file_api_app_proto_enumTypes,
		MessageInfos:      file_api_app_proto_msgTypes,
	}.Build()
	File_api_app_proto = out.File
	file_api_app_proto_rawDesc = nil
	file_api_app_proto_goTypes = nil
	file_api_app_proto_depIdxs = nil
}
