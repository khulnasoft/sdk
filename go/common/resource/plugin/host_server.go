// Copyright 2016-2018, KhulnaSoft, Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"fmt"
	"sync/atomic"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/khulnasoft/sdk/go/common/diag"
	"github.com/khulnasoft/sdk/go/common/resource"
	"github.com/khulnasoft/sdk/go/common/util/rpcutil"
	khulnasoftrpc "github.com/khulnasoft/sdk/proto/go"
)

// hostServer is the server side of the host RPC machinery.
type hostServer struct {
	khulnasoftrpc.UnsafeEngineServer // opt out of forward compat

	host   Host         // the host for this RPC server.
	ctx    *Context     // the associated plugin context.
	addr   string       // the address the host is listening on.
	cancel chan bool    // a channel that can cancel the server.
	done   <-chan error // a channel that resolves when the server completes.

	// hostServer contains little bits of state that can't be saved in the language host.
	rootUrn atomic.Value // a root resource URN that has been saved via SetRootResource
}

// newHostServer creates a new host server wired up to the given host and context.
func newHostServer(host Host, ctx *Context) (*hostServer, error) {
	// New up an engine RPC server.
	engine := &hostServer{
		host:   host,
		ctx:    ctx,
		cancel: make(chan bool),
	}

	// Fire up a gRPC server and start listening for incomings.
	handle, err := rpcutil.ServeWithOptions(rpcutil.ServeOptions{
		Cancel: engine.cancel,
		Init: func(srv *grpc.Server) error {
			khulnasoftrpc.RegisterEngineServer(srv, engine)
			return nil
		},
		Options: rpcutil.OpenTracingServerInterceptorOptions(ctx.tracingSpan),
	})
	if err != nil {
		return nil, err
	}

	port := handle.Port
	done := handle.Done

	engine.addr = fmt.Sprintf("127.0.0.1:%d", port)
	engine.done = done
	engine.rootUrn.Store("")

	return engine, nil
}

// Address returns the address at which the engine's RPC server may be reached.
func (eng *hostServer) Address() string {
	return eng.addr
}

// Cancel signals that the engine should be terminated, awaits its termination, and returns any errors that result.
func (eng *hostServer) Cancel() error {
	eng.cancel <- true
	return <-eng.done
}

// Log logs a global message in the engine, including errors and warnings.
func (eng *hostServer) Log(ctx context.Context, req *khulnasoftrpc.LogRequest) (*emptypb.Empty, error) {
	var sev diag.Severity
	switch req.Severity {
	case khulnasoftrpc.LogSeverity_DEBUG:
		sev = diag.Debug
	case khulnasoftrpc.LogSeverity_INFO:
		sev = diag.Info
	case khulnasoftrpc.LogSeverity_WARNING:
		sev = diag.Warning
	case khulnasoftrpc.LogSeverity_ERROR:
		sev = diag.Error
	default:
		return nil, fmt.Errorf("Unrecognized logging severity: %v", req.Severity)
	}

	if req.Ephemeral {
		eng.host.LogStatus(sev, resource.URN(req.Urn), req.Message, req.StreamId)
	} else {
		eng.host.Log(sev, resource.URN(req.Urn), req.Message, req.StreamId)
	}
	return &emptypb.Empty{}, nil
}

// GetRootResource returns the current root resource's URN, which will serve as the parent of resources that are
// otherwise left unparented.
func (eng *hostServer) GetRootResource(ctx context.Context,
	req *khulnasoftrpc.GetRootResourceRequest,
) (*khulnasoftrpc.GetRootResourceResponse, error) {
	var response khulnasoftrpc.GetRootResourceResponse
	response.Urn = eng.rootUrn.Load().(string)
	return &response, nil
}

// SetRootResource sets the current root resource's URN. Generally only called on startup when the Stack resource is
// registered.
func (eng *hostServer) SetRootResource(ctx context.Context,
	req *khulnasoftrpc.SetRootResourceRequest,
) (*khulnasoftrpc.SetRootResourceResponse, error) {
	var response khulnasoftrpc.SetRootResourceResponse
	eng.rootUrn.Store(req.GetUrn())
	return &response, nil
}

func (eng *hostServer) StartDebugging(ctx context.Context,
	req *khulnasoftrpc.StartDebuggingRequest,
) (*emptypb.Empty, error) {
	// fire an engine event to start the debugger
	info := DebuggingInfo{
		Config: req.Config.AsMap(),
	}
	// log a status message
	eng.host.LogStatus(
		diag.Info, resource.URN(""),
		fmt.Sprintf("Waiting for debugger to attach (%v)...", req.GetMessage()), 0)

	err := eng.host.StartDebugging(info)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
