// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: trading.proto

package trading_service

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

type Share struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Share string  `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Share) Reset() {
	*x = Share{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Share) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Share) ProtoMessage() {}

func (x *Share) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Share.ProtoReflect.Descriptor instead.
func (*Share) Descriptor() ([]byte, []int) {
	return file_trading_proto_rawDescGZIP(), []int{0}
}

func (x *Share) GetShare() string {
	if x != nil {
		return x.Share
	}
	return ""
}

func (x *Share) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsLong      bool    `protobuf:"varint,2,opt,name=isLong,proto3" json:"isLong,omitempty"`
	ShareName   string  `protobuf:"bytes,3,opt,name=shareName,proto3" json:"shareName,omitempty"`
	SharePrice  float64 `protobuf:"fixed64,4,opt,name=sharePrice,proto3" json:"sharePrice,omitempty"`
	Total       float64 `protobuf:"fixed64,5,opt,name=total,proto3" json:"total,omitempty"`
	ShareAmount float64 `protobuf:"fixed64,6,opt,name=shareAmount,proto3" json:"shareAmount,omitempty"`
	StopLoss    float64 `protobuf:"fixed64,7,opt,name=stopLoss,proto3" json:"stopLoss,omitempty"`
	TakeProfit  float64 `protobuf:"fixed64,8,opt,name=takeProfit,proto3" json:"takeProfit,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_trading_proto_rawDescGZIP(), []int{1}
}

func (x *Position) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Position) GetIsLong() bool {
	if x != nil {
		return x.IsLong
	}
	return false
}

func (x *Position) GetShareName() string {
	if x != nil {
		return x.ShareName
	}
	return ""
}

func (x *Position) GetSharePrice() float64 {
	if x != nil {
		return x.SharePrice
	}
	return 0
}

func (x *Position) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Position) GetShareAmount() float64 {
	if x != nil {
		return x.ShareAmount
	}
	return 0
}

func (x *Position) GetStopLoss() float64 {
	if x != nil {
		return x.StopLoss
	}
	return 0
}

func (x *Position) GetTakeProfit() float64 {
	if x != nil {
		return x.TakeProfit
	}
	return 0
}

type OpenPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *OpenPositionRequest) Reset() {
	*x = OpenPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenPositionRequest) ProtoMessage() {}

func (x *OpenPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenPositionRequest.ProtoReflect.Descriptor instead.
func (*OpenPositionRequest) Descriptor() ([]byte, []int) {
	return file_trading_proto_rawDescGZIP(), []int{2}
}

func (x *OpenPositionRequest) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type OpenPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *OpenPositionResponse) Reset() {
	*x = OpenPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenPositionResponse) ProtoMessage() {}

func (x *OpenPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenPositionResponse.ProtoReflect.Descriptor instead.
func (*OpenPositionResponse) Descriptor() ([]byte, []int) {
	return file_trading_proto_rawDescGZIP(), []int{3}
}

func (x *OpenPositionResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ClosePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ClosePositionRequest) Reset() {
	*x = ClosePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClosePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosePositionRequest) ProtoMessage() {}

func (x *ClosePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosePositionRequest.ProtoReflect.Descriptor instead.
func (*ClosePositionRequest) Descriptor() ([]byte, []int) {
	return file_trading_proto_rawDescGZIP(), []int{4}
}

func (x *ClosePositionRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ClosePositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PnL float64 `protobuf:"fixed64,1,opt,name=PnL,proto3" json:"PnL,omitempty"`
}

func (x *ClosePositionResponse) Reset() {
	*x = ClosePositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trading_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClosePositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosePositionResponse) ProtoMessage() {}

func (x *ClosePositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trading_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosePositionResponse.ProtoReflect.Descriptor instead.
func (*ClosePositionResponse) Descriptor() ([]byte, []int) {
	return file_trading_proto_rawDescGZIP(), []int{5}
}

func (x *ClosePositionResponse) GetPnL() float64 {
	if x != nil {
		return x.PnL
	}
	return 0
}

var File_trading_proto protoreflect.FileDescriptor

var file_trading_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x33, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0xe4, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x74, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x22, 0x3c, 0x0a, 0x13, 0x4f,
	0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x14, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x15, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x6e, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x50, 0x6e, 0x4c, 0x32, 0x8d, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x75, 0x67, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x6d, 0x61, 0x2f, 0x74, 0x72,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trading_proto_rawDescOnce sync.Once
	file_trading_proto_rawDescData = file_trading_proto_rawDesc
)

func file_trading_proto_rawDescGZIP() []byte {
	file_trading_proto_rawDescOnce.Do(func() {
		file_trading_proto_rawDescData = protoimpl.X.CompressGZIP(file_trading_proto_rawDescData)
	})
	return file_trading_proto_rawDescData
}

var file_trading_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_trading_proto_goTypes = []interface{}{
	(*Share)(nil),                 // 0: Share
	(*Position)(nil),              // 1: Position
	(*OpenPositionRequest)(nil),   // 2: OpenPositionRequest
	(*OpenPositionResponse)(nil),  // 3: OpenPositionResponse
	(*ClosePositionRequest)(nil),  // 4: ClosePositionRequest
	(*ClosePositionResponse)(nil), // 5: ClosePositionResponse
}
var file_trading_proto_depIdxs = []int32{
	1, // 0: OpenPositionRequest.position:type_name -> Position
	2, // 1: TradingService.OpenPosition:input_type -> OpenPositionRequest
	4, // 2: TradingService.ClosePosition:input_type -> ClosePositionRequest
	3, // 3: TradingService.OpenPosition:output_type -> OpenPositionResponse
	5, // 4: TradingService.ClosePosition:output_type -> ClosePositionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trading_proto_init() }
func file_trading_proto_init() {
	if File_trading_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trading_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Share); i {
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
		file_trading_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_trading_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenPositionRequest); i {
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
		file_trading_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenPositionResponse); i {
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
		file_trading_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClosePositionRequest); i {
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
		file_trading_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClosePositionResponse); i {
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
			RawDescriptor: file_trading_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trading_proto_goTypes,
		DependencyIndexes: file_trading_proto_depIdxs,
		MessageInfos:      file_trading_proto_msgTypes,
	}.Build()
	File_trading_proto = out.File
	file_trading_proto_rawDesc = nil
	file_trading_proto_goTypes = nil
	file_trading_proto_depIdxs = nil
}
