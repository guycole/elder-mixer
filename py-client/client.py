import logging
import random

import grpc
import mixer_pb2
import mixer_pb2_grpc

def execute():
    with grpc.insecure_channel("localhost:50053") as channel:
        stub = mixer_pb2_grpc.MixerStub(channel)
        response = stub.EnqueueCommand(mixer_pb2.EnqueueRequest(command="test", client_id="client"))
        print("Response: " + response.receipt_id)

if __name__ == "__main__":
    logging.basicConfig()
    execute()
