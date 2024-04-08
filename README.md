# elder-mixer
AWS lambda to gRPC demonstration

## Python lambda invokes gRPC server on EC2 instance.

1.  Lambda and gRPC server must route (i.e. same VPC, etc)
    1.  Security group must allow gRPC (in this case port 50053)
1.  LambdaBasicExecution (i.e. CloudWatch logs) is enough after installation.
1.  Developed on AWS EC2 (ubuntu server 20.04 LTS) 
    1.  python 3.8 runtime

## Repository tour

1.  proto (directory)
    1. gRPC protobuf definitions
    1. Invoke "proto.sh" to regenerate
1.  server (directory)
    1. gRPC server, "go build" to create
    1. Invoke "run.sh" to run (use a dedicated terminal for the server)
1.  go-client (directory)
    1. gRPC client in go, configured for localhost
    1. Invoke "go run" for diagnostic (with server from previous step)
1.  py-client (directory)
    1. gRPC client in python, configured for localhost
    1. Uses virtualenv
    1. Invoke "proto.sh" to generate gRPC stubs
    1. Invoke "python client.py" to run
1.  py-lambda (directory)
    1. gRPC client as a AWS python lambda
    1. Uses virtualenv
    1. Invoke "lambda-build.sh" to create deployment zip
        1. Two zip files: grpc-layer.zip and elder-mixer.zip
        1. Deploys from s3 (because gRPC libraries are large)
        1. Upload both zip files to s3 bucket 
    1. From AWS console
        1. Create python lambda function "elder-mixer" using python 3.8 runtime
            1. Select VPC, subnet and security group to match server EC2 instance
        1. Upload elder-mixer.zip as lambda function (from s3 bucket)
        1. Upload grpc-layer.zip as lambda layer (from s3 bucket)
        1. Configure elder-mixer to use grpc-layer
        1. Invoke "test" on elder-mixer lambda, the "hello, world" test is OK
            1. The lambda does not care, but need an event which causes gRPC write to server
        1. Success means the lambda ran without errors and the server wrote a log message
