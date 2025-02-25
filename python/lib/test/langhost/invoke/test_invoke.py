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
from os import path
from ..util import LanghostTest


class TestInvoke(LanghostTest):
    def test_invoke_success(self):
        self.run_test(
            program=path.join(self.base_path(), "invoke"), expected_resource_count=3
        )

    def invoke(self, _ctx, token, args, provider, _version):
        self.assertEqual("test:index:MyFunction", token)
        self.assertEqual("", provider)
        self.assertDictEqual(
            {
                "value": 41,
                "value2": 42,
            },
            args,
        )

        return [], {"value": args["value"] + 1}

    def register_resource(
        self,
        _ctx,
        _dry_run,
        ty,
        name,
        _resource,
        _dependencies,
        _parent,
        _custom,
        protect,
        _provider,
        _property_deps,
        _delete_before_replace,
        _ignore_changes,
        _version,
        _import,
        _replace_on_changes,
        _providers,
        source_position,
    ):
        self.assertEqual("test:index:MyResource", ty)
        self.assertEqual(_resource["value"], 42)

        return {
            "urn": self.make_urn(ty, name),
            "id": name,
            "object": _resource,
        }


class TestInvokeWithFailures(LanghostTest):
    def test_invoke_failure(self):
        self.run_test(
            program=path.join(self.base_path(), "invoke"),
            expected_resource_count=0,
            expected_bail=True,
        )

    def invoke(self, _ctx, token, args, _provider, _version):
        self.assertEqual("test:index:MyFunction", token)
        self.assertDictEqual(
            {
                "value": 41,
                "value2": 42,
            },
            args,
        )

        return [{"property": "value", "reason": "the invoke failed"}], {}
