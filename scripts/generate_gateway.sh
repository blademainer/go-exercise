#!/usr/bin/env bash
#set -x
set -e
cur_script_dir="$(cd $(dirname $0) && pwd)"
WORK_HOME="${cur_script_dir}/.."
IMPORT_HOME="${WORK_HOME}/../../../"
echo "dirname $WORK_HOME"
echo "IMPORT_HOME: $IMPORT_HOME"
echo "WORK_HOME = $WORK_HOME"
find $WORK_HOME -name "*.proto" | while read proto; do
  dir="$(dirname $proto)"
  dir="$(cd $dir && pwd)"
  # parse file name without directory and suffix
  # parse "./proto/adn.proto" to "adn"
  file_name="${proto##*/}"
  proto_name="${file_name%.*}"
  echo "proto file: $proto dir: ${dir} file_name: $proto_name"
  echo "generating gateway..."
  docker run --rm -v $dir:/defs -v ${IMPORT_HOME}:/input blademainer/gen-grpc-gateway:latest -i /defs -i /input -f /defs/$proto_name -s Service -o ${proto_name}
done

