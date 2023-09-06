gen-examples:
	mkdir -p examples/default examples/slim
	cd examples/def && protoc --go_out=. --go_opt=paths=source_relative def.proto
	cd examples/slim && protoc --go-protoslim_out=. --go-protoslim_opt=paths=source_relative slim.proto
