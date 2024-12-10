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
var khulnasoft_analyzer_pb = require('./analyzer_pb.js');
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

function serialize_khulnasoftrpc_AnalyzeRequest(arg) {
  if (!(arg instanceof khulnasoft_analyzer_pb.AnalyzeRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.AnalyzeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_AnalyzeRequest(buffer_arg) {
  return khulnasoft_analyzer_pb.AnalyzeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_AnalyzeResponse(arg) {
  if (!(arg instanceof khulnasoft_analyzer_pb.AnalyzeResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.AnalyzeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_AnalyzeResponse(buffer_arg) {
  return khulnasoft_analyzer_pb.AnalyzeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_AnalyzeStackRequest(arg) {
  if (!(arg instanceof khulnasoft_analyzer_pb.AnalyzeStackRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.AnalyzeStackRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_AnalyzeStackRequest(buffer_arg) {
  return khulnasoft_analyzer_pb.AnalyzeStackRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_AnalyzerInfo(arg) {
  if (!(arg instanceof khulnasoft_analyzer_pb.AnalyzerInfo)) {
    throw new Error('Expected argument of type khulnasoftrpc.AnalyzerInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_AnalyzerInfo(buffer_arg) {
  return khulnasoft_analyzer_pb.AnalyzerInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ConfigureAnalyzerRequest(arg) {
  if (!(arg instanceof khulnasoft_analyzer_pb.ConfigureAnalyzerRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConfigureAnalyzerRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConfigureAnalyzerRequest(buffer_arg) {
  return khulnasoft_analyzer_pb.ConfigureAnalyzerRequest.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_khulnasoftrpc_RemediateResponse(arg) {
  if (!(arg instanceof khulnasoft_analyzer_pb.RemediateResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.RemediateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RemediateResponse(buffer_arg) {
  return khulnasoft_analyzer_pb.RemediateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Analyzer provides a pluggable interface for checking resource definitions against some number of
// resource policies. It is intentionally open-ended, allowing for implementations that check
// everything from raw resource definitions to entire projects/stacks/snapshots for arbitrary
// issues -- style, policy, correctness, security, and so on.
var AnalyzerService = exports.AnalyzerService = {
  // Analyze analyzes a single resource object, and returns any errors that it finds.
// Called with the "inputs" to the resource, before it is updated.
analyze: {
    path: '/khulnasoftrpc.Analyzer/Analyze',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_analyzer_pb.AnalyzeRequest,
    responseType: khulnasoft_analyzer_pb.AnalyzeResponse,
    requestSerialize: serialize_khulnasoftrpc_AnalyzeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_AnalyzeRequest,
    responseSerialize: serialize_khulnasoftrpc_AnalyzeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_AnalyzeResponse,
  },
  // AnalyzeStack analyzes all resources within a stack, at the end of a successful
// preview or update. The provided resources are the "outputs", after any mutations
// have taken place.
analyzeStack: {
    path: '/khulnasoftrpc.Analyzer/AnalyzeStack',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_analyzer_pb.AnalyzeStackRequest,
    responseType: khulnasoft_analyzer_pb.AnalyzeResponse,
    requestSerialize: serialize_khulnasoftrpc_AnalyzeStackRequest,
    requestDeserialize: deserialize_khulnasoftrpc_AnalyzeStackRequest,
    responseSerialize: serialize_khulnasoftrpc_AnalyzeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_AnalyzeResponse,
  },
  // Remediate optionally transforms a single resource object. This effectively rewrites
// a single resource object's properties instead of using what was generated by the program.
remediate: {
    path: '/khulnasoftrpc.Analyzer/Remediate',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_analyzer_pb.AnalyzeRequest,
    responseType: khulnasoft_analyzer_pb.RemediateResponse,
    requestSerialize: serialize_khulnasoftrpc_AnalyzeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_AnalyzeRequest,
    responseSerialize: serialize_khulnasoftrpc_RemediateResponse,
    responseDeserialize: deserialize_khulnasoftrpc_RemediateResponse,
  },
  // GetAnalyzerInfo returns metadata about the analyzer (e.g., list of policies contained).
getAnalyzerInfo: {
    path: '/khulnasoftrpc.Analyzer/GetAnalyzerInfo',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: khulnasoft_analyzer_pb.AnalyzerInfo,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_khulnasoftrpc_AnalyzerInfo,
    responseDeserialize: deserialize_khulnasoftrpc_AnalyzerInfo,
  },
  // GetPluginInfo returns generic information about this plugin, like its version.
getPluginInfo: {
    path: '/khulnasoftrpc.Analyzer/GetPluginInfo',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: khulnasoft_plugin_pb.PluginInfo,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_khulnasoftrpc_PluginInfo,
    responseDeserialize: deserialize_khulnasoftrpc_PluginInfo,
  },
  // Configure configures the analyzer, passing configuration properties for each policy.
configure: {
    path: '/khulnasoftrpc.Analyzer/Configure',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_analyzer_pb.ConfigureAnalyzerRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_khulnasoftrpc_ConfigureAnalyzerRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ConfigureAnalyzerRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.AnalyzerClient = grpc.makeGenericClientConstructor(AnalyzerService);
