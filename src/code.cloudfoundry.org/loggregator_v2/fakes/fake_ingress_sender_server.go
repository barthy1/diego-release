// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/loggregator_v2"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type FakeIngress_SenderServer struct {
	SendAndCloseStub        func(*loggregator_v2.IngressResponse) error
	sendAndCloseMutex       sync.RWMutex
	sendAndCloseArgsForCall []struct {
		arg1 *loggregator_v2.IngressResponse
	}
	sendAndCloseReturns struct {
		result1 error
	}
	RecvStub        func() (*loggregator_v2.Envelope, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct{}
	recvReturns     struct {
		result1 *loggregator_v2.Envelope
		result2 error
	}
	SetHeaderStub        func(metadata.MD) error
	setHeaderMutex       sync.RWMutex
	setHeaderArgsForCall []struct {
		arg1 metadata.MD
	}
	setHeaderReturns struct {
		result1 error
	}
	SendHeaderStub        func(metadata.MD) error
	sendHeaderMutex       sync.RWMutex
	sendHeaderArgsForCall []struct {
		arg1 metadata.MD
	}
	sendHeaderReturns struct {
		result1 error
	}
	SetTrailerStub        func(metadata.MD)
	setTrailerMutex       sync.RWMutex
	setTrailerArgsForCall []struct {
		arg1 metadata.MD
	}
	ContextStub        func() context.Context
	contextMutex       sync.RWMutex
	contextArgsForCall []struct{}
	contextReturns     struct {
		result1 context.Context
	}
	SendMsgStub        func(m interface{}) error
	sendMsgMutex       sync.RWMutex
	sendMsgArgsForCall []struct {
		m interface{}
	}
	sendMsgReturns struct {
		result1 error
	}
	RecvMsgStub        func(m interface{}) error
	recvMsgMutex       sync.RWMutex
	recvMsgArgsForCall []struct {
		m interface{}
	}
	recvMsgReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIngress_SenderServer) SendAndClose(arg1 *loggregator_v2.IngressResponse) error {
	fake.sendAndCloseMutex.Lock()
	fake.sendAndCloseArgsForCall = append(fake.sendAndCloseArgsForCall, struct {
		arg1 *loggregator_v2.IngressResponse
	}{arg1})
	fake.recordInvocation("SendAndClose", []interface{}{arg1})
	fake.sendAndCloseMutex.Unlock()
	if fake.SendAndCloseStub != nil {
		return fake.SendAndCloseStub(arg1)
	} else {
		return fake.sendAndCloseReturns.result1
	}
}

func (fake *FakeIngress_SenderServer) SendAndCloseCallCount() int {
	fake.sendAndCloseMutex.RLock()
	defer fake.sendAndCloseMutex.RUnlock()
	return len(fake.sendAndCloseArgsForCall)
}

func (fake *FakeIngress_SenderServer) SendAndCloseArgsForCall(i int) *loggregator_v2.IngressResponse {
	fake.sendAndCloseMutex.RLock()
	defer fake.sendAndCloseMutex.RUnlock()
	return fake.sendAndCloseArgsForCall[i].arg1
}

func (fake *FakeIngress_SenderServer) SendAndCloseReturns(result1 error) {
	fake.SendAndCloseStub = nil
	fake.sendAndCloseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngress_SenderServer) Recv() (*loggregator_v2.Envelope, error) {
	fake.recvMutex.Lock()
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct{}{})
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if fake.RecvStub != nil {
		return fake.RecvStub()
	} else {
		return fake.recvReturns.result1, fake.recvReturns.result2
	}
}

func (fake *FakeIngress_SenderServer) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *FakeIngress_SenderServer) RecvReturns(result1 *loggregator_v2.Envelope, result2 error) {
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *loggregator_v2.Envelope
		result2 error
	}{result1, result2}
}

func (fake *FakeIngress_SenderServer) SetHeader(arg1 metadata.MD) error {
	fake.setHeaderMutex.Lock()
	fake.setHeaderArgsForCall = append(fake.setHeaderArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SetHeader", []interface{}{arg1})
	fake.setHeaderMutex.Unlock()
	if fake.SetHeaderStub != nil {
		return fake.SetHeaderStub(arg1)
	} else {
		return fake.setHeaderReturns.result1
	}
}

func (fake *FakeIngress_SenderServer) SetHeaderCallCount() int {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return len(fake.setHeaderArgsForCall)
}

func (fake *FakeIngress_SenderServer) SetHeaderArgsForCall(i int) metadata.MD {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return fake.setHeaderArgsForCall[i].arg1
}

func (fake *FakeIngress_SenderServer) SetHeaderReturns(result1 error) {
	fake.SetHeaderStub = nil
	fake.setHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngress_SenderServer) SendHeader(arg1 metadata.MD) error {
	fake.sendHeaderMutex.Lock()
	fake.sendHeaderArgsForCall = append(fake.sendHeaderArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SendHeader", []interface{}{arg1})
	fake.sendHeaderMutex.Unlock()
	if fake.SendHeaderStub != nil {
		return fake.SendHeaderStub(arg1)
	} else {
		return fake.sendHeaderReturns.result1
	}
}

func (fake *FakeIngress_SenderServer) SendHeaderCallCount() int {
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	return len(fake.sendHeaderArgsForCall)
}

func (fake *FakeIngress_SenderServer) SendHeaderArgsForCall(i int) metadata.MD {
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	return fake.sendHeaderArgsForCall[i].arg1
}

func (fake *FakeIngress_SenderServer) SendHeaderReturns(result1 error) {
	fake.SendHeaderStub = nil
	fake.sendHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngress_SenderServer) SetTrailer(arg1 metadata.MD) {
	fake.setTrailerMutex.Lock()
	fake.setTrailerArgsForCall = append(fake.setTrailerArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SetTrailer", []interface{}{arg1})
	fake.setTrailerMutex.Unlock()
	if fake.SetTrailerStub != nil {
		fake.SetTrailerStub(arg1)
	}
}

func (fake *FakeIngress_SenderServer) SetTrailerCallCount() int {
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	return len(fake.setTrailerArgsForCall)
}

func (fake *FakeIngress_SenderServer) SetTrailerArgsForCall(i int) metadata.MD {
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	return fake.setTrailerArgsForCall[i].arg1
}

func (fake *FakeIngress_SenderServer) Context() context.Context {
	fake.contextMutex.Lock()
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct{}{})
	fake.recordInvocation("Context", []interface{}{})
	fake.contextMutex.Unlock()
	if fake.ContextStub != nil {
		return fake.ContextStub()
	} else {
		return fake.contextReturns.result1
	}
}

func (fake *FakeIngress_SenderServer) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *FakeIngress_SenderServer) ContextReturns(result1 context.Context) {
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 context.Context
	}{result1}
}

func (fake *FakeIngress_SenderServer) SendMsg(m interface{}) error {
	fake.sendMsgMutex.Lock()
	fake.sendMsgArgsForCall = append(fake.sendMsgArgsForCall, struct {
		m interface{}
	}{m})
	fake.recordInvocation("SendMsg", []interface{}{m})
	fake.sendMsgMutex.Unlock()
	if fake.SendMsgStub != nil {
		return fake.SendMsgStub(m)
	} else {
		return fake.sendMsgReturns.result1
	}
}

func (fake *FakeIngress_SenderServer) SendMsgCallCount() int {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	return len(fake.sendMsgArgsForCall)
}

func (fake *FakeIngress_SenderServer) SendMsgArgsForCall(i int) interface{} {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	return fake.sendMsgArgsForCall[i].m
}

func (fake *FakeIngress_SenderServer) SendMsgReturns(result1 error) {
	fake.SendMsgStub = nil
	fake.sendMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngress_SenderServer) RecvMsg(m interface{}) error {
	fake.recvMsgMutex.Lock()
	fake.recvMsgArgsForCall = append(fake.recvMsgArgsForCall, struct {
		m interface{}
	}{m})
	fake.recordInvocation("RecvMsg", []interface{}{m})
	fake.recvMsgMutex.Unlock()
	if fake.RecvMsgStub != nil {
		return fake.RecvMsgStub(m)
	} else {
		return fake.recvMsgReturns.result1
	}
}

func (fake *FakeIngress_SenderServer) RecvMsgCallCount() int {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return len(fake.recvMsgArgsForCall)
}

func (fake *FakeIngress_SenderServer) RecvMsgArgsForCall(i int) interface{} {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return fake.recvMsgArgsForCall[i].m
}

func (fake *FakeIngress_SenderServer) RecvMsgReturns(result1 error) {
	fake.RecvMsgStub = nil
	fake.recvMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngress_SenderServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendAndCloseMutex.RLock()
	defer fake.sendAndCloseMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeIngress_SenderServer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ loggregator_v2.Ingress_SenderServer = new(FakeIngress_SenderServer)
