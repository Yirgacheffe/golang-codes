// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: proto/person/person.proto

package person

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

type Person_Gender int32

const (
	Person_MALE    Person_Gender = 0
	Person_FEMAILE Person_Gender = 1
)

// Enum value maps for Person_Gender.
var (
	Person_Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMAILE",
	}
	Person_Gender_value = map[string]int32{
		"MALE":    0,
		"FEMAILE": 1,
	}
)

func (x Person_Gender) Enum() *Person_Gender {
	p := new(Person_Gender)
	*p = x
	return p
}

func (x Person_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_person_person_proto_enumTypes[0].Descriptor()
}

func (Person_Gender) Type() protoreflect.EnumType {
	return &file_proto_person_person_proto_enumTypes[0]
}

func (x Person_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_Gender.Descriptor instead.
func (Person_Gender) EnumDescriptor() ([]byte, []int) {
	return file_proto_person_person_proto_rawDescGZIP(), []int{0, 0}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age    uint32        `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Gender Person_Gender `protobuf:"varint,3,opt,name=gender,proto3,enum=person.Person_Gender" json:"gender,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_person_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_person_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetGender() Person_Gender {
	if x != nil {
		return x.Gender
	}
	return Person_MALE
}

var File_proto_person_person_proto protoreflect.FileDescriptor

var file_proto_person_person_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x22, 0x7e, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x22, 0x1f, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x45, 0x10, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2d, 0x74, 0x6c, 0x73,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_person_person_proto_rawDescOnce sync.Once
	file_proto_person_person_proto_rawDescData = file_proto_person_person_proto_rawDesc
)

func file_proto_person_person_proto_rawDescGZIP() []byte {
	file_proto_person_person_proto_rawDescOnce.Do(func() {
		file_proto_person_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_person_person_proto_rawDescData)
	})
	return file_proto_person_person_proto_rawDescData
}

var file_proto_person_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_person_person_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_person_person_proto_goTypes = []interface{}{
	(Person_Gender)(0), // 0: person.Person.Gender
	(*Person)(nil),     // 1: person.Person
}
var file_proto_person_person_proto_depIdxs = []int32{
	0, // 0: person.Person.gender:type_name -> person.Person.Gender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_person_person_proto_init() }
func file_proto_person_person_proto_init() {
	if File_proto_person_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_person_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
			RawDescriptor: file_proto_person_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_person_person_proto_goTypes,
		DependencyIndexes: file_proto_person_person_proto_depIdxs,
		EnumInfos:         file_proto_person_person_proto_enumTypes,
		MessageInfos:      file_proto_person_person_proto_msgTypes,
	}.Build()
	File_proto_person_person_proto = out.File
	file_proto_person_person_proto_rawDesc = nil
	file_proto_person_person_proto_goTypes = nil
	file_proto_person_person_proto_depIdxs = nil
}
