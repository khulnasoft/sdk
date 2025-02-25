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
from pathlib import Path

from khulnasoft import CustomResource
from khulnasoft.asset import FileAsset, StringAsset, RemoteAsset


class MyResource(CustomResource):
    def __init__(self, name, asset):
        CustomResource.__init__(
            self, "test:index:MyResource", name, props={"asset": asset}
        )


MyResource("file", FileAsset("./testfile.txt"))
MyResource("file", FileAsset(Path(".") / "testfile.txt"))
MyResource("string", StringAsset("its a string"))
MyResource("remote", RemoteAsset("https://khulnasoft.com"))
