from concurrent import futures
import logging

import grpc
import rpc.greeting_pb2 as greeting_pb2
import rpc.greeting_pb2_grpc as greeting_pb2_grpc


class GreetingProtoServicer(greeting_pb2_grpc.GreetingProtoServicer):
    def Greet(self, request, context):
        return greeting_pb2.Message(value = request.value);

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  greeting_pb2_grpc.add_GreetingProtoServicer_to_server(
      GreetingProtoServicer(), server)
  server.add_insecure_port('[::]:50051')
  server.start()
  server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()
