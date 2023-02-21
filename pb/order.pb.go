// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: pb/order.proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string        `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Buyer     *OrderBuyer   `protobuf:"bytes,2,opt,name=buyer,proto3" json:"buyer,omitempty"`
	Product   *OrderProduct `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	Payment   *OrderPayment `protobuf:"bytes,4,opt,name=payment,proto3" json:"payment,omitempty"`
	Status    string        `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt int64         `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64         `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[0]
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
	return file_pb_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetBuyer() *OrderBuyer {
	if x != nil {
		return x.Buyer
	}
	return nil
}

func (x *Order) GetProduct() *OrderProduct {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *Order) GetPayment() *OrderPayment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Order) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type OrderProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price       int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Duration    int64  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *OrderProduct) Reset() {
	*x = OrderProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProduct) ProtoMessage() {}

func (x *OrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProduct.ProtoReflect.Descriptor instead.
func (*OrderProduct) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderProduct) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderProduct) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *OrderProduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type OrderBuyer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	User       string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *OrderBuyer) Reset() {
	*x = OrderBuyer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBuyer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBuyer) ProtoMessage() {}

func (x *OrderBuyer) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBuyer.ProtoReflect.Descriptor instead.
func (*OrderBuyer) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderBuyer) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *OrderBuyer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderBuyer) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type OrderPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentType string `protobuf:"bytes,1,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	OrderId     string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Bank        string `protobuf:"bytes,3,opt,name=bank,proto3" json:"bank,omitempty"`
	VaNumber    string `protobuf:"bytes,4,opt,name=va_number,json=vaNumber,proto3" json:"va_number,omitempty"`
	GrossAmount int64  `protobuf:"varint,5,opt,name=gross_amount,json=grossAmount,proto3" json:"gross_amount,omitempty"`
}

func (x *OrderPayment) Reset() {
	*x = OrderPayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayment) ProtoMessage() {}

func (x *OrderPayment) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayment.ProtoReflect.Descriptor instead.
func (*OrderPayment) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderPayment) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *OrderPayment) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderPayment) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *OrderPayment) GetVaNumber() string {
	if x != nil {
		return x.VaNumber
	}
	return ""
}

func (x *OrderPayment) GetGrossAmount() int64 {
	if x != nil {
		return x.GrossAmount
	}
	return 0
}

type OrderCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buyer   *OrderBuyer   `protobuf:"bytes,1,opt,name=buyer,proto3" json:"buyer,omitempty"`
	Product *OrderProduct `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Payment *OrderPayment `protobuf:"bytes,3,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *OrderCreateRequest) Reset() {
	*x = OrderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateRequest) ProtoMessage() {}

func (x *OrderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateRequest.ProtoReflect.Descriptor instead.
func (*OrderCreateRequest) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderCreateRequest) GetBuyer() *OrderBuyer {
	if x != nil {
		return x.Buyer
	}
	return nil
}

func (x *OrderCreateRequest) GetProduct() *OrderProduct {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *OrderCreateRequest) GetPayment() *OrderPayment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type OrderChangeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId        string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status         string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	SettlementTime int64  `protobuf:"varint,3,opt,name=settlement_time,json=settlementTime,proto3" json:"settlement_time,omitempty"`
}

func (x *OrderChangeStatus) Reset() {
	*x = OrderChangeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderChangeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderChangeStatus) ProtoMessage() {}

func (x *OrderChangeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderChangeStatus.ProtoReflect.Descriptor instead.
func (*OrderChangeStatus) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderChangeStatus) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderChangeStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderChangeStatus) GetSettlementTime() int64 {
	if x != nil {
		return x.SettlementTime
	}
	return 0
}

type OrderFindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderFindOneRequest) Reset() {
	*x = OrderFindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFindOneRequest) ProtoMessage() {}

func (x *OrderFindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFindOneRequest.ProtoReflect.Descriptor instead.
func (*OrderFindOneRequest) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderFindOneRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OrderFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sort      string `protobuf:"bytes,1,opt,name=sort,proto3" json:"sort,omitempty"`
	Page      int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage   int64  `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Search    string `protobuf:"bytes,4,opt,name=search,proto3" json:"search,omitempty"`
	Status    string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	UserId    string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CountOnly bool   `protobuf:"varint,7,opt,name=count_only,json=countOnly,proto3" json:"count_only,omitempty"`
}

func (x *OrderFindAllRequest) Reset() {
	*x = OrderFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFindAllRequest) ProtoMessage() {}

func (x *OrderFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFindAllRequest.ProtoReflect.Descriptor instead.
func (*OrderFindAllRequest) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{7}
}

func (x *OrderFindAllRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *OrderFindAllRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OrderFindAllRequest) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *OrderFindAllRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *OrderFindAllRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderFindAllRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderFindAllRequest) GetCountOnly() bool {
	if x != nil {
		return x.CountOnly
	}
	return false
}

type OrderFindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders     []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Rows       int64    `protobuf:"varint,2,opt,name=rows,proto3" json:"rows,omitempty"`
	Pages      int64    `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	PerPage    int64    `protobuf:"varint,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	ActivePage int64    `protobuf:"varint,5,opt,name=active_page,json=activePage,proto3" json:"active_page,omitempty"`
	Total      int64    `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *OrderFindAllResponse) Reset() {
	*x = OrderFindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFindAllResponse) ProtoMessage() {}

func (x *OrderFindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFindAllResponse.ProtoReflect.Descriptor instead.
func (*OrderFindAllResponse) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{8}
}

func (x *OrderFindAllResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *OrderFindAllResponse) GetRows() int64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *OrderFindAllResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *OrderFindAllResponse) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *OrderFindAllResponse) GetActivePage() int64 {
	if x != nil {
		return x.ActivePage
	}
	return 0
}

func (x *OrderFindAllResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type OrderSumIncomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderSumIncomeRequest) Reset() {
	*x = OrderSumIncomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderSumIncomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderSumIncomeRequest) ProtoMessage() {}

func (x *OrderSumIncomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderSumIncomeRequest.ProtoReflect.Descriptor instead.
func (*OrderSumIncomeRequest) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{9}
}

func (x *OrderSumIncomeRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OrderSumPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *OrderSumPayload) Reset() {
	*x = OrderSumPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderSumPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderSumPayload) ProtoMessage() {}

func (x *OrderSumPayload) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderSumPayload.ProtoReflect.Descriptor instead.
func (*OrderSumPayload) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{10}
}

func (x *OrderSumPayload) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type OrderSumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEmpty bool             `protobuf:"varint,1,opt,name=is_empty,json=isEmpty,proto3" json:"is_empty,omitempty"`
	Payload *OrderSumPayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *OrderSumResponse) Reset() {
	*x = OrderSumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_order_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderSumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderSumResponse) ProtoMessage() {}

func (x *OrderSumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderSumResponse.ProtoReflect.Descriptor instead.
func (*OrderSumResponse) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{11}
}

func (x *OrderSumResponse) GetIsEmpty() bool {
	if x != nil {
		return x.IsEmpty
	}
	return false
}

func (x *OrderSumResponse) GetPayload() *OrderSumPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_pb_order_proto protoreflect.FileDescriptor

var file_pb_order_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x75, 0x79, 0x65, 0x72, 0x52, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x0a, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x75, 0x79, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0xa0, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x61, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x61, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x12,
	0x27, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x6f, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x22, 0xbf, 0x01, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xb2, 0x01, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2f, 0x0a, 0x15, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x27, 0x0a, 0x0f,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x59, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x75, 0x6d,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x32, 0x90, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x27, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x12,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12,
	0x14, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x75, 0x6d,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x75,
	0x6d, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_order_proto_rawDescOnce sync.Once
	file_pb_order_proto_rawDescData = file_pb_order_proto_rawDesc
)

func file_pb_order_proto_rawDescGZIP() []byte {
	file_pb_order_proto_rawDescOnce.Do(func() {
		file_pb_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_order_proto_rawDescData)
	})
	return file_pb_order_proto_rawDescData
}

var file_pb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pb_order_proto_goTypes = []interface{}{
	(*Order)(nil),                 // 0: Order
	(*OrderProduct)(nil),          // 1: OrderProduct
	(*OrderBuyer)(nil),            // 2: OrderBuyer
	(*OrderPayment)(nil),          // 3: OrderPayment
	(*OrderCreateRequest)(nil),    // 4: OrderCreateRequest
	(*OrderChangeStatus)(nil),     // 5: OrderChangeStatus
	(*OrderFindOneRequest)(nil),   // 6: OrderFindOneRequest
	(*OrderFindAllRequest)(nil),   // 7: OrderFindAllRequest
	(*OrderFindAllResponse)(nil),  // 8: OrderFindAllResponse
	(*OrderSumIncomeRequest)(nil), // 9: OrderSumIncomeRequest
	(*OrderSumPayload)(nil),       // 10: OrderSumPayload
	(*OrderSumResponse)(nil),      // 11: OrderSumResponse
	(*Empty)(nil),                 // 12: Empty
	(*OperationResponse)(nil),     // 13: OperationResponse
}
var file_pb_order_proto_depIdxs = []int32{
	2,  // 0: Order.buyer:type_name -> OrderBuyer
	1,  // 1: Order.product:type_name -> OrderProduct
	3,  // 2: Order.payment:type_name -> OrderPayment
	2,  // 3: OrderCreateRequest.buyer:type_name -> OrderBuyer
	1,  // 4: OrderCreateRequest.product:type_name -> OrderProduct
	3,  // 5: OrderCreateRequest.payment:type_name -> OrderPayment
	0,  // 6: OrderFindAllResponse.orders:type_name -> Order
	10, // 7: OrderSumResponse.payload:type_name -> OrderSumPayload
	4,  // 8: OrderService.Create:input_type -> OrderCreateRequest
	5,  // 9: OrderService.ChangeStatus:input_type -> OrderChangeStatus
	6,  // 10: OrderService.FindOne:input_type -> OrderFindOneRequest
	7,  // 11: OrderService.FindAll:input_type -> OrderFindAllRequest
	9,  // 12: OrderService.SumIncome:input_type -> OrderSumIncomeRequest
	12, // 13: OrderService.Create:output_type -> Empty
	13, // 14: OrderService.ChangeStatus:output_type -> OperationResponse
	0,  // 15: OrderService.FindOne:output_type -> Order
	8,  // 16: OrderService.FindAll:output_type -> OrderFindAllResponse
	11, // 17: OrderService.SumIncome:output_type -> OrderSumResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pb_order_proto_init() }
func file_pb_order_proto_init() {
	if File_pb_order_proto != nil {
		return
	}
	file_pb_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProduct); i {
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
		file_pb_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBuyer); i {
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
		file_pb_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayment); i {
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
		file_pb_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateRequest); i {
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
		file_pb_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderChangeStatus); i {
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
		file_pb_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderFindOneRequest); i {
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
		file_pb_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderFindAllRequest); i {
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
		file_pb_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderFindAllResponse); i {
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
		file_pb_order_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderSumIncomeRequest); i {
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
		file_pb_order_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderSumPayload); i {
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
		file_pb_order_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderSumResponse); i {
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
			RawDescriptor: file_pb_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_order_proto_goTypes,
		DependencyIndexes: file_pb_order_proto_depIdxs,
		MessageInfos:      file_pb_order_proto_msgTypes,
	}.Build()
	File_pb_order_proto = out.File
	file_pb_order_proto_rawDesc = nil
	file_pb_order_proto_goTypes = nil
	file_pb_order_proto_depIdxs = nil
}
