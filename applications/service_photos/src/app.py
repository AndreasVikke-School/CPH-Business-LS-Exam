from concurrent import futures
import urllib.request
import logging
import json

import grpc
import rpc.photo_pb2 as photo_pb2
import rpc.photo_pb2_grpc as photo_pb2_grpc


class PhotoProtoServicer(photo_pb2_grpc.PhotoProtoServicer):
    def GetPhotoById(self, request, context):
        photo = json.loads(urllib.request.urlopen(
            "https://jsonplaceholder.typicode.com/photos/{0}".format(request.value)).read())

        return photo_pb2.Photo(
            albumId=photo['albumId'],
            id=photo['id'],
            title=photo['title'],
            url=photo['url']
        )

    def GetAllPhotos(self, request, context):
        photos = json.loads(urllib.request.urlopen(
            "https://jsonplaceholder.typicode.com/photos").read())

        return photo_pb2.Photos(photos=[
            photo_pb2.Photo(
                albumId=photo['albumId'],
                id=photo['id'],
                title=photo['title'],
                url=photo['url']
            ) for photo in photos
        ])

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

  photo_pb2_grpc.add_PhotoProtoServicer_to_server(
      PhotoProtoServicer(), server)

  server.add_insecure_port('[::]:50051')
  server.start()
  server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()
