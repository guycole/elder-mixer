#
# Title: lambda_driver.py
# Description: sns listener to gRPC adapter
# Development Environment:OS X 11.0.1/Python 3.8.2
# Lambda Environment: Python 3.9 runtime
#
# import boto3
import logging
import json

import grpc
print(grpc)
import mixer_pb2
import mixer_pb2_grpc

logger = logging.getLogger()
logger.setLevel(logging.INFO)

def grpc_driver():
    logger.info("grpc_driver")

    with grpc.insecure_channel("192.168.170.239:50053") as channel:
        stub = mixer_pb2_grpc.MixerStub(channel)
        response = stub.EnqueueCommand(mixer_pb2.EnqueueRequest(command="test", client_id="client"))
        print("Response: " + response.receipt_id)


def dispatcher(payload: dict):
    logger.info("dispatcher")

    #    logger.info(payload)

    if "Records" in payload:
        records = payload["Records"]
        for record in records:
            message = json.loads(record["Sns"]["Message"])

            page = message["page"]
            values = message["values"]
            visited_at = message["visitedAt"]

            logger.info(message)
            logger.info(page)
            logger.info(values)
            logger.info(visited_at)
        else:
            logger.info("no records")

    grpc_driver()

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
