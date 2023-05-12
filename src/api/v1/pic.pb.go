// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: src/api/v1/pic.proto

package v1

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

type UploadPicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File     string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`                         // 文件
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"` // 用户
	Channel  string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`                   // 图床
	ImgKey   string `protobuf:"bytes,4,opt,name=img_key,json=imgKey,proto3" json:"img_key,omitempty"`
}

func (x *UploadPicReq) Reset() {
	*x = UploadPicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_v1_pic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadPicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadPicReq) ProtoMessage() {}

func (x *UploadPicReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_v1_pic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadPicReq.ProtoReflect.Descriptor instead.
func (*UploadPicReq) Descriptor() ([]byte, []int) {
	return file_src_api_v1_pic_proto_rawDescGZIP(), []int{0}
}

func (x *UploadPicReq) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *UploadPicReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UploadPicReq) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *UploadPicReq) GetImgKey() string {
	if x != nil {
		return x.ImgKey
	}
	return ""
}

type UploadPicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img      string `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
	ImgKey   string `protobuf:"bytes,2,opt,name=img_key,json=imgKey,proto3" json:"img_key,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *UploadPicReply) Reset() {
	*x = UploadPicReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_v1_pic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadPicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadPicReply) ProtoMessage() {}

func (x *UploadPicReply) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_v1_pic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadPicReply.ProtoReflect.Descriptor instead.
func (*UploadPicReply) Descriptor() ([]byte, []int) {
	return file_src_api_v1_pic_proto_rawDescGZIP(), []int{1}
}

func (x *UploadPicReply) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *UploadPicReply) GetImgKey() string {
	if x != nil {
		return x.ImgKey
	}
	return ""
}

func (x *UploadPicReply) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

var File_src_api_v1_pic_proto protoreflect.FileDescriptor

var file_src_api_v1_pic_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66, 0x65, 0x69, 0x73, 0x68, 0x75, 0x50, 0x69,
	0x63, 0x4c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x72, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x50, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x4b, 0x65, 0x79, 0x22, 0x58, 0x0a, 0x0e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x5b, 0x0a, 0x0a, 0x50, 0x69, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x09, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x69,
	0x63, 0x12, 0x1e, 0x2e, 0x66, 0x65, 0x69, 0x73, 0x68, 0x75, 0x50, 0x69, 0x63, 0x4c, 0x6f, 0x61,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x69, 0x63, 0x52, 0x65,
	0x71, 0x1a, 0x20, 0x2e, 0x66, 0x65, 0x69, 0x73, 0x68, 0x75, 0x50, 0x69, 0x63, 0x4c, 0x6f, 0x61,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x69, 0x63, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x58, 0x69, 0x65, 0x57, 0x65, 0x69, 0x58, 0x69, 0x65, 0x2f, 0x66, 0x65, 0x69, 0x73,
	0x68, 0x75, 0x50, 0x69, 0x63, 0x4c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_api_v1_pic_proto_rawDescOnce sync.Once
	file_src_api_v1_pic_proto_rawDescData = file_src_api_v1_pic_proto_rawDesc
)

func file_src_api_v1_pic_proto_rawDescGZIP() []byte {
	file_src_api_v1_pic_proto_rawDescOnce.Do(func() {
		file_src_api_v1_pic_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_api_v1_pic_proto_rawDescData)
	})
	return file_src_api_v1_pic_proto_rawDescData
}

var file_src_api_v1_pic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_api_v1_pic_proto_goTypes = []interface{}{
	(*UploadPicReq)(nil),   // 0: feishuPicLoad.v1.UploadPicReq
	(*UploadPicReply)(nil), // 1: feishuPicLoad.v1.UploadPicReply
}
var file_src_api_v1_pic_proto_depIdxs = []int32{
	0, // 0: feishuPicLoad.v1.PicService.UploadPic:input_type -> feishuPicLoad.v1.UploadPicReq
	1, // 1: feishuPicLoad.v1.PicService.UploadPic:output_type -> feishuPicLoad.v1.UploadPicReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_api_v1_pic_proto_init() }
func file_src_api_v1_pic_proto_init() {
	if File_src_api_v1_pic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_api_v1_pic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadPicReq); i {
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
		file_src_api_v1_pic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadPicReply); i {
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
			RawDescriptor: file_src_api_v1_pic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_api_v1_pic_proto_goTypes,
		DependencyIndexes: file_src_api_v1_pic_proto_depIdxs,
		MessageInfos:      file_src_api_v1_pic_proto_msgTypes,
	}.Build()
	File_src_api_v1_pic_proto = out.File
	file_src_api_v1_pic_proto_rawDesc = nil
	file_src_api_v1_pic_proto_goTypes = nil
	file_src_api_v1_pic_proto_depIdxs = nil
}