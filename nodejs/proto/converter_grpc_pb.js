// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2016-2023, KhulnaSoft, Ltd.
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
var khulnasoft_converter_pb = require('./converter_pb.js');
var khulnasoft_codegen_hcl_pb = require('./codegen/hcl_pb.js');

function serialize_khulnasoftrpc_ConvertProgramRequest(arg) {
  if (!(arg instanceof khulnasoft_converter_pb.ConvertProgramRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConvertProgramRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConvertProgramRequest(buffer_arg) {
  return khulnasoft_converter_pb.ConvertProgramRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ConvertProgramResponse(arg) {
  if (!(arg instanceof khulnasoft_converter_pb.ConvertProgramResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConvertProgramResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConvertProgramResponse(buffer_arg) {
  return khulnasoft_converter_pb.ConvertProgramResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ConvertStateRequest(arg) {
  if (!(arg instanceof khulnasoft_converter_pb.ConvertStateRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConvertStateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConvertStateRequest(buffer_arg) {
  return khulnasoft_converter_pb.ConvertStateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_ConvertStateResponse(arg) {
  if (!(arg instanceof khulnasoft_converter_pb.ConvertStateResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.ConvertStateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_ConvertStateResponse(buffer_arg) {
  return khulnasoft_converter_pb.ConvertStateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Converter is a service for converting between other ecosystems and Khulnasoft.
// This is currently unstable and experimental.
var ConverterService = exports.ConverterService = {
  // ConvertState converts state from the target ecosystem into a form that can be imported into Khulnasoft.
convertState: {
    path: '/khulnasoftrpc.Converter/ConvertState',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_converter_pb.ConvertStateRequest,
    responseType: khulnasoft_converter_pb.ConvertStateResponse,
    requestSerialize: serialize_khulnasoftrpc_ConvertStateRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ConvertStateRequest,
    responseSerialize: serialize_khulnasoftrpc_ConvertStateResponse,
    responseDeserialize: deserialize_khulnasoftrpc_ConvertStateResponse,
  },
  // ConvertProgram converts a program from the target ecosystem into a form that can be used with Khulnasoft.
convertProgram: {
    path: '/khulnasoftrpc.Converter/ConvertProgram',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_converter_pb.ConvertProgramRequest,
    responseType: khulnasoft_converter_pb.ConvertProgramResponse,
    requestSerialize: serialize_khulnasoftrpc_ConvertProgramRequest,
    requestDeserialize: deserialize_khulnasoftrpc_ConvertProgramRequest,
    responseSerialize: serialize_khulnasoftrpc_ConvertProgramResponse,
    responseDeserialize: deserialize_khulnasoftrpc_ConvertProgramResponse,
  },
};

exports.ConverterClient = grpc.makeGenericClientConstructor(ConverterService);
