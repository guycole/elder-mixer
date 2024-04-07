from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class EnqueueRequest(_message.Message):
    __slots__ = ("client_id", "command")
    CLIENT_ID_FIELD_NUMBER: _ClassVar[int]
    COMMAND_FIELD_NUMBER: _ClassVar[int]
    client_id: str
    command: str
    def __init__(self, client_id: _Optional[str] = ..., command: _Optional[str] = ...) -> None: ...

class EnqueueResponse(_message.Message):
    __slots__ = ("client_id", "receipt_id")
    CLIENT_ID_FIELD_NUMBER: _ClassVar[int]
    RECEIPT_ID_FIELD_NUMBER: _ClassVar[int]
    client_id: str
    receipt_id: str
    def __init__(self, client_id: _Optional[str] = ..., receipt_id: _Optional[str] = ...) -> None: ...
