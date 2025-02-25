# Copyright 2016-2020, KhulnaSoft, Ltd.
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

import khulnasoft


@khulnasoft.input_type
class AdditionalArgs:
    first_value: khulnasoft.Input[str] = khulnasoft.property("firstValue")
    second_value: Optional[khulnasoft.Input[float]] = khulnasoft.property("secondValue")

    def __init__(
        self,
        first_value: khulnasoft.Input[str],
        second_value: Optional[khulnasoft.Input[float]] = None,
    ):
        khulnasoft.set(self, "first_value", first_value)
        khulnasoft.set(self, "second_value", second_value)


@khulnasoft.output_type
class Additional(dict):
    first_value: str = khulnasoft.property("firstValue")
    second_value: Optional[float] = khulnasoft.property("secondValue")


class AdditionalResource(khulnasoft.CustomResource):
    additional: khulnasoft.Output[Additional]

    def __init__(self, name: str, additional: khulnasoft.InputType[AdditionalArgs]):
        super().__init__(
            "test:index:AdditionalResource", name, {"additional": additional}
        )


# Create a resource with input object.
res = AdditionalResource(
    "testres", additional=AdditionalArgs(first_value="hello", second_value=42)
)

# Create a resource using the output object of another resource.
res2 = AdditionalResource(
    "testres2",
    additional=AdditionalArgs(
        first_value=res.additional.first_value,
        second_value=res.additional.second_value,
    ),
)

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = AdditionalResource(
    "testres3",
    additional=AdditionalArgs(
        first_value=res.additional["first_value"],
        second_value=res.additional["second_value"],
    ),
)

# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res4 = AdditionalResource(
    "testres4",
    additional={
        "firstValue": "hello",
        "secondValue": 42,
    },
)


# Now, test some resources that use property translations.

SNAKE_TO_CAMEL_CASE_TABLE = {
    "first_value": "firstValue",
    "second_value": "secondValue",
}

CAMEL_TO_SNAKE_CASE_TABLE = {
    "firstValue": "first_value",
    "secondValue": "second_value",
}


@khulnasoft.input_type
class ExtraArgs:
    first_value: khulnasoft.Input[str] = khulnasoft.property("firstValue")
    second_value: Optional[khulnasoft.Input[float]] = khulnasoft.property("secondValue")

    def __init__(
        self,
        first_value: khulnasoft.Input[str],
        second_value: Optional[khulnasoft.Input[float]] = None,
    ):
        khulnasoft.set(self, "first_value", first_value)
        khulnasoft.set(self, "second_value", second_value)


@khulnasoft.output_type
class Extra(dict):
    first_value: str = khulnasoft.property("firstValue")
    second_value: Optional[float] = khulnasoft.property("secondValue")

    def _translate_property(self, prop):
        return CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


class ExtraResource(khulnasoft.CustomResource):
    extra: khulnasoft.Output[Extra]

    def __init__(self, name: str, extra: khulnasoft.InputType[ExtraArgs]):
        super().__init__("test:index:ExtraResource", name, {"extra": extra})

    def translate_output_property(self, prop):
        return CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# Create a resource with input object.
res5 = ExtraResource("testres5", extra=ExtraArgs(first_value="foo", second_value=100))

# Create a resource using the output object of another resource.
res6 = ExtraResource(
    "testres6",
    extra=ExtraArgs(
        first_value=res5.extra.first_value,
        second_value=res5.extra.second_value,
    ),
)

# Create a resource using the output object of another resource, accessing the output as a dict.
# Note: the output dict's keys are translated keys.
res7 = ExtraResource(
    "testres7",
    extra=ExtraArgs(
        first_value=res5.extra["first_value"],
        second_value=res5.extra["second_value"],
    ),
)

# Create a resource using a dict as the input.
# Note: these are specified as snake_case, and the resource will translate to camelCase.
res8 = ExtraResource(
    "testres8",
    extra={
        "first_value": res5.extra["first_value"],
        "second_value": res5.extra["second_value"],
    },
)


# Now test some resources that use explicitly declared properties.


@khulnasoft.input_type
class SupplementaryArgs:
    def __init__(
        self,
        first_value: khulnasoft.Input[str],
        second_value: Optional[khulnasoft.Input[float]] = None,
        third: Optional[khulnasoft.Input[str]] = None,
        fourth: Optional[khulnasoft.Input[str]] = None,
    ):
        khulnasoft.set(self, "first_value", first_value)
        khulnasoft.set(self, "second_value", second_value)
        khulnasoft.set(self, "third", third)
        khulnasoft.set(self, "fourth", fourth)

    # Property with empty getter/setter bodies.
    @property
    @khulnasoft.getter(name="firstValue")
    def first_value(self) -> khulnasoft.Input[str]: ...  # type: ignore

    @first_value.setter
    def first_value(self, value: khulnasoft.Input[str]):
        khulnasoft.set(self, "first_value", value)

    # Property with explicitly specified getter/setter bodies.
    @property
    @khulnasoft.getter(name="secondValue")
    def second_value(self) -> Optional[khulnasoft.Input[float]]:
        return khulnasoft.get(self, "second_value")

    @second_value.setter
    def second_value(self, value: Optional[khulnasoft.Input[float]]):
        khulnasoft.set(self, "second_value", value)

    # Single word property name that doesn't require a name to be
    # passed to the getter decorator.
    @property
    @khulnasoft.getter
    def third(self) -> Optional[khulnasoft.Input[str]]: ...  # type: ignore

    @third.setter
    def third(self, value: Optional[khulnasoft.Input[str]]): ...  # type: ignore

    # Another single word property name that doesn't require a name to be
    # passed to the getter decorator, this time using the decorator with
    # parens.
    @property
    @khulnasoft.getter()
    def fourth(self) -> Optional[khulnasoft.Input[str]]: ...

    @fourth.setter
    def fourth(self, value: Optional[khulnasoft.Input[str]]): ...


@khulnasoft.output_type
class Supplementary(dict):
    def __init__(
        self, first_value: str, second_value: Optional[float], third: str, fourth: str
    ):
        khulnasoft.set(self, "first_value", first_value)
        khulnasoft.set(self, "second_value", second_value)
        khulnasoft.set(self, "third", third)
        khulnasoft.set(self, "fourth", fourth)

    # Property with empty getter/setter bodies.
    @property
    @khulnasoft.getter(name="firstValue")
    def first_value(self) -> str: ...  # type: ignore

    # Property with explicitly specified getter/setter bodies.
    @property
    @khulnasoft.getter(name="secondValue")
    def second_value(self) -> Optional[float]:
        return khulnasoft.get(self, "second_value")

    # Single word property name that doesn't require a name to be
    # passed to the getter decorator.
    @property
    @khulnasoft.getter
    def third(self) -> str: ...  # type: ignore

    # Another single word property name that doesn't require a name to be
    # passed to the getter decorator, this time using the decorator with
    # parens.
    @property
    @khulnasoft.getter
    def fourth(self) -> str: ...  # type: ignore


class SupplementaryResource(khulnasoft.CustomResource):
    supplementary: khulnasoft.Output[Supplementary]

    def __init__(self, name: str, supplementary: khulnasoft.InputType[SupplementaryArgs]):
        super().__init__(
            "test:index:SupplementaryResource", name, {"supplementary": supplementary}
        )


# Create a resource with input object.
res9 = SupplementaryResource(
    "testres9",
    supplementary=SupplementaryArgs(
        first_value="bar",
        second_value=200,
        third="third value",
        fourth="fourth value",
    ),
)

# Create a resource using the output object of another resource.
res10 = SupplementaryResource(
    "testres10",
    supplementary=SupplementaryArgs(
        first_value=res9.supplementary.first_value,
        second_value=res9.supplementary.second_value,
        third=res9.supplementary.third,
        fourth=res9.supplementary.fourth,
    ),
)

# Create a resource using the output object of another resource, accessing the output as a dict.
res11 = SupplementaryResource(
    "testres11",
    supplementary=SupplementaryArgs(
        first_value=res9.supplementary["first_value"],
        second_value=res9.supplementary["second_value"],
        third=res9.supplementary["third"],
        fourth=res9.supplementary["fourth"],
    ),
)

# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res12 = SupplementaryResource(
    "testres12",
    supplementary={
        "firstValue": "bar",
        "secondValue": 200,
        "third": "third value",
        "fourth": "fourth value",
    },
)


# Now, test some resources that use property translations and explicitly declared properties.


@khulnasoft.input_type
class AncillaryArgs:
    def __init__(
        self,
        first_value: khulnasoft.Input[str],
        second_value: Optional[khulnasoft.Input[float]] = None,
        third: Optional[khulnasoft.Input[str]] = None,
        fourth: Optional[khulnasoft.Input[str]] = None,
    ):
        khulnasoft.set(self, "first_value", first_value)
        khulnasoft.set(self, "second_value", second_value)
        khulnasoft.set(self, "third", third)
        khulnasoft.set(self, "fourth", fourth)

    # Property with empty getter/setter bodies.
    @property
    @khulnasoft.getter(name="firstValue")
    def first_value(self) -> khulnasoft.Input[str]: ...  # type: ignore

    @first_value.setter
    def first_value(self, value: khulnasoft.Input[str]):
        khulnasoft.set(self, "first_value", value)

    # Property with explicitly specified getter/setter bodies.
    @property
    @khulnasoft.getter(name="secondValue")
    def second_value(self) -> Optional[khulnasoft.Input[float]]:
        return khulnasoft.get(self, "second_value")

    @second_value.setter
    def second_value(self, value: Optional[khulnasoft.Input[float]]):
        khulnasoft.set(self, "second_value", value)

    # Single word property name that doesn't require a name to be
    # passed to the getter decorator.
    @property
    @khulnasoft.getter
    def third(self) -> Optional[khulnasoft.Input[str]]: ...  # type: ignore

    @third.setter
    def third(self, value: Optional[khulnasoft.Input[str]]): ...  # type: ignore

    # Another single word property name that doesn't require a name to be
    # passed to the getter decorator, this time using the decorator with
    # parens.
    @property
    @khulnasoft.getter()
    def fourth(self) -> Optional[khulnasoft.Input[str]]: ...  # type: ignore

    @fourth.setter
    def fourth(self, value: Optional[khulnasoft.Input[str]]): ...  # type: ignore


@khulnasoft.output_type
class Ancillary(dict):
    def __init__(
        self, first_value: str, second_value: Optional[float], third: str, fourth: str
    ):
        khulnasoft.set(self, "first_value", first_value)
        khulnasoft.set(self, "second_value", second_value)
        khulnasoft.set(self, "third", third)
        khulnasoft.set(self, "fourth", fourth)

    # Property with empty getter/setter bodies.
    @property
    @khulnasoft.getter(name="firstValue")
    def first_value(self) -> str: ...  # type: ignore

    # Property with explicitly specified getter/setter bodies.
    @property
    @khulnasoft.getter(name="secondValue")
    def second_value(self) -> Optional[float]:
        return khulnasoft.get(self, "second_value")

    # Single word property name that doesn't require a name to be
    # passed to the getter decorator.
    @property
    @khulnasoft.getter
    def third(self) -> str: ...  # type: ignore

    # Another single word property name that doesn't require a name to be
    # passed to the getter decorator, this time using the decorator with
    # parens.
    @property
    @khulnasoft.getter()
    def fourth(self) -> str: ...  # type: ignore

    def _translate_property(self, prop):
        return CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


class AncillaryResource(khulnasoft.CustomResource):
    ancillary: khulnasoft.Output[Ancillary]

    def __init__(self, name: str, ancillary: khulnasoft.InputType[AncillaryArgs]):
        super().__init__("test:index:AncillaryResource", name, {"ancillary": ancillary})

    def translate_output_property(self, prop):
        return CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# Create a resource with input object.
res13 = AncillaryResource(
    "testres13",
    ancillary=AncillaryArgs(
        first_value="baz",
        second_value=500,
        third="third value!",
        fourth="fourth!",
    ),
)

# Create a resource using the output object of another resource.
res14 = AncillaryResource(
    "testres14",
    ancillary=AncillaryArgs(
        first_value=res13.ancillary.first_value,
        second_value=res13.ancillary.second_value,
        third=res13.ancillary.third,
        fourth=res13.ancillary.fourth,
    ),
)

# Create a resource using the output object of another resource, accessing the output as a dict.
# Note: the output dict's keys are translated keys.
res15 = AncillaryResource(
    "testres15",
    ancillary=AncillaryArgs(
        first_value=res13.ancillary["first_value"],
        second_value=res13.ancillary["second_value"],
        third=res13.ancillary["third"],
        fourth=res13.ancillary["fourth"],
    ),
)

# Create a resource using a dict as the input.
# Note: these are specified as snake_case, and the resource will translate to camelCase.
res16 = AncillaryResource(
    "testres16",
    ancillary={
        "first_value": "baz",
        "second_value": 500,
        "third": "third value!",
        "fourth": "fourth!",
    },
)
