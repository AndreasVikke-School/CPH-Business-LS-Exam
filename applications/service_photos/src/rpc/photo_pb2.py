# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: photo.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='photo.proto',
  package='photo',
  syntax='proto3',
  serialized_options=b'\252\002\006Protos',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0bphoto.proto\x12\x05photo\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1bgoogle/protobuf/empty.proto\"&\n\x06Photos\x12\x1c\n\x06photos\x18\x01 \x03(\x0b\x32\x0c.photo.Photo\"@\n\x05Photo\x12\x0f\n\x07\x61lbumId\x18\x01 \x01(\x05\x12\n\n\x02id\x18\x02 \x01(\x05\x12\r\n\x05title\x18\x03 \x01(\t\x12\x0b\n\x03url\x18\x04 \x01(\t2~\n\nPhotoProto\x12\x39\n\x0cGetPhotoById\x12\x1b.google.protobuf.Int32Value\x1a\x0c.photo.Photo\x12\x35\n\x0cGetAllPhotos\x12\x16.google.protobuf.Empty\x1a\r.photo.PhotosB\t\xaa\x02\x06Protosb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])




_PHOTOS = _descriptor.Descriptor(
  name='Photos',
  full_name='photo.Photos',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='photos', full_name='photo.Photos.photos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=83,
  serialized_end=121,
)


_PHOTO = _descriptor.Descriptor(
  name='Photo',
  full_name='photo.Photo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='albumId', full_name='photo.Photo.albumId', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='photo.Photo.id', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='title', full_name='photo.Photo.title', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='url', full_name='photo.Photo.url', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=123,
  serialized_end=187,
)

_PHOTOS.fields_by_name['photos'].message_type = _PHOTO
DESCRIPTOR.message_types_by_name['Photos'] = _PHOTOS
DESCRIPTOR.message_types_by_name['Photo'] = _PHOTO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Photos = _reflection.GeneratedProtocolMessageType('Photos', (_message.Message,), {
  'DESCRIPTOR' : _PHOTOS,
  '__module__' : 'photo_pb2'
  # @@protoc_insertion_point(class_scope:photo.Photos)
  })
_sym_db.RegisterMessage(Photos)

Photo = _reflection.GeneratedProtocolMessageType('Photo', (_message.Message,), {
  'DESCRIPTOR' : _PHOTO,
  '__module__' : 'photo_pb2'
  # @@protoc_insertion_point(class_scope:photo.Photo)
  })
_sym_db.RegisterMessage(Photo)


DESCRIPTOR._options = None

_PHOTOPROTO = _descriptor.ServiceDescriptor(
  name='PhotoProto',
  full_name='photo.PhotoProto',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=189,
  serialized_end=315,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetPhotoById',
    full_name='photo.PhotoProto.GetPhotoById',
    index=0,
    containing_service=None,
    input_type=google_dot_protobuf_dot_wrappers__pb2._INT32VALUE,
    output_type=_PHOTO,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetAllPhotos',
    full_name='photo.PhotoProto.GetAllPhotos',
    index=1,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=_PHOTOS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_PHOTOPROTO)

DESCRIPTOR.services_by_name['PhotoProto'] = _PHOTOPROTO

# @@protoc_insertion_point(module_scope)
