# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: khulnasoft/errors.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13khulnasoft/errors.proto\x12\tkhulnasoftrpc\"1\n\nErrorCause\x12\x0f\n\x07message\x18\x01 \x01(\t\x12\x12\n\nstackTrace\x18\x02 \x01(\t\"\x8d\x01\n\x14InputPropertiesError\x12=\n\x06\x65rrors\x18\x01 \x03(\x0b\x32-.khulnasoftrpc.InputPropertiesError.PropertyError\x1a\x36\n\rPropertyError\x12\x15\n\rproperty_path\x18\x01 \x01(\t\x12\x0e\n\x06reason\x18\x02 \x01(\tB4Z2github.com/khulnasoft/sdk/proto/go;khulnasoftrpcb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'khulnasoft.errors_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z2github.com/khulnasoft/sdk/proto/go;khulnasoftrpc'
  _ERRORCAUSE._serialized_start=34
  _ERRORCAUSE._serialized_end=83
  _INPUTPROPERTIESERROR._serialized_start=86
  _INPUTPROPERTIESERROR._serialized_end=227
  _INPUTPROPERTIESERROR_PROPERTYERROR._serialized_start=173
  _INPUTPROPERTIESERROR_PROPERTYERROR._serialized_end=227
# @@protoc_insertion_point(module_scope)
