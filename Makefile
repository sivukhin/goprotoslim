gen-examples:
	mkdir -p examples/default examples/slim
	cd examples/def && protoc --go_out=. --go_opt=paths=source_relative def.proto
	cd examples/slim && protoc --goprotoslim_out=. --goprotoslim_opt=paths=source_relative slim.proto
