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
from khulnasoft import CustomResource, ProviderResource, ResourceOptions, Output


class MyResource(CustomResource):
    value: Output[int]

    def __init__(self, name, value, opts=None):
        CustomResource.__init__(
            self,
            "test:index:MyResource",
            name,
            props={
                "value": value,
            },
            opts=opts,
        )


class MyProvider(ProviderResource):
    def __init__(self, name, opts=None):
        ProviderResource.__init__(self, "test", name, {}, opts)


prov = MyProvider("testprov")
res = MyResource("res", 42, ResourceOptions(provider=prov))
