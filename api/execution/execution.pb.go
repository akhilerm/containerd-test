// Code generated by protoc-gen-gogo.
// source: execution.proto
// DO NOT EDIT!

/*
	Package execution is a generated protocol buffer package.

	It is generated from these files:
		execution.proto
		container.proto

	It has these top-level messages:
		CreateContainerRequest
		CreateContainerResponse
		DeleteContainerRequest
		DeleteContainerResponse
		ListContainersRequest
		ListContainersResponse
		CreateProcessRequest
		CreateProcessResponse
		Container
		Process
		User
		GetContainerRequest
		GetContainerResponse
		UpdateContainerRequest
		PauseContainerRequest
		ResumeContainerRequest
		StartProcessRequest
		StartProcessResponse
		GetProcessRequest
		GetProcessResponse
		SignalProcessRequest
		DeleteProcessRequest
		ListProcessesRequest
		ListProcessesResponse
*/
package execution

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateContainerRequest struct {
	ID         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BundlePath string `protobuf:"bytes,2,opt,name=bundle_path,json=bundlePath,proto3" json:"bundle_path,omitempty"`
	Stdin      string `protobuf:"bytes,3,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout     string `protobuf:"bytes,4,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr     string `protobuf:"bytes,5,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (m *CreateContainerRequest) Reset()                    { *m = CreateContainerRequest{} }
func (*CreateContainerRequest) ProtoMessage()               {}
func (*CreateContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptorExecution, []int{0} }

type CreateContainerResponse struct {
	Container *Container `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
}

func (m *CreateContainerResponse) Reset()                    { *m = CreateContainerResponse{} }
func (*CreateContainerResponse) ProtoMessage()               {}
func (*CreateContainerResponse) Descriptor() ([]byte, []int) { return fileDescriptorExecution, []int{1} }

type DeleteContainerRequest struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *DeleteContainerRequest) Reset()                    { *m = DeleteContainerRequest{} }
func (*DeleteContainerRequest) ProtoMessage()               {}
func (*DeleteContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptorExecution, []int{2} }

type DeleteContainerResponse struct {
}

func (m *DeleteContainerResponse) Reset()                    { *m = DeleteContainerResponse{} }
func (*DeleteContainerResponse) ProtoMessage()               {}
func (*DeleteContainerResponse) Descriptor() ([]byte, []int) { return fileDescriptorExecution, []int{3} }

type ListContainersRequest struct {
	Owner []string `protobuf:"bytes,1,rep,name=owner" json:"owner,omitempty"`
}

func (m *ListContainersRequest) Reset()                    { *m = ListContainersRequest{} }
func (*ListContainersRequest) ProtoMessage()               {}
func (*ListContainersRequest) Descriptor() ([]byte, []int) { return fileDescriptorExecution, []int{4} }

type ListContainersResponse struct {
	Containers []*Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
}

func (m *ListContainersResponse) Reset()                    { *m = ListContainersResponse{} }
func (*ListContainersResponse) ProtoMessage()               {}
func (*ListContainersResponse) Descriptor() ([]byte, []int) { return fileDescriptorExecution, []int{5} }

func init() {
	proto.RegisterType((*CreateContainerRequest)(nil), "containerd.v1.CreateContainerRequest")
	proto.RegisterType((*CreateContainerResponse)(nil), "containerd.v1.CreateContainerResponse")
	proto.RegisterType((*DeleteContainerRequest)(nil), "containerd.v1.DeleteContainerRequest")
	proto.RegisterType((*DeleteContainerResponse)(nil), "containerd.v1.DeleteContainerResponse")
	proto.RegisterType((*ListContainersRequest)(nil), "containerd.v1.ListContainersRequest")
	proto.RegisterType((*ListContainersResponse)(nil), "containerd.v1.ListContainersResponse")
}
func (this *CreateContainerRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&execution.CreateContainerRequest{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "BundlePath: "+fmt.Sprintf("%#v", this.BundlePath)+",\n")
	s = append(s, "Stdin: "+fmt.Sprintf("%#v", this.Stdin)+",\n")
	s = append(s, "Stdout: "+fmt.Sprintf("%#v", this.Stdout)+",\n")
	s = append(s, "Stderr: "+fmt.Sprintf("%#v", this.Stderr)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateContainerResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&execution.CreateContainerResponse{")
	if this.Container != nil {
		s = append(s, "Container: "+fmt.Sprintf("%#v", this.Container)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DeleteContainerRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&execution.DeleteContainerRequest{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DeleteContainerResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&execution.DeleteContainerResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListContainersRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&execution.ListContainersRequest{")
	s = append(s, "Owner: "+fmt.Sprintf("%#v", this.Owner)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListContainersResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&execution.ListContainersResponse{")
	if this.Containers != nil {
		s = append(s, "Containers: "+fmt.Sprintf("%#v", this.Containers)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringExecution(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringExecution(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ExecutionService service

type ExecutionServiceClient interface {
	Create(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error)
	Delete(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	List(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error)
}

type executionServiceClient struct {
	cc *grpc.ClientConn
}

func NewExecutionServiceClient(cc *grpc.ClientConn) ExecutionServiceClient {
	return &executionServiceClient{cc}
}

func (c *executionServiceClient) Create(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error) {
	out := new(CreateContainerResponse)
	err := grpc.Invoke(ctx, "/containerd.v1.ExecutionService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionServiceClient) Delete(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/containerd.v1.ExecutionService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionServiceClient) List(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error) {
	out := new(ListContainersResponse)
	err := grpc.Invoke(ctx, "/containerd.v1.ExecutionService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExecutionService service

type ExecutionServiceServer interface {
	Create(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error)
	Delete(context.Context, *DeleteContainerRequest) (*google_protobuf.Empty, error)
	List(context.Context, *ListContainersRequest) (*ListContainersResponse, error)
}

func RegisterExecutionServiceServer(s *grpc.Server, srv ExecutionServiceServer) {
	s.RegisterService(&_ExecutionService_serviceDesc, srv)
}

func _ExecutionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.ExecutionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServiceServer).Create(ctx, req.(*CreateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.ExecutionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServiceServer).Delete(ctx, req.(*DeleteContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutionService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.ExecutionService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServiceServer).List(ctx, req.(*ListContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecutionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.v1.ExecutionService",
	HandlerType: (*ExecutionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ExecutionService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ExecutionService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ExecutionService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "execution.proto",
}

func (m *CreateContainerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateContainerRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExecution(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.BundlePath) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExecution(dAtA, i, uint64(len(m.BundlePath)))
		i += copy(dAtA[i:], m.BundlePath)
	}
	if len(m.Stdin) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintExecution(dAtA, i, uint64(len(m.Stdin)))
		i += copy(dAtA[i:], m.Stdin)
	}
	if len(m.Stdout) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintExecution(dAtA, i, uint64(len(m.Stdout)))
		i += copy(dAtA[i:], m.Stdout)
	}
	if len(m.Stderr) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintExecution(dAtA, i, uint64(len(m.Stderr)))
		i += copy(dAtA[i:], m.Stderr)
	}
	return i, nil
}

func (m *CreateContainerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateContainerResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Container != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExecution(dAtA, i, uint64(m.Container.Size()))
		n1, err := m.Container.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *DeleteContainerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteContainerRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExecution(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	return i, nil
}

func (m *DeleteContainerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteContainerResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ListContainersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListContainersRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		for _, s := range m.Owner {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *ListContainersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListContainersResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for _, msg := range m.Containers {
			dAtA[i] = 0xa
			i++
			i = encodeVarintExecution(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Execution(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Execution(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintExecution(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreateContainerRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.BundlePath)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.Stdin)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.Stdout)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.Stderr)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	return n
}

func (m *CreateContainerResponse) Size() (n int) {
	var l int
	_ = l
	if m.Container != nil {
		l = m.Container.Size()
		n += 1 + l + sovExecution(uint64(l))
	}
	return n
}

func (m *DeleteContainerRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	return n
}

func (m *DeleteContainerResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ListContainersRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Owner) > 0 {
		for _, s := range m.Owner {
			l = len(s)
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	return n
}

func (m *ListContainersResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for _, e := range m.Containers {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	return n
}

func sovExecution(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExecution(x uint64) (n int) {
	return sovExecution(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CreateContainerRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateContainerRequest{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`BundlePath:` + fmt.Sprintf("%v", this.BundlePath) + `,`,
		`Stdin:` + fmt.Sprintf("%v", this.Stdin) + `,`,
		`Stdout:` + fmt.Sprintf("%v", this.Stdout) + `,`,
		`Stderr:` + fmt.Sprintf("%v", this.Stderr) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateContainerResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateContainerResponse{`,
		`Container:` + strings.Replace(fmt.Sprintf("%v", this.Container), "Container", "Container", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DeleteContainerRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DeleteContainerRequest{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DeleteContainerResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DeleteContainerResponse{`,
		`}`,
	}, "")
	return s
}
func (this *ListContainersRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListContainersRequest{`,
		`Owner:` + fmt.Sprintf("%v", this.Owner) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListContainersResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListContainersResponse{`,
		`Containers:` + strings.Replace(fmt.Sprintf("%v", this.Containers), "Container", "Container", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringExecution(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CreateContainerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateContainerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateContainerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundlePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BundlePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stderr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stderr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExecution
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateContainerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateContainerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateContainerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Container", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Container == nil {
				m.Container = &Container{}
			}
			if err := m.Container.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExecution
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeleteContainerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeleteContainerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteContainerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExecution
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeleteContainerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeleteContainerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteContainerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExecution
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListContainersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListContainersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListContainersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExecution
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListContainersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListContainersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListContainersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Containers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Containers = append(m.Containers, &Container{})
			if err := m.Containers[len(m.Containers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExecution
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExecution(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExecution
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthExecution
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExecution
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipExecution(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthExecution = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExecution   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("execution.proto", fileDescriptorExecution) }

var fileDescriptorExecution = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x7d, 0x93, 0xf7, 0x1a, 0x78, 0xf7, 0x21, 0x95, 0xa1, 0xa6, 0x31, 0x4a, 0x2a, 0xc1, 0x8a,
	0x1b, 0x53, 0xad, 0x20, 0xae, 0xfb, 0x81, 0x08, 0x2e, 0x34, 0x5d, 0xb8, 0x94, 0xb4, 0xb9, 0xa6,
	0x81, 0x9a, 0x89, 0x33, 0x93, 0xaa, 0x3b, 0xb7, 0xfe, 0x02, 0xff, 0x52, 0x97, 0x2e, 0x5d, 0x89,
	0xcd, 0x2f, 0xf0, 0x27, 0x48, 0x66, 0x92, 0x54, 0xdb, 0x62, 0xdd, 0xcd, 0xbd, 0xf7, 0xdc, 0xc3,
	0x39, 0x67, 0x2e, 0xb4, 0xf1, 0x23, 0x2e, 0x72, 0x99, 0xb0, 0xd4, 0xcf, 0x38, 0x93, 0x8c, 0x5e,
	0x5b, 0xb0, 0x54, 0x86, 0x49, 0x8a, 0x3c, 0xf2, 0xd7, 0x8f, 0x9c, 0x76, 0x53, 0xea, 0xb9, 0x73,
	0x2b, 0x66, 0x2c, 0x5e, 0xe1, 0x40, 0x55, 0xf3, 0xfc, 0xed, 0x00, 0xdf, 0x65, 0xf2, 0x53, 0x35,
	0xec, 0xc4, 0x2c, 0x66, 0xea, 0x39, 0x28, 0x5f, 0xba, 0xeb, 0x7d, 0x25, 0x60, 0x8d, 0x39, 0x86,
	0x12, 0xc7, 0x35, 0x59, 0x80, 0xef, 0x73, 0x14, 0x92, 0x5a, 0x60, 0x24, 0x91, 0x4d, 0xee, 0x90,
	0xfb, 0x97, 0x23, 0xb3, 0xf8, 0xd1, 0x33, 0x9e, 0x4f, 0x02, 0x23, 0x89, 0x68, 0x0f, 0xae, 0xe6,
	0x79, 0x1a, 0xad, 0xf0, 0x4d, 0x16, 0xca, 0xa5, 0x6d, 0x94, 0x80, 0x00, 0x74, 0xeb, 0x65, 0x28,
	0x97, 0xb4, 0x03, 0x2d, 0x21, 0xa3, 0x24, 0xb5, 0xcf, 0xd5, 0x48, 0x17, 0xd4, 0x02, 0x53, 0xc8,
	0x88, 0xe5, 0xd2, 0xbe, 0x50, 0xed, 0xaa, 0xaa, 0xfa, 0xc8, 0xb9, 0xdd, 0x6a, 0xfa, 0xc8, 0xb9,
	0xf7, 0x0a, 0xba, 0x07, 0xc2, 0x44, 0xc6, 0x52, 0x81, 0xf4, 0x09, 0x5c, 0x36, 0xd6, 0x95, 0xc0,
	0xab, 0xa1, 0xed, 0xff, 0x95, 0x8d, 0xbf, 0x5b, 0xda, 0x41, 0xbd, 0x87, 0x60, 0x4d, 0x70, 0x85,
	0xff, 0xef, 0xd5, 0xbb, 0x09, 0xdd, 0x83, 0x0d, 0x2d, 0xc2, 0x7b, 0x00, 0x37, 0x5e, 0x24, 0x42,
	0x36, 0x03, 0x51, 0x73, 0x75, 0xa0, 0xc5, 0x3e, 0x68, 0x65, 0xe7, 0xa5, 0x7d, 0x55, 0x78, 0x01,
	0x58, 0xfb, 0xf0, 0xca, 0xcd, 0x53, 0x80, 0x46, 0xa2, 0x50, 0x4b, 0xff, 0xb2, 0xf3, 0x07, 0x76,
	0xf8, 0xc5, 0x80, 0xeb, 0xd3, 0xfa, 0x46, 0x66, 0xc8, 0xd7, 0xc9, 0x02, 0xe9, 0x6b, 0x30, 0x75,
	0x6e, 0xb4, 0xbf, 0x4f, 0x72, 0xf4, 0x9f, 0x9d, 0x7b, 0xa7, 0x60, 0x95, 0xce, 0x67, 0x60, 0xea,
	0x2c, 0x0e, 0x88, 0x8f, 0x87, 0xea, 0x58, 0xbe, 0xbe, 0x47, 0xbf, 0xbe, 0x47, 0x7f, 0x5a, 0xde,
	0x23, 0x9d, 0xc1, 0x45, 0x19, 0x05, 0xbd, 0xbb, 0x47, 0x73, 0x34, 0x4e, 0xa7, 0x7f, 0x02, 0xa5,
	0xd5, 0x8d, 0x6e, 0x6f, 0xb6, 0xee, 0xd9, 0xf7, 0xad, 0x7b, 0xf6, 0x6b, 0xeb, 0x92, 0xcf, 0x85,
	0x4b, 0x36, 0x85, 0x4b, 0xbe, 0x15, 0x2e, 0xf9, 0x59, 0xb8, 0x64, 0x6e, 0x2a, 0x09, 0x8f, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x29, 0x12, 0x21, 0x53, 0x03, 0x00, 0x00,
}
