#!/bin/bash
set -ex
protoc -I . config.proto --go_out=plugins=grpc:.
