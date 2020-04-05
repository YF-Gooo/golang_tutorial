#!/usr/bin/env python

from __future__ import print_function

import grpc
import sys

import helloworld_pb2,helloworld_pb2_grpc


def run():
  if len(sys.argv) >= 2:
      grpc_server = sys.argv[1]
  else:
      grpc_server = 'localhost:50051'
  channel = grpc.insecure_channel(grpc_server)
  stub = helloworld_pb2_grpc.GreeterStub(channel)
  response = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))
  print("Greeter client received: " + response.message)


if __name__ == '__main__':
  run()