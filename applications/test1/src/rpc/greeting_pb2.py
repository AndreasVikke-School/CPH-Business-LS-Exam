# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: greeting.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='greeting.proto',
  package='greeting',
  syntax='proto3',
  serialized_options=b'\252\002\006Protos',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0egreeting.proto\x12\x08greeting\x1a\x1egoogle/protobuf/wrappers.proto\"\x18\n\x07Message\x12\r\n\x05value\x18\x01 \x01(\t2I\n\rGreetingProto\x12\x38\n\x05Greet\x12\x1c.google.protobuf.StringValue\x1a\x11.greeting.MessageB\t\xaa\x02\x06Protosb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,])




_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='greeting.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='greeting.Message.value', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=60,
  serialized_end=84,
)

DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGE,
  '__module__' : 'greeting_pb2'
  # @@protoc_insertion_point(class_scope:greeting.Message)
  })
_sym_db.RegisterMessage(Message)


DESCRIPTOR._options = None

_GREETINGPROTO = _descriptor.ServiceDescriptor(
  name='GreetingProto',
  full_name='greeting.GreetingProto',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=86,
  serialized_end=159,
  methods=[
  _descriptor.MethodDescriptor(
    name='Greet',
    full_name='greeting.GreetingProto.Greet',
    index=0,
    containing_service=None,
    input_type=google_dot_protobuf_dot_wrappers__pb2._STRINGVALUE,
    output_type=_MESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_GREETINGPROTO)

DESCRIPTOR.services_by_name['GreetingProto'] = _GREETINGPROTO

# @@protoc_insertion_point(module_scope)
