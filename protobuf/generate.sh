# 產生 simple protocol buffer 的 golang code
# -I: proto 文件的目錄位置
# --go_out: 產生的 golang code 的位置
# src/simple/simple.proto: 要被產生的 protobuf 的所在地
# 最後會產生 src/simple/simple_pb2.py
protoc -I src/ --go_out=src/ src/simple/simple.proto