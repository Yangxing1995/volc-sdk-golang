// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.4
// source: volcengine/vod/business/vod_trade.proto

package business

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TradeConfigurationInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basic    []*TradeProduct `protobuf:"bytes,1,rep,name=Basic,proto3" json:"Basic,omitempty"`
	Addition []*TradeProduct `protobuf:"bytes,2,rep,name=Addition,proto3" json:"Addition,omitempty"`
}

func (x *TradeConfigurationInfoResult) Reset() {
	*x = TradeConfigurationInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeConfigurationInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeConfigurationInfoResult) ProtoMessage() {}

func (x *TradeConfigurationInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeConfigurationInfoResult.ProtoReflect.Descriptor instead.
func (*TradeConfigurationInfoResult) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_trade_proto_rawDescGZIP(), []int{0}
}

func (x *TradeConfigurationInfoResult) GetBasic() []*TradeProduct {
	if x != nil {
		return x.Basic
	}
	return nil
}

func (x *TradeConfigurationInfoResult) GetAddition() []*TradeProduct {
	if x != nil {
		return x.Addition
	}
	return nil
}

type TradeProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configuration          string        `protobuf:"bytes,1,opt,name=Configuration,proto3" json:"Configuration,omitempty"`
	ConfigurationName      string        `protobuf:"bytes,2,opt,name=ConfigurationName,proto3" json:"ConfigurationName,omitempty"`
	Product                string        `protobuf:"bytes,3,opt,name=Product,proto3" json:"Product,omitempty"`
	SettlementPeriod       int32         `protobuf:"varint,4,opt,name=SettlementPeriod,proto3" json:"SettlementPeriod,omitempty"`
	CanPrePay              int32         `protobuf:"varint,5,opt,name=CanPrePay,proto3" json:"CanPrePay,omitempty"`
	ChargeItems            []*ChargeItem `protobuf:"bytes,6,rep,name=ChargeItems,proto3" json:"ChargeItems,omitempty"`
	InstanceID             string        `protobuf:"bytes,7,opt,name=InstanceID,proto3" json:"InstanceID,omitempty"`
	InstanceStatus         int32         `protobuf:"varint,8,opt,name=InstanceStatus,proto3" json:"InstanceStatus,omitempty"`
	InstanceCreateTime     string        `protobuf:"bytes,9,opt,name=InstanceCreateTime,proto3" json:"InstanceCreateTime,omitempty"`
	NextEventEffectiveTime string        `protobuf:"bytes,10,opt,name=NextEventEffectiveTime,proto3" json:"NextEventEffectiveTime,omitempty"`
	NextEventType          string        `protobuf:"bytes,11,opt,name=NextEventType,proto3" json:"NextEventType,omitempty"`
}

func (x *TradeProduct) Reset() {
	*x = TradeProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_trade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeProduct) ProtoMessage() {}

func (x *TradeProduct) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_trade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeProduct.ProtoReflect.Descriptor instead.
func (*TradeProduct) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_trade_proto_rawDescGZIP(), []int{1}
}

func (x *TradeProduct) GetConfiguration() string {
	if x != nil {
		return x.Configuration
	}
	return ""
}

func (x *TradeProduct) GetConfigurationName() string {
	if x != nil {
		return x.ConfigurationName
	}
	return ""
}

func (x *TradeProduct) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *TradeProduct) GetSettlementPeriod() int32 {
	if x != nil {
		return x.SettlementPeriod
	}
	return 0
}

func (x *TradeProduct) GetCanPrePay() int32 {
	if x != nil {
		return x.CanPrePay
	}
	return 0
}

func (x *TradeProduct) GetChargeItems() []*ChargeItem {
	if x != nil {
		return x.ChargeItems
	}
	return nil
}

func (x *TradeProduct) GetInstanceID() string {
	if x != nil {
		return x.InstanceID
	}
	return ""
}

func (x *TradeProduct) GetInstanceStatus() int32 {
	if x != nil {
		return x.InstanceStatus
	}
	return 0
}

func (x *TradeProduct) GetInstanceCreateTime() string {
	if x != nil {
		return x.InstanceCreateTime
	}
	return ""
}

func (x *TradeProduct) GetNextEventEffectiveTime() string {
	if x != nil {
		return x.NextEventEffectiveTime
	}
	return ""
}

func (x *TradeProduct) GetNextEventType() string {
	if x != nil {
		return x.NextEventType
	}
	return ""
}

type ChargeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChargeItemCode string  `protobuf:"bytes,1,opt,name=ChargeItemCode,proto3" json:"ChargeItemCode,omitempty"`
	Price          float64 `protobuf:"fixed64,2,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *ChargeItem) Reset() {
	*x = ChargeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_trade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeItem) ProtoMessage() {}

func (x *ChargeItem) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_trade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeItem.ProtoReflect.Descriptor instead.
func (*ChargeItem) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_trade_proto_rawDescGZIP(), []int{2}
}

func (x *ChargeItem) GetChargeItemCode() string {
	if x != nil {
		return x.ChargeItemCode
	}
	return ""
}

func (x *ChargeItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_volcengine_vod_business_vod_trade_proto protoreflect.FileDescriptor

var file_volcengine_vod_business_vod_trade_proto_rawDesc = []byte{
	0x0a, 0x27, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x64,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x5f, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x56, 0x6f, 0x6c, 0x63, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x1c, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x42, 0x0a, 0x05, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x56, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x48,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f,
	0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xea, 0x03, 0x0a, 0x0c, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x50, 0x72, 0x65, 0x50, 0x61, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x61, 0x6e, 0x50, 0x72, 0x65, 0x50, 0x61,
	0x79, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0b, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x4e, 0x65, 0x78, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x4e, 0x65, 0x78, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x4e, 0x65, 0x78, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4e, 0x65, 0x78, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x42, 0xcc, 0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x6f, 0x64,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42,
	0x08, 0x56, 0x6f, 0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01,
	0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x20, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x23, 0x56, 0x6f, 0x6c,
	0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_volcengine_vod_business_vod_trade_proto_rawDescOnce sync.Once
	file_volcengine_vod_business_vod_trade_proto_rawDescData = file_volcengine_vod_business_vod_trade_proto_rawDesc
)

func file_volcengine_vod_business_vod_trade_proto_rawDescGZIP() []byte {
	file_volcengine_vod_business_vod_trade_proto_rawDescOnce.Do(func() {
		file_volcengine_vod_business_vod_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_volcengine_vod_business_vod_trade_proto_rawDescData)
	})
	return file_volcengine_vod_business_vod_trade_proto_rawDescData
}

var file_volcengine_vod_business_vod_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_volcengine_vod_business_vod_trade_proto_goTypes = []interface{}{
	(*TradeConfigurationInfoResult)(nil), // 0: Volcengine.Vod.Models.Business.TradeConfigurationInfoResult
	(*TradeProduct)(nil),                 // 1: Volcengine.Vod.Models.Business.TradeProduct
	(*ChargeItem)(nil),                   // 2: Volcengine.Vod.Models.Business.ChargeItem
}
var file_volcengine_vod_business_vod_trade_proto_depIdxs = []int32{
	1, // 0: Volcengine.Vod.Models.Business.TradeConfigurationInfoResult.Basic:type_name -> Volcengine.Vod.Models.Business.TradeProduct
	1, // 1: Volcengine.Vod.Models.Business.TradeConfigurationInfoResult.Addition:type_name -> Volcengine.Vod.Models.Business.TradeProduct
	2, // 2: Volcengine.Vod.Models.Business.TradeProduct.ChargeItems:type_name -> Volcengine.Vod.Models.Business.ChargeItem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_volcengine_vod_business_vod_trade_proto_init() }
func file_volcengine_vod_business_vod_trade_proto_init() {
	if File_volcengine_vod_business_vod_trade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_volcengine_vod_business_vod_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeConfigurationInfoResult); i {
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
		file_volcengine_vod_business_vod_trade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeProduct); i {
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
		file_volcengine_vod_business_vod_trade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeItem); i {
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
			RawDescriptor: file_volcengine_vod_business_vod_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_volcengine_vod_business_vod_trade_proto_goTypes,
		DependencyIndexes: file_volcengine_vod_business_vod_trade_proto_depIdxs,
		MessageInfos:      file_volcengine_vod_business_vod_trade_proto_msgTypes,
	}.Build()
	File_volcengine_vod_business_vod_trade_proto = out.File
	file_volcengine_vod_business_vod_trade_proto_rawDesc = nil
	file_volcengine_vod_business_vod_trade_proto_goTypes = nil
	file_volcengine_vod_business_vod_trade_proto_depIdxs = nil
}
