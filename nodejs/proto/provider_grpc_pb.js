// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
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
//
'use strict';
var grpc = require('@grpc/grpc-js');
var khulnasoft_provider_pb = require('./provider_pb.js');
var khulnasoft_plugin_pb = require('./plugin_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_CallRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.CallRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.CallRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_CallRequest(buffer_arg) {
  return khulnasoft_provider_pb.CallRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_CallResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.CallResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.CallResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_CallResponse(buffer_arg) {
  return khulnasoft_provider_pb.CallResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_CheckRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.CheckRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.CheckRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_CheckRequest(buffer_arg) {
  return khulnasoft_provider_pb.CheckRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_CheckResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.CheckResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.CheckResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_CheckResponse(buffer_arg) {
  return khulnasoft_provider_pb.CheckResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ConfigureRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ConfigureRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConfigureRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConfigureRequest(buffer_arg) {
  return khulnasoft_provider_pb.ConfigureRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ConfigureResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ConfigureResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConfigureResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConfigureResponse(buffer_arg) {
  return khulnasoft_provider_pb.ConfigureResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ConstructRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ConstructRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConstructRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConstructRequest(buffer_arg) {
  return khulnasoft_provider_pb.ConstructRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ConstructResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ConstructResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConstructResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConstructResponse(buffer_arg) {
  return khulnasoft_provider_pb.ConstructResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_CreateRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.CreateRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.CreateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_CreateRequest(buffer_arg) {
  return khulnasoft_provider_pb.CreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_CreateResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.CreateResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.CreateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_CreateResponse(buffer_arg) {
  return khulnasoft_provider_pb.CreateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_DeleteRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.DeleteRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.DeleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_DeleteRequest(buffer_arg) {
  return khulnasoft_provider_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_DiffRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.DiffRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.DiffRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_DiffRequest(buffer_arg) {
  return khulnasoft_provider_pb.DiffRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_DiffResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.DiffResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.DiffResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_DiffResponse(buffer_arg) {
  return khulnasoft_provider_pb.DiffResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetMappingRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.GetMappingRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetMappingRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetMappingRequest(buffer_arg) {
  return khulnasoft_provider_pb.GetMappingRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetMappingResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.GetMappingResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetMappingResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetMappingResponse(buffer_arg) {
  return khulnasoft_provider_pb.GetMappingResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetMappingsRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.GetMappingsRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetMappingsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetMappingsRequest(buffer_arg) {
  return khulnasoft_provider_pb.GetMappingsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetMappingsResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.GetMappingsResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetMappingsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetMappingsResponse(buffer_arg) {
  return khulnasoft_provider_pb.GetMappingsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetSchemaRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.GetSchemaRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetSchemaRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetSchemaRequest(buffer_arg) {
  return khulnasoft_provider_pb.GetSchemaRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetSchemaResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.GetSchemaResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetSchemaResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetSchemaResponse(buffer_arg) {
  return khulnasoft_provider_pb.GetSchemaResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_InvokeRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.InvokeRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.InvokeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_InvokeRequest(buffer_arg) {
  return khulnasoft_provider_pb.InvokeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_InvokeResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.InvokeResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.InvokeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_InvokeResponse(buffer_arg) {
  return khulnasoft_provider_pb.InvokeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ParameterizeRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ParameterizeRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ParameterizeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ParameterizeRequest(buffer_arg) {
  return khulnasoft_provider_pb.ParameterizeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ParameterizeResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ParameterizeResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.ParameterizeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ParameterizeResponse(buffer_arg) {
  return khulnasoft_provider_pb.ParameterizeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_PluginAttach(arg) {
  if (!(arg instanceof khulnasoft_plugin_pb.PluginAttach)) {
    throw new Error('Expected argument of type khulnasoftrpc.PluginAttach');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_PluginAttach(buffer_arg) {
  return khulnasoft_plugin_pb.PluginAttach.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_PluginInfo(arg) {
  if (!(arg instanceof khulnasoft_plugin_pb.PluginInfo)) {
    throw new Error('Expected argument of type khulnasoftrpc.PluginInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_PluginInfo(buffer_arg) {
  return khulnasoft_plugin_pb.PluginInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ProviderHandshakeRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ProviderHandshakeRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ProviderHandshakeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ProviderHandshakeRequest(buffer_arg) {
  return khulnasoft_provider_pb.ProviderHandshakeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ProviderHandshakeResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ProviderHandshakeResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.ProviderHandshakeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ProviderHandshakeResponse(buffer_arg) {
  return khulnasoft_provider_pb.ProviderHandshakeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ReadRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ReadRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ReadRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ReadRequest(buffer_arg) {
  return khulnasoft_provider_pb.ReadRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ReadResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.ReadResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.ReadResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ReadResponse(buffer_arg) {
  return khulnasoft_provider_pb.ReadResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_UpdateRequest(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.UpdateRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.UpdateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_UpdateRequest(buffer_arg) {
  return khulnasoft_provider_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_UpdateResponse(arg) {
  if (!(arg instanceof khulnasoft_provider_pb.UpdateResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.UpdateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_UpdateResponse(buffer_arg) {
  return khulnasoft_provider_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The ResourceProvider service defines a standard interface for [resource providers](providers). A resource provider
// manages a set of configuration, resources, functions and so on in a single package, and offers methods such as CRUD
// operations on resources and invocations of functions. Resource providers are primarily managed by the Khulnasoft engine
// as part of a deployment in order to interact with the cloud providers underpinning a Khulnasoft application.
var ResourceProviderService = exports.ResourceProviderService = {
  // `Handshake` is the first call made by the engine to a provider. It is used to pass the engine's address to the
// provider so that it may establish its own connections back, and to establish protocol configuration that will be
// used to communicate between the two parties. Providers that support `Handshake` implicitly support the set of
// feature flags previously handled by `Configure` prior to `Handshake`'s introduction, such as secrets and resource
// references.
handshake: {
    path: '/khulnasoftrpc.ResourceProvider/Handshake',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.ProviderHandshakeRequest,
    responseType: khulnasoft_provider_pb.ProviderHandshakeResponse,
    requestSerialize: serialize_khulnasoftrpc_ProviderHandshakeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ProviderHandshakeRequest,
    responseSerialize: serialize_khulnasoftrpc_ProviderHandshakeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_ProviderHandshakeResponse,
  },
  // `Parameterize` is the primary means of supporting [parameterized providers](parameterized-providers), which allow
// a caller to change a provider's behavior ahead of its [configuration](khulnasoftrpc.ResourceProvider.Configure) and
// subsequent use. Where a [](khulnasoftrpc.ResourceProvider.Configure) call allows a caller to influence provider
// behaviour at a high level (e.g. by specifying the region in which an AWS provider should operate), a
// `Parameterize` call may change the set of resources and functions that a provider offers (that is, its schema).
// This is useful in any case where some "set" of providers can be captured by a single implementation that may
// power fundamentally different schemata -- dynamically bridging Terraform providers, or managing Kubernetes
// clusters with custom resource definitions, for instance, are good examples. The parameterized package that
// `Parameterize` yields is known as a *sub-package* of the original (unparameterized) package.
//
// `Parameterize` supports two types of parameterization:
//
// * *Replacement parameterization*, whereby a `Parameterize` call results in a schema that completely replaces the
//   original provider schema. Bridging a Terraform provider dynamically might be an example of this -- following
//   the call to `Parameterize`, the provider's schema will become that of the Terraform provider that was bridged.
//   Providers that implement replacement parameterization expect a *single* call to `Parameterize`.
//
// * *Extension parameterization*, in which a `Parameterize` call results in a schema that is a superset of the
//   original. This is useful in cases where a provider can be extended with additional resources or functions, such
//   as a Kubernetes provider that can be extended with resources representing custom resource definitions.
//   Providers that implement extension parameterization should accept multiple calls to `Parameterize`. Extension
//   packages may even be called multiple times with the same package name, but with different versions. The CRUD
//   operations of extension resources must include the version of which sub-package they correspond to.
//
// `Parameterize` should work the same whether it is provided with `ParametersArgs` or `ParametersValue` input. In
// each case it should return the sub-package name and version (which when a `ParametersValue` is supplied should
// match the given input).
parameterize: {
    path: '/khulnasoftrpc.ResourceProvider/Parameterize',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.ParameterizeRequest,
    responseType: khulnasoft_provider_pb.ParameterizeResponse,
    requestSerialize: serialize_khulnasoftrpc_ParameterizeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ParameterizeRequest,
    responseSerialize: serialize_khulnasoftrpc_ParameterizeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_ParameterizeResponse,
  },
  // GetSchema fetches the schema for this resource provider.
getSchema: {
    path: '/khulnasoftrpc.ResourceProvider/GetSchema',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.GetSchemaRequest,
    responseType: khulnasoft_provider_pb.GetSchemaResponse,
    requestSerialize: serialize_khulnasoftrpc_GetSchemaRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GetSchemaRequest,
    responseSerialize: serialize_khulnasoftrpc_GetSchemaResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GetSchemaResponse,
  },
  // `CheckConfig` validates a set of configuration inputs that will be passed to this provider instance.
// `CheckConfig` is to provider resources what [](khulnasoftrpc.ResourceProvider.Check) is to individual resources, and
// is the first stage in configuring (that is, eventually executing a [](khulnasoftrpc.ResourceProvider.Configure) call)
// a provider using user-supplied values. In the case that provider inputs are coming from some source that has been
// checked previously (e.g. a Khulnasoft state), it is not necessary to call `CheckConfig`.
//
// A `CheckConfig` call returns either a set of checked, known-valid inputs that may subsequently be passed to
// [](khulnasoftrpc.ResourceProvider.DiffConfig) and/or [](khulnasoftrpc.ResourceProvider.Configure), or a set of errors
// explaining why the inputs are invalid. In the case that a set of inputs are successfully validated and returned,
// `CheckConfig` *may also populate default values* for provider configuration, returning them so that they may be
// passed to a subsequent [](khulnasoftrpc.ResourceProvider.Configure) call and persisted in the Khulnasoft state. In the
// case that `CheckConfig` fails and returns a set of errors, it is expected that the caller (typically the Khulnasoft
// engine) will fail provider registration.
//
// As a rule, the provider inputs returned by a call to `CheckConfig` should preserve the original representation of
// the properties as present in the program inputs. Though this rule is not required for correctness, violations
// thereof can negatively impact the end-user experience, as the provider inputs are used for detecting and
// rendering diffs.
checkConfig: {
    path: '/khulnasoftrpc.ResourceProvider/CheckConfig',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.CheckRequest,
    responseType: khulnasoft_provider_pb.CheckResponse,
    requestSerialize: serialize_khulnasoftrpc_CheckRequest,
    requestDeserialize: deserialize_khulnasoftrpc_CheckRequest,
    responseSerialize: serialize_khulnasoftrpc_CheckResponse,
    responseDeserialize: deserialize_khulnasoftrpc_CheckResponse,
  },
  // `DiffConfig` compares an existing ("old") provider configuration with a new configuration and computes the
// difference (if any) between them. `DiffConfig` is to provider resources what [](khulnasoftrpc.ResourceProvider.Diff)
// is to individual resources. `DiffConfig` should only be called with values that have at some point been validated
// by a [](khulnasoftrpc.ResourceProvider.CheckConfig) call. The [](khulnasoftrpc.DiffResponse) returned by a `DiffConfig`
// call is used primarily to determine whether or not the newly configured provider is capable of managing resources
// owned by the old provider. If `DiffConfig` indicates that the provider resource needs to be replaced, for
// instance, then all resources owned by that provider will *also* need to be replaced. Replacement semantics should
// thus be reserved for changes to configuration properties that are guaranteed to make old resources unmanageable.
// Changes to an AWS region, for example, will almost certainly require a provider replacement, but changes to an
// AWS access key, should almost certainly not.
diffConfig: {
    path: '/khulnasoftrpc.ResourceProvider/DiffConfig',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.DiffRequest,
    responseType: khulnasoft_provider_pb.DiffResponse,
    requestSerialize: serialize_khulnasoftrpc_DiffRequest,
    requestDeserialize: deserialize_khulnasoftrpc_DiffRequest,
    responseSerialize: serialize_khulnasoftrpc_DiffResponse,
    responseDeserialize: deserialize_khulnasoftrpc_DiffResponse,
  },
  // `Configure` is the final stage in configuring a provider instance. Callers may supply two sets of data:
//
// * Provider-specific configuration, which is the set of inputs that have been validated by a previous
//   [](khulnasoftrpc.ResourceProvider.CheckConfig) call.
// * Provider-agnostic ("protocol") configuration, such as whether or not the caller supports secrets.
//
// The provider is expected to return its own set of protocol configuration, indicating which features it supports
// in turn so that the caller and the provider can interact appropriately.
//
// Providers may expect a *single* call to `Configure`. If a call to `Configure` is missing required configuration,
// the provider may return a set of error details containing [](khulnasoftrpc.ConfigureErrorMissingKeys) values to
// indicate which keys are missing.
//
// :::{important}
// The use of `Configure` to configure protocol features is deprecated in favour of the
// [](khulnasoftrpc.ResourceProvider.Handshake) method, which should be implemented by newer providers. To enable
// compatibility between older engines and providers:
//
// * Callers which call `Handshake` *must* call `Configure` with flags such as `acceptSecrets` and `acceptResources`
//   set to `true`, since these features predate the introduction of `Handshake` and thus `Handshake`-aware callers
//   must support them. See [](khulnasoftrpc.ConfigureRequest) for more information.
// * Providers which implement `Handshake` *must* support flags such as `acceptSecrets` and `acceptResources`, and
//   indicate as such by always returning `true` for these fields in [](khulnasoftrpc.ConfigureResponse). See
//   [](khulnasoftrpc.ConfigureResponse) for more information.
// :::
configure: {
    path: '/khulnasoftrpc.ResourceProvider/Configure',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.ConfigureRequest,
    responseType: khulnasoft_provider_pb.ConfigureResponse,
    requestSerialize: serialize_khulnasoftrpc_ConfigureRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ConfigureRequest,
    responseSerialize: serialize_khulnasoftrpc_ConfigureResponse,
    responseDeserialize: deserialize_khulnasoftrpc_ConfigureResponse,
  },
  // Invoke dynamically executes a built-in function in the provider.
invoke: {
    path: '/khulnasoftrpc.ResourceProvider/Invoke',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.InvokeRequest,
    responseType: khulnasoft_provider_pb.InvokeResponse,
    requestSerialize: serialize_khulnasoftrpc_InvokeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_InvokeRequest,
    responseSerialize: serialize_khulnasoftrpc_InvokeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_InvokeResponse,
  },
  // StreamInvoke dynamically executes a built-in function in the provider, which returns a stream
// of responses.
streamInvoke: {
    path: '/khulnasoftrpc.ResourceProvider/StreamInvoke',
    requestStream: false,
    responseStream: true,
    requestType: khulnasoft_provider_pb.InvokeRequest,
    responseType: khulnasoft_provider_pb.InvokeResponse,
    requestSerialize: serialize_khulnasoftrpc_InvokeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_InvokeRequest,
    responseSerialize: serialize_khulnasoftrpc_InvokeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_InvokeResponse,
  },
  // Call dynamically executes a method in the provider associated with a component resource.
call: {
    path: '/khulnasoftrpc.ResourceProvider/Call',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.CallRequest,
    responseType: khulnasoft_provider_pb.CallResponse,
    requestSerialize: serialize_khulnasoftrpc_CallRequest,
    requestDeserialize: deserialize_khulnasoftrpc_CallRequest,
    responseSerialize: serialize_khulnasoftrpc_CallResponse,
    responseDeserialize: deserialize_khulnasoftrpc_CallResponse,
  },
  // `Check` validates a set of input properties against a given resource type. A `Check` call returns either a set of
// checked, known-valid inputs that may subsequently be passed to [](khulnasoftrpc.ResourceProvider.Diff),
// [](khulnasoftrpc.ResourceProvider.Create), or [](khulnasoftrpc.ResourceProvider.Update); or a set of errors explaining
// why the inputs are invalid. In the case that a set of inputs are successfully validated and returned, `Check`
// *may also populate default values* for resource inputs, returning them so that they may be passed to a subsequent
// call and persisted in the Khulnasoft state. In the case that `Check` fails and returns a set of errors, it is
// expected that the caller (typically the Khulnasoft engine) will fail resource registration.
//
// As a rule, the provider inputs returned by a call to `Check` should preserve the original representation of the
// properties as present in the program inputs. Though this rule is not required for correctness, violations thereof
// can negatively impact the end-user experience, as the provider inputs are used for detecting and rendering
// diffs.
check: {
    path: '/khulnasoftrpc.ResourceProvider/Check',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.CheckRequest,
    responseType: khulnasoft_provider_pb.CheckResponse,
    requestSerialize: serialize_khulnasoftrpc_CheckRequest,
    requestDeserialize: deserialize_khulnasoftrpc_CheckRequest,
    responseSerialize: serialize_khulnasoftrpc_CheckResponse,
    responseDeserialize: deserialize_khulnasoftrpc_CheckResponse,
  },
  // `Diff` compares an existing ("old") set of resource properties with a new set of properties and computes the
// difference (if any) between them. `Diff` should only be called with values that have at some point been validated
// by a [](khulnasoftrpc.ResourceProvider.Check) call.
diff: {
    path: '/khulnasoftrpc.ResourceProvider/Diff',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.DiffRequest,
    responseType: khulnasoft_provider_pb.DiffResponse,
    requestSerialize: serialize_khulnasoftrpc_DiffRequest,
    requestDeserialize: deserialize_khulnasoftrpc_DiffRequest,
    responseSerialize: serialize_khulnasoftrpc_DiffResponse,
    responseDeserialize: deserialize_khulnasoftrpc_DiffResponse,
  },
  // `Create` provisions a new instance of the specified [(custom) resource](custom-resources). It returns a
// provider-assigned ID for the resource as well as the output properties that arose from the creation properties.
// Output properties are typically the union of the resource's input properties and any additional values that were
// computed or made available during creation.
//
// If creation fails, `Create` may return an [](khulnasoftrpc.ErrorResourceInitFailed) error detail explaining why.
// Moreover, if `Create` does return an error, it must be the case that the resource was *not* created (that is,
// `Create` can be thought of as transactional or atomic).
create: {
    path: '/khulnasoftrpc.ResourceProvider/Create',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.CreateRequest,
    responseType: khulnasoft_provider_pb.CreateResponse,
    requestSerialize: serialize_khulnasoftrpc_CreateRequest,
    requestDeserialize: deserialize_khulnasoftrpc_CreateRequest,
    responseSerialize: serialize_khulnasoftrpc_CreateResponse,
    responseDeserialize: deserialize_khulnasoftrpc_CreateResponse,
  },
  // `Read` reads the current live state associated with a resource identified by the supplied state. The given state
// must be sufficient to uniquely identify the resource. This is typically just the resource ID, but may also
// include other properties.
read: {
    path: '/khulnasoftrpc.ResourceProvider/Read',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.ReadRequest,
    responseType: khulnasoft_provider_pb.ReadResponse,
    requestSerialize: serialize_khulnasoftrpc_ReadRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ReadRequest,
    responseSerialize: serialize_khulnasoftrpc_ReadResponse,
    responseDeserialize: deserialize_khulnasoftrpc_ReadResponse,
  },
  // `Update` updates an existing resource according to a new set of inputs, returning a new set of output properties.
update: {
    path: '/khulnasoftrpc.ResourceProvider/Update',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.UpdateRequest,
    responseType: khulnasoft_provider_pb.UpdateResponse,
    requestSerialize: serialize_khulnasoftrpc_UpdateRequest,
    requestDeserialize: deserialize_khulnasoftrpc_UpdateRequest,
    responseSerialize: serialize_khulnasoftrpc_UpdateResponse,
    responseDeserialize: deserialize_khulnasoftrpc_UpdateResponse,
  },
  // `Delete` deprovisions an existing resource as specified by its ID. `Delete` should be transactional/atomic -- if
// a call to `Delete` fails, it must be the case that the resource was *not* deleted and can be assumed to still
// exist.
delete: {
    path: '/khulnasoftrpc.ResourceProvider/Delete',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.DeleteRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_khulnasoftrpc_DeleteRequest,
    requestDeserialize: deserialize_khulnasoftrpc_DeleteRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // `Construct` provisions a new [component resource](component-resources). Providers that implement `Construct` are
// referred to as [component providers](component-providers). `Construct` is to component resources what
// [](khulnasoftrpc.ResourceProvider.Create) is to [custom resources](custom-resources). Components do not have any
// lifecycle of their own, and instead embody the lifecycles of the resources that they are composed of. As such,
// `Construct` is effectively a subprogram whose resources will be persisted in the caller's state. It is
// consequently passed enough information to manage fully these resources. At a high level, this comprises:
//
// * A [](khulnasoftrpc.ResourceMonitor) endpoint which the provider can use to [register](resource-registration) nested
//   custom or component resources that belong to the component.
//
// * A set of input properties.
//
// * A full set of [resource options](https://www.khulnasoft.com/docs/iac/concepts/options/) that the component should
//   propagate to resources it registers against the supplied resource monitor.
construct: {
    path: '/khulnasoftrpc.ResourceProvider/Construct',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.ConstructRequest,
    responseType: khulnasoft_provider_pb.ConstructResponse,
    requestSerialize: serialize_khulnasoftrpc_ConstructRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ConstructRequest,
    responseSerialize: serialize_khulnasoftrpc_ConstructResponse,
    responseDeserialize: deserialize_khulnasoftrpc_ConstructResponse,
  },
  // Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either return a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
cancel: {
    path: '/khulnasoftrpc.ResourceProvider/Cancel',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // GetPluginInfo returns generic information about this plugin, like its version.
getPluginInfo: {
    path: '/khulnasoftrpc.ResourceProvider/GetPluginInfo',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: khulnasoft_plugin_pb.PluginInfo,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_khulnasoftrpc_PluginInfo,
    responseDeserialize: deserialize_khulnasoftrpc_PluginInfo,
  },
  // Attach sends the engine address to an already running plugin.
attach: {
    path: '/khulnasoftrpc.ResourceProvider/Attach',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_plugin_pb.PluginAttach,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_khulnasoftrpc_PluginAttach,
    requestDeserialize: deserialize_khulnasoftrpc_PluginAttach,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // GetMapping fetches the mapping for this resource provider, if any. A provider should return an empty
// response (not an error) if it doesn't have a mapping for the given key.
getMapping: {
    path: '/khulnasoftrpc.ResourceProvider/GetMapping',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.GetMappingRequest,
    responseType: khulnasoft_provider_pb.GetMappingResponse,
    requestSerialize: serialize_khulnasoftrpc_GetMappingRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GetMappingRequest,
    responseSerialize: serialize_khulnasoftrpc_GetMappingResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GetMappingResponse,
  },
  // GetMappings is an optional method that returns what mappings (if any) a provider supports. If a provider does not
// implement this method the engine falls back to the old behaviour of just calling GetMapping without a name.
// If this method is implemented than the engine will then call GetMapping only with the names returned from this method.
getMappings: {
    path: '/khulnasoftrpc.ResourceProvider/GetMappings',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_provider_pb.GetMappingsRequest,
    responseType: khulnasoft_provider_pb.GetMappingsResponse,
    requestSerialize: serialize_khulnasoftrpc_GetMappingsRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GetMappingsRequest,
    responseSerialize: serialize_khulnasoftrpc_GetMappingsResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GetMappingsResponse,
  },
};

exports.ResourceProviderClient = grpc.makeGenericClientConstructor(ResourceProviderService);
