// Copyright 2016-2020, KhulnaSoft, Ltd.
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
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/blang/semver"
	"github.com/khulnasoft/sdk/go/common/resource"
	"github.com/khulnasoft/sdk/go/common/resource/config"
	"github.com/khulnasoft/sdk/go/common/tokens"
	"github.com/khulnasoft/sdk/go/common/util/rpcutil/rpcerror"
	khulnasoftrpc "github.com/khulnasoft/sdk/proto/go"
)

type providerServer struct {
	khulnasoftrpc.UnsafeResourceProviderServer // opt out of forward compat

	provider      Provider
	keepSecrets   bool
	keepResources bool
}

func NewProviderServer(provider Provider) khulnasoftrpc.ResourceProviderServer {
	return &providerServer{provider: provider}
}

func (p *providerServer) unmarshalOptions(label string, keepOutputValues bool) MarshalOptions {
	return MarshalOptions{
		Label:            label,
		KeepUnknowns:     true,
		KeepSecrets:      true,
		KeepResources:    true,
		KeepOutputValues: keepOutputValues,
	}
}

func (p *providerServer) marshalOptions(label string) MarshalOptions {
	return MarshalOptions{
		Label:         label,
		KeepUnknowns:  true,
		KeepSecrets:   p.keepSecrets,
		KeepResources: p.keepResources,
	}
}

func (p *providerServer) checkNYI(method string, err error) error {
	if err == ErrNotYetImplemented {
		return status.Error(codes.Unimplemented, fmt.Sprintf("%v is not yet implemented", method))
	}
	return err
}

func (p *providerServer) marshalDiff(diff DiffResult) (*khulnasoftrpc.DiffResponse, error) {
	var changes khulnasoftrpc.DiffResponse_DiffChanges
	switch diff.Changes {
	case DiffNone:
		changes = khulnasoftrpc.DiffResponse_DIFF_NONE
	case DiffSome:
		changes = khulnasoftrpc.DiffResponse_DIFF_SOME
	case DiffUnknown:
		changes = khulnasoftrpc.DiffResponse_DIFF_UNKNOWN
	}

	// Infer the result from the detailed diff.
	var diffs, replaces []string
	var detailedDiff map[string]*khulnasoftrpc.PropertyDiff
	if len(diff.DetailedDiff) == 0 {
		diffs = make([]string, len(diff.ChangedKeys))
		for i, k := range diff.ChangedKeys {
			diffs[i] = string(k)
		}
		replaces = make([]string, len(diff.ReplaceKeys))
		for i, k := range diff.ReplaceKeys {
			replaces[i] = string(k)
		}
	} else {
		changes = khulnasoftrpc.DiffResponse_DIFF_SOME

		detailedDiff = make(map[string]*khulnasoftrpc.PropertyDiff)
		for path, diff := range diff.DetailedDiff {
			diffs = append(diffs, path)

			var kind khulnasoftrpc.PropertyDiff_Kind
			switch diff.Kind {
			case DiffAdd:
				kind = khulnasoftrpc.PropertyDiff_ADD
			case DiffAddReplace:
				kind, replaces = khulnasoftrpc.PropertyDiff_ADD_REPLACE, append(replaces, path)
			case DiffDelete:
				kind = khulnasoftrpc.PropertyDiff_DELETE
			case DiffDeleteReplace:
				kind, replaces = khulnasoftrpc.PropertyDiff_DELETE, append(replaces, path)
			case DiffUpdate:
				kind = khulnasoftrpc.PropertyDiff_UPDATE
			case DiffUpdateReplace:
				kind, replaces = khulnasoftrpc.PropertyDiff_UPDATE_REPLACE, append(replaces, path)
			}

			detailedDiff[path] = &khulnasoftrpc.PropertyDiff{
				Kind:      kind,
				InputDiff: diff.InputDiff,
			}
		}
	}

	return &khulnasoftrpc.DiffResponse{
		Replaces:            replaces,
		DeleteBeforeReplace: diff.DeleteBeforeReplace,
		Changes:             changes,
		Diffs:               diffs,
		DetailedDiff:        detailedDiff,
	}, nil
}

func (p *providerServer) Handshake(
	ctx context.Context,
	req *khulnasoftrpc.ProviderHandshakeRequest,
) (*khulnasoftrpc.ProviderHandshakeResponse, error) {
	_, err := p.provider.Handshake(ctx, ProviderHandshakeRequest{
		EngineAddress:    req.EngineAddress,
		RootDirectory:    req.RootDirectory,
		ProgramDirectory: req.ProgramDirectory,
	})
	if err != nil {
		return nil, err
	}

	return &khulnasoftrpc.ProviderHandshakeResponse{}, nil
}

func (p *providerServer) Parameterize(
	ctx context.Context, req *khulnasoftrpc.ParameterizeRequest,
) (*khulnasoftrpc.ParameterizeResponse, error) {
	var params ParameterizeParameters
	switch p := req.Parameters.(type) {
	case *khulnasoftrpc.ParameterizeRequest_Args:
		params = &ParameterizeArgs{Args: p.Args.GetArgs()}
	case *khulnasoftrpc.ParameterizeRequest_Value:
		version, err := semver.Parse(p.Value.GetVersion())
		if err != nil {
			return nil, err
		}
		params = &ParameterizeValue{
			Name:    p.Value.GetName(),
			Version: version,
			Value:   p.Value.Value,
		}
	}
	resp, err := p.provider.Parameterize(ctx, ParameterizeRequest{Parameters: params})
	if err != nil {
		return nil, err
	}
	return &khulnasoftrpc.ParameterizeResponse{
		Name:    resp.Name,
		Version: resp.Version.String(),
	}, nil
}

func (p *providerServer) GetSchema(ctx context.Context,
	req *khulnasoftrpc.GetSchemaRequest,
) (*khulnasoftrpc.GetSchemaResponse, error) {
	var subpackageVersion *semver.Version
	if req.SubpackageVersion != "" {
		v, err := semver.ParseTolerant(req.SubpackageVersion)
		if err != nil {
			return nil, err
		}
		subpackageVersion = &v
	}

	schema, err := p.provider.GetSchema(ctx, GetSchemaRequest{
		Version:           req.Version,
		SubpackageName:    req.SubpackageName,
		SubpackageVersion: subpackageVersion,
	})
	if err != nil {
		return nil, err
	}
	return &khulnasoftrpc.GetSchemaResponse{Schema: string(schema.Schema)}, nil
}

func (p *providerServer) GetPluginInfo(ctx context.Context, req *emptypb.Empty) (*khulnasoftrpc.PluginInfo, error) {
	info, err := p.provider.GetPluginInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &khulnasoftrpc.PluginInfo{Version: info.Version.String()}, nil
}

func (p *providerServer) Attach(ctx context.Context, req *khulnasoftrpc.PluginAttach) (*emptypb.Empty, error) {
	// NewProviderServer should take a GrpcProvider instead of Provider, but that's a breaking change
	// so for now we type test here
	if grpcProvider, ok := p.provider.(GrpcProvider); ok {
		err := grpcProvider.Attach(req.GetAddress())
		if err != nil {
			return nil, err
		}
		return &emptypb.Empty{}, nil
	}
	// Else report this is unsupported
	return nil, status.Error(codes.Unimplemented, "Attach is not yet implemented")
}

func (p *providerServer) Cancel(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	if err := p.provider.SignalCancellation(ctx); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (p *providerServer) CheckConfig(ctx context.Context,
	req *khulnasoftrpc.CheckRequest,
) (*khulnasoftrpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())

	// To support old engines fill in Name/Type if the engine didn't send them
	if req.Name == "" {
		req.Name = urn.Name()
	}
	if req.Name != urn.Name() {
		return nil, status.Error(codes.InvalidArgument, "name in request does not match URN")
	}
	if req.Type == "" {
		req.Type = string(urn.Type())
	}
	if req.Type != string(urn.Type()) {
		return nil, status.Error(codes.InvalidArgument, "type in request does not match URN")
	}

	state, err := UnmarshalProperties(req.GetOlds(), p.unmarshalOptions("olds", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	inputs, err := UnmarshalProperties(req.GetNews(), p.unmarshalOptions("news", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	resp, err := p.provider.CheckConfig(ctx, CheckConfigRequest{
		URN:           urn,
		Name:          req.Name,
		Type:          tokens.Type(req.Type),
		Olds:          state,
		News:          inputs,
		AllowUnknowns: true,
	})
	if err != nil {
		return nil, p.checkNYI("CheckConfig", err)
	}

	rpcInputs, err := MarshalProperties(resp.Properties, p.marshalOptions("inputs"))
	if err != nil {
		return nil, err
	}

	rpcFailures := make([]*khulnasoftrpc.CheckFailure, len(resp.Failures))
	for i, f := range resp.Failures {
		rpcFailures[i] = &khulnasoftrpc.CheckFailure{Property: string(f.Property), Reason: f.Reason}
	}

	return &khulnasoftrpc.CheckResponse{Inputs: rpcInputs, Failures: rpcFailures}, nil
}

func (p *providerServer) DiffConfig(ctx context.Context, req *khulnasoftrpc.DiffRequest) (*khulnasoftrpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())

	// To support old engines fill in Name/Type if the engine didn't send them
	if req.Name == "" {
		req.Name = urn.Name()
	}
	if req.Name != urn.Name() {
		return nil, status.Error(codes.InvalidArgument, "name in request does not match URN")
	}
	if req.Type == "" {
		req.Type = string(urn.Type())
	}
	if req.Type != string(urn.Type()) {
		return nil, status.Error(codes.InvalidArgument, "type in request does not match URN")
	}

	oldInputs, err := UnmarshalProperties(
		req.GetOldInputs(), p.unmarshalOptions("oldInputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	oldOutputs, err := UnmarshalProperties(
		req.GetOlds(), p.unmarshalOptions("oldOutputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	newInputs, err := UnmarshalProperties(
		req.GetNews(), p.unmarshalOptions("newInputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	diff, err := p.provider.DiffConfig(ctx, DiffConfigRequest{
		URN:           urn,
		Name:          req.Name,
		Type:          tokens.Type(req.Type),
		OldInputs:     oldInputs,
		OldOutputs:    oldOutputs,
		NewInputs:     newInputs,
		AllowUnknowns: true,
		IgnoreChanges: req.GetIgnoreChanges(),
	})
	if err != nil {
		return nil, p.checkNYI("DiffConfig", err)
	}
	return p.marshalDiff(diff)
}

func (p *providerServer) Configure(ctx context.Context,
	req *khulnasoftrpc.ConfigureRequest,
) (*khulnasoftrpc.ConfigureResponse, error) {
	var inputs resource.PropertyMap
	if req.GetArgs() != nil {
		args, err := UnmarshalProperties(req.GetArgs(), p.unmarshalOptions("args", false /* keepOutputValues */))
		if err != nil {
			return nil, err
		}
		inputs = args
	} else {
		inputs = make(resource.PropertyMap)
		for k, v := range req.GetVariables() {
			key, err := config.ParseKey(k)
			if err != nil {
				return nil, err
			}

			var value interface{}
			if err = json.Unmarshal([]byte(v), &value); err != nil {
				// If we couldn't unmarshal a JSON value, just pass the raw string through.
				value = v
			}

			inputs[resource.PropertyKey(key.Name())] = resource.NewPropertyValue(value)
		}
	}

	if _, err := p.provider.Configure(ctx, ConfigureRequest{inputs}); err != nil {
		return nil, err
	}

	p.keepSecrets = req.GetAcceptSecrets()
	p.keepResources = req.GetAcceptResources()
	return &khulnasoftrpc.ConfigureResponse{
		AcceptSecrets: true, SupportsPreview: true, AcceptResources: true, AcceptOutputs: true,
	}, nil
}

func (p *providerServer) Check(ctx context.Context, req *khulnasoftrpc.CheckRequest) (*khulnasoftrpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())

	// To support old engines fill in Name/Type if the engine didn't send them
	if req.Name == "" {
		req.Name = urn.Name()
	}
	if req.Name != urn.Name() {
		return nil, status.Error(codes.InvalidArgument, "name in request does not match URN")
	}
	if req.Type == "" {
		req.Type = string(urn.Type())
	}
	if req.Type != string(urn.Type()) {
		return nil, status.Error(codes.InvalidArgument, "type in request does not match URN")
	}

	state, err := UnmarshalProperties(req.GetOlds(), p.unmarshalOptions("state", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	inputs, err := UnmarshalProperties(req.GetNews(), p.unmarshalOptions("inputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	var autonaming *AutonamingOptions
	if req.Autonaming != nil {
		autonaming = &AutonamingOptions{
			ProposedName: req.Autonaming.ProposedName,
			Mode:         AutonamingMode(req.Autonaming.Mode),
		}
	}

	resp, err := p.provider.Check(ctx, CheckRequest{
		URN:           urn,
		Name:          req.Name,
		Type:          tokens.Type(req.Type),
		Olds:          state,
		News:          inputs,
		AllowUnknowns: true,
		RandomSeed:    req.RandomSeed,
		Autonaming:    autonaming,
	})
	if err != nil {
		return nil, err
	}

	rpcInputs, err := MarshalProperties(resp.Properties, p.marshalOptions("newInputs"))
	if err != nil {
		return nil, err
	}

	rpcFailures := make([]*khulnasoftrpc.CheckFailure, len(resp.Failures))
	for i, f := range resp.Failures {
		rpcFailures[i] = &khulnasoftrpc.CheckFailure{Property: string(f.Property), Reason: f.Reason}
	}

	return &khulnasoftrpc.CheckResponse{Inputs: rpcInputs, Failures: rpcFailures}, nil
}

func (p *providerServer) Diff(ctx context.Context, req *khulnasoftrpc.DiffRequest) (*khulnasoftrpc.DiffResponse, error) {
	urn, id := resource.URN(req.GetUrn()), resource.ID(req.GetId())

	// To support old engines fill in Name/Type if the engine didn't send them
	if req.Name == "" {
		req.Name = urn.Name()
	}
	if req.Name != urn.Name() {
		return nil, status.Error(codes.InvalidArgument, "name in request does not match URN")
	}
	if req.Type == "" {
		req.Type = string(urn.Type())
	}
	if req.Type != string(urn.Type()) {
		return nil, status.Error(codes.InvalidArgument, "type in request does not match URN")
	}

	oldInputs, err := UnmarshalProperties(
		req.GetOldInputs(), p.unmarshalOptions("oldInputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	oldOutputs, err := UnmarshalProperties(
		req.GetOlds(), p.unmarshalOptions("oldOutputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	newInputs, err := UnmarshalProperties(
		req.GetNews(), p.unmarshalOptions("newInputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	diff, err := p.provider.Diff(ctx, DiffRequest{
		URN:           urn,
		Name:          req.Name,
		Type:          tokens.Type(req.Type),
		ID:            id,
		OldInputs:     oldInputs,
		OldOutputs:    oldOutputs,
		NewInputs:     newInputs,
		AllowUnknowns: true,
		IgnoreChanges: req.GetIgnoreChanges(),
	})
	if err != nil {
		return nil, err
	}
	return p.marshalDiff(diff)
}

func (p *providerServer) Create(ctx context.Context, req *khulnasoftrpc.CreateRequest) (*khulnasoftrpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())

	// To support old engines fill in Name/Type if the engine didn't send them
	if req.Name == "" {
		req.Name = urn.Name()
	}
	if req.Name != urn.Name() {
		return nil, status.Error(codes.InvalidArgument, "name in request does not match URN")
	}
	if req.Type == "" {
		req.Type = string(urn.Type())
	}
	if req.Type != string(urn.Type()) {
		return nil, status.Error(codes.InvalidArgument, "type in request does not match URN")
	}

	inputs, err := UnmarshalProperties(req.GetProperties(), p.unmarshalOptions("inputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	resp, err := p.provider.Create(ctx, CreateRequest{
		URN:        urn,
		Name:       req.Name,
		Type:       tokens.Type(req.Type),
		Properties: inputs,
		Timeout:    req.GetTimeout(),
		Preview:    req.GetPreview(),
	})
	if err != nil {
		return nil, err
	}

	rpcState, err := MarshalProperties(resp.Properties, p.marshalOptions("newState"))
	if err != nil {
		return nil, err
	}

	return &khulnasoftrpc.CreateResponse{
		Id:         string(resp.ID),
		Properties: rpcState,
	}, nil
}

func (p *providerServer) Read(ctx context.Context, req *khulnasoftrpc.ReadRequest) (*khulnasoftrpc.ReadResponse, error) {
	urn, requestID := resource.URN(req.GetUrn()), resource.ID(req.GetId())

	// To support old engines fill in Name/Type if the engine didn't send them
	if req.Name == "" {
		req.Name = urn.Name()
	}
	if req.Name != urn.Name() {
		return nil, status.Error(codes.InvalidArgument, "name in request does not match URN")
	}
	if req.Type == "" {
		req.Type = string(urn.Type())
	}
	if req.Type != string(urn.Type()) {
		return nil, status.Error(codes.InvalidArgument, "type in request does not match URN")
	}

	state, err := UnmarshalProperties(req.GetProperties(), p.unmarshalOptions("state", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	inputs, err := UnmarshalProperties(req.GetInputs(), p.unmarshalOptions("inputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	resp, err := p.provider.Read(ctx, ReadRequest{
		URN:    urn,
		Name:   req.Name,
		Type:   tokens.Type(req.Type),
		ID:     requestID,
		Inputs: inputs,
		State:  state,
	})
	if err != nil {
		return nil, err
	}

	rpcState, err := MarshalProperties(resp.Outputs, p.marshalOptions("newState"))
	if err != nil {
		return nil, err
	}

	rpcInputs, err := MarshalProperties(resp.Inputs, p.marshalOptions("newInputs"))
	if err != nil {
		return nil, err
	}

	return &khulnasoftrpc.ReadResponse{
		Id:         string(resp.ID),
		Properties: rpcState,
		Inputs:     rpcInputs,
	}, nil
}

func (p *providerServer) Update(ctx context.Context, req *khulnasoftrpc.UpdateRequest) (*khulnasoftrpc.UpdateResponse, error) {
	urn, id := resource.URN(req.GetUrn()), resource.ID(req.GetId())

	// To support old engines fill in Name/Type if the engine didn't send them
	if req.Name == "" {
		req.Name = urn.Name()
	}
	if req.Name != urn.Name() {
		return nil, status.Error(codes.InvalidArgument, "name in request does not match URN")
	}
	if req.Type == "" {
		req.Type = string(urn.Type())
	}
	if req.Type != string(urn.Type()) {
		return nil, status.Error(codes.InvalidArgument, "type in request does not match URN")
	}

	oldOutputs, err := UnmarshalProperties(
		req.GetOlds(), p.unmarshalOptions("oldOutputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	oldInputs, err := UnmarshalProperties(
		req.GetOldInputs(), p.unmarshalOptions("oldInputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	newInputs, err := UnmarshalProperties(
		req.GetNews(), p.unmarshalOptions("newInputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	resp, err := p.provider.Update(ctx, UpdateRequest{
		URN:           urn,
		Name:          req.Name,
		Type:          tokens.Type(req.Type),
		ID:            id,
		OldInputs:     oldInputs,
		OldOutputs:    oldOutputs,
		NewInputs:     newInputs,
		Timeout:       req.GetTimeout(),
		IgnoreChanges: req.GetIgnoreChanges(),
		Preview:       req.GetPreview(),
	})
	if err != nil {
		return nil, err
	}

	rpcState, err := MarshalProperties(resp.Properties, p.marshalOptions("newState"))
	if err != nil {
		return nil, err
	}

	return &khulnasoftrpc.UpdateResponse{Properties: rpcState}, nil
}

func (p *providerServer) Delete(ctx context.Context, req *khulnasoftrpc.DeleteRequest) (*emptypb.Empty, error) {
	urn, id := resource.URN(req.GetUrn()), resource.ID(req.GetId())

	// To support old engines fill in Name/Type if the engine didn't send them
	if req.Name == "" {
		req.Name = urn.Name()
	}
	if req.Name != urn.Name() {
		return nil, status.Error(codes.InvalidArgument, "name in request does not match URN")
	}
	if req.Type == "" {
		req.Type = string(urn.Type())
	}
	if req.Type != string(urn.Type()) {
		return nil, status.Error(codes.InvalidArgument, "type in request does not match URN")
	}

	inputs, err := UnmarshalProperties(req.GetOldInputs(), p.unmarshalOptions("inputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	outputs, err := UnmarshalProperties(req.GetProperties(), p.unmarshalOptions("outputs", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	if _, err = p.provider.Delete(ctx, DeleteRequest{
		URN:     urn,
		Name:    req.Name,
		Type:    tokens.Type(req.Type),
		ID:      id,
		Inputs:  inputs,
		Outputs: outputs,
		Timeout: req.GetTimeout(),
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (p *providerServer) Construct(ctx context.Context,
	req *khulnasoftrpc.ConstructRequest,
) (*khulnasoftrpc.ConstructResponse, error) {
	typ, name, parent := tokens.Type(req.GetType()), req.GetName(), resource.URN(req.GetParent())

	inputs, err := UnmarshalProperties(req.GetInputs(), p.unmarshalOptions("inputs", true /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	cfg := map[config.Key]string{}
	for k, v := range req.GetConfig() {
		configKey, err := config.ParseKey(k)
		if err != nil {
			return nil, err
		}
		cfg[configKey] = v
	}

	cfgSecretKeys := []config.Key{}
	for _, k := range req.GetConfigSecretKeys() {
		key, err := config.ParseKey(k)
		if err != nil {
			return nil, err
		}
		cfgSecretKeys = append(cfgSecretKeys, key)
	}

	info := ConstructInfo{
		Project:          req.GetProject(),
		Stack:            req.GetStack(),
		Config:           cfg,
		ConfigSecretKeys: cfgSecretKeys,
		DryRun:           req.GetDryRun(),
		Parallel:         req.GetParallel(),
		MonitorAddress:   req.GetMonitorEndpoint(),
	}

	aliases := make([]resource.Alias, len(req.GetAliases()))
	for i, urn := range req.GetAliases() {
		aliases[i] = resource.Alias{URN: resource.URN(urn)}
	}
	dependencies := make([]resource.URN, len(req.GetDependencies()))
	for i, urn := range req.GetDependencies() {
		dependencies[i] = resource.URN(urn)
	}
	propertyDependencies := map[resource.PropertyKey][]resource.URN{}
	for name, deps := range req.GetInputDependencies() {
		urns := make([]resource.URN, len(deps.Urns))
		for i, urn := range deps.Urns {
			urns[i] = resource.URN(urn)
		}
		propertyDependencies[resource.PropertyKey(name)] = urns
	}
	options := ConstructOptions{
		Aliases:              aliases,
		Dependencies:         dependencies,
		Protect:              req.GetProtect(),
		Providers:            req.GetProviders(),
		PropertyDependencies: propertyDependencies,
	}

	resp, err := p.provider.Construct(ctx, ConstructRequest{
		Info:    info,
		Type:    typ,
		Name:    name,
		Parent:  parent,
		Inputs:  inputs,
		Options: options,
	})
	if err != nil {
		return nil, rpcerror.WrapDetailedError(err)
	}

	opts := p.marshalOptions("outputs")
	opts.KeepOutputValues = req.AcceptsOutputValues
	outputs, err := MarshalProperties(resp.Outputs, opts)
	if err != nil {
		return nil, err
	}

	outputDependencies := map[string]*khulnasoftrpc.ConstructResponse_PropertyDependencies{}
	for name, deps := range resp.OutputDependencies {
		urns := make([]string, len(deps))
		for i, urn := range deps {
			urns[i] = string(urn)
		}
		outputDependencies[string(name)] = &khulnasoftrpc.ConstructResponse_PropertyDependencies{Urns: urns}
	}

	return &khulnasoftrpc.ConstructResponse{
		Urn:               string(resp.URN),
		State:             outputs,
		StateDependencies: outputDependencies,
	}, nil
}

func (p *providerServer) Invoke(ctx context.Context, req *khulnasoftrpc.InvokeRequest) (*khulnasoftrpc.InvokeResponse, error) {
	args, err := UnmarshalProperties(req.GetArgs(), p.unmarshalOptions("args", false /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	resp, err := p.provider.Invoke(ctx, InvokeRequest{
		Tok:  tokens.ModuleMember(req.GetTok()),
		Args: args,
	})
	if err != nil {
		return nil, err
	}

	rpcResult, err := MarshalProperties(resp.Properties, p.marshalOptions("result"))
	if err != nil {
		return nil, err
	}

	rpcFailures := make([]*khulnasoftrpc.CheckFailure, len(resp.Failures))
	for i, f := range resp.Failures {
		rpcFailures[i] = &khulnasoftrpc.CheckFailure{Property: string(f.Property), Reason: f.Reason}
	}

	return &khulnasoftrpc.InvokeResponse{
		Return:   rpcResult,
		Failures: rpcFailures,
	}, nil
}

func (p *providerServer) StreamInvoke(req *khulnasoftrpc.InvokeRequest,
	server khulnasoftrpc.ResourceProvider_StreamInvokeServer,
) error {
	args, err := UnmarshalProperties(req.GetArgs(), p.unmarshalOptions("args", false /* keepOutputValues */))
	if err != nil {
		return err
	}

	resp, err := p.provider.StreamInvoke(context.TODO(), StreamInvokeRequest{
		Tok:  tokens.ModuleMember(req.GetTok()),
		Args: args,
		OnNext: func(item resource.PropertyMap) error {
			rpcItem, err := MarshalProperties(item, p.marshalOptions("item"))
			if err != nil {
				return err
			}

			return server.Send(&khulnasoftrpc.InvokeResponse{Return: rpcItem})
		},
	})
	if err != nil {
		return err
	}
	if len(resp.Failures) == 0 {
		return nil
	}

	rpcFailures := make([]*khulnasoftrpc.CheckFailure, len(resp.Failures))
	for i, f := range resp.Failures {
		rpcFailures[i] = &khulnasoftrpc.CheckFailure{Property: string(f.Property), Reason: f.Reason}
	}

	return server.Send(&khulnasoftrpc.InvokeResponse{Failures: rpcFailures})
}

func (p *providerServer) Call(ctx context.Context, req *khulnasoftrpc.CallRequest) (*khulnasoftrpc.CallResponse, error) {
	args, err := UnmarshalProperties(req.GetArgs(), p.unmarshalOptions("args", true /* keepOutputValues */))
	if err != nil {
		return nil, err
	}

	cfg := map[config.Key]string{}
	for k, v := range req.GetConfig() {
		configKey, err := config.ParseKey(k)
		if err != nil {
			return nil, err
		}
		cfg[configKey] = v
	}
	info := CallInfo{
		Project:        req.GetProject(),
		Stack:          req.GetStack(),
		Config:         cfg,
		DryRun:         req.GetDryRun(),
		Parallel:       req.GetParallel(),
		MonitorAddress: req.GetMonitorEndpoint(),
	}
	argDependencies := map[resource.PropertyKey][]resource.URN{}
	for name, deps := range req.GetArgDependencies() {
		urns := make([]resource.URN, len(deps.Urns))
		for i, urn := range deps.Urns {
			urns[i] = resource.URN(urn)
		}
		argDependencies[resource.PropertyKey(name)] = urns
	}
	options := CallOptions{
		ArgDependencies: argDependencies,
	}

	result, err := p.provider.Call(ctx, CallRequest{
		Tok:     tokens.ModuleMember(req.GetTok()),
		Args:    args,
		Info:    info,
		Options: options,
	})
	if err != nil {
		err = rpcerror.WrapDetailedError(err)
		return nil, err
	}

	opts := p.marshalOptions("return")
	opts.KeepOutputValues = req.AcceptsOutputValues
	rpcResult, err := MarshalProperties(result.Return, opts)
	if err != nil {
		return nil, err
	}

	returnDependencies := map[string]*khulnasoftrpc.CallResponse_ReturnDependencies{}
	for name, deps := range result.ReturnDependencies {
		urns := make([]string, len(deps))
		for i, urn := range deps {
			urns[i] = string(urn)
		}
		returnDependencies[string(name)] = &khulnasoftrpc.CallResponse_ReturnDependencies{Urns: urns}
	}

	rpcFailures := make([]*khulnasoftrpc.CheckFailure, len(result.Failures))
	for i, f := range result.Failures {
		rpcFailures[i] = &khulnasoftrpc.CheckFailure{Property: string(f.Property), Reason: f.Reason}
	}

	return &khulnasoftrpc.CallResponse{
		Return:             rpcResult,
		ReturnDependencies: returnDependencies,
		Failures:           rpcFailures,
	}, nil
}

func (p *providerServer) GetMapping(ctx context.Context,
	req *khulnasoftrpc.GetMappingRequest,
) (*khulnasoftrpc.GetMappingResponse, error) {
	resp, err := p.provider.GetMapping(ctx, GetMappingRequest{
		Key:      req.Key,
		Provider: req.Provider,
	})
	if err != nil {
		return nil, err
	}
	return &khulnasoftrpc.GetMappingResponse{Data: resp.Data, Provider: resp.Provider}, nil
}

func (p *providerServer) GetMappings(ctx context.Context,
	req *khulnasoftrpc.GetMappingsRequest,
) (*khulnasoftrpc.GetMappingsResponse, error) {
	providers, err := p.provider.GetMappings(ctx, GetMappingsRequest{
		Key: req.Key,
	})
	if err != nil {
		return nil, err
	}
	return &khulnasoftrpc.GetMappingsResponse{Providers: providers.Keys}, nil
}
