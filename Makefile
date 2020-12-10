generate:
	rm -rf ./gen/*
	find ${PWD}/mservices -type f -iname "*.proto" -exec \
	protoc --proto_path ${GOPATH}/src --go_out=plugins=grpc:./gen --govalidators_out=.  {} \;
