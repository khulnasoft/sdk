# Copyright 2016-2021, KhulnaSoft, Ltd.
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

from typing import Optional
import pytest

from khulnasoft.runtime import settings, mocks
import khulnasoft


@pytest.mark.timeout(10)
@khulnasoft.runtime.test
def test_khulnasoft_does_not_hang_on_dependency_cycle(my_mocks):
    c = MockComponentResource(name="c")
    r = MockResource(name="r", input1=c.output1, opts=khulnasoft.ResourceOptions(parent=c))
    return khulnasoft.Output.all(c.urn, r.urn).apply(print)


class MockResource(khulnasoft.CustomResource):
    def __init__(
        self,
        name: str,
        input1: khulnasoft.Input[str],
        opts: Optional[khulnasoft.ResourceOptions] = None,
    ):
        super().__init__(
            "python:test_resource_dep_cycle:MockResource",
            name,
            {"input1": input1},
            opts,
        )


class MockComponentResource(khulnasoft.ComponentResource):
    output1: khulnasoft.Output[str]

    def __init__(self, name: str, opts: Optional[khulnasoft.ResourceOptions] = None):
        super().__init__(
            "python:test_resource_dep_cycle:MockComponentResource",
            name,
            props=None,
            opts=opts,
            remote=True,
        )
        self.output1 = self.urn
        self.register_outputs({"output1": self.output1})


@pytest.fixture
def my_mocks():
    old_settings = settings.SETTINGS
    mm = MyMocks()
    mocks.set_mocks(mm, preview=True)
    try:
        yield mm
    finally:
        settings.configure(old_settings)


class MyMocks(khulnasoft.runtime.Mocks):
    def new_resource(self, args: khulnasoft.runtime.MockResourceArgs):
        return [args.name + "_id", args.inputs]

    def call(self, args: khulnasoft.runtime.MockCallArgs):
        return {}
