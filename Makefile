gen-examples:
	mkdir -p examples/default examples/slim
	cd examples && protoc --go_out=default --go_opt=paths=source_relative message.proto
	cd examples && protoc --goprotoslim_out=slim --goprotoslim_opt=paths=source_relative message.proto
