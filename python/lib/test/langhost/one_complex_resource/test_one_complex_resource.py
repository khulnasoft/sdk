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
import unittest
from ..util import LanghostTest


class OneComplexResourceTest(LanghostTest):
    def test_one_complex_resource(self):
        self.run_test(
            program=path.join(self.base_path(), "one_complex_resource"),
            expected_resource_count=1,
        )

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
        self.assertEqual(ty, "test:index:MyResource")
        self.assertEqual(name, "testres")
        self.assertEqual(_resource["falseprop"], False)
        self.assertEqual(_resource["trueprop"], True)
        self.assertEqual(_resource["intprop"], 42)
        self.assertListEqual(_resource["listprop"], [1, 2, "string", False])
        self.assertDictEqual(_resource["mapprop"], {"foo": ["bar", "baz"]})

        return {
            "urn": self.make_urn(ty, name),
            "id": name,
            "object": {
                "outprop": "output properties ftw",
                "outintprop": 99,
            },
        }
