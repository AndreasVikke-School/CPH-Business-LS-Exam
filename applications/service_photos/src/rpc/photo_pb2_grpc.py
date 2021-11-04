# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
import rpc.photo_pb2 as photo__pb2


class PhotoProtoStub(object):
    """The greeting service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetPhotoById = channel.unary_unary(
                '/photo.PhotoProto/GetPhotoById',
                request_serializer=google_dot_protobuf_dot_wrappers__pb2.Int32Value.SerializeToString,
                response_deserializer=photo__pb2.Photo.FromString,
                )
        self.GetAllPhotos = channel.unary_unary(
                '/photo.PhotoProto/GetAllPhotos',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=photo__pb2.Photos.FromString,
                )


class PhotoProtoServicer(object):
    """The greeting service definition.
    """

    def GetPhotoById(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllPhotos(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PhotoProtoServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetPhotoById': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPhotoById,
                    request_deserializer=google_dot_protobuf_dot_wrappers__pb2.Int32Value.FromString,
                    response_serializer=photo__pb2.Photo.SerializeToString,
            ),
            'GetAllPhotos': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAllPhotos,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=photo__pb2.Photos.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'photo.PhotoProto', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PhotoProto(object):
    """The greeting service definition.
    """

    @staticmethod
    def GetPhotoById(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/photo.PhotoProto/GetPhotoById',
            google_dot_protobuf_dot_wrappers__pb2.Int32Value.SerializeToString,
            photo__pb2.Photo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllPhotos(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/photo.PhotoProto/GetAllPhotos',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            photo__pb2.Photos.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
