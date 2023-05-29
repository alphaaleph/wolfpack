#!/usr/bin/bash

# oss=(linux darwin windows android)
# archs=(amd64 arm64 "386")
oss=(linux windows)
archs=(amd64)

# shellcheck disable=SC2068
for os in ${oss[@]}
do
  for arch in ${archs[@]}
  do
	  env GOOS=${os} GOARCH=${arch} go build -o ./bin/wolfpack_${os}-${arch}
	done
done