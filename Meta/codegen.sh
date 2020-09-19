#!/bin/bash

# 服务器golang代码
protoc --go_out=plugins=grpc:../Server/ -I ../Meta/  ../Meta/*.proto

# 客户端python代码
python -m grpc_tools.protoc -I ../Meta/ --python_out=../Client/ --grpc_python_out=../Client/ ../Meta/*.proto