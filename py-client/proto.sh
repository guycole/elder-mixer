#
PATH=/bin:/usr/bin:/etc:/usr/local/bin; export PATH
#
rm mixer_pb2.py
rm mixer_pb2.pyi
rm mixer_pb2_grpc.py
#
source venv/bin/activate
python -m grpc_tools.protoc -I../proto --python_out=. --pyi_out=. --grpc_python_out=. ../proto/mixer.proto
#