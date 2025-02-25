# Copyright 2016-2018, KhulnaSoft, Ltd.
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
from os import path
from ..util import LanghostTest


class StackOutputTest(LanghostTest):
    """
    Test that tests Khulnasoft's ability to register resource outputs.
    """

    def test_stack_outputs(self):
        self.run_test(
            program=path.join(self.base_path(), "stack_output"),
            expected_resource_count=0,
        )

    def register_resource_outputs(
        self, _ctx, _dry_run, _urn, ty, _name, _resource, outputs
    ):
        self.assertEqual(ty, "khulnasoft:khulnasoft:Stack")
        self.assertDictEqual(
            {
                "string": "khulnasoft",
                "number": 1.0,
                "boolean": True,
                "list": [],
                "list_with_none": [None],
                "list_of_lists": [[], []],
                "list_of_outputs": [[1], [2]],
                "set": ["val"],
                "dict": {"a": 1.0},
                "output": 1.0,
                "class": {"num": 1.0},
                "recursive": {"a": 1.0, "b": 2.0},
                "duplicate_output_0": {"num": 1.0},
                "duplicate_output_1": {"num": 1.0},
            },
            outputs,
        )
