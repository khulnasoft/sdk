# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import khulnasoft
import khulnasoft.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'DataArgs',
    'InnerDataArgs',
]

@khulnasoft.input_type
class DataArgs:
    def __init__(__self__, *,
                 bool_array: Sequence[bool],
                 boolean: bool,
                 float: float,
                 inner_data: 'InnerDataArgs',
                 integer: int,
                 string: str,
                 string_map: Mapping[str, str]):
        khulnasoft.set(__self__, "bool_array", bool_array)
        khulnasoft.set(__self__, "boolean", boolean)
        khulnasoft.set(__self__, "float", float)
        khulnasoft.set(__self__, "inner_data", inner_data)
        khulnasoft.set(__self__, "integer", integer)
        khulnasoft.set(__self__, "string", string)
        khulnasoft.set(__self__, "string_map", string_map)

    @property
    @khulnasoft.getter(name="boolArray")
    def bool_array(self) -> Sequence[bool]:
        return khulnasoft.get(self, "bool_array")

    @bool_array.setter
    def bool_array(self, value: Sequence[bool]):
        khulnasoft.set(self, "bool_array", value)

    @property
    @khulnasoft.getter
    def boolean(self) -> bool:
        return khulnasoft.get(self, "boolean")

    @boolean.setter
    def boolean(self, value: bool):
        khulnasoft.set(self, "boolean", value)

    @property
    @khulnasoft.getter
    def float(self) -> float:
        return khulnasoft.get(self, "float")

    @float.setter
    def float(self, value: float):
        khulnasoft.set(self, "float", value)

    @property
    @khulnasoft.getter(name="innerData")
    def inner_data(self) -> 'InnerDataArgs':
        return khulnasoft.get(self, "inner_data")

    @inner_data.setter
    def inner_data(self, value: 'InnerDataArgs'):
        khulnasoft.set(self, "inner_data", value)

    @property
    @khulnasoft.getter
    def integer(self) -> int:
        return khulnasoft.get(self, "integer")

    @integer.setter
    def integer(self, value: int):
        khulnasoft.set(self, "integer", value)

    @property
    @khulnasoft.getter
    def string(self) -> str:
        return khulnasoft.get(self, "string")

    @string.setter
    def string(self, value: str):
        khulnasoft.set(self, "string", value)

    @property
    @khulnasoft.getter(name="stringMap")
    def string_map(self) -> Mapping[str, str]:
        return khulnasoft.get(self, "string_map")

    @string_map.setter
    def string_map(self, value: Mapping[str, str]):
        khulnasoft.set(self, "string_map", value)


@khulnasoft.input_type
class InnerDataArgs:
    def __init__(__self__, *,
                 bool_array: Sequence[bool],
                 boolean: bool,
                 float: float,
                 integer: int,
                 string: str,
                 string_map: Mapping[str, str]):
        khulnasoft.set(__self__, "bool_array", bool_array)
        khulnasoft.set(__self__, "boolean", boolean)
        khulnasoft.set(__self__, "float", float)
        khulnasoft.set(__self__, "integer", integer)
        khulnasoft.set(__self__, "string", string)
        khulnasoft.set(__self__, "string_map", string_map)

    @property
    @khulnasoft.getter(name="boolArray")
    def bool_array(self) -> Sequence[bool]:
        return khulnasoft.get(self, "bool_array")

    @bool_array.setter
    def bool_array(self, value: Sequence[bool]):
        khulnasoft.set(self, "bool_array", value)

    @property
    @khulnasoft.getter
    def boolean(self) -> bool:
        return khulnasoft.get(self, "boolean")

    @boolean.setter
    def boolean(self, value: bool):
        khulnasoft.set(self, "boolean", value)

    @property
    @khulnasoft.getter
    def float(self) -> float:
        return khulnasoft.get(self, "float")

    @float.setter
    def float(self, value: float):
        khulnasoft.set(self, "float", value)

    @property
    @khulnasoft.getter
    def integer(self) -> int:
        return khulnasoft.get(self, "integer")

    @integer.setter
    def integer(self, value: int):
        khulnasoft.set(self, "integer", value)

    @property
    @khulnasoft.getter
    def string(self) -> str:
        return khulnasoft.get(self, "string")

    @string.setter
    def string(self, value: str):
        khulnasoft.set(self, "string", value)

    @property
    @khulnasoft.getter(name="stringMap")
    def string_map(self) -> Mapping[str, str]:
        return khulnasoft.get(self, "string_map")

    @string_map.setter
    def string_map(self, value: Mapping[str, str]):
        khulnasoft.set(self, "string_map", value)


