// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: shopify-service/service.proto

package flashsale

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
	FlashSaleService_CreateFlashSale_FullMethodName             = "/flashsale.FlashSaleService/CreateFlashSale"
	FlashSaleService_UpdateFlashSale_FullMethodName             = "/flashsale.FlashSaleService/UpdateFlashSale"
	FlashSaleService_GetFlashSale_FullMethodName                = "/flashsale.FlashSaleService/GetFlashSale"
	FlashSaleService_ListFlashSales_FullMethodName              = "/flashsale.FlashSaleService/ListFlashSales"
	FlashSaleService_CancelFlashSale_FullMethodName             = "/flashsale.FlashSaleService/CancelFlashSale"
	FlashSaleService_AddProductToSale_FullMethodName            = "/flashsale.FlashSaleService/AddProductToSale"
	FlashSaleService_RemoveProductFromSale_FullMethodName       = "/flashsale.FlashSaleService/RemoveProductFromSale"
	FlashSaleService_UpdateSaleProduct_FullMethodName           = "/flashsale.FlashSaleService/UpdateSaleProduct"
	FlashSaleService_ListSaleProducts_FullMethodName            = "/flashsale.FlashSaleService/ListSaleProducts"
	FlashSaleService_GetActiveFlashSales_FullMethodName         = "/flashsale.FlashSaleService/GetActiveFlashSales"
	FlashSaleService_ReserveProduct_FullMethodName              = "/flashsale.FlashSaleService/ReserveProduct"
	FlashSaleService_ConfirmReservation_FullMethodName          = "/flashsale.FlashSaleService/ConfirmReservation"
	FlashSaleService_CancelReservation_FullMethodName           = "/flashsale.FlashSaleService/CancelReservation"
	FlashSaleService_SubscribeToSaleUpdates_FullMethodName      = "/flashsale.FlashSaleService/SubscribeToSaleUpdates"
	FlashSaleService_SubscribeToInventoryUpdates_FullMethodName = "/flashsale.FlashSaleService/SubscribeToInventoryUpdates"
)

// FlashSaleServiceClient is the client API for FlashSaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlashSaleServiceClient interface {
	CreateFlashSale(ctx context.Context, in *CreateFlashSaleRequest, opts ...grpc.CallOption) (*FlashSaleResponse, error)
	UpdateFlashSale(ctx context.Context, in *UpdateFlashSaleRequest, opts ...grpc.CallOption) (*FlashSaleResponse, error)
	GetFlashSale(ctx context.Context, in *GetFlashSaleRequest, opts ...grpc.CallOption) (*FlashSaleResponse, error)
	ListFlashSales(ctx context.Context, in *ListFlashSalesRequest, opts ...grpc.CallOption) (*ListFlashSalesResponse, error)
	CancelFlashSale(ctx context.Context, in *CancelFlashSaleRequest, opts ...grpc.CallOption) (*FlashSaleResponse, error)
	AddProductToSale(ctx context.Context, in *AddProductToSaleRequest, opts ...grpc.CallOption) (*SaleProductResponse, error)
	RemoveProductFromSale(ctx context.Context, in *RemoveProductFromSaleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateSaleProduct(ctx context.Context, in *UpdateSaleProductRequest, opts ...grpc.CallOption) (*SaleProductResponse, error)
	ListSaleProducts(ctx context.Context, in *ListSaleProductsRequest, opts ...grpc.CallOption) (*ListSaleProductsResponse, error)
	GetActiveFlashSales(ctx context.Context, in *GetActiveFlashSalesRequest, opts ...grpc.CallOption) (*ListFlashSalesResponse, error)
	ReserveProduct(ctx context.Context, in *ReserveProductRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	ConfirmReservation(ctx context.Context, in *ConfirmReservationRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SubscribeToSaleUpdates(ctx context.Context, in *SubscribeSaleRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SaleUpdateEvent], error)
	SubscribeToInventoryUpdates(ctx context.Context, in *SubscribeInventoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[InventoryUpdateEvent], error)
}

type flashSaleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlashSaleServiceClient(cc grpc.ClientConnInterface) FlashSaleServiceClient {
	return &flashSaleServiceClient{cc}
}

func (c *flashSaleServiceClient) CreateFlashSale(ctx context.Context, in *CreateFlashSaleRequest, opts ...grpc.CallOption) (*FlashSaleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlashSaleResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_CreateFlashSale_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) UpdateFlashSale(ctx context.Context, in *UpdateFlashSaleRequest, opts ...grpc.CallOption) (*FlashSaleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlashSaleResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_UpdateFlashSale_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) GetFlashSale(ctx context.Context, in *GetFlashSaleRequest, opts ...grpc.CallOption) (*FlashSaleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlashSaleResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_GetFlashSale_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) ListFlashSales(ctx context.Context, in *ListFlashSalesRequest, opts ...grpc.CallOption) (*ListFlashSalesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFlashSalesResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_ListFlashSales_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) CancelFlashSale(ctx context.Context, in *CancelFlashSaleRequest, opts ...grpc.CallOption) (*FlashSaleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlashSaleResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_CancelFlashSale_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) AddProductToSale(ctx context.Context, in *AddProductToSaleRequest, opts ...grpc.CallOption) (*SaleProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaleProductResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_AddProductToSale_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) RemoveProductFromSale(ctx context.Context, in *RemoveProductFromSaleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FlashSaleService_RemoveProductFromSale_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) UpdateSaleProduct(ctx context.Context, in *UpdateSaleProductRequest, opts ...grpc.CallOption) (*SaleProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaleProductResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_UpdateSaleProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) ListSaleProducts(ctx context.Context, in *ListSaleProductsRequest, opts ...grpc.CallOption) (*ListSaleProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSaleProductsResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_ListSaleProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) GetActiveFlashSales(ctx context.Context, in *GetActiveFlashSalesRequest, opts ...grpc.CallOption) (*ListFlashSalesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFlashSalesResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_GetActiveFlashSales_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) ReserveProduct(ctx context.Context, in *ReserveProductRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_ReserveProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) ConfirmReservation(ctx context.Context, in *ConfirmReservationRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, FlashSaleService_ConfirmReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FlashSaleService_CancelReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashSaleServiceClient) SubscribeToSaleUpdates(ctx context.Context, in *SubscribeSaleRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SaleUpdateEvent], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlashSaleService_ServiceDesc.Streams[0], FlashSaleService_SubscribeToSaleUpdates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeSaleRequest, SaleUpdateEvent]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlashSaleService_SubscribeToSaleUpdatesClient = grpc.ServerStreamingClient[SaleUpdateEvent]

func (c *flashSaleServiceClient) SubscribeToInventoryUpdates(ctx context.Context, in *SubscribeInventoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[InventoryUpdateEvent], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FlashSaleService_ServiceDesc.Streams[1], FlashSaleService_SubscribeToInventoryUpdates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeInventoryRequest, InventoryUpdateEvent]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlashSaleService_SubscribeToInventoryUpdatesClient = grpc.ServerStreamingClient[InventoryUpdateEvent]

// FlashSaleServiceServer is the server API for FlashSaleService service.
// All implementations must embed UnimplementedFlashSaleServiceServer
// for forward compatibility.
type FlashSaleServiceServer interface {
	CreateFlashSale(context.Context, *CreateFlashSaleRequest) (*FlashSaleResponse, error)
	UpdateFlashSale(context.Context, *UpdateFlashSaleRequest) (*FlashSaleResponse, error)
	GetFlashSale(context.Context, *GetFlashSaleRequest) (*FlashSaleResponse, error)
	ListFlashSales(context.Context, *ListFlashSalesRequest) (*ListFlashSalesResponse, error)
	CancelFlashSale(context.Context, *CancelFlashSaleRequest) (*FlashSaleResponse, error)
	AddProductToSale(context.Context, *AddProductToSaleRequest) (*SaleProductResponse, error)
	RemoveProductFromSale(context.Context, *RemoveProductFromSaleRequest) (*emptypb.Empty, error)
	UpdateSaleProduct(context.Context, *UpdateSaleProductRequest) (*SaleProductResponse, error)
	ListSaleProducts(context.Context, *ListSaleProductsRequest) (*ListSaleProductsResponse, error)
	GetActiveFlashSales(context.Context, *GetActiveFlashSalesRequest) (*ListFlashSalesResponse, error)
	ReserveProduct(context.Context, *ReserveProductRequest) (*ReservationResponse, error)
	ConfirmReservation(context.Context, *ConfirmReservationRequest) (*OrderResponse, error)
	CancelReservation(context.Context, *CancelReservationRequest) (*emptypb.Empty, error)
	SubscribeToSaleUpdates(*SubscribeSaleRequest, grpc.ServerStreamingServer[SaleUpdateEvent]) error
	SubscribeToInventoryUpdates(*SubscribeInventoryRequest, grpc.ServerStreamingServer[InventoryUpdateEvent]) error
	mustEmbedUnimplementedFlashSaleServiceServer()
}

// UnimplementedFlashSaleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFlashSaleServiceServer struct{}

func (UnimplementedFlashSaleServiceServer) CreateFlashSale(context.Context, *CreateFlashSaleRequest) (*FlashSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlashSale not implemented")
}
func (UnimplementedFlashSaleServiceServer) UpdateFlashSale(context.Context, *UpdateFlashSaleRequest) (*FlashSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlashSale not implemented")
}
func (UnimplementedFlashSaleServiceServer) GetFlashSale(context.Context, *GetFlashSaleRequest) (*FlashSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlashSale not implemented")
}
func (UnimplementedFlashSaleServiceServer) ListFlashSales(context.Context, *ListFlashSalesRequest) (*ListFlashSalesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFlashSales not implemented")
}
func (UnimplementedFlashSaleServiceServer) CancelFlashSale(context.Context, *CancelFlashSaleRequest) (*FlashSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelFlashSale not implemented")
}
func (UnimplementedFlashSaleServiceServer) AddProductToSale(context.Context, *AddProductToSaleRequest) (*SaleProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductToSale not implemented")
}
func (UnimplementedFlashSaleServiceServer) RemoveProductFromSale(context.Context, *RemoveProductFromSaleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProductFromSale not implemented")
}
func (UnimplementedFlashSaleServiceServer) UpdateSaleProduct(context.Context, *UpdateSaleProductRequest) (*SaleProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSaleProduct not implemented")
}
func (UnimplementedFlashSaleServiceServer) ListSaleProducts(context.Context, *ListSaleProductsRequest) (*ListSaleProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSaleProducts not implemented")
}
func (UnimplementedFlashSaleServiceServer) GetActiveFlashSales(context.Context, *GetActiveFlashSalesRequest) (*ListFlashSalesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveFlashSales not implemented")
}
func (UnimplementedFlashSaleServiceServer) ReserveProduct(context.Context, *ReserveProductRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveProduct not implemented")
}
func (UnimplementedFlashSaleServiceServer) ConfirmReservation(context.Context, *ConfirmReservationRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReservation not implemented")
}
func (UnimplementedFlashSaleServiceServer) CancelReservation(context.Context, *CancelReservationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedFlashSaleServiceServer) SubscribeToSaleUpdates(*SubscribeSaleRequest, grpc.ServerStreamingServer[SaleUpdateEvent]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToSaleUpdates not implemented")
}
func (UnimplementedFlashSaleServiceServer) SubscribeToInventoryUpdates(*SubscribeInventoryRequest, grpc.ServerStreamingServer[InventoryUpdateEvent]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToInventoryUpdates not implemented")
}
func (UnimplementedFlashSaleServiceServer) mustEmbedUnimplementedFlashSaleServiceServer() {}
func (UnimplementedFlashSaleServiceServer) testEmbeddedByValue()                          {}

// UnsafeFlashSaleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlashSaleServiceServer will
// result in compilation errors.
type UnsafeFlashSaleServiceServer interface {
	mustEmbedUnimplementedFlashSaleServiceServer()
}

func RegisterFlashSaleServiceServer(s grpc.ServiceRegistrar, srv FlashSaleServiceServer) {
	// If the following call pancis, it indicates UnimplementedFlashSaleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FlashSaleService_ServiceDesc, srv)
}

func _FlashSaleService_CreateFlashSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlashSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).CreateFlashSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_CreateFlashSale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).CreateFlashSale(ctx, req.(*CreateFlashSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_UpdateFlashSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlashSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).UpdateFlashSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_UpdateFlashSale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).UpdateFlashSale(ctx, req.(*UpdateFlashSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_GetFlashSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlashSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).GetFlashSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_GetFlashSale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).GetFlashSale(ctx, req.(*GetFlashSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_ListFlashSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFlashSalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).ListFlashSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_ListFlashSales_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).ListFlashSales(ctx, req.(*ListFlashSalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_CancelFlashSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelFlashSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).CancelFlashSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_CancelFlashSale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).CancelFlashSale(ctx, req.(*CancelFlashSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_AddProductToSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductToSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).AddProductToSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_AddProductToSale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).AddProductToSale(ctx, req.(*AddProductToSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_RemoveProductFromSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProductFromSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).RemoveProductFromSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_RemoveProductFromSale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).RemoveProductFromSale(ctx, req.(*RemoveProductFromSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_UpdateSaleProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSaleProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).UpdateSaleProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_UpdateSaleProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).UpdateSaleProduct(ctx, req.(*UpdateSaleProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_ListSaleProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSaleProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).ListSaleProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_ListSaleProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).ListSaleProducts(ctx, req.(*ListSaleProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_GetActiveFlashSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveFlashSalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).GetActiveFlashSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_GetActiveFlashSales_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).GetActiveFlashSales(ctx, req.(*GetActiveFlashSalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_ReserveProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).ReserveProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_ReserveProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).ReserveProduct(ctx, req.(*ReserveProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_ConfirmReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).ConfirmReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_ConfirmReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).ConfirmReservation(ctx, req.(*ConfirmReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashSaleServiceServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlashSaleService_CancelReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashSaleServiceServer).CancelReservation(ctx, req.(*CancelReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashSaleService_SubscribeToSaleUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeSaleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlashSaleServiceServer).SubscribeToSaleUpdates(m, &grpc.GenericServerStream[SubscribeSaleRequest, SaleUpdateEvent]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlashSaleService_SubscribeToSaleUpdatesServer = grpc.ServerStreamingServer[SaleUpdateEvent]

func _FlashSaleService_SubscribeToInventoryUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeInventoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlashSaleServiceServer).SubscribeToInventoryUpdates(m, &grpc.GenericServerStream[SubscribeInventoryRequest, InventoryUpdateEvent]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FlashSaleService_SubscribeToInventoryUpdatesServer = grpc.ServerStreamingServer[InventoryUpdateEvent]

// FlashSaleService_ServiceDesc is the grpc.ServiceDesc for FlashSaleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlashSaleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flashsale.FlashSaleService",
	HandlerType: (*FlashSaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFlashSale",
			Handler:    _FlashSaleService_CreateFlashSale_Handler,
		},
		{
			MethodName: "UpdateFlashSale",
			Handler:    _FlashSaleService_UpdateFlashSale_Handler,
		},
		{
			MethodName: "GetFlashSale",
			Handler:    _FlashSaleService_GetFlashSale_Handler,
		},
		{
			MethodName: "ListFlashSales",
			Handler:    _FlashSaleService_ListFlashSales_Handler,
		},
		{
			MethodName: "CancelFlashSale",
			Handler:    _FlashSaleService_CancelFlashSale_Handler,
		},
		{
			MethodName: "AddProductToSale",
			Handler:    _FlashSaleService_AddProductToSale_Handler,
		},
		{
			MethodName: "RemoveProductFromSale",
			Handler:    _FlashSaleService_RemoveProductFromSale_Handler,
		},
		{
			MethodName: "UpdateSaleProduct",
			Handler:    _FlashSaleService_UpdateSaleProduct_Handler,
		},
		{
			MethodName: "ListSaleProducts",
			Handler:    _FlashSaleService_ListSaleProducts_Handler,
		},
		{
			MethodName: "GetActiveFlashSales",
			Handler:    _FlashSaleService_GetActiveFlashSales_Handler,
		},
		{
			MethodName: "ReserveProduct",
			Handler:    _FlashSaleService_ReserveProduct_Handler,
		},
		{
			MethodName: "ConfirmReservation",
			Handler:    _FlashSaleService_ConfirmReservation_Handler,
		},
		{
			MethodName: "CancelReservation",
			Handler:    _FlashSaleService_CancelReservation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToSaleUpdates",
			Handler:       _FlashSaleService_SubscribeToSaleUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToInventoryUpdates",
			Handler:       _FlashSaleService_SubscribeToInventoryUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "shopify-service/service.proto",
}
