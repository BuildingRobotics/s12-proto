// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: complete.proto

package s12_complete

import (
	fmt "fmt"
	_ "github.com/BuildingRobotics/s12-proto/protobuf/s12proto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRequest) Reset()         { *m = MsgRequest{} }
func (m *MsgRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRequest) ProtoMessage()    {}
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07686b0939113b7, []int{0}
}
func (m *MsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRequest.Unmarshal(m, b)
}
func (m *MsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRequest.Marshal(b, m, deterministic)
}
func (m *MsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequest.Merge(m, src)
}
func (m *MsgRequest) XXX_Size() int {
	return xxx_messageInfo_MsgRequest.Size(m)
}
func (m *MsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequest proto.InternalMessageInfo

type MsgResponse struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Completed            []*Complete `protobuf:"bytes,2,rep,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MsgResponse) Reset()         { *m = MsgResponse{} }
func (m *MsgResponse) String() string { return proto.CompactTextString(m) }
func (*MsgResponse) ProtoMessage()    {}
func (*MsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07686b0939113b7, []int{1}
}
func (m *MsgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgResponse.Unmarshal(m, b)
}
func (m *MsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgResponse.Marshal(b, m, deterministic)
}
func (m *MsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgResponse.Merge(m, src)
}
func (m *MsgResponse) XXX_Size() int {
	return xxx_messageInfo_MsgResponse.Size(m)
}
func (m *MsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgResponse proto.InternalMessageInfo

func (m *MsgResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgResponse) GetCompleted() []*Complete {
	if m != nil {
		return m.Completed
	}
	return nil
}

type Complete struct {
	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status         string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Description    string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Email          string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	NotEmail       string   `protobuf:"bytes,5,opt,name=not_email,json=notEmail,proto3" json:"not_email,omitempty"`
	Phone          string   `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Word           []string `protobuf:"bytes,7,rep,name=word,proto3" json:"word,omitempty"`
	Url            string   `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	SingleNumber   int32    `protobuf:"varint,9,opt,name=single_number,json=singleNumber,proto3" json:"single_number,omitempty"`
	RepeatedNumber []int64  `protobuf:"varint,10,rep,packed,name=repeated_number,json=repeatedNumber,proto3" json:"repeated_number,omitempty"`
	Lat            int32    `protobuf:"varint,11,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng            int32    `protobuf:"varint,12,opt,name=lng,proto3" json:"lng,omitempty"`
	Words          string   `protobuf:"bytes,13,opt,name=words,proto3" json:"words,omitempty"`
	Wordsn         string   `protobuf:"bytes,14,opt,name=wordsn,proto3" json:"wordsn,omitempty"`
	Intn           int32    `protobuf:"varint,15,opt,name=intn,proto3" json:"intn,omitempty"`
	Fullname       string   `protobuf:"bytes,16,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Firstname      string   `protobuf:"bytes,17,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname       string   `protobuf:"bytes,18,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Paragraph      string   `protobuf:"bytes,19,opt,name=paragraph,proto3" json:"paragraph,omitempty"`
	Paragraphs     string   `protobuf:"bytes,20,opt,name=paragraphs,proto3" json:"paragraphs,omitempty"`
	Paragraphsn    string   `protobuf:"bytes,21,opt,name=paragraphsn,proto3" json:"paragraphsn,omitempty"`
	Uuid           string   `protobuf:"bytes,22,opt,name=uuid,proto3" json:"uuid,omitempty"`
	EmailAddress   string   `protobuf:"bytes,23,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	PhoneNumber    string   `protobuf:"bytes,24,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Company        string   `protobuf:"bytes,25,opt,name=company,proto3" json:"company,omitempty"`
	Brand          string   `protobuf:"bytes,26,opt,name=brand,proto3" json:"brand,omitempty"`
	Product        string   `protobuf:"bytes,27,opt,name=product,proto3" json:"product,omitempty"`
	Color          string   `protobuf:"bytes,28,opt,name=color,proto3" json:"color,omitempty"`
	Hexcolor       string   `protobuf:"bytes,29,opt,name=hexcolor,proto3" json:"hexcolor,omitempty"`
	Latitude       float64  `protobuf:"fixed64,30,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude      float64  `protobuf:"fixed64,31,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Floatn         float32  `protobuf:"fixed32,32,opt,name=floatn,proto3" json:"floatn,omitempty"`
	Boolean        bool     `protobuf:"varint,33,opt,name=boolean,proto3" json:"boolean,omitempty"`
	// Types that are valid to be assigned to SomeOneof:
	//	*Complete_CouldBe
	//	*Complete_OrCouldBe
	SomeOneof isComplete_SomeOneof `protobuf_oneof:"some_oneof"`
	// Types that are valid to be assigned to AnotherOneof:
	//	*Complete_This
	//	*Complete_That
	AnotherOneof         isComplete_AnotherOneof `protobuf_oneof:"another_oneof"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Complete) Reset()         { *m = Complete{} }
func (m *Complete) String() string { return proto.CompactTextString(m) }
func (*Complete) ProtoMessage()    {}
func (*Complete) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07686b0939113b7, []int{2}
}
func (m *Complete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Complete.Unmarshal(m, b)
}
func (m *Complete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Complete.Marshal(b, m, deterministic)
}
func (m *Complete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Complete.Merge(m, src)
}
func (m *Complete) XXX_Size() int {
	return xxx_messageInfo_Complete.Size(m)
}
func (m *Complete) XXX_DiscardUnknown() {
	xxx_messageInfo_Complete.DiscardUnknown(m)
}

var xxx_messageInfo_Complete proto.InternalMessageInfo

type isComplete_SomeOneof interface {
	isComplete_SomeOneof()
}
type isComplete_AnotherOneof interface {
	isComplete_AnotherOneof()
}

type Complete_CouldBe struct {
	CouldBe string `protobuf:"bytes,34,opt,name=could_be,json=couldBe,proto3,oneof" json:"could_be,omitempty"`
}
type Complete_OrCouldBe struct {
	OrCouldBe *SubMessage `protobuf:"bytes,35,opt,name=or_could_be,json=orCouldBe,proto3,oneof" json:"or_could_be,omitempty"`
}
type Complete_This struct {
	This string `protobuf:"bytes,36,opt,name=this,proto3,oneof" json:"this,omitempty"`
}
type Complete_That struct {
	That string `protobuf:"bytes,37,opt,name=that,proto3,oneof" json:"that,omitempty"`
}

func (*Complete_CouldBe) isComplete_SomeOneof()   {}
func (*Complete_OrCouldBe) isComplete_SomeOneof() {}
func (*Complete_This) isComplete_AnotherOneof()   {}
func (*Complete_That) isComplete_AnotherOneof()   {}

func (m *Complete) GetSomeOneof() isComplete_SomeOneof {
	if m != nil {
		return m.SomeOneof
	}
	return nil
}
func (m *Complete) GetAnotherOneof() isComplete_AnotherOneof {
	if m != nil {
		return m.AnotherOneof
	}
	return nil
}

func (m *Complete) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Complete) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Complete) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Complete) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Complete) GetNotEmail() string {
	if m != nil {
		return m.NotEmail
	}
	return ""
}

func (m *Complete) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Complete) GetWord() []string {
	if m != nil {
		return m.Word
	}
	return nil
}

func (m *Complete) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Complete) GetSingleNumber() int32 {
	if m != nil {
		return m.SingleNumber
	}
	return 0
}

func (m *Complete) GetRepeatedNumber() []int64 {
	if m != nil {
		return m.RepeatedNumber
	}
	return nil
}

func (m *Complete) GetLat() int32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Complete) GetLng() int32 {
	if m != nil {
		return m.Lng
	}
	return 0
}

func (m *Complete) GetWords() string {
	if m != nil {
		return m.Words
	}
	return ""
}

func (m *Complete) GetWordsn() string {
	if m != nil {
		return m.Wordsn
	}
	return ""
}

func (m *Complete) GetIntn() int32 {
	if m != nil {
		return m.Intn
	}
	return 0
}

func (m *Complete) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *Complete) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Complete) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *Complete) GetParagraph() string {
	if m != nil {
		return m.Paragraph
	}
	return ""
}

func (m *Complete) GetParagraphs() string {
	if m != nil {
		return m.Paragraphs
	}
	return ""
}

func (m *Complete) GetParagraphsn() string {
	if m != nil {
		return m.Paragraphsn
	}
	return ""
}

func (m *Complete) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Complete) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *Complete) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Complete) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Complete) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Complete) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Complete) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Complete) GetHexcolor() string {
	if m != nil {
		return m.Hexcolor
	}
	return ""
}

func (m *Complete) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Complete) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Complete) GetFloatn() float32 {
	if m != nil {
		return m.Floatn
	}
	return 0
}

func (m *Complete) GetBoolean() bool {
	if m != nil {
		return m.Boolean
	}
	return false
}

func (m *Complete) GetCouldBe() string {
	if x, ok := m.GetSomeOneof().(*Complete_CouldBe); ok {
		return x.CouldBe
	}
	return ""
}

func (m *Complete) GetOrCouldBe() *SubMessage {
	if x, ok := m.GetSomeOneof().(*Complete_OrCouldBe); ok {
		return x.OrCouldBe
	}
	return nil
}

func (m *Complete) GetThis() string {
	if x, ok := m.GetAnotherOneof().(*Complete_This); ok {
		return x.This
	}
	return ""
}

func (m *Complete) GetThat() string {
	if x, ok := m.GetAnotherOneof().(*Complete_That); ok {
		return x.That
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Complete) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Complete_CouldBe)(nil),
		(*Complete_OrCouldBe)(nil),
		(*Complete_This)(nil),
		(*Complete_That)(nil),
	}
}

type SubMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubMessage) Reset()         { *m = SubMessage{} }
func (m *SubMessage) String() string { return proto.CompactTextString(m) }
func (*SubMessage) ProtoMessage()    {}
func (*SubMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e07686b0939113b7, []int{3}
}
func (m *SubMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubMessage.Unmarshal(m, b)
}
func (m *SubMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubMessage.Marshal(b, m, deterministic)
}
func (m *SubMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubMessage.Merge(m, src)
}
func (m *SubMessage) XXX_Size() int {
	return xxx_messageInfo_SubMessage.Size(m)
}
func (m *SubMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubMessage proto.InternalMessageInfo

func (m *SubMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRequest)(nil), "s12.complete.MsgRequest")
	proto.RegisterType((*MsgResponse)(nil), "s12.complete.MsgResponse")
	proto.RegisterType((*Complete)(nil), "s12.complete.Complete")
	proto.RegisterType((*SubMessage)(nil), "s12.complete.SubMessage")
}

func init() { proto.RegisterFile("complete.proto", fileDescriptor_e07686b0939113b7) }

var fileDescriptor_e07686b0939113b7 = []byte{
	// 813 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x4f, 0x73, 0x1b, 0x35,
	0x18, 0xc6, 0x23, 0x6f, 0xfc, 0xef, 0x5d, 0xe7, 0x4f, 0x45, 0x68, 0x54, 0xc7, 0x6d, 0x55, 0xa7,
	0xc0, 0x32, 0xcc, 0x84, 0xd6, 0xe5, 0xc0, 0x70, 0x02, 0x67, 0x18, 0x72, 0x09, 0xc3, 0x28, 0x97,
	0xde, 0x8c, 0xec, 0x55, 0xd6, 0x3b, 0xac, 0xa5, 0x45, 0xd2, 0x42, 0x7b, 0xe3, 0xc8, 0x91, 0x23,
	0x67, 0x4e, 0xdc, 0xc2, 0x67, 0xca, 0xf0, 0x41, 0x18, 0x69, 0x77, 0x2d, 0x53, 0xb8, 0xe9, 0xfd,
	0x3d, 0x8f, 0x1e, 0xc9, 0xab, 0xd7, 0x2f, 0x1c, 0xae, 0xd4, 0xa6, 0x2c, 0x84, 0x15, 0x17, 0xa5,
	0x56, 0x56, 0xe1, 0x91, 0x79, 0x39, 0xbb, 0x68, 0xd9, 0xf8, 0xd4, 0xbc, 0x9c, 0x79, 0xfe, 0x69,
	0xa6, 0xcb, 0xd5, 0x46, 0xad, 0x7e, 0xa8, 0x6d, 0xd3, 0x11, 0xc0, 0xb5, 0xc9, 0x98, 0xf8, 0xb1,
	0x12, 0xc6, 0x4e, 0x6f, 0x20, 0xf6, 0x95, 0x29, 0x95, 0x34, 0x02, 0x1f, 0x42, 0x27, 0x4f, 0x09,
	0xa2, 0x28, 0x19, 0xb2, 0x4e, 0x9e, 0xe2, 0xcf, 0x60, 0xd8, 0x26, 0xa6, 0xa4, 0x43, 0xa3, 0x24,
	0x9e, 0x3d, 0xbc, 0xd8, 0x3d, 0xe7, 0xe2, 0xb2, 0x59, 0xb0, 0x60, 0x9c, 0xfe, 0x3d, 0x84, 0x41,
	0xcb, 0xff, 0x13, 0xf9, 0x0a, 0x7a, 0xc6, 0x72, 0x5b, 0x19, 0xd2, 0x71, 0x6c, 0x7e, 0x76, 0x7f,
	0x47, 0x4f, 0x21, 0xce, 0x25, 0x2d, 0xb5, 0xca, 0xb4, 0x30, 0x06, 0x06, 0x6d, 0x14, 0x6b, 0xac,
	0x98, 0x42, 0x9c, 0x0a, 0xb3, 0xd2, 0x79, 0x69, 0x73, 0x25, 0x49, 0xe4, 0xd3, 0x76, 0x11, 0x3e,
	0x81, 0xae, 0xd8, 0xf0, 0xbc, 0x20, 0xfb, 0x5e, 0xab, 0x0b, 0x7c, 0x0e, 0x43, 0xa9, 0xec, 0xa2,
	0x56, 0xba, 0xfe, 0xbc, 0xde, 0xfd, 0x1d, 0xed, 0x1c, 0x23, 0x36, 0x90, 0xca, 0x7e, 0xed, 0x4d,
	0x27, 0xd0, 0x2d, 0xd7, 0x4a, 0x0a, 0xd2, 0xab, 0xb7, 0xfa, 0x02, 0x63, 0xd8, 0xff, 0x59, 0xe9,
	0x94, 0xf4, 0x69, 0x94, 0x0c, 0x99, 0x5f, 0xe3, 0x63, 0x88, 0x2a, 0x5d, 0x90, 0x81, 0xf7, 0xb9,
	0x25, 0x3e, 0x87, 0x03, 0x93, 0xcb, 0xac, 0x10, 0x0b, 0x59, 0x6d, 0x96, 0x42, 0x93, 0x21, 0x45,
	0x49, 0x97, 0x8d, 0x6a, 0xf8, 0xad, 0x67, 0xf8, 0x05, 0x1c, 0x69, 0x51, 0x0a, 0x6e, 0x45, 0xda,
	0xda, 0x80, 0x46, 0x49, 0x34, 0xef, 0xdf, 0xdf, 0xd1, 0xe8, 0x2f, 0x14, 0xb1, 0xc3, 0x56, 0x6f,
	0x76, 0x1c, 0x43, 0x54, 0x70, 0x4b, 0x62, 0x1f, 0xe6, 0x96, 0x9e, 0xc8, 0x8c, 0x8c, 0x1a, 0x22,
	0x33, 0x3c, 0x81, 0xae, 0xbb, 0x94, 0x21, 0x07, 0xe1, 0x77, 0x11, 0xc4, 0x6a, 0x88, 0x9f, 0x40,
	0xcf, 0x2f, 0x24, 0x39, 0x0c, 0x32, 0x05, 0xd6, 0x50, 0x3c, 0x86, 0xfd, 0x5c, 0x5a, 0x49, 0x8e,
	0x5c, 0x60, 0xad, 0x26, 0x5d, 0xe6, 0x19, 0x9e, 0xc2, 0xe0, 0xb6, 0x2a, 0x0a, 0xc9, 0x37, 0x82,
	0x1c, 0x87, 0xdd, 0x2f, 0x10, 0xdb, 0x72, 0xfc, 0x1c, 0x86, 0xb7, 0xb9, 0x36, 0xd6, 0x9b, 0x1e,
	0x04, 0xd3, 0xe7, 0x88, 0x05, 0xc1, 0x25, 0x15, 0xbc, 0x31, 0xe1, 0x60, 0xfa, 0x12, 0xb1, 0x2d,
	0x77, 0x49, 0x25, 0xd7, 0x3c, 0xd3, 0xbc, 0x5c, 0x93, 0xf7, 0x82, 0xe9, 0x0a, 0xb1, 0x20, 0xe0,
	0x0f, 0x01, 0xb6, 0x85, 0x21, 0x27, 0xc1, 0xf6, 0x1d, 0x62, 0x3b, 0x0a, 0x4e, 0x20, 0x0e, 0x95,
	0x24, 0xef, 0x07, 0xe3, 0xeb, 0x0e, 0xdb, 0x95, 0xdc, 0x17, 0xa8, 0xaa, 0x3c, 0x25, 0x0f, 0x83,
	0xe5, 0x7b, 0xc4, 0x3c, 0xc3, 0x9f, 0xc0, 0x81, 0xef, 0x99, 0x05, 0x4f, 0x53, 0xd7, 0x92, 0xe4,
	0x34, 0x98, 0xd6, 0x88, 0x8d, 0xbc, 0xf8, 0x55, 0xad, 0xe1, 0x8f, 0x61, 0xe4, 0x5b, 0xa6, 0x7d,
	0x5b, 0x12, 0xbc, 0x25, 0x62, 0xb1, 0xd7, 0x9a, 0x77, 0xa5, 0xd0, 0x77, 0xbd, 0xcd, 0xe5, 0x5b,
	0xf2, 0x28, 0xb8, 0xde, 0x20, 0xd6, 0x62, 0xfc, 0x18, 0xba, 0x4b, 0xcd, 0x65, 0x4a, 0xc6, 0x5e,
	0xf7, 0x1d, 0xf2, 0x0b, 0x42, 0xac, 0xa6, 0xf8, 0x19, 0xf4, 0x4b, 0xad, 0xd2, 0x6a, 0x65, 0xc9,
	0x59, 0x30, 0xfc, 0x8a, 0x10, 0x6b, 0xb9, 0x4b, 0x58, 0xa9, 0x42, 0x69, 0x32, 0x09, 0x86, 0xdf,
	0x5c, 0x82, 0xa7, 0xf8, 0x23, 0x18, 0xac, 0xc5, 0x9b, 0xda, 0xf1, 0xd8, 0x3b, 0xe2, 0xfb, 0x3b,
	0xda, 0xff, 0x1d, 0xa1, 0x3f, 0x10, 0x3a, 0x67, 0x5b, 0x11, 0x8f, 0xdd, 0xdb, 0xd9, 0xdc, 0x56,
	0xa9, 0x20, 0x4f, 0x28, 0x4a, 0xfc, 0x9b, 0xd5, 0x35, 0x9e, 0xc0, 0xb0, 0x50, 0x32, 0xab, 0xc5,
	0xa7, 0x5e, 0x0c, 0x00, 0x3f, 0x85, 0xde, 0x6d, 0xa1, 0xb8, 0x95, 0x84, 0x52, 0x94, 0x74, 0xea,
	0x2b, 0xfc, 0x89, 0x22, 0xd6, 0x60, 0x4c, 0xa0, 0xbf, 0x54, 0xaa, 0x10, 0x5c, 0x92, 0x67, 0x14,
	0x25, 0x03, 0xd6, 0x96, 0xf8, 0xcc, 0xfd, 0xf9, 0xab, 0x22, 0x5d, 0x2c, 0x05, 0x99, 0xba, 0xdb,
	0x5d, 0xed, 0xb9, 0x6f, 0x53, 0x15, 0xe9, 0x5c, 0xe0, 0x2f, 0x20, 0x56, 0x7a, 0xb1, 0xd5, 0xcf,
	0x29, 0x4a, 0xe2, 0x19, 0xf9, 0xf7, 0x3c, 0xba, 0xa9, 0x96, 0xd7, 0xc2, 0x18, 0x9e, 0x89, 0xab,
	0x3d, 0x36, 0x54, 0xfa, 0xb2, 0xd9, 0x7b, 0x02, 0xfb, 0x76, 0x9d, 0x1b, 0xf2, 0xdc, 0x87, 0x22,
	0xe6, 0x2b, 0x3c, 0x71, 0x94, 0x5b, 0xf2, 0xc1, 0x6e, 0xdb, 0xd5, 0x2a, 0xb7, 0xf3, 0x11, 0x80,
	0x51, 0x1b, 0xb1, 0x50, 0x52, 0xa8, 0xdb, 0xf9, 0x11, 0x1c, 0x70, 0xa9, 0xec, 0x5a, 0xe8, 0x1a,
	0x4c, 0x27, 0x00, 0xe1, 0xb4, 0x77, 0xe7, 0xdc, 0xec, 0x35, 0x3c, 0xb8, 0x7e, 0xdb, 0x4e, 0xc1,
	0x1b, 0xa1, 0x7f, 0xca, 0x57, 0x02, 0x5f, 0x02, 0x7c, 0x23, 0x6c, 0xbb, 0xe5, 0x9d, 0xab, 0x87,
	0xb1, 0x3c, 0x7e, 0xf4, 0x3f, 0x4a, 0x3d, 0xa2, 0xa7, 0x7b, 0xcb, 0x9e, 0x1f, 0xe4, 0xaf, 0xfe,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xcf, 0xaa, 0x51, 0x01, 0x06, 0x00, 0x00,
}
