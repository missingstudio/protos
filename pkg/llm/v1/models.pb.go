// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: llm/v1/models.proto

package llmv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FinishReason int32

const (
	FinishReason_FINISH_REASON_UNSPECIFIED FinishReason = 0
	FinishReason_FINISH_REASON_LENGTH      FinishReason = 1
	FinishReason_FINISH_REASON_STOP        FinishReason = 2
	FinishReason_FINISH_REASON_ERROR       FinishReason = 3
)

// Enum value maps for FinishReason.
var (
	FinishReason_name = map[int32]string{
		0: "FINISH_REASON_UNSPECIFIED",
		1: "FINISH_REASON_LENGTH",
		2: "FINISH_REASON_STOP",
		3: "FINISH_REASON_ERROR",
	}
	FinishReason_value = map[string]int32{
		"FINISH_REASON_UNSPECIFIED": 0,
		"FINISH_REASON_LENGTH":      1,
		"FINISH_REASON_STOP":        2,
		"FINISH_REASON_ERROR":       3,
	}
)

func (x FinishReason) Enum() *FinishReason {
	p := new(FinishReason)
	*p = x
	return p
}

func (x FinishReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FinishReason) Descriptor() protoreflect.EnumDescriptor {
	return file_llm_v1_models_proto_enumTypes[0].Descriptor()
}

func (FinishReason) Type() protoreflect.EnumType {
	return &file_llm_v1_models_proto_enumTypes[0]
}

func (x FinishReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FinishReason.Descriptor instead.
func (FinishReason) EnumDescriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{0}
}

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string           `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Name        string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Config      *structpb.Struct `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{0}
}

func (x *Provider) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Provider) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Provider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Provider) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Provider) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{1}
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Model) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ProviderModels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Models []*Model `protobuf:"bytes,2,rep,name=models,proto3" json:"models,omitempty"`
}

func (x *ProviderModels) Reset() {
	*x = ProviderModels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderModels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderModels) ProtoMessage() {}

func (x *ProviderModels) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderModels.ProtoReflect.Descriptor instead.
func (*ProviderModels) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{2}
}

func (x *ProviderModels) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProviderModels) GetModels() []*Model {
	if x != nil {
		return x.Models
	}
	return nil
}

type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number of tokens in the prompt.
	PromptTokens *int32 `protobuf:"varint,1,opt,name=prompt_tokens,json=promptTokens,proto3,oneof" json:"prompt_tokens,omitempty"`
	// number of tokens in the generated completion.
	CompletionTokens *int32 `protobuf:"varint,2,opt,name=completion_tokens,json=completionTokens,proto3,oneof" json:"completion_tokens,omitempty"`
	// total number of tokens used in the request (prompt + completion).
	TotalTokens *int32 `protobuf:"varint,3,opt,name=total_tokens,json=totalTokens,proto3,oneof" json:"total_tokens,omitempty"`
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{3}
}

func (x *Usage) GetPromptTokens() int32 {
	if x != nil && x.PromptTokens != nil {
		return *x.PromptTokens
	}
	return 0
}

func (x *Usage) GetCompletionTokens() int32 {
	if x != nil && x.CompletionTokens != nil {
		return *x.CompletionTokens
	}
	return 0
}

func (x *Usage) GetTotalTokens() int32 {
	if x != nil && x.TotalTokens != nil {
		return *x.TotalTokens
	}
	return 0
}

type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arguments string `protobuf:"bytes,2,opt,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{4}
}

func (x *Function) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Function) GetArguments() string {
	if x != nil {
		return x.Arguments
	}
	return ""
}

type ToolCallMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type     string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Function *Function `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
}

func (x *ToolCallMessage) Reset() {
	*x = ToolCallMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolCallMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolCallMessage) ProtoMessage() {}

func (x *ToolCallMessage) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolCallMessage.ProtoReflect.Descriptor instead.
func (*ToolCallMessage) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{5}
}

func (x *ToolCallMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ToolCallMessage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ToolCallMessage) GetFunction() *Function {
	if x != nil {
		return x.Function
	}
	return nil
}

type ChatCompletionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// role of the message author. One of "system", "user", "assistant".
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// content of the message
	Content   string             `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Logprobs  *structpb.Struct   `protobuf:"bytes,3,opt,name=logprobs,proto3" json:"logprobs,omitempty"`
	ToolCalls []*ToolCallMessage `protobuf:"bytes,4,rep,name=tool_calls,json=toolCalls,proto3" json:"tool_calls,omitempty"`
}

func (x *ChatCompletionMessage) Reset() {
	*x = ChatCompletionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionMessage) ProtoMessage() {}

func (x *ChatCompletionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionMessage.ProtoReflect.Descriptor instead.
func (*ChatCompletionMessage) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{6}
}

func (x *ChatCompletionMessage) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ChatCompletionMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatCompletionMessage) GetLogprobs() *structpb.Struct {
	if x != nil {
		return x.Logprobs
	}
	return nil
}

func (x *ChatCompletionMessage) GetToolCalls() []*ToolCallMessage {
	if x != nil {
		return x.ToolCalls
	}
	return nil
}

type CompletionChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// index of the choice in the list of choices.
	Index uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// message generated by the model.
	Message      *ChatCompletionMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Logprobs     *structpb.Struct       `protobuf:"bytes,3,opt,name=logprobs,proto3" json:"logprobs,omitempty"`
	FinishReason string                 `protobuf:"bytes,4,opt,name=finish_reason,json=finishReason,proto3" json:"finish_reason,omitempty"`
}

func (x *CompletionChoice) Reset() {
	*x = CompletionChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionChoice) ProtoMessage() {}

func (x *CompletionChoice) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionChoice.ProtoReflect.Descriptor instead.
func (*CompletionChoice) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{7}
}

func (x *CompletionChoice) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CompletionChoice) GetMessage() *ChatCompletionMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *CompletionChoice) GetLogprobs() *structpb.Struct {
	if x != nil {
		return x.Logprobs
	}
	return nil
}

func (x *CompletionChoice) GetFinishReason() string {
	if x != nil {
		return x.FinishReason
	}
	return ""
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Role:
	//
	//	*Role_System
	//	*Role_User
	//	*Role_Assistant
	Role isRole_Role `protobuf_oneof:"role"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{8}
}

func (m *Role) GetRole() isRole_Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (x *Role) GetSystem() string {
	if x, ok := x.GetRole().(*Role_System); ok {
		return x.System
	}
	return ""
}

func (x *Role) GetUser() string {
	if x, ok := x.GetRole().(*Role_User); ok {
		return x.User
	}
	return ""
}

func (x *Role) GetAssistant() string {
	if x, ok := x.GetRole().(*Role_Assistant); ok {
		return x.Assistant
	}
	return ""
}

type isRole_Role interface {
	isRole_Role()
}

type Role_System struct {
	System string `protobuf:"bytes,1,opt,name=system,proto3,oneof"`
}

type Role_User struct {
	User string `protobuf:"bytes,2,opt,name=user,proto3,oneof"`
}

type Role_Assistant struct {
	Assistant string `protobuf:"bytes,3,opt,name=assistant,proto3,oneof"`
}

func (*Role_System) isRole_Role() {}

func (*Role_User) isRole_Role() {}

func (*Role_Assistant) isRole_Role() {}

type ResponseFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ResponseFormat) Reset() {
	*x = ResponseFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFormat) ProtoMessage() {}

func (x *ResponseFormat) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFormat.ProtoReflect.Descriptor instead.
func (*ResponseFormat) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{9}
}

func (x *ResponseFormat) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type APIKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value       string                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	MaskedValue string                 `protobuf:"bytes,4,opt,name=masked_value,json=maskedValue,proto3" json:"masked_value,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastUsedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_used_at,json=lastUsedAt,proto3" json:"last_used_at,omitempty"`
}

func (x *APIKey) Reset() {
	*x = APIKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_models_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKey) ProtoMessage() {}

func (x *APIKey) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_models_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKey.ProtoReflect.Descriptor instead.
func (*APIKey) Descriptor() ([]byte, []int) {
	return file_llm_v1_models_proto_rawDescGZIP(), []int{10}
}

func (x *APIKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *APIKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *APIKey) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *APIKey) GetMaskedValue() string {
	if x != nil {
		return x.MaskedValue
	}
	return ""
}

func (x *APIKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *APIKey) GetLastUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUsedAt
	}
	return nil
}

var File_llm_v1_models_proto protoreflect.FileDescriptor

var file_llm_v1_models_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6c, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x31, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4b, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x06,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c,
	0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x01, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x08, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x63, 0x0a, 0x0f, 0x54, 0x6f, 0x6f, 0x6c,
	0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc2, 0x01,
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x62,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x08, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x62, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x74, 0x6f,
	0x6f, 0x6c, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x43, 0x61, 0x6c, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x43, 0x61, 0x6c,
	0x6c, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x37, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f,
	0x62, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x70, 0x72, 0x6f, 0x62, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x22, 0x5e, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x14, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x22, 0x24, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x61, 0x73, 0x6b, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x78, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x19, 0x46, 0x49, 0x4e, 0x49, 0x53,
	0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x4e, 0x49,
	0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x42, 0x84, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6c, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6c, 0x6d, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x4c, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x4c, 0x6c, 0x6d, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x06, 0x4c, 0x6c, 0x6d, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x4c, 0x6c, 0x6d, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x07, 0x4c, 0x6c, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_llm_v1_models_proto_rawDescOnce sync.Once
	file_llm_v1_models_proto_rawDescData = file_llm_v1_models_proto_rawDesc
)

func file_llm_v1_models_proto_rawDescGZIP() []byte {
	file_llm_v1_models_proto_rawDescOnce.Do(func() {
		file_llm_v1_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_llm_v1_models_proto_rawDescData)
	})
	return file_llm_v1_models_proto_rawDescData
}

var file_llm_v1_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_llm_v1_models_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_llm_v1_models_proto_goTypes = []interface{}{
	(FinishReason)(0),             // 0: llm.v1.FinishReason
	(*Provider)(nil),              // 1: llm.v1.Provider
	(*Model)(nil),                 // 2: llm.v1.Model
	(*ProviderModels)(nil),        // 3: llm.v1.ProviderModels
	(*Usage)(nil),                 // 4: llm.v1.Usage
	(*Function)(nil),              // 5: llm.v1.Function
	(*ToolCallMessage)(nil),       // 6: llm.v1.ToolCallMessage
	(*ChatCompletionMessage)(nil), // 7: llm.v1.ChatCompletionMessage
	(*CompletionChoice)(nil),      // 8: llm.v1.CompletionChoice
	(*Role)(nil),                  // 9: llm.v1.Role
	(*ResponseFormat)(nil),        // 10: llm.v1.ResponseFormat
	(*APIKey)(nil),                // 11: llm.v1.APIKey
	(*structpb.Struct)(nil),       // 12: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_llm_v1_models_proto_depIdxs = []int32{
	12, // 0: llm.v1.Provider.config:type_name -> google.protobuf.Struct
	2,  // 1: llm.v1.ProviderModels.models:type_name -> llm.v1.Model
	5,  // 2: llm.v1.ToolCallMessage.function:type_name -> llm.v1.Function
	12, // 3: llm.v1.ChatCompletionMessage.logprobs:type_name -> google.protobuf.Struct
	6,  // 4: llm.v1.ChatCompletionMessage.tool_calls:type_name -> llm.v1.ToolCallMessage
	7,  // 5: llm.v1.CompletionChoice.message:type_name -> llm.v1.ChatCompletionMessage
	12, // 6: llm.v1.CompletionChoice.logprobs:type_name -> google.protobuf.Struct
	13, // 7: llm.v1.APIKey.created_at:type_name -> google.protobuf.Timestamp
	13, // 8: llm.v1.APIKey.last_used_at:type_name -> google.protobuf.Timestamp
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_llm_v1_models_proto_init() }
func file_llm_v1_models_proto_init() {
	if File_llm_v1_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_llm_v1_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
		file_llm_v1_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_llm_v1_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderModels); i {
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
		file_llm_v1_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
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
		file_llm_v1_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
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
		file_llm_v1_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolCallMessage); i {
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
		file_llm_v1_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionMessage); i {
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
		file_llm_v1_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletionChoice); i {
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
		file_llm_v1_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_llm_v1_models_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFormat); i {
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
		file_llm_v1_models_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKey); i {
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
	file_llm_v1_models_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_llm_v1_models_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Role_System)(nil),
		(*Role_User)(nil),
		(*Role_Assistant)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_llm_v1_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_llm_v1_models_proto_goTypes,
		DependencyIndexes: file_llm_v1_models_proto_depIdxs,
		EnumInfos:         file_llm_v1_models_proto_enumTypes,
		MessageInfos:      file_llm_v1_models_proto_msgTypes,
	}.Build()
	File_llm_v1_models_proto = out.File
	file_llm_v1_models_proto_rawDesc = nil
	file_llm_v1_models_proto_goTypes = nil
	file_llm_v1_models_proto_depIdxs = nil
}
