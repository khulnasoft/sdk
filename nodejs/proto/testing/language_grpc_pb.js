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
var khulnasoft_testing_language_pb = require('../testing/language_pb.js');

function serialize_khulnasoftrpc_testing_GetLanguageTestsRequest(arg) {
  if (!(arg instanceof khulnasoft_testing_language_pb.GetLanguageTestsRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.testing.GetLanguageTestsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_testing_GetLanguageTestsRequest(buffer_arg) {
  return khulnasoft_testing_language_pb.GetLanguageTestsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_testing_GetLanguageTestsResponse(arg) {
  if (!(arg instanceof khulnasoft_testing_language_pb.GetLanguageTestsResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.testing.GetLanguageTestsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_testing_GetLanguageTestsResponse(buffer_arg) {
  return khulnasoft_testing_language_pb.GetLanguageTestsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_testing_PrepareLanguageTestsRequest(arg) {
  if (!(arg instanceof khulnasoft_testing_language_pb.PrepareLanguageTestsRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.testing.PrepareLanguageTestsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_testing_PrepareLanguageTestsRequest(buffer_arg) {
  return khulnasoft_testing_language_pb.PrepareLanguageTestsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_testing_PrepareLanguageTestsResponse(arg) {
  if (!(arg instanceof khulnasoft_testing_language_pb.PrepareLanguageTestsResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.testing.PrepareLanguageTestsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_testing_PrepareLanguageTestsResponse(buffer_arg) {
  return khulnasoft_testing_language_pb.PrepareLanguageTestsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_testing_RunLanguageTestRequest(arg) {
  if (!(arg instanceof khulnasoft_testing_language_pb.RunLanguageTestRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.testing.RunLanguageTestRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_testing_RunLanguageTestRequest(buffer_arg) {
  return khulnasoft_testing_language_pb.RunLanguageTestRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_testing_RunLanguageTestResponse(arg) {
  if (!(arg instanceof khulnasoft_testing_language_pb.RunLanguageTestResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.testing.RunLanguageTestResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_testing_RunLanguageTestResponse(buffer_arg) {
  return khulnasoft_testing_language_pb.RunLanguageTestResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// LanguageTest is the interface to the khulnasoft language test framework. This is _highly_ experimental and
// currently subject to breaking changes without warning.
var LanguageTestService = exports.LanguageTestService = {
  // GetLanguageTests returns a list of all the language tests.
getLanguageTests: {
    path: '/khulnasoftrpc.testing.LanguageTest/GetLanguageTests',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_testing_language_pb.GetLanguageTestsRequest,
    responseType: khulnasoft_testing_language_pb.GetLanguageTestsResponse,
    requestSerialize: serialize_khulnasoftrpc_testing_GetLanguageTestsRequest,
    requestDeserialize: deserialize_khulnasoftrpc_testing_GetLanguageTestsRequest,
    responseSerialize: serialize_khulnasoftrpc_testing_GetLanguageTestsResponse,
    responseDeserialize: deserialize_khulnasoftrpc_testing_GetLanguageTestsResponse,
  },
  // PrepareLanguageTests prepares the engine to run language tests. It sets up a stable artifacts folder
// (which should be .gitignore'd) and fills it with the core SDK artifact.
prepareLanguageTests: {
    path: '/khulnasoftrpc.testing.LanguageTest/PrepareLanguageTests',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_testing_language_pb.PrepareLanguageTestsRequest,
    responseType: khulnasoft_testing_language_pb.PrepareLanguageTestsResponse,
    requestSerialize: serialize_khulnasoftrpc_testing_PrepareLanguageTestsRequest,
    requestDeserialize: deserialize_khulnasoftrpc_testing_PrepareLanguageTestsRequest,
    responseSerialize: serialize_khulnasoftrpc_testing_PrepareLanguageTestsResponse,
    responseDeserialize: deserialize_khulnasoftrpc_testing_PrepareLanguageTestsResponse,
  },
  // RunLanguageTest runs a single test of the language plugin.
runLanguageTest: {
    path: '/khulnasoftrpc.testing.LanguageTest/RunLanguageTest',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_testing_language_pb.RunLanguageTestRequest,
    responseType: khulnasoft_testing_language_pb.RunLanguageTestResponse,
    requestSerialize: serialize_khulnasoftrpc_testing_RunLanguageTestRequest,
    requestDeserialize: deserialize_khulnasoftrpc_testing_RunLanguageTestRequest,
    responseSerialize: serialize_khulnasoftrpc_testing_RunLanguageTestResponse,
    responseDeserialize: deserialize_khulnasoftrpc_testing_RunLanguageTestResponse,
  },
};

exports.LanguageTestClient = grpc.makeGenericClientConstructor(LanguageTestService);
