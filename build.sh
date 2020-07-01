#!/bin/bash
export SRC_DIR=${PWD}/destroyer/proto DST_DIR=${PWD}/destroyer/proto/pb
mkdir $DST_DIR
protoc -I=${SRC_DIR} --go_out=${DST_DIR} ${SRC_DIR}/destroyer.proto
