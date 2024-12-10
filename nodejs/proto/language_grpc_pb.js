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
var khulnasoft_language_pb = require('./language_pb.js');
var khulnasoft_codegen_hcl_pb = require('./codegen/hcl_pb.js');
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

function serialize_khulnasoftrpc_AboutRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.AboutRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.AboutRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_AboutRequest(buffer_arg) {
  return khulnasoft_language_pb.AboutRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_AboutResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.AboutResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.AboutResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_AboutResponse(buffer_arg) {
  return khulnasoft_language_pb.AboutResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GeneratePackageRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GeneratePackageRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GeneratePackageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GeneratePackageRequest(buffer_arg) {
  return khulnasoft_language_pb.GeneratePackageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GeneratePackageResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GeneratePackageResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GeneratePackageResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GeneratePackageResponse(buffer_arg) {
  return khulnasoft_language_pb.GeneratePackageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GenerateProgramRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GenerateProgramRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GenerateProgramRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GenerateProgramRequest(buffer_arg) {
  return khulnasoft_language_pb.GenerateProgramRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GenerateProgramResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GenerateProgramResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GenerateProgramResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GenerateProgramResponse(buffer_arg) {
  return khulnasoft_language_pb.GenerateProgramResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GenerateProjectRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GenerateProjectRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GenerateProjectRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GenerateProjectRequest(buffer_arg) {
  return khulnasoft_language_pb.GenerateProjectRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GenerateProjectResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GenerateProjectResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GenerateProjectResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GenerateProjectResponse(buffer_arg) {
  return khulnasoft_language_pb.GenerateProjectResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetProgramDependenciesRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GetProgramDependenciesRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetProgramDependenciesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetProgramDependenciesRequest(buffer_arg) {
  return khulnasoft_language_pb.GetProgramDependenciesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetProgramDependenciesResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GetProgramDependenciesResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetProgramDependenciesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetProgramDependenciesResponse(buffer_arg) {
  return khulnasoft_language_pb.GetProgramDependenciesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetRequiredPackagesRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GetRequiredPackagesRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetRequiredPackagesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetRequiredPackagesRequest(buffer_arg) {
  return khulnasoft_language_pb.GetRequiredPackagesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetRequiredPackagesResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GetRequiredPackagesResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetRequiredPackagesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetRequiredPackagesResponse(buffer_arg) {
  return khulnasoft_language_pb.GetRequiredPackagesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetRequiredPluginsRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GetRequiredPluginsRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetRequiredPluginsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetRequiredPluginsRequest(buffer_arg) {
  return khulnasoft_language_pb.GetRequiredPluginsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_GetRequiredPluginsResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.GetRequiredPluginsResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.GetRequiredPluginsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_GetRequiredPluginsResponse(buffer_arg) {
  return khulnasoft_language_pb.GetRequiredPluginsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_InstallDependenciesRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.InstallDependenciesRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.InstallDependenciesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_InstallDependenciesRequest(buffer_arg) {
  return khulnasoft_language_pb.InstallDependenciesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_InstallDependenciesResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.InstallDependenciesResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.InstallDependenciesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_InstallDependenciesResponse(buffer_arg) {
  return khulnasoft_language_pb.InstallDependenciesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_PackRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.PackRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.PackRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_PackRequest(buffer_arg) {
  return khulnasoft_language_pb.PackRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_PackResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.PackResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.PackResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_PackResponse(buffer_arg) {
  return khulnasoft_language_pb.PackResponse.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_khulnasoftrpc_RunPluginRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.RunPluginRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.RunPluginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RunPluginRequest(buffer_arg) {
  return khulnasoft_language_pb.RunPluginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RunPluginResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.RunPluginResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.RunPluginResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RunPluginResponse(buffer_arg) {
  return khulnasoft_language_pb.RunPluginResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RunRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.RunRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.RunRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RunRequest(buffer_arg) {
  return khulnasoft_language_pb.RunRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RunResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.RunResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.RunResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RunResponse(buffer_arg) {
  return khulnasoft_language_pb.RunResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RuntimeOptionsRequest(arg) {
  if (!(arg instanceof khulnasoft_language_pb.RuntimeOptionsRequest)) {
    throw new Error('Expected argument of type khulnasoftrpc.RuntimeOptionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RuntimeOptionsRequest(buffer_arg) {
  return khulnasoft_language_pb.RuntimeOptionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_khulnasoftrpc_RuntimeOptionsResponse(arg) {
  if (!(arg instanceof khulnasoft_language_pb.RuntimeOptionsResponse)) {
    throw new Error('Expected argument of type khulnasoftrpc.RuntimeOptionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_khulnasoftrpc_RuntimeOptionsResponse(buffer_arg) {
  return khulnasoft_language_pb.RuntimeOptionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// LanguageRuntime is the interface that the planning monitor uses to drive execution of an interpreter responsible
// for confguring and creating resource objects.
var LanguageRuntimeService = exports.LanguageRuntimeService = {
  // GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
getRequiredPlugins: {
    path: '/khulnasoftrpc.LanguageRuntime/GetRequiredPlugins',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.GetRequiredPluginsRequest,
    responseType: khulnasoft_language_pb.GetRequiredPluginsResponse,
    requestSerialize: serialize_khulnasoftrpc_GetRequiredPluginsRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GetRequiredPluginsRequest,
    responseSerialize: serialize_khulnasoftrpc_GetRequiredPluginsResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GetRequiredPluginsResponse,
  },
  // GetRequiredPackages computes the complete set of anticipated packages required by a program.
getRequiredPackages: {
    path: '/khulnasoftrpc.LanguageRuntime/GetRequiredPackages',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.GetRequiredPackagesRequest,
    responseType: khulnasoft_language_pb.GetRequiredPackagesResponse,
    requestSerialize: serialize_khulnasoftrpc_GetRequiredPackagesRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GetRequiredPackagesRequest,
    responseSerialize: serialize_khulnasoftrpc_GetRequiredPackagesResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GetRequiredPackagesResponse,
  },
  // Run executes a program and returns its result.
run: {
    path: '/khulnasoftrpc.LanguageRuntime/Run',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.RunRequest,
    responseType: khulnasoft_language_pb.RunResponse,
    requestSerialize: serialize_khulnasoftrpc_RunRequest,
    requestDeserialize: deserialize_khulnasoftrpc_RunRequest,
    responseSerialize: serialize_khulnasoftrpc_RunResponse,
    responseDeserialize: deserialize_khulnasoftrpc_RunResponse,
  },
  // GetPluginInfo returns generic information about this plugin, like its version.
getPluginInfo: {
    path: '/khulnasoftrpc.LanguageRuntime/GetPluginInfo',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: khulnasoft_plugin_pb.PluginInfo,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_khulnasoftrpc_PluginInfo,
    responseDeserialize: deserialize_khulnasoftrpc_PluginInfo,
  },
  // InstallDependencies will install dependencies for the project, e.g. by running `npm install` for nodejs projects.
installDependencies: {
    path: '/khulnasoftrpc.LanguageRuntime/InstallDependencies',
    requestStream: false,
    responseStream: true,
    requestType: khulnasoft_language_pb.InstallDependenciesRequest,
    responseType: khulnasoft_language_pb.InstallDependenciesResponse,
    requestSerialize: serialize_khulnasoftrpc_InstallDependenciesRequest,
    requestDeserialize: deserialize_khulnasoftrpc_InstallDependenciesRequest,
    responseSerialize: serialize_khulnasoftrpc_InstallDependenciesResponse,
    responseDeserialize: deserialize_khulnasoftrpc_InstallDependenciesResponse,
  },
  // RuntimeOptionsPrompts returns a list of additional prompts to ask during `khulnasoft new`.
runtimeOptionsPrompts: {
    path: '/khulnasoftrpc.LanguageRuntime/RuntimeOptionsPrompts',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.RuntimeOptionsRequest,
    responseType: khulnasoft_language_pb.RuntimeOptionsResponse,
    requestSerialize: serialize_khulnasoftrpc_RuntimeOptionsRequest,
    requestDeserialize: deserialize_khulnasoftrpc_RuntimeOptionsRequest,
    responseSerialize: serialize_khulnasoftrpc_RuntimeOptionsResponse,
    responseDeserialize: deserialize_khulnasoftrpc_RuntimeOptionsResponse,
  },
  // About returns information about the runtime for this language.
about: {
    path: '/khulnasoftrpc.LanguageRuntime/About',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.AboutRequest,
    responseType: khulnasoft_language_pb.AboutResponse,
    requestSerialize: serialize_khulnasoftrpc_AboutRequest,
    requestDeserialize: deserialize_khulnasoftrpc_AboutRequest,
    responseSerialize: serialize_khulnasoftrpc_AboutResponse,
    responseDeserialize: deserialize_khulnasoftrpc_AboutResponse,
  },
  // GetProgramDependencies returns the set of dependencies required by the program.
getProgramDependencies: {
    path: '/khulnasoftrpc.LanguageRuntime/GetProgramDependencies',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.GetProgramDependenciesRequest,
    responseType: khulnasoft_language_pb.GetProgramDependenciesResponse,
    requestSerialize: serialize_khulnasoftrpc_GetProgramDependenciesRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GetProgramDependenciesRequest,
    responseSerialize: serialize_khulnasoftrpc_GetProgramDependenciesResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GetProgramDependenciesResponse,
  },
  // RunPlugin executes a plugin program and returns its result asynchronously.
runPlugin: {
    path: '/khulnasoftrpc.LanguageRuntime/RunPlugin',
    requestStream: false,
    responseStream: true,
    requestType: khulnasoft_language_pb.RunPluginRequest,
    responseType: khulnasoft_language_pb.RunPluginResponse,
    requestSerialize: serialize_khulnasoftrpc_RunPluginRequest,
    requestDeserialize: deserialize_khulnasoftrpc_RunPluginRequest,
    responseSerialize: serialize_khulnasoftrpc_RunPluginResponse,
    responseDeserialize: deserialize_khulnasoftrpc_RunPluginResponse,
  },
  // GenerateProgram generates a given PCL program into a program for this language.
generateProgram: {
    path: '/khulnasoftrpc.LanguageRuntime/GenerateProgram',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.GenerateProgramRequest,
    responseType: khulnasoft_language_pb.GenerateProgramResponse,
    requestSerialize: serialize_khulnasoftrpc_GenerateProgramRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GenerateProgramRequest,
    responseSerialize: serialize_khulnasoftrpc_GenerateProgramResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GenerateProgramResponse,
  },
  // GenerateProject generates a given PCL program into a project for this language.
generateProject: {
    path: '/khulnasoftrpc.LanguageRuntime/GenerateProject',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.GenerateProjectRequest,
    responseType: khulnasoft_language_pb.GenerateProjectResponse,
    requestSerialize: serialize_khulnasoftrpc_GenerateProjectRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GenerateProjectRequest,
    responseSerialize: serialize_khulnasoftrpc_GenerateProjectResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GenerateProjectResponse,
  },
  // GeneratePackage generates a given khulnasoft package into a package for this language.
generatePackage: {
    path: '/khulnasoftrpc.LanguageRuntime/GeneratePackage',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.GeneratePackageRequest,
    responseType: khulnasoft_language_pb.GeneratePackageResponse,
    requestSerialize: serialize_khulnasoftrpc_GeneratePackageRequest,
    requestDeserialize: deserialize_khulnasoftrpc_GeneratePackageRequest,
    responseSerialize: serialize_khulnasoftrpc_GeneratePackageResponse,
    responseDeserialize: deserialize_khulnasoftrpc_GeneratePackageResponse,
  },
  // Pack packs a package into a language specific artifact.
pack: {
    path: '/khulnasoftrpc.LanguageRuntime/Pack',
    requestStream: false,
    responseStream: false,
    requestType: khulnasoft_language_pb.PackRequest,
    responseType: khulnasoft_language_pb.PackResponse,
    requestSerialize: serialize_khulnasoftrpc_PackRequest,
    requestDeserialize: deserialize_khulnasoftrpc_PackRequest,
    responseSerialize: serialize_khulnasoftrpc_PackResponse,
    responseDeserialize: deserialize_khulnasoftrpc_PackResponse,
  },
};

exports.LanguageRuntimeClient = grpc.makeGenericClientConstructor(LanguageRuntimeService);
