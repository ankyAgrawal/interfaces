# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: system/ack.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
import extensions_pb2 as extensions__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='system/ack.proto',
  package='system',
  syntax='proto3',
  serialized_options=_b('Z*github.com/airmap/interfaces/src/go/system'),
  serialized_pb=_b('\n\x10system/ack.proto\x12\x06system\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x10\x65xtensions.proto\"C\n\x03\x41\x63k\x12\r\n\x05\x63ount\x18\x01 \x01(\x04\x12-\n\tsubmitted\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampB,Z*github.com/airmap/interfaces/src/go/systemb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,extensions__pb2.DESCRIPTOR,])




_ACK = _descriptor.Descriptor(
  name='Ack',
  full_name='system.Ack',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='count', full_name='system.Ack.count', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='submitted', full_name='system.Ack.submitted', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=79,
  serialized_end=146,
)

_ACK.fields_by_name['submitted'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['Ack'] = _ACK
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Ack = _reflection.GeneratedProtocolMessageType('Ack', (_message.Message,), dict(
  DESCRIPTOR = _ACK,
  __module__ = 'system.ack_pb2'
  # @@protoc_insertion_point(class_scope:system.Ack)
  ))
_sym_db.RegisterMessage(Ack)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
