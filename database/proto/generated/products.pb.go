// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: products.proto

package generated

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Upc   string `protobuf:"bytes,2,opt,name=upc,proto3" json:"upc,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetUpc() string {
	if x != nil {
		return x.Upc
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ProductsResponse) Reset() {
	*x = ProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsResponse) ProtoMessage() {}

func (x *ProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsResponse.ProtoReflect.Descriptor instead.
func (*ProductsResponse) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{1}
}

func (x *ProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Upc string `protobuf:"bytes,1,opt,name=upc,proto3" json:"upc,omitempty"`
}

func (x *ProductQuery) Reset() {
	*x = ProductQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductQuery) ProtoMessage() {}

func (x *ProductQuery) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductQuery.ProtoReflect.Descriptor instead.
func (*ProductQuery) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{2}
}

func (x *ProductQuery) GetUpc() string {
	if x != nil {
		return x.Upc
	}
	return ""
}

type TopProductsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First int64 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
}

func (x *TopProductsQuery) Reset() {
	*x = TopProductsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopProductsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopProductsQuery) ProtoMessage() {}

func (x *TopProductsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopProductsQuery.ProtoReflect.Descriptor instead.
func (*TopProductsQuery) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{3}
}

func (x *TopProductsQuery) GetFirst() int64 {
	if x != nil {
		return x.First
	}
	return 0
}

var File_products_proto protoreflect.FileDescriptor

var file_products_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x55, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x70, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x41, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x22, 0x20, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x70, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x70, 0x63, 0x22, 0x28, 0x0a, 0x10, 0x54, 0x6f, 0x70, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x32, 0x9a, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x79, 0x55, 0x70, 0x63, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x54,
	0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a,
	0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x3b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_products_proto_rawDescOnce sync.Once
	file_products_proto_rawDescData = file_products_proto_rawDesc
)

func file_products_proto_rawDescGZIP() []byte {
	file_products_proto_rawDescOnce.Do(func() {
		file_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_products_proto_rawDescData)
	})
	return file_products_proto_rawDescData
}

var file_products_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_products_proto_goTypes = []interface{}{
	(*Product)(nil),          // 0: products.Product
	(*ProductsResponse)(nil), // 1: products.ProductsResponse
	(*ProductQuery)(nil),     // 2: products.ProductQuery
	(*TopProductsQuery)(nil), // 3: products.TopProductsQuery
}
var file_products_proto_depIdxs = []int32{
	0, // 0: products.ProductsResponse.products:type_name -> products.Product
	2, // 1: products.ProductService.FindProductByUpc:input_type -> products.ProductQuery
	3, // 2: products.ProductService.TopProducts:input_type -> products.TopProductsQuery
	0, // 3: products.ProductService.FindProductByUpc:output_type -> products.Product
	1, // 4: products.ProductService.TopProducts:output_type -> products.ProductsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_products_proto_init() }
func file_products_proto_init() {
	if File_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsResponse); i {
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
		file_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductQuery); i {
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
		file_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopProductsQuery); i {
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
			RawDescriptor: file_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_products_proto_goTypes,
		DependencyIndexes: file_products_proto_depIdxs,
		MessageInfos:      file_products_proto_msgTypes,
	}.Build()
	File_products_proto = out.File
	file_products_proto_rawDesc = nil
	file_products_proto_goTypes = nil
	file_products_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	FindProductByUpc(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*Product, error)
	TopProducts(ctx context.Context, in *TopProductsQuery, opts ...grpc.CallOption) (*ProductsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) FindProductByUpc(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/products.ProductService/FindProductByUpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) TopProducts(ctx context.Context, in *TopProductsQuery, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, "/products.ProductService/TopProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	FindProductByUpc(context.Context, *ProductQuery) (*Product, error)
	TopProducts(context.Context, *TopProductsQuery) (*ProductsResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) FindProductByUpc(context.Context, *ProductQuery) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProductByUpc not implemented")
}
func (*UnimplementedProductServiceServer) TopProducts(context.Context, *TopProductsQuery) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopProducts not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_FindProductByUpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindProductByUpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products.ProductService/FindProductByUpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindProductByUpc(ctx, req.(*ProductQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_TopProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopProductsQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).TopProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products.ProductService/TopProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).TopProducts(ctx, req.(*TopProductsQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "products.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindProductByUpc",
			Handler:    _ProductService_FindProductByUpc_Handler,
		},
		{
			MethodName: "TopProducts",
			Handler:    _ProductService_TopProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products.proto",
}
