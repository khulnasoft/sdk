"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
Copyright 2016-2023, KhulnaSoft, Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""
import abc
import collections.abc
import google.protobuf.empty_pb2
import grpc
import grpc.aio
import typing
import khulnasoft.language_pb2
import khulnasoft.plugin_pb2

class LanguageRuntimeStub:
    """LanguageRuntime is the interface that the planning monitor uses to drive execution of an interpreter responsible
    for confguring and creating resource objects.
    """

    def __init__(self, channel: grpc.Channel) -> None: ...
    GetRequiredPlugins: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.GetRequiredPluginsRequest,
        khulnasoft.language_pb2.GetRequiredPluginsResponse,
    ]
    """GetRequiredPlugins computes the complete set of anticipated plugins required by a program."""
    GetRequiredPackages: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.GetRequiredPackagesRequest,
        khulnasoft.language_pb2.GetRequiredPackagesResponse,
    ]
    """GetRequiredPackages computes the complete set of anticipated packages required by a program."""
    Run: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.RunRequest,
        khulnasoft.language_pb2.RunResponse,
    ]
    """Run executes a program and returns its result."""
    GetPluginInfo: grpc.UnaryUnaryMultiCallable[
        google.protobuf.empty_pb2.Empty,
        khulnasoft.plugin_pb2.PluginInfo,
    ]
    """GetPluginInfo returns generic information about this plugin, like its version."""
    InstallDependencies: grpc.UnaryStreamMultiCallable[
        khulnasoft.language_pb2.InstallDependenciesRequest,
        khulnasoft.language_pb2.InstallDependenciesResponse,
    ]
    """InstallDependencies will install dependencies for the project, e.g. by running `npm install` for nodejs projects."""
    RuntimeOptionsPrompts: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.RuntimeOptionsRequest,
        khulnasoft.language_pb2.RuntimeOptionsResponse,
    ]
    """RuntimeOptionsPrompts returns a list of additional prompts to ask during `khulnasoft new`."""
    About: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.AboutRequest,
        khulnasoft.language_pb2.AboutResponse,
    ]
    """About returns information about the runtime for this language."""
    GetProgramDependencies: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.GetProgramDependenciesRequest,
        khulnasoft.language_pb2.GetProgramDependenciesResponse,
    ]
    """GetProgramDependencies returns the set of dependencies required by the program."""
    RunPlugin: grpc.UnaryStreamMultiCallable[
        khulnasoft.language_pb2.RunPluginRequest,
        khulnasoft.language_pb2.RunPluginResponse,
    ]
    """RunPlugin executes a plugin program and returns its result asynchronously."""
    GenerateProgram: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.GenerateProgramRequest,
        khulnasoft.language_pb2.GenerateProgramResponse,
    ]
    """GenerateProgram generates a given PCL program into a program for this language."""
    GenerateProject: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.GenerateProjectRequest,
        khulnasoft.language_pb2.GenerateProjectResponse,
    ]
    """GenerateProject generates a given PCL program into a project for this language."""
    GeneratePackage: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.GeneratePackageRequest,
        khulnasoft.language_pb2.GeneratePackageResponse,
    ]
    """GeneratePackage generates a given khulnasoft package into a package for this language."""
    Pack: grpc.UnaryUnaryMultiCallable[
        khulnasoft.language_pb2.PackRequest,
        khulnasoft.language_pb2.PackResponse,
    ]
    """Pack packs a package into a language specific artifact."""

class LanguageRuntimeServicer(metaclass=abc.ABCMeta):
    """LanguageRuntime is the interface that the planning monitor uses to drive execution of an interpreter responsible
    for confguring and creating resource objects.
    """

    
    def GetRequiredPlugins(
        self,
        request: khulnasoft.language_pb2.GetRequiredPluginsRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.GetRequiredPluginsResponse:
        """GetRequiredPlugins computes the complete set of anticipated plugins required by a program."""
    
    def GetRequiredPackages(
        self,
        request: khulnasoft.language_pb2.GetRequiredPackagesRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.GetRequiredPackagesResponse:
        """GetRequiredPackages computes the complete set of anticipated packages required by a program."""
    
    def Run(
        self,
        request: khulnasoft.language_pb2.RunRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.RunResponse:
        """Run executes a program and returns its result."""
    
    def GetPluginInfo(
        self,
        request: google.protobuf.empty_pb2.Empty,
        context: grpc.ServicerContext,
    ) -> khulnasoft.plugin_pb2.PluginInfo:
        """GetPluginInfo returns generic information about this plugin, like its version."""
    
    def InstallDependencies(
        self,
        request: khulnasoft.language_pb2.InstallDependenciesRequest,
        context: grpc.ServicerContext,
    ) -> collections.abc.Iterator[khulnasoft.language_pb2.InstallDependenciesResponse]:
        """InstallDependencies will install dependencies for the project, e.g. by running `npm install` for nodejs projects."""
    
    def RuntimeOptionsPrompts(
        self,
        request: khulnasoft.language_pb2.RuntimeOptionsRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.RuntimeOptionsResponse:
        """RuntimeOptionsPrompts returns a list of additional prompts to ask during `khulnasoft new`."""
    
    def About(
        self,
        request: khulnasoft.language_pb2.AboutRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.AboutResponse:
        """About returns information about the runtime for this language."""
    
    def GetProgramDependencies(
        self,
        request: khulnasoft.language_pb2.GetProgramDependenciesRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.GetProgramDependenciesResponse:
        """GetProgramDependencies returns the set of dependencies required by the program."""
    
    def RunPlugin(
        self,
        request: khulnasoft.language_pb2.RunPluginRequest,
        context: grpc.ServicerContext,
    ) -> collections.abc.Iterator[khulnasoft.language_pb2.RunPluginResponse]:
        """RunPlugin executes a plugin program and returns its result asynchronously."""
    
    def GenerateProgram(
        self,
        request: khulnasoft.language_pb2.GenerateProgramRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.GenerateProgramResponse:
        """GenerateProgram generates a given PCL program into a program for this language."""
    
    def GenerateProject(
        self,
        request: khulnasoft.language_pb2.GenerateProjectRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.GenerateProjectResponse:
        """GenerateProject generates a given PCL program into a project for this language."""
    
    def GeneratePackage(
        self,
        request: khulnasoft.language_pb2.GeneratePackageRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.GeneratePackageResponse:
        """GeneratePackage generates a given khulnasoft package into a package for this language."""
    
    def Pack(
        self,
        request: khulnasoft.language_pb2.PackRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.language_pb2.PackResponse:
        """Pack packs a package into a language specific artifact."""

def add_LanguageRuntimeServicer_to_server(servicer: LanguageRuntimeServicer, server: typing.Union[grpc.Server, grpc.aio.Server]) -> None: ...
