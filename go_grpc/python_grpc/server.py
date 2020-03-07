#!/usr/bin/env python

from concurrent import futures
import time

import grpc

import helloworld_pb2

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class Greeter(helloworld_pb2.GreeterServicer):

  def SayHello(self, request, context):
    return helloworld_pb2.HelloReply(message='Hello, %s!' % request.name)


def serve():
  #server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  #helloworld_pb2.add_GreeterServicer_to_server(Greeter(), server)
  server = helloworld_pb2.beta_create_Greeter_server(Greeter())
  server.add_insecure_port('[::]:50051')
  server.start()
  try:
    while True:
      time.sleep(_ONE_DAY_IN_SECONDS)
  except KeyboardInterrupt:
    server.stop(0)

if __name__ == '__main__':
  serve()
