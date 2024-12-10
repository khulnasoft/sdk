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

import asyncio
import unittest
import pytest

import khulnasoft
from khulnasoft.runtime import mocks
from khulnasoft import StackReference, StackReferenceOutputDetails


@pytest.mark.asyncio
async def test_stack_reference_output_details(simple_mock):
    ref = StackReference("ref")

    non_secret = await ref.get_output_details("bucket")
    assert StackReferenceOutputDetails(value="mybucket-1234"), non_secret

    secret = await ref.get_output_details("password")
    assert StackReferenceOutputDetails(secret_value="mypassword"), non_secret

    unknown = await ref.get_output_details("does-not-exist")
    assert StackReferenceOutputDetails(), non_secret


@pytest.fixture
def simple_mock():
    mock = StackReferenceOutputMock()
    mocks.set_mocks(mock)
    yield mock


class StackReferenceOutputMock(khulnasoft.runtime.Mocks):
    def new_resource(self, args: khulnasoft.runtime.MockResourceArgs):
        assert "khulnasoft:khulnasoft:StackReference" == args.typ
        return [
            args.name + "_id",
            {
                "outputs": {
                    "bucket": "mybucket-1234",
                    "password": khulnasoft.Output.secret("mypassword"),
                },
            },
        ]

    def call(self, args: khulnasoft.runtime.MockCallArgs):
        return {}
