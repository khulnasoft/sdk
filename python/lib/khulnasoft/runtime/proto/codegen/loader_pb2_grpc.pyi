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
import grpc
import khulnasoft.codegen.loader_pb2

class LoaderStub:
    """Loader is a service for getting schemas from the Khulnasoft engine for use in code generators and other tools.
    This is currently unstable and experimental.
    """

    def __init__(self, channel: grpc.Channel) -> None: ...
    GetSchema: grpc.UnaryUnaryMultiCallable[
        khulnasoft.codegen.loader_pb2.GetSchemaRequest,
        khulnasoft.codegen.loader_pb2.GetSchemaResponse,
    ]
    """GetSchema tries to find a schema for the given package and version."""

class LoaderServicer(metaclass=abc.ABCMeta):
    """Loader is a service for getting schemas from the Khulnasoft engine for use in code generators and other tools.
    This is currently unstable and experimental.
    """

    @abc.abstractmethod
    def GetSchema(
        self,
        request: khulnasoft.codegen.loader_pb2.GetSchemaRequest,
        context: grpc.ServicerContext,
    ) -> khulnasoft.codegen.loader_pb2.GetSchemaResponse:
        """GetSchema tries to find a schema for the given package and version."""

def add_LoaderServicer_to_server(servicer: LoaderServicer, server: grpc.Server) -> None: ...
