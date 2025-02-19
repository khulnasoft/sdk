# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import khulnasoft
import khulnasoft.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'DataArgs',
    'DataArgsDict',
]

MYPY = False

if not MYPY:
    class DataArgsDict(TypedDict):
        private: khulnasoft.Input[str]
        public: khulnasoft.Input[str]
elif False:
    DataArgsDict: TypeAlias = Mapping[str, Any]

@khulnasoft.input_type
class DataArgs:
    def __init__(__self__, *,
                 private: khulnasoft.Input[str],
                 public: khulnasoft.Input[str]):
        khulnasoft.set(__self__, "private", private)
        khulnasoft.set(__self__, "public", public)

    @property
    @khulnasoft.getter
    def private(self) -> khulnasoft.Input[str]:
        return khulnasoft.get(self, "private")

    @private.setter
    def private(self, value: khulnasoft.Input[str]):
        khulnasoft.set(self, "private", value)

    @property
    @khulnasoft.getter
    def public(self) -> khulnasoft.Input[str]:
        return khulnasoft.get(self, "public")

    @public.setter
    def public(self, value: khulnasoft.Input[str]):
        khulnasoft.set(self, "public", value)


