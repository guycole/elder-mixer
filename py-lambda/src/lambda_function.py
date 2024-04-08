#
# Title: lambda_driver.py
# Description: sns listener to gRPC adapter
# Lambda Environment: Python 3.8 runtime
#
import logging
import os

import grpc
import mixer_pb2
import mixer_pb2_grpc

logger = logging.getLogger()
logger.setLevel(logging.INFO)

def grpc_driver(server_address):
    logger.info("grpc_driver")

    target = server_address + ":50053"

    with grpc.insecure_channel(target) as channel:
        stub = mixer_pb2_grpc.MixerStub(channel)
        response = stub.EnqueueCommand(mixer_pb2.EnqueueRequest(command="test", client_id="client"))
        print("Response: " + response.receipt_id)


def dispatcher(payload: dict):
    logger.info("dispatcher")

    server_address = os.environ.get("SERVER_ADDRESS")
    if server_address is None:
        logger.error("SERVER_ADDRESS not set")
        return

    grpc_driver(server_address)

def lambda_handler(event, context):

    logger.info("lambda_handler")

    dispatcher(event)

if __name__ == "__main__":
    print("main noted")

    logger = logging.getLogger()
    #    logger.setLevel(Personality.get_logging_level())

    dispatcher({})

# ;;; Local Variables: ***
# ;;; mode:python ***
# ;;; End: ***
