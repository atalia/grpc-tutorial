#!/bin/bash

# 服务器golang代码
protoc --go_out=plugins=grpc:../Server/ -I ../Meta/  ../Meta/*.proto

# 客户端python代码
protoc --python_out=../Client/ -I ../Meta/ ../Meta/*.proto