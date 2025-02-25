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

import khulnasoft


@khulnasoft.output_type
class MyResourceFoo(dict):
    def __init__(self, *, bar: str, baz: str):
        khulnasoft.set(self, "bar", bar)
        khulnasoft.set(self, "baz", baz)

    @property
    @khulnasoft.getter
    def bar(self) -> str:
        return khulnasoft.get(self, "bar")

    @property
    @khulnasoft.getter
    def baz(self) -> str:
        return khulnasoft.get(self, "baz")
