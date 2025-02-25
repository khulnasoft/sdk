# Copyright 2016-2023, KhulnaSoft, Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import pytest

from typing import Awaitable, Optional, Union
from concurrent.futures import ThreadPoolExecutor
import asyncio
import time

import khulnasoft
from khulnasoft import InvokeOptions, InvokeOutputOptions


class MyMocks(khulnasoft.runtime.Mocks):
    def new_resource(self, args: khulnasoft.runtime.MockResourceArgs):
        return [args.name + "_id", args.inputs]

    def call(self, args: khulnasoft.runtime.MockCallArgs):
        return {} if args.args.get("empty") else {"result": "mock"}


class MockResource(khulnasoft.CustomResource):
    def __init__(self, name: str, opts: Optional[khulnasoft.ResourceOptions] = None):
        super().__init__("python:test:MockResource", name, {}, opts)


@pytest.mark.parametrize(
    "tok,version,empty,expected",
    [
        ("test:index:MyFunction", "", True, {}),
        ("test:index:MyFunction", "invalid", True, {}),
        ("test:index:MyFunction", "1.0.0", True, {}),
        ("test:index:MyFunction", "", False, {"result": "mock"}),
        ("test:index:MyFunction", "invalid", False, {"result": "mock"}),
        ("test:index:MyFunction", "1.0.0", False, {"result": "mock"}),
        ("kubernetes:something:new", "", True, {}),
        ("kubernetes:something:new", "invalid", True, {}),
        ("kubernetes:something:new", "1.0.0", True, {}),
        ("kubernetes:something:new", "4.5.3", True, {}),
        ("kubernetes:something:new", "4.5.4", True, {}),
        ("kubernetes:something:new", "4.5.5", True, {}),
        ("kubernetes:something:new", "4.6.0", True, {}),
        ("kubernetes:something:new", "5.0.0", True, {}),
        ("kubernetes:something:new", "", False, {"result": "mock"}),
        ("kubernetes:something:new", "invalid", False, {"result": "mock"}),
        ("kubernetes:something:new", "1.0.0", False, {"result": "mock"}),
        # Expect the legacy behavior of getting None for empty results for these Kubernetes
        # function tokens when the version is unspecified, invalid, or <= 4.5.4.
        ("kubernetes:yaml:decode", "", True, None),
        ("kubernetes:yaml:decode", "invalid", True, None),
        ("kubernetes:yaml:decode", "1.0.0", True, None),
        ("kubernetes:yaml:decode", "4.5.3", True, None),
        ("kubernetes:yaml:decode", "4.5.4", True, None),
        ("kubernetes:yaml:decode", "4.5.5", True, {}),
        ("kubernetes:yaml:decode", "4.6.0", True, {}),
        ("kubernetes:yaml:decode", "5.0.0", True, {}),
        ("kubernetes:yaml:decode", "", False, {"result": "mock"}),
        ("kubernetes:yaml:decode", "invalid", False, {"result": "mock"}),
        ("kubernetes:yaml:decode", "1.0.0", False, {"result": "mock"}),
        ("kubernetes:helm:template", "", True, None),
        ("kubernetes:helm:template", "invalid", True, None),
        ("kubernetes:helm:template", "1.0.0", True, None),
        ("kubernetes:helm:template", "4.5.3", True, None),
        ("kubernetes:helm:template", "4.5.4", True, None),
        ("kubernetes:helm:template", "4.5.5", True, {}),
        ("kubernetes:helm:template", "4.6.0", True, {}),
        ("kubernetes:helm:template", "5.0.0", True, {}),
        ("kubernetes:helm:template", "", False, {"result": "mock"}),
        ("kubernetes:helm:template", "invalid", False, {"result": "mock"}),
        ("kubernetes:helm:template", "1.0.0", False, {"result": "mock"}),
        ("kubernetes:kustomize:directory", "", True, None),
        ("kubernetes:kustomize:directory", "invalid", True, None),
        ("kubernetes:kustomize:directory", "1.0.0", True, None),
        ("kubernetes:kustomize:directory", "4.5.3", True, None),
        ("kubernetes:kustomize:directory", "4.5.4", True, None),
        ("kubernetes:kustomize:directory", "4.5.5", True, {}),
        ("kubernetes:kustomize:directory", "4.6.0", True, {}),
        ("kubernetes:kustomize:directory", "5.0.0", True, {}),
        ("kubernetes:kustomize:directory", "", False, {"result": "mock"}),
        ("kubernetes:kustomize:directory", "invalid", False, {"result": "mock"}),
        ("kubernetes:kustomize:directory", "1.0.0", False, {"result": "mock"}),
    ],
)
@khulnasoft.runtime.test
def test_invoke_empty_return(tok: str, version: str, empty: bool, expected) -> None:
    khulnasoft.runtime.mocks.set_mocks(MyMocks())

    props = {"empty": True} if empty else {}
    opts = khulnasoft.InvokeOptions(version=version) if version else None
    assert khulnasoft.runtime.invoke(tok, props, opts).value == expected


@khulnasoft.runtime.test
async def test_invoke_depends_on() -> None:
    khulnasoft.runtime.mocks.set_mocks(MyMocks())

    dep: khulnasoft.Resource = MockResource(name="dep1")
    dep_future = asyncio.Future()
    dep_output = khulnasoft.Output.from_input(dep_future)

    # This function will sleep for 1 second and then resolve dep_future with the
    # resource, which in turn resolves the output.
    def resolve():
        time.sleep(1)
        dep_future.set_result(dep)

    # Run the resolve function in a separate thread.
    with ThreadPoolExecutor(max_workers=2) as executor:
        f = executor.submit(resolve)

        assert not dep_future.done(), "dep_future should not be done yet"

        opts = khulnasoft.InvokeOutputOptions(depends_on=[dep_output])
        o = khulnasoft.runtime.invoke_output("test::MyFunction", {}, opts)

        def check(invoke_result):
            assert (
                dep_future.done()
            ), "invoke_output should wait for dep_future to be done"
            assert invoke_result == {"result": "mock"}

        return o.apply(check)


@pytest.mark.parametrize(
    "a,b,expected",
    [
        (InvokeOptions(version="1.0.0"), InvokeOptions(version="2.0.0"), "2.0.0"),
        (None, InvokeOptions(version="2.0.0"), "2.0.0"),
        (InvokeOptions(version="1.0.0"), None, "1.0.0"),
        (InvokeOptions(version="1.0.0"), InvokeOptions(version=""), ""),
        (InvokeOutputOptions(version="1.0.0"), InvokeOptions(version="2.0.0"), "2.0.0"),
        (
            InvokeOutputOptions(version="1.0.0"),
            InvokeOutputOptions(version="2.0.0"),
            "2.0.0",
        ),
        (None, InvokeOutputOptions(version="2.0.0"), "2.0.0"),
        (InvokeOutputOptions(version="1.0.0"), None, "1.0.0"),
    ],
)
def test_invoke_merge(
    a: Union[InvokeOptions, InvokeOutputOptions],
    b: Union[InvokeOptions, InvokeOutputOptions],
    expected: str,
) -> None:
    if a is not None:
        assert a.merge(b).version == expected
    if isinstance(a, InvokeOutputOptions):
        assert InvokeOutputOptions.merge(a, b).version == expected
    else:
        assert InvokeOptions.merge(a, b).version == expected
