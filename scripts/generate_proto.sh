#!/usr/bin/env bash
#set -x
set -e
cur_script_dir="$(cd $(dirname "$0") && pwd)"
WORK_HOME="${cur_script_dir}/.."
IMPORT_HOME="${WORK_HOME}/../../../"

echo "dirname $WORK_HOME"
echo "IMPORT_HOME: $(cd "$IMPORT_HOME" && pwd)"
echo "WORK_HOME = $(cd "$WORK_HOME" && pwd)"

# swagger bindata generator
# see https://segmentfault.com/a/1190000013513469
# TODO move to docker
#go get -u github.com/go-bindata/go-bindata/...
#go get -u github.com/elazarl/go-bindata-assetfs/...
#mkdir -p swagger

find $WORK_HOME -name "*.proto" | while read proto; do
  dir="$(dirname "$proto")"
  dir="$(cd "$dir" && pwd)"
  # parse file name without directory and suffix
  # parse "./proto/adn.proto" to "adn"
  file_name="${proto##*/}"
  proto_name="${file_name%.*}"
  echo "proto file: $proto dir: ${dir} file_name: $proto_name"
  echo "generating proto..."
  docker run --rm -v "$dir":/defs -v "${IMPORT_HOME}":/input blademainer/protoc-all:latest -i /defs -i /input -i /go/src/ -d /defs/ -l go -o /defs --validate-out "lang=go:/defs" --with-gateway --lint $addition
#  go-bindata --nocompress -pkg swagger -o swagger/${proto_name}/${proto_name}.swagger.go ${dir}/${proto_name}.swagger.json
#  echo "generating gateway..."
#  docker run --rm -v $dir:/defs -v ${IMPORT_HOME}:/input blademainer/gen-grpc-gateway:latest -i /defs -i /input -f /defs/$proto_name -s Service -o ${proto_name}
done

# generage js protos
find $WORK_HOME -name "generage.sh" | while read script; do
  #  dir="`dirname $proto`"
  #  echo "dir: `cd $dir && pwd`"
  #  docker run --rm -v $dir:/defs -v ${IMPORT_HOME}:/input blademainer/protoc-all:1.23_v0.0.3 -i /defs -i /input -d /defs/ -l node -o /defs  --lint $addition;
  sh $script
done
