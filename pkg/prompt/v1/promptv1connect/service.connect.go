// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: prompt/v1/service.proto

package promptv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/missingstudio/protos/pkg/prompt/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PromptRegistryServiceName is the fully-qualified name of the PromptRegistryService service.
	PromptRegistryServiceName = "prompt.v1.PromptRegistryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PromptRegistryServiceListPromptsProcedure is the fully-qualified name of the
	// PromptRegistryService's ListPrompts RPC.
	PromptRegistryServiceListPromptsProcedure = "/prompt.v1.PromptRegistryService/ListPrompts"
	// PromptRegistryServiceCreatePromptProcedure is the fully-qualified name of the
	// PromptRegistryService's CreatePrompt RPC.
	PromptRegistryServiceCreatePromptProcedure = "/prompt.v1.PromptRegistryService/CreatePrompt"
	// PromptRegistryServiceGetPromptProcedure is the fully-qualified name of the
	// PromptRegistryService's GetPrompt RPC.
	PromptRegistryServiceGetPromptProcedure = "/prompt.v1.PromptRegistryService/GetPrompt"
	// PromptRegistryServiceGetPromptValueProcedure is the fully-qualified name of the
	// PromptRegistryService's GetPromptValue RPC.
	PromptRegistryServiceGetPromptValueProcedure = "/prompt.v1.PromptRegistryService/GetPromptValue"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	promptRegistryServiceServiceDescriptor              = v1.File_prompt_v1_service_proto.Services().ByName("PromptRegistryService")
	promptRegistryServiceListPromptsMethodDescriptor    = promptRegistryServiceServiceDescriptor.Methods().ByName("ListPrompts")
	promptRegistryServiceCreatePromptMethodDescriptor   = promptRegistryServiceServiceDescriptor.Methods().ByName("CreatePrompt")
	promptRegistryServiceGetPromptMethodDescriptor      = promptRegistryServiceServiceDescriptor.Methods().ByName("GetPrompt")
	promptRegistryServiceGetPromptValueMethodDescriptor = promptRegistryServiceServiceDescriptor.Methods().ByName("GetPromptValue")
)

// PromptRegistryServiceClient is a client for the prompt.v1.PromptRegistryService service.
type PromptRegistryServiceClient interface {
	// Prompt managment
	ListPrompts(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListPromptsResponse], error)
	CreatePrompt(context.Context, *connect.Request[v1.CreatePromptRequest]) (*connect.Response[v1.CreatePromptResponse], error)
	GetPrompt(context.Context, *connect.Request[v1.GetPromptRequest]) (*connect.Response[v1.GetPromptResponse], error)
	GetPromptValue(context.Context, *connect.Request[v1.GetPromptValueRequest]) (*connect.Response[v1.GetPromptValueResponse], error)
}

// NewPromptRegistryServiceClient constructs a client for the prompt.v1.PromptRegistryService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPromptRegistryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PromptRegistryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &promptRegistryServiceClient{
		listPrompts: connect.NewClient[emptypb.Empty, v1.ListPromptsResponse](
			httpClient,
			baseURL+PromptRegistryServiceListPromptsProcedure,
			connect.WithSchema(promptRegistryServiceListPromptsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createPrompt: connect.NewClient[v1.CreatePromptRequest, v1.CreatePromptResponse](
			httpClient,
			baseURL+PromptRegistryServiceCreatePromptProcedure,
			connect.WithSchema(promptRegistryServiceCreatePromptMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPrompt: connect.NewClient[v1.GetPromptRequest, v1.GetPromptResponse](
			httpClient,
			baseURL+PromptRegistryServiceGetPromptProcedure,
			connect.WithSchema(promptRegistryServiceGetPromptMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPromptValue: connect.NewClient[v1.GetPromptValueRequest, v1.GetPromptValueResponse](
			httpClient,
			baseURL+PromptRegistryServiceGetPromptValueProcedure,
			connect.WithSchema(promptRegistryServiceGetPromptValueMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// promptRegistryServiceClient implements PromptRegistryServiceClient.
type promptRegistryServiceClient struct {
	listPrompts    *connect.Client[emptypb.Empty, v1.ListPromptsResponse]
	createPrompt   *connect.Client[v1.CreatePromptRequest, v1.CreatePromptResponse]
	getPrompt      *connect.Client[v1.GetPromptRequest, v1.GetPromptResponse]
	getPromptValue *connect.Client[v1.GetPromptValueRequest, v1.GetPromptValueResponse]
}

// ListPrompts calls prompt.v1.PromptRegistryService.ListPrompts.
func (c *promptRegistryServiceClient) ListPrompts(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListPromptsResponse], error) {
	return c.listPrompts.CallUnary(ctx, req)
}

// CreatePrompt calls prompt.v1.PromptRegistryService.CreatePrompt.
func (c *promptRegistryServiceClient) CreatePrompt(ctx context.Context, req *connect.Request[v1.CreatePromptRequest]) (*connect.Response[v1.CreatePromptResponse], error) {
	return c.createPrompt.CallUnary(ctx, req)
}

// GetPrompt calls prompt.v1.PromptRegistryService.GetPrompt.
func (c *promptRegistryServiceClient) GetPrompt(ctx context.Context, req *connect.Request[v1.GetPromptRequest]) (*connect.Response[v1.GetPromptResponse], error) {
	return c.getPrompt.CallUnary(ctx, req)
}

// GetPromptValue calls prompt.v1.PromptRegistryService.GetPromptValue.
func (c *promptRegistryServiceClient) GetPromptValue(ctx context.Context, req *connect.Request[v1.GetPromptValueRequest]) (*connect.Response[v1.GetPromptValueResponse], error) {
	return c.getPromptValue.CallUnary(ctx, req)
}

// PromptRegistryServiceHandler is an implementation of the prompt.v1.PromptRegistryService service.
type PromptRegistryServiceHandler interface {
	// Prompt managment
	ListPrompts(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListPromptsResponse], error)
	CreatePrompt(context.Context, *connect.Request[v1.CreatePromptRequest]) (*connect.Response[v1.CreatePromptResponse], error)
	GetPrompt(context.Context, *connect.Request[v1.GetPromptRequest]) (*connect.Response[v1.GetPromptResponse], error)
	GetPromptValue(context.Context, *connect.Request[v1.GetPromptValueRequest]) (*connect.Response[v1.GetPromptValueResponse], error)
}

// NewPromptRegistryServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPromptRegistryServiceHandler(svc PromptRegistryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	promptRegistryServiceListPromptsHandler := connect.NewUnaryHandler(
		PromptRegistryServiceListPromptsProcedure,
		svc.ListPrompts,
		connect.WithSchema(promptRegistryServiceListPromptsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	promptRegistryServiceCreatePromptHandler := connect.NewUnaryHandler(
		PromptRegistryServiceCreatePromptProcedure,
		svc.CreatePrompt,
		connect.WithSchema(promptRegistryServiceCreatePromptMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	promptRegistryServiceGetPromptHandler := connect.NewUnaryHandler(
		PromptRegistryServiceGetPromptProcedure,
		svc.GetPrompt,
		connect.WithSchema(promptRegistryServiceGetPromptMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	promptRegistryServiceGetPromptValueHandler := connect.NewUnaryHandler(
		PromptRegistryServiceGetPromptValueProcedure,
		svc.GetPromptValue,
		connect.WithSchema(promptRegistryServiceGetPromptValueMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/prompt.v1.PromptRegistryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PromptRegistryServiceListPromptsProcedure:
			promptRegistryServiceListPromptsHandler.ServeHTTP(w, r)
		case PromptRegistryServiceCreatePromptProcedure:
			promptRegistryServiceCreatePromptHandler.ServeHTTP(w, r)
		case PromptRegistryServiceGetPromptProcedure:
			promptRegistryServiceGetPromptHandler.ServeHTTP(w, r)
		case PromptRegistryServiceGetPromptValueProcedure:
			promptRegistryServiceGetPromptValueHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPromptRegistryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPromptRegistryServiceHandler struct{}

func (UnimplementedPromptRegistryServiceHandler) ListPrompts(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListPromptsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("prompt.v1.PromptRegistryService.ListPrompts is not implemented"))
}

func (UnimplementedPromptRegistryServiceHandler) CreatePrompt(context.Context, *connect.Request[v1.CreatePromptRequest]) (*connect.Response[v1.CreatePromptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("prompt.v1.PromptRegistryService.CreatePrompt is not implemented"))
}

func (UnimplementedPromptRegistryServiceHandler) GetPrompt(context.Context, *connect.Request[v1.GetPromptRequest]) (*connect.Response[v1.GetPromptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("prompt.v1.PromptRegistryService.GetPrompt is not implemented"))
}

func (UnimplementedPromptRegistryServiceHandler) GetPromptValue(context.Context, *connect.Request[v1.GetPromptValueRequest]) (*connect.Response[v1.GetPromptValueResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("prompt.v1.PromptRegistryService.GetPromptValue is not implemented"))
}
