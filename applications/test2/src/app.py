from flask import Flask
app = Flask(__name__)

import google.protobuf.wrappers_pb2 as wrappers

import grpc
import rpc.greeting_pb2 as greeting_pb2
import rpc.greeting_pb2_grpc as greeting_pb2_grpc

@app.route('/')
def hello_world():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = greeting_pb2_grpc.GreetingProtoStub(channel)
        return stub.Greet(wrappers.StringValue(value='Hello World!')).value
