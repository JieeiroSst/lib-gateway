// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: catalogues-service/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CatalogService_GetProduct_FullMethodName            = "/catalog.CatalogService/GetProduct"
	CatalogService_ListProducts_FullMethodName          = "/catalog.CatalogService/ListProducts"
	CatalogService_CreateProduct_FullMethodName         = "/catalog.CatalogService/CreateProduct"
	CatalogService_UpdateProduct_FullMethodName         = "/catalog.CatalogService/UpdateProduct"
	CatalogService_DeleteProduct_FullMethodName         = "/catalog.CatalogService/DeleteProduct"
	CatalogService_GetCategory_FullMethodName           = "/catalog.CatalogService/GetCategory"
	CatalogService_ListCategories_FullMethodName        = "/catalog.CatalogService/ListCategories"
	CatalogService_GetProductAttribute_FullMethodName   = "/catalog.CatalogService/GetProductAttribute"
	CatalogService_ListProductAttributes_FullMethodName = "/catalog.CatalogService/ListProductAttributes"
	CatalogService_GetProductReview_FullMethodName      = "/catalog.CatalogService/GetProductReview"
	CatalogService_ListProductReviews_FullMethodName    = "/catalog.CatalogService/ListProductReviews"
	CatalogService_CreateProductReview_FullMethodName   = "/catalog.CatalogService/CreateProductReview"
	CatalogService_GetProductInventory_FullMethodName   = "/catalog.CatalogService/GetProductInventory"
)

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
	ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*Product, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error)
	GetProductAttribute(ctx context.Context, in *GetProductAttributeRequest, opts ...grpc.CallOption) (*ProductAttribute, error)
	ListProductAttributes(ctx context.Context, in *ListProductAttributesRequest, opts ...grpc.CallOption) (*ListProductAttributesResponse, error)
	GetProductReview(ctx context.Context, in *GetProductReviewRequest, opts ...grpc.CallOption) (*ProductReview, error)
	ListProductReviews(ctx context.Context, in *ListProductReviewsRequest, opts ...grpc.CallOption) (*ListProductReviewsResponse, error)
	CreateProductReview(ctx context.Context, in *CreateProductReviewRequest, opts ...grpc.CallOption) (*ProductReview, error)
	GetProductInventory(ctx context.Context, in *GetProductInventoryRequest, opts ...grpc.CallOption) (*ProductInventory, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Product)
	err := c.cc.Invoke(ctx, CatalogService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Product)
	err := c.cc.Invoke(ctx, CatalogService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*Product, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Product)
	err := c.cc.Invoke(ctx, CatalogService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CatalogService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Category)
	err := c.cc.Invoke(ctx, CatalogService_GetCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCategoriesResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProductAttribute(ctx context.Context, in *GetProductAttributeRequest, opts ...grpc.CallOption) (*ProductAttribute, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductAttribute)
	err := c.cc.Invoke(ctx, CatalogService_GetProductAttribute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListProductAttributes(ctx context.Context, in *ListProductAttributesRequest, opts ...grpc.CallOption) (*ListProductAttributesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductAttributesResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListProductAttributes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProductReview(ctx context.Context, in *GetProductReviewRequest, opts ...grpc.CallOption) (*ProductReview, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductReview)
	err := c.cc.Invoke(ctx, CatalogService_GetProductReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListProductReviews(ctx context.Context, in *ListProductReviewsRequest, opts ...grpc.CallOption) (*ListProductReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductReviewsResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListProductReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateProductReview(ctx context.Context, in *CreateProductReviewRequest, opts ...grpc.CallOption) (*ProductReview, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductReview)
	err := c.cc.Invoke(ctx, CatalogService_CreateProductReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProductInventory(ctx context.Context, in *GetProductInventoryRequest, opts ...grpc.CallOption) (*ProductInventory, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductInventory)
	err := c.cc.Invoke(ctx, CatalogService_GetProductInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations must embed UnimplementedCatalogServiceServer
// for forward compatibility.
type CatalogServiceServer interface {
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*Product, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*Product, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error)
	GetCategory(context.Context, *GetCategoryRequest) (*Category, error)
	ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
	GetProductAttribute(context.Context, *GetProductAttributeRequest) (*ProductAttribute, error)
	ListProductAttributes(context.Context, *ListProductAttributesRequest) (*ListProductAttributesResponse, error)
	GetProductReview(context.Context, *GetProductReviewRequest) (*ProductReview, error)
	ListProductReviews(context.Context, *ListProductReviewsRequest) (*ListProductReviewsResponse, error)
	CreateProductReview(context.Context, *CreateProductReviewRequest) (*ProductReview, error)
	GetProductInventory(context.Context, *GetProductInventoryRequest) (*ProductInventory, error)
	mustEmbedUnimplementedCatalogServiceServer()
}

// UnimplementedCatalogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCatalogServiceServer struct{}

func (UnimplementedCatalogServiceServer) GetProduct(context.Context, *GetProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedCatalogServiceServer) ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedCatalogServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedCatalogServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedCatalogServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedCatalogServiceServer) ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedCatalogServiceServer) GetProductAttribute(context.Context, *GetProductAttributeRequest) (*ProductAttribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductAttribute not implemented")
}
func (UnimplementedCatalogServiceServer) ListProductAttributes(context.Context, *ListProductAttributesRequest) (*ListProductAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProductAttributes not implemented")
}
func (UnimplementedCatalogServiceServer) GetProductReview(context.Context, *GetProductReviewRequest) (*ProductReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductReview not implemented")
}
func (UnimplementedCatalogServiceServer) ListProductReviews(context.Context, *ListProductReviewsRequest) (*ListProductReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProductReviews not implemented")
}
func (UnimplementedCatalogServiceServer) CreateProductReview(context.Context, *CreateProductReviewRequest) (*ProductReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductReview not implemented")
}
func (UnimplementedCatalogServiceServer) GetProductInventory(context.Context, *GetProductInventoryRequest) (*ProductInventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductInventory not implemented")
}
func (UnimplementedCatalogServiceServer) mustEmbedUnimplementedCatalogServiceServer() {}
func (UnimplementedCatalogServiceServer) testEmbeddedByValue()                        {}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	// If the following call pancis, it indicates UnimplementedCatalogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListProducts(ctx, req.(*ListProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListCategories(ctx, req.(*ListCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProductAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProductAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetProductAttribute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProductAttribute(ctx, req.(*GetProductAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListProductAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListProductAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListProductAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListProductAttributes(ctx, req.(*ListProductAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProductReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProductReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetProductReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProductReview(ctx, req.(*GetProductReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListProductReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListProductReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListProductReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListProductReviews(ctx, req.(*ListProductReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateProductReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateProductReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateProductReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateProductReview(ctx, req.(*CreateProductReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProductInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProductInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetProductInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProductInventory(ctx, req.(*GetProductInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _CatalogService_GetProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _CatalogService_ListProducts_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _CatalogService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _CatalogService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _CatalogService_DeleteProduct_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _CatalogService_GetCategory_Handler,
		},
		{
			MethodName: "ListCategories",
			Handler:    _CatalogService_ListCategories_Handler,
		},
		{
			MethodName: "GetProductAttribute",
			Handler:    _CatalogService_GetProductAttribute_Handler,
		},
		{
			MethodName: "ListProductAttributes",
			Handler:    _CatalogService_ListProductAttributes_Handler,
		},
		{
			MethodName: "GetProductReview",
			Handler:    _CatalogService_GetProductReview_Handler,
		},
		{
			MethodName: "ListProductReviews",
			Handler:    _CatalogService_ListProductReviews_Handler,
		},
		{
			MethodName: "CreateProductReview",
			Handler:    _CatalogService_CreateProductReview_Handler,
		},
		{
			MethodName: "GetProductInventory",
			Handler:    _CatalogService_GetProductInventory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalogues-service/service.proto",
}
