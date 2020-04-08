.DEFAULT: protos

protos: dependency generate

dependency:
	sh scripts/download_dep.sh

generate:
	sh scripts/generate_proto.sh

gateway:
	sh scripts/generate_gateway.sh

build_go:
	sh build/build.sh

build_edgelet:
	sh build/build_edgelet.sh
