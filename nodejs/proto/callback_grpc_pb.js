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
var khulnasoft_callback_pb = require('./callback_pb.js');

function serialize_khulnasoftrpc_CallbackInvokeRequest(arg) {
  if (!(arg instanceof khulnasoft_callback_pb.CallbackInvokeRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.CallbackInvokeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_CallbackInvokeRequest(buffer_arg) {
  return khulnasoft_callback_pb.CallbackInvokeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_CallbackInvokeResponse(arg) {
  if (!(arg instanceof khulnasoft_callback_pb.CallbackInvokeResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.CallbackInvokeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_CallbackInvokeResponse(buffer_arg) {
  return khulnasoft_callback_pb.CallbackInvokeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Callbacks is a service for invoking functions in one runtime from other processes.
var CallbacksService = exports.CallbacksService = {
  // Invoke invokes a given callback, identified by its token.
invoke: {
    path: '/khulnasoftrpc.Callbacks/Invoke',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_callback_pb.CallbackInvokeRequest,
    responseType: khulnasoft_callback_pb.CallbackInvokeResponse,
    requestSerialize: serialize_khulnasoftrpc_CallbackInvokeRequest,
    requestDeserialize: deserialize_khulnasoftrpc_CallbackInvokeRequest,
    responseSerialize: serialize_khulnasoftrpc_CallbackInvokeResponse,
    responseDeserialize: deserialize_khulnasoftrpc_CallbackInvokeResponse,
  },
};

exports.CallbacksClient = grpc.makeGenericClientConstructor(CallbacksService);
