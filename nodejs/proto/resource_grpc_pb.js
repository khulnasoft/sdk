// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2016-2022, KhulnaSoft, Ltd.
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
var khulnasoft_resource_pb = require('./resource_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');
var khulnasoft_provider_pb = require('./provider_pb.js');
var khulnasoft_alias_pb = require('./alias_pb.js');
var khulnasoft_source_pb = require('./source_pb.js');
var khulnasoft_callback_pb = require('./callback_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_khulnasoftrpc_Callback(arg) {
  if (!(arg instanceof khulnasoft_callback_pb.Callback)) {
    throw new Error('Expected argument of type khulnasoftrpc.Callback');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_Callback(buffer_arg) {
  return khulnasoft_callback_pb.Callback.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_khulnasoftrpc_ReadResourceRequest(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.ReadResourceRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ReadResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ReadResourceRequest(buffer_arg) {
  return khulnasoft_resource_pb.ReadResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ReadResourceResponse(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.ReadResourceResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.ReadResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ReadResourceResponse(buffer_arg) {
  return khulnasoft_resource_pb.ReadResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RegisterPackageRequest(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.RegisterPackageRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.RegisterPackageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RegisterPackageRequest(buffer_arg) {
  return khulnasoft_resource_pb.RegisterPackageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RegisterPackageResponse(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.RegisterPackageResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.RegisterPackageResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RegisterPackageResponse(buffer_arg) {
  return khulnasoft_resource_pb.RegisterPackageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RegisterResourceOutputsRequest(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.RegisterResourceOutputsRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.RegisterResourceOutputsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RegisterResourceOutputsRequest(buffer_arg) {
  return khulnasoft_resource_pb.RegisterResourceOutputsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RegisterResourceRequest(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.RegisterResourceRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.RegisterResourceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RegisterResourceRequest(buffer_arg) {
  return khulnasoft_resource_pb.RegisterResourceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RegisterResourceResponse(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.RegisterResourceResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.RegisterResourceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RegisterResourceResponse(buffer_arg) {
  return khulnasoft_resource_pb.RegisterResourceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ResourceCallRequest(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.ResourceCallRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ResourceCallRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ResourceCallRequest(buffer_arg) {
  return khulnasoft_resource_pb.ResourceCallRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ResourceInvokeRequest(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.ResourceInvokeRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ResourceInvokeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ResourceInvokeRequest(buffer_arg) {
  return khulnasoft_resource_pb.ResourceInvokeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_SupportsFeatureRequest(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.SupportsFeatureRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.SupportsFeatureRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_SupportsFeatureRequest(buffer_arg) {
  return khulnasoft_resource_pb.SupportsFeatureRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_SupportsFeatureResponse(arg) {
  if (!(arg instanceof khulnasoft_resource_pb.SupportsFeatureResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.SupportsFeatureResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_SupportsFeatureResponse(buffer_arg) {
  return khulnasoft_resource_pb.SupportsFeatureResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// ResourceMonitor is the interface a source uses to talk back to the planning monitor orchestrating the execution.
var ResourceMonitorService = exports.ResourceMonitorService = {
  supportsFeature: {
    path: '/khulnasoftrpc.ResourceMonitor/SupportsFeature',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_resource_pb.SupportsFeatureRequest,
    responseType: khulnasoft_resource_pb.SupportsFeatureResponse,
    requestSerialize: serialize_khulnasoftrpc_SupportsFeatureRequest,
    requestDeserialize: deserialize_khulnasoftrpc_SupportsFeatureRequest,
    responseSerialize: serialize_khulnasoftrpc_SupportsFeatureResponse,
    responseDeserialize: deserialize_khulnasoftrpc_SupportsFeatureResponse,
  },
  invoke: {
    path: '/khulnasoftrpc.ResourceMonitor/Invoke',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_resource_pb.ResourceInvokeRequest,
    responseType: khulnasoft_provider_pb.InvokeResponse,
    requestSerialize: serialize_khulnasoftrpc_ResourceInvokeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ResourceInvokeRequest,
    responseSerialize: serialize_khulnasoftrpc_InvokeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_InvokeResponse,
  },
  streamInvoke: {
    path: '/khulnasoftrpc.ResourceMonitor/StreamInvoke',
    requestStream: false,
    responseStream: true,
    requestType: khulnasoft_resource_pb.ResourceInvokeRequest,
    responseType: khulnasoft_provider_pb.InvokeResponse,
    requestSerialize: serialize_khulnasoftrpc_ResourceInvokeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ResourceInvokeRequest,
    responseSerialize: serialize_khulnasoftrpc_InvokeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_InvokeResponse,
  },
  call: {
    path: '/khulnasoftrpc.ResourceMonitor/Call',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_resource_pb.ResourceCallRequest,
    responseType: khulnasoft_provider_pb.CallResponse,
    requestSerialize: serialize_khulnasoftrpc_ResourceCallRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ResourceCallRequest,
    responseSerialize: serialize_khulnasoftrpc_CallResponse,
    responseDeserialize: deserialize_khulnasoftrpc_CallResponse,
  },
  readResource: {
    path: '/khulnasoftrpc.ResourceMonitor/ReadResource',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_resource_pb.ReadResourceRequest,
    responseType: khulnasoft_resource_pb.ReadResourceResponse,
    requestSerialize: serialize_khulnasoftrpc_ReadResourceRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ReadResourceRequest,
    responseSerialize: serialize_khulnasoftrpc_ReadResourceResponse,
    responseDeserialize: deserialize_khulnasoftrpc_ReadResourceResponse,
  },
  registerResource: {
    path: '/khulnasoftrpc.ResourceMonitor/RegisterResource',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_resource_pb.RegisterResourceRequest,
    responseType: khulnasoft_resource_pb.RegisterResourceResponse,
    requestSerialize: serialize_khulnasoftrpc_RegisterResourceRequest,
    requestDeserialize: deserialize_khulnasoftrpc_RegisterResourceRequest,
    responseSerialize: serialize_khulnasoftrpc_RegisterResourceResponse,
    responseDeserialize: deserialize_khulnasoftrpc_RegisterResourceResponse,
  },
  registerResourceOutputs: {
    path: '/khulnasoftrpc.ResourceMonitor/RegisterResourceOutputs',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_resource_pb.RegisterResourceOutputsRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_khulnasoftrpc_RegisterResourceOutputsRequest,
    requestDeserialize: deserialize_khulnasoftrpc_RegisterResourceOutputsRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Register a resource transform for the stack
registerStackTransform: {
    path: '/khulnasoftrpc.ResourceMonitor/RegisterStackTransform',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_callback_pb.Callback,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_khulnasoftrpc_Callback,
    requestDeserialize: deserialize_khulnasoftrpc_Callback,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Register an invoke transform for the stack
registerStackInvokeTransform: {
    path: '/khulnasoftrpc.ResourceMonitor/RegisterStackInvokeTransform',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_callback_pb.Callback,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_khulnasoftrpc_Callback,
    requestDeserialize: deserialize_khulnasoftrpc_Callback,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  registerPackage: {
    path: '/khulnasoftrpc.ResourceMonitor/RegisterPackage',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_resource_pb.RegisterPackageRequest,
    responseType: khulnasoft_resource_pb.RegisterPackageResponse,
    requestSerialize: serialize_khulnasoftrpc_RegisterPackageRequest,
    requestDeserialize: deserialize_khulnasoftrpc_RegisterPackageRequest,
    responseSerialize: serialize_khulnasoftrpc_RegisterPackageResponse,
    responseDeserialize: deserialize_khulnasoftrpc_RegisterPackageResponse,
  },
};

exports.ResourceMonitorClient = grpc.makeGenericClientConstructor(ResourceMonitorService);
