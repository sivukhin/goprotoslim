# goprotoslim

**goprotoslim** is a Go package that provides a code generator for optimizing Google Protocol Buffers (protobuf) generated Go code. It is designed to be used as a plugin with the `protoc` compiler. **goprotoslim** replaces certain types in the generated Go code with a more efficient `struct{}` representation, reducing memory overhead and improving performance.

## Installation

To use **goprotoslim**, you need to have Go and the Protocol Buffers compiler (protoc) installed on your system. You can install **goprotoslim** using go install:

```bash
go install github.com/sivukhin/goprotoslim/cmd/protoc-gen-go-protoslim@latest 
```

## Usage

**goprotoslim** should be used as a `protoc` plugin. You can use it by specifying the `--go-protoslim_out` option instead of `--go_out` option. Also, you can specify additional parameter for **goprotoslim** plugin (now only `types` parameter is supported):

```bash
protoc \
  --go-protoslim_out=. \
  --go-protoslim_opt=paths=source_relative \
  --go-protoslim_opt=types=protoimpl.MessageState+protoimpl.SizeCache \
  contracts/message.proto
```

The `types` optional parameter allows you to specify the types that you want to make slim (replace with `struct{}` for now) with `+` (plus sign) as a delimiter.

By default, **goprotoslim** replaces the following types:

```
protoimpl.MessageState
protoimpl.SizeCache
protoimpl.UnknownFields
```

## Example

You can find example of generated source code in the [./examples/slim/slim.pb.go](./examples/slim/message.pb.go) file.

Couple of tests also written in order to test message size for slim version of generated contracts:
```bash
$> go test -v ./...
=== RUN   TestMessageSize
    message_test.go:13: slimSize: 48, defSize: 88
--- PASS: TestMessageSize (0.00s)
=== RUN   TestAddressSize
    message_test.go:22: slimSize: 64, defSize: 104
--- PASS: TestAddressSize (0.00s)
PASS
ok  	github.com/sivukhin/goprotoslim/examples	0.002s
```