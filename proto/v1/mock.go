// Code generated by MockGen. DO NOT EDIT.
// Source: stream_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -destination=mock.go -source=stream_grpc.pb.go -package=v1 -mock_names=StreamServiceServer=MockStreamServiceServer StreamServiceServer
//

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	opencdcv1 "github.com/conduitio/conduit-connector-protocol/proto/opencdc/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockSourceServiceClient is a mock of SourceServiceClient interface.
type MockSourceServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockSourceServiceClientMockRecorder
	isgomock struct{}
}

// MockSourceServiceClientMockRecorder is the mock recorder for MockSourceServiceClient.
type MockSourceServiceClientMockRecorder struct {
	mock *MockSourceServiceClient
}

// NewMockSourceServiceClient creates a new mock instance.
func NewMockSourceServiceClient(ctrl *gomock.Controller) *MockSourceServiceClient {
	mock := &MockSourceServiceClient{ctrl: ctrl}
	mock.recorder = &MockSourceServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceServiceClient) EXPECT() *MockSourceServiceClientMockRecorder {
	return m.recorder
}

// Stream mocks base method.
func (m *MockSourceServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (SourceService_StreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stream", varargs...)
	ret0, _ := ret[0].(SourceService_StreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stream indicates an expected call of Stream.
func (mr *MockSourceServiceClientMockRecorder) Stream(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockSourceServiceClient)(nil).Stream), varargs...)
}

// MockSourceService_StreamClient is a mock of SourceService_StreamClient interface.
type MockSourceService_StreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockSourceService_StreamClientMockRecorder
	isgomock struct{}
}

// MockSourceService_StreamClientMockRecorder is the mock recorder for MockSourceService_StreamClient.
type MockSourceService_StreamClientMockRecorder struct {
	mock *MockSourceService_StreamClient
}

// NewMockSourceService_StreamClient creates a new mock instance.
func NewMockSourceService_StreamClient(ctrl *gomock.Controller) *MockSourceService_StreamClient {
	mock := &MockSourceService_StreamClient{ctrl: ctrl}
	mock.recorder = &MockSourceService_StreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceService_StreamClient) EXPECT() *MockSourceService_StreamClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockSourceService_StreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockSourceService_StreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockSourceService_StreamClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockSourceService_StreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSourceService_StreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSourceService_StreamClient)(nil).Context))
}

// Header mocks base method.
func (m *MockSourceService_StreamClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockSourceService_StreamClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockSourceService_StreamClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockSourceService_StreamClient) Recv() (*Ack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*Ack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockSourceService_StreamClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockSourceService_StreamClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockSourceService_StreamClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockSourceService_StreamClientMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSourceService_StreamClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockSourceService_StreamClient) Send(arg0 *opencdcv1.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockSourceService_StreamClientMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSourceService_StreamClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockSourceService_StreamClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockSourceService_StreamClientMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSourceService_StreamClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockSourceService_StreamClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockSourceService_StreamClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockSourceService_StreamClient)(nil).Trailer))
}

// MockSourceServiceServer is a mock of SourceServiceServer interface.
type MockSourceServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockSourceServiceServerMockRecorder
	isgomock struct{}
}

// MockSourceServiceServerMockRecorder is the mock recorder for MockSourceServiceServer.
type MockSourceServiceServerMockRecorder struct {
	mock *MockSourceServiceServer
}

// NewMockSourceServiceServer creates a new mock instance.
func NewMockSourceServiceServer(ctrl *gomock.Controller) *MockSourceServiceServer {
	mock := &MockSourceServiceServer{ctrl: ctrl}
	mock.recorder = &MockSourceServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceServiceServer) EXPECT() *MockSourceServiceServerMockRecorder {
	return m.recorder
}

// Stream mocks base method.
func (m *MockSourceServiceServer) Stream(arg0 SourceService_StreamServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockSourceServiceServerMockRecorder) Stream(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockSourceServiceServer)(nil).Stream), arg0)
}

// mustEmbedUnimplementedSourceServiceServer mocks base method.
func (m *MockSourceServiceServer) mustEmbedUnimplementedSourceServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSourceServiceServer")
}

// mustEmbedUnimplementedSourceServiceServer indicates an expected call of mustEmbedUnimplementedSourceServiceServer.
func (mr *MockSourceServiceServerMockRecorder) mustEmbedUnimplementedSourceServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSourceServiceServer", reflect.TypeOf((*MockSourceServiceServer)(nil).mustEmbedUnimplementedSourceServiceServer))
}

// MockUnsafeSourceServiceServer is a mock of UnsafeSourceServiceServer interface.
type MockUnsafeSourceServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSourceServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeSourceServiceServerMockRecorder is the mock recorder for MockUnsafeSourceServiceServer.
type MockUnsafeSourceServiceServerMockRecorder struct {
	mock *MockUnsafeSourceServiceServer
}

// NewMockUnsafeSourceServiceServer creates a new mock instance.
func NewMockUnsafeSourceServiceServer(ctrl *gomock.Controller) *MockUnsafeSourceServiceServer {
	mock := &MockUnsafeSourceServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSourceServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSourceServiceServer) EXPECT() *MockUnsafeSourceServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSourceServiceServer mocks base method.
func (m *MockUnsafeSourceServiceServer) mustEmbedUnimplementedSourceServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSourceServiceServer")
}

// mustEmbedUnimplementedSourceServiceServer indicates an expected call of mustEmbedUnimplementedSourceServiceServer.
func (mr *MockUnsafeSourceServiceServerMockRecorder) mustEmbedUnimplementedSourceServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSourceServiceServer", reflect.TypeOf((*MockUnsafeSourceServiceServer)(nil).mustEmbedUnimplementedSourceServiceServer))
}

// MockSourceService_StreamServer is a mock of SourceService_StreamServer interface.
type MockSourceService_StreamServer struct {
	ctrl     *gomock.Controller
	recorder *MockSourceService_StreamServerMockRecorder
	isgomock struct{}
}

// MockSourceService_StreamServerMockRecorder is the mock recorder for MockSourceService_StreamServer.
type MockSourceService_StreamServerMockRecorder struct {
	mock *MockSourceService_StreamServer
}

// NewMockSourceService_StreamServer creates a new mock instance.
func NewMockSourceService_StreamServer(ctrl *gomock.Controller) *MockSourceService_StreamServer {
	mock := &MockSourceService_StreamServer{ctrl: ctrl}
	mock.recorder = &MockSourceService_StreamServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceService_StreamServer) EXPECT() *MockSourceService_StreamServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockSourceService_StreamServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSourceService_StreamServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSourceService_StreamServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockSourceService_StreamServer) Recv() (*opencdcv1.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*opencdcv1.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockSourceService_StreamServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockSourceService_StreamServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockSourceService_StreamServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockSourceService_StreamServerMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSourceService_StreamServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockSourceService_StreamServer) Send(arg0 *Ack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockSourceService_StreamServerMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSourceService_StreamServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockSourceService_StreamServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockSourceService_StreamServerMockRecorder) SendHeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockSourceService_StreamServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockSourceService_StreamServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockSourceService_StreamServerMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSourceService_StreamServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockSourceService_StreamServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockSourceService_StreamServerMockRecorder) SetHeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockSourceService_StreamServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockSourceService_StreamServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockSourceService_StreamServerMockRecorder) SetTrailer(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockSourceService_StreamServer)(nil).SetTrailer), arg0)
}
