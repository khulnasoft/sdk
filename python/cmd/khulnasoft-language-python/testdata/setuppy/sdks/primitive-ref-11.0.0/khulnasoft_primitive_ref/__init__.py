# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .resource import *
from ._inputs import *
from . import outputs
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "primitive-ref",
  "mod": "index",
  "fqn": "khulnasoft_primitive_ref",
  "classes": {
   "primitive-ref:index:Resource": "Resource"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "primitive-ref",
  "token": "khulnasoft:providers:primitive-ref",
  "fqn": "khulnasoft_primitive_ref",
  "class": "Provider"
 }
]
"""
)
