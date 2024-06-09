// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: proto/ocr_service.proto

package proto

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

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ocr_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ocr_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_proto_ocr_service_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ocr_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ocr_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_proto_ocr_service_proto_rawDescGZIP(), []int{1}
}

func (x *TestResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type ExtractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Binary []byte `protobuf:"bytes,1,opt,name=binary,proto3" json:"binary,omitempty"`
}

func (x *ExtractRequest) Reset() {
	*x = ExtractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ocr_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractRequest) ProtoMessage() {}

func (x *ExtractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ocr_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractRequest.ProtoReflect.Descriptor instead.
func (*ExtractRequest) Descriptor() ([]byte, []int) {
	return file_proto_ocr_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExtractRequest) GetBinary() []byte {
	if x != nil {
		return x.Binary
	}
	return nil
}

type ExpenseField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string  `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Confidence float64 `protobuf:"fixed64,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *ExpenseField) Reset() {
	*x = ExpenseField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ocr_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpenseField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpenseField) ProtoMessage() {}

func (x *ExpenseField) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ocr_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpenseField.ProtoReflect.Descriptor instead.
func (*ExpenseField) Descriptor() ([]byte, []int) {
	return file_proto_ocr_service_proto_rawDescGZIP(), []int{3}
}

func (x *ExpenseField) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ExpenseField) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

type ExtractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FILE_PAGE            *ExpenseField `protobuf:"bytes,1,opt,name=FILE_PAGE,json=FILEPAGE,proto3" json:"FILE_PAGE,omitempty"`
	FILE_NAME            *ExpenseField `protobuf:"bytes,2,opt,name=FILE_NAME,json=FILENAME,proto3" json:"FILE_NAME,omitempty"`
	INVOICE_RECEIPT_DATE *ExpenseField `protobuf:"bytes,3,opt,name=INVOICE_RECEIPT_DATE,json=INVOICERECEIPTDATE,proto3" json:"INVOICE_RECEIPT_DATE,omitempty"`
	VENDOR_NAME          *ExpenseField `protobuf:"bytes,4,opt,name=VENDOR_NAME,json=VENDORNAME,proto3" json:"VENDOR_NAME,omitempty"`
	VENDOR_ADDRESS       *ExpenseField `protobuf:"bytes,5,opt,name=VENDOR_ADDRESS,json=VENDORADDRESS,proto3" json:"VENDOR_ADDRESS,omitempty"`
	TOTAL                *ExpenseField `protobuf:"bytes,6,opt,name=TOTAL,proto3" json:"TOTAL,omitempty"`
	SUBTOTAL             *ExpenseField `protobuf:"bytes,7,opt,name=SUBTOTAL,proto3" json:"SUBTOTAL,omitempty"`
	TAX                  *ExpenseField `protobuf:"bytes,8,opt,name=TAX,proto3" json:"TAX,omitempty"`
	VENDOR_PHONE         *ExpenseField `protobuf:"bytes,9,opt,name=VENDOR_PHONE,json=VENDORPHONE,proto3" json:"VENDOR_PHONE,omitempty"`
	STREET               *ExpenseField `protobuf:"bytes,10,opt,name=STREET,proto3" json:"STREET,omitempty"`
	GRATUITY             *ExpenseField `protobuf:"bytes,11,opt,name=GRATUITY,proto3" json:"GRATUITY,omitempty"`
	CITY                 *ExpenseField `protobuf:"bytes,12,opt,name=CITY,proto3" json:"CITY,omitempty"`
	STATE                *ExpenseField `protobuf:"bytes,13,opt,name=STATE,proto3" json:"STATE,omitempty"`
	COUNTRY              *ExpenseField `protobuf:"bytes,14,opt,name=COUNTRY,proto3" json:"COUNTRY,omitempty"`
	ZIP_CODE             *ExpenseField `protobuf:"bytes,15,opt,name=ZIP_CODE,json=ZIPCODE,proto3" json:"ZIP_CODE,omitempty"`
	CATEGORY             *ExpenseField `protobuf:"bytes,16,opt,name=CATEGORY,proto3" json:"CATEGORY,omitempty"`
}

func (x *ExtractResponse) Reset() {
	*x = ExtractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ocr_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractResponse) ProtoMessage() {}

func (x *ExtractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ocr_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractResponse.ProtoReflect.Descriptor instead.
func (*ExtractResponse) Descriptor() ([]byte, []int) {
	return file_proto_ocr_service_proto_rawDescGZIP(), []int{4}
}

func (x *ExtractResponse) GetFILE_PAGE() *ExpenseField {
	if x != nil {
		return x.FILE_PAGE
	}
	return nil
}

func (x *ExtractResponse) GetFILE_NAME() *ExpenseField {
	if x != nil {
		return x.FILE_NAME
	}
	return nil
}

func (x *ExtractResponse) GetINVOICE_RECEIPT_DATE() *ExpenseField {
	if x != nil {
		return x.INVOICE_RECEIPT_DATE
	}
	return nil
}

func (x *ExtractResponse) GetVENDOR_NAME() *ExpenseField {
	if x != nil {
		return x.VENDOR_NAME
	}
	return nil
}

func (x *ExtractResponse) GetVENDOR_ADDRESS() *ExpenseField {
	if x != nil {
		return x.VENDOR_ADDRESS
	}
	return nil
}

func (x *ExtractResponse) GetTOTAL() *ExpenseField {
	if x != nil {
		return x.TOTAL
	}
	return nil
}

func (x *ExtractResponse) GetSUBTOTAL() *ExpenseField {
	if x != nil {
		return x.SUBTOTAL
	}
	return nil
}

func (x *ExtractResponse) GetTAX() *ExpenseField {
	if x != nil {
		return x.TAX
	}
	return nil
}

func (x *ExtractResponse) GetVENDOR_PHONE() *ExpenseField {
	if x != nil {
		return x.VENDOR_PHONE
	}
	return nil
}

func (x *ExtractResponse) GetSTREET() *ExpenseField {
	if x != nil {
		return x.STREET
	}
	return nil
}

func (x *ExtractResponse) GetGRATUITY() *ExpenseField {
	if x != nil {
		return x.GRATUITY
	}
	return nil
}

func (x *ExtractResponse) GetCITY() *ExpenseField {
	if x != nil {
		return x.CITY
	}
	return nil
}

func (x *ExtractResponse) GetSTATE() *ExpenseField {
	if x != nil {
		return x.STATE
	}
	return nil
}

func (x *ExtractResponse) GetCOUNTRY() *ExpenseField {
	if x != nil {
		return x.COUNTRY
	}
	return nil
}

func (x *ExtractResponse) GetZIP_CODE() *ExpenseField {
	if x != nil {
		return x.ZIP_CODE
	}
	return nil
}

func (x *ExtractResponse) GetCATEGORY() *ExpenseField {
	if x != nil {
		return x.CATEGORY
	}
	return nil
}

var File_proto_ocr_service_proto protoreflect.FileDescriptor

var file_proto_ocr_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x63, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x2a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x42, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x8b, 0x07, 0x0a, 0x0f, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x09, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x08, 0x46, 0x49, 0x4c,
	0x45, 0x50, 0x41, 0x47, 0x45, 0x12, 0x36, 0x0a, 0x09, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x08, 0x46, 0x49, 0x4c, 0x45, 0x4e, 0x41, 0x4d, 0x45, 0x12, 0x4b, 0x0a,
	0x14, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x50, 0x54,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x12, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x52,
	0x45, 0x43, 0x45, 0x49, 0x50, 0x54, 0x44, 0x41, 0x54, 0x45, 0x12, 0x3a, 0x0a, 0x0b, 0x56, 0x45,
	0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0a, 0x56, 0x45, 0x4e, 0x44,
	0x4f, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x12, 0x40, 0x0a, 0x0e, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52,
	0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70,
	0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0d, 0x56, 0x45, 0x4e, 0x44, 0x4f,
	0x52, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x12, 0x2f, 0x0a, 0x05, 0x54, 0x4f, 0x54, 0x41,
	0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x05, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x55, 0x42,
	0x54, 0x4f, 0x54, 0x41, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x08, 0x53, 0x55, 0x42, 0x54, 0x4f, 0x54, 0x41, 0x4c,
	0x12, 0x2b, 0x0a, 0x03, 0x54, 0x41, 0x58, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x03, 0x54, 0x41, 0x58, 0x12, 0x3c, 0x0a,
	0x0c, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b,
	0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x12, 0x31, 0x0a, 0x06, 0x53,
	0x54, 0x52, 0x45, 0x45, 0x54, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x53, 0x54, 0x52, 0x45, 0x45, 0x54, 0x12, 0x35,
	0x0a, 0x08, 0x47, 0x52, 0x41, 0x54, 0x55, 0x49, 0x54, 0x59, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x08, 0x47, 0x52, 0x41,
	0x54, 0x55, 0x49, 0x54, 0x59, 0x12, 0x2d, 0x0a, 0x04, 0x43, 0x49, 0x54, 0x59, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04,
	0x43, 0x49, 0x54, 0x59, 0x12, 0x2f, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x45, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x12, 0x33, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x12, 0x34, 0x0a, 0x08, 0x5a, 0x49,
	0x50, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f,
	0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x07, 0x5a, 0x49, 0x50, 0x43, 0x4f, 0x44, 0x45,
	0x12, 0x35, 0x0a, 0x08, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x08, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x32, 0xa3, 0x01, 0x0a, 0x0a, 0x4f, 0x63, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x0f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1b, 0x2e, 0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6f, 0x63, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x10, 0x5a,
	0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ocr_service_proto_rawDescOnce sync.Once
	file_proto_ocr_service_proto_rawDescData = file_proto_ocr_service_proto_rawDesc
)

func file_proto_ocr_service_proto_rawDescGZIP() []byte {
	file_proto_ocr_service_proto_rawDescOnce.Do(func() {
		file_proto_ocr_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ocr_service_proto_rawDescData)
	})
	return file_proto_ocr_service_proto_rawDescData
}

var file_proto_ocr_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_ocr_service_proto_goTypes = []interface{}{
	(*TestRequest)(nil),     // 0: ocr_service.TestRequest
	(*TestResponse)(nil),    // 1: ocr_service.TestResponse
	(*ExtractRequest)(nil),  // 2: ocr_service.ExtractRequest
	(*ExpenseField)(nil),    // 3: ocr_service.ExpenseField
	(*ExtractResponse)(nil), // 4: ocr_service.ExtractResponse
}
var file_proto_ocr_service_proto_depIdxs = []int32{
	3,  // 0: ocr_service.ExtractResponse.FILE_PAGE:type_name -> ocr_service.ExpenseField
	3,  // 1: ocr_service.ExtractResponse.FILE_NAME:type_name -> ocr_service.ExpenseField
	3,  // 2: ocr_service.ExtractResponse.INVOICE_RECEIPT_DATE:type_name -> ocr_service.ExpenseField
	3,  // 3: ocr_service.ExtractResponse.VENDOR_NAME:type_name -> ocr_service.ExpenseField
	3,  // 4: ocr_service.ExtractResponse.VENDOR_ADDRESS:type_name -> ocr_service.ExpenseField
	3,  // 5: ocr_service.ExtractResponse.TOTAL:type_name -> ocr_service.ExpenseField
	3,  // 6: ocr_service.ExtractResponse.SUBTOTAL:type_name -> ocr_service.ExpenseField
	3,  // 7: ocr_service.ExtractResponse.TAX:type_name -> ocr_service.ExpenseField
	3,  // 8: ocr_service.ExtractResponse.VENDOR_PHONE:type_name -> ocr_service.ExpenseField
	3,  // 9: ocr_service.ExtractResponse.STREET:type_name -> ocr_service.ExpenseField
	3,  // 10: ocr_service.ExtractResponse.GRATUITY:type_name -> ocr_service.ExpenseField
	3,  // 11: ocr_service.ExtractResponse.CITY:type_name -> ocr_service.ExpenseField
	3,  // 12: ocr_service.ExtractResponse.STATE:type_name -> ocr_service.ExpenseField
	3,  // 13: ocr_service.ExtractResponse.COUNTRY:type_name -> ocr_service.ExpenseField
	3,  // 14: ocr_service.ExtractResponse.ZIP_CODE:type_name -> ocr_service.ExpenseField
	3,  // 15: ocr_service.ExtractResponse.CATEGORY:type_name -> ocr_service.ExpenseField
	0,  // 16: ocr_service.OcrService.TestConnection:input_type -> ocr_service.TestRequest
	2,  // 17: ocr_service.OcrService.ExtractFileData:input_type -> ocr_service.ExtractRequest
	1,  // 18: ocr_service.OcrService.TestConnection:output_type -> ocr_service.TestResponse
	4,  // 19: ocr_service.OcrService.ExtractFileData:output_type -> ocr_service.ExtractResponse
	18, // [18:20] is the sub-list for method output_type
	16, // [16:18] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_proto_ocr_service_proto_init() }
func file_proto_ocr_service_proto_init() {
	if File_proto_ocr_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ocr_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_proto_ocr_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
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
		file_proto_ocr_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractRequest); i {
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
		file_proto_ocr_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpenseField); i {
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
		file_proto_ocr_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractResponse); i {
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
			RawDescriptor: file_proto_ocr_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ocr_service_proto_goTypes,
		DependencyIndexes: file_proto_ocr_service_proto_depIdxs,
		MessageInfos:      file_proto_ocr_service_proto_msgTypes,
	}.Build()
	File_proto_ocr_service_proto = out.File
	file_proto_ocr_service_proto_rawDesc = nil
	file_proto_ocr_service_proto_goTypes = nil
	file_proto_ocr_service_proto_depIdxs = nil
}
