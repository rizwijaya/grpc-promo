// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: promo.proto

package pb

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

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromoCode string `protobuf:"bytes,1,opt,name=promo_code,json=promoCode,proto3" json:"promo_code,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetPromoCode() string {
	if x != nil {
		return x.PromoCode
	}
	return ""
}

type Promo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromoId   string `protobuf:"bytes,1,opt,name=promo_id,json=promoId,proto3" json:"promo_id,omitempty"`
	PromoAttr string `protobuf:"bytes,2,opt,name=promo_attr,json=promoAttr,proto3" json:"promo_attr,omitempty"`
	PromoCode string `protobuf:"bytes,3,opt,name=promo_code,json=promoCode,proto3" json:"promo_code,omitempty"`
}

func (x *Promo) Reset() {
	*x = Promo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promo) ProtoMessage() {}

func (x *Promo) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promo.ProtoReflect.Descriptor instead.
func (*Promo) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{1}
}

func (x *Promo) GetPromoId() string {
	if x != nil {
		return x.PromoId
	}
	return ""
}

func (x *Promo) GetPromoAttr() string {
	if x != nil {
		return x.PromoAttr
	}
	return ""
}

func (x *Promo) GetPromoCode() string {
	if x != nil {
		return x.PromoCode
	}
	return ""
}

type Pages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Pages) Reset() {
	*x = Pages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pages) ProtoMessage() {}

func (x *Pages) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pages.ProtoReflect.Descriptor instead.
func (*Pages) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{2}
}

func (x *Pages) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pages) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderAttr string   `protobuf:"bytes,2,opt,name=order_attr,json=orderAttr,proto3" json:"order_attr,omitempty"`
	PromoCode string   `protobuf:"bytes,3,opt,name=promo_code,json=promoCode,proto3" json:"promo_code,omitempty"`
	Promos    []*Promo `protobuf:"bytes,4,rep,name=promos,proto3" json:"promos,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetOrderAttr() string {
	if x != nil {
		return x.OrderAttr
	}
	return ""
}

func (x *Order) GetPromoCode() string {
	if x != nil {
		return x.PromoCode
	}
	return ""
}

func (x *Order) GetPromos() []*Promo {
	if x != nil {
		return x.Promos
	}
	return nil
}

type ListOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderAttr string `protobuf:"bytes,2,opt,name=order_attr,json=orderAttr,proto3" json:"order_attr,omitempty"`
	PromoCode string `protobuf:"bytes,3,opt,name=promo_code,json=promoCode,proto3" json:"promo_code,omitempty"`
}

func (x *ListOrder) Reset() {
	*x = ListOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrder) ProtoMessage() {}

func (x *ListOrder) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrder.ProtoReflect.Descriptor instead.
func (*ListOrder) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{4}
}

func (x *ListOrder) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ListOrder) GetOrderAttr() string {
	if x != nil {
		return x.OrderAttr
	}
	return ""
}

func (x *ListOrder) GetPromoCode() string {
	if x != nil {
		return x.PromoCode
	}
	return ""
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination  `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Orders     []*ListOrder `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{5}
}

func (x *Orders) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *Orders) GetOrders() []*ListOrder {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderId) Reset() {
	*x = OrderId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderId) ProtoMessage() {}

func (x *OrderId) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderId.ProtoReflect.Descriptor instead.
func (*OrderId) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{6}
}

func (x *OrderId) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderAttr string `protobuf:"bytes,1,opt,name=order_attr,json=orderAttr,proto3" json:"order_attr,omitempty"`
	PromoCode string `protobuf:"bytes,2,opt,name=promo_code,json=promoCode,proto3" json:"promo_code,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{7}
}

func (x *CreateOrderRequest) GetOrderAttr() string {
	if x != nil {
		return x.OrderAttr
	}
	return ""
}

func (x *CreateOrderRequest) GetPromoCode() string {
	if x != nil {
		return x.PromoCode
	}
	return ""
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total       uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PerPage     uint32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	CurrentPage uint32 `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	LastPage    uint32 `protobuf:"varint,4,opt,name=last_page,json=lastPage,proto3" json:"last_page,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{8}
}

func (x *Pagination) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetPerPage() uint32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *Pagination) GetCurrentPage() uint32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Pagination) GetLastPage() uint32 {
	if x != nil {
		return x.LastPage
	}
	return 0
}

var File_promo_proto protoreflect.FileDescriptor

var file_promo_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x22, 0x25, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x60, 0x0a, 0x05, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x41, 0x74, 0x74, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x31, 0x0a,
	0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x86, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x74, 0x74, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x22, 0x64, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x65, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x24, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x74, 0x74,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x7d, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x32,
	0xd0, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x42, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x12, 0x2b, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x1b, 0x5a, 0x19, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_promo_proto_rawDescOnce sync.Once
	file_promo_proto_rawDescData = file_promo_proto_rawDesc
)

func file_promo_proto_rawDescGZIP() []byte {
	file_promo_proto_rawDescOnce.Do(func() {
		file_promo_proto_rawDescData = protoimpl.X.CompressGZIP(file_promo_proto_rawDescData)
	})
	return file_promo_proto_rawDescData
}

var file_promo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_promo_proto_goTypes = []interface{}{
	(*Code)(nil),               // 0: promo.Code
	(*Promo)(nil),              // 1: promo.Promo
	(*Pages)(nil),              // 2: promo.Pages
	(*Order)(nil),              // 3: promo.Order
	(*ListOrder)(nil),          // 4: promo.ListOrder
	(*Orders)(nil),             // 5: promo.Orders
	(*OrderId)(nil),            // 6: promo.OrderId
	(*CreateOrderRequest)(nil), // 7: promo.CreateOrderRequest
	(*Pagination)(nil),         // 8: promo.Pagination
}
var file_promo_proto_depIdxs = []int32{
	1, // 0: promo.Order.promos:type_name -> promo.Promo
	8, // 1: promo.Orders.pagination:type_name -> promo.Pagination
	4, // 2: promo.Orders.orders:type_name -> promo.ListOrder
	0, // 3: promo.PromoService.GetPromoByCode:input_type -> promo.Code
	2, // 4: promo.PromoService.GetOrderList:input_type -> promo.Pages
	6, // 5: promo.PromoService.GetOrderDetail:input_type -> promo.OrderId
	7, // 6: promo.PromoService.CreateOrder:input_type -> promo.CreateOrderRequest
	1, // 7: promo.PromoService.GetPromoByCode:output_type -> promo.Promo
	5, // 8: promo.PromoService.GetOrderList:output_type -> promo.Orders
	3, // 9: promo.PromoService.GetOrderDetail:output_type -> promo.Order
	3, // 10: promo.PromoService.CreateOrder:output_type -> promo.Order
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_promo_proto_init() }
func file_promo_proto_init() {
	if File_promo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_promo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_promo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Promo); i {
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
		file_promo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pages); i {
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
		file_promo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_promo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrder); i {
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
		file_promo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
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
		file_promo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderId); i {
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
		file_promo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_promo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
			RawDescriptor: file_promo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_promo_proto_goTypes,
		DependencyIndexes: file_promo_proto_depIdxs,
		MessageInfos:      file_promo_proto_msgTypes,
	}.Build()
	File_promo_proto = out.File
	file_promo_proto_rawDesc = nil
	file_promo_proto_goTypes = nil
	file_promo_proto_depIdxs = nil
}
