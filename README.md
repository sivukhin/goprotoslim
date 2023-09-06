# goprotoslim

**goprotoslim** is a Go package that provides a code generator for optimizing Google Protocol Buffers (protobuf) generated Go code. It is designed to be used as a plugin with the `protoc` compiler. **goprotoslim** replaces certain types in the generated Go code with a more efficient `struct{}` representation, reducing memory overhead and improving performance.

## Installation

To use **goprotoslim**, you need to have Go and the Protocol Buffers compiler (protoc) installed on your system. You can install **goprotoslim** using go install:

```bash
go install github.com/sivukhin/goprotoslim/cmd/protoc-gen-go-goprotoslim 
```

## Usage

**goprotoslim** should be used as a `protoc` plugin. You can use it by specifying the `--goprotoslim_out` option instead of `--go_out` option. Also, you can specify additional parameter for **goprotoslim** plugin (now only `types` parameter is supported):

```bash
protoc \
  --goprotoslim_out=. \
  --goprotoslim_opt=paths=source_relative \
  --goprotoslim_opt=types=protoimpl.MessageState+protoimpl.SizeCache \
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

You can find example of generated source code in the [./examples/slim/message.pb.go](./examples/slim/message.pb.go) file.

```bash
$> diff examples/default/message.pb.go examples/slim/message.pb.go 
3c3
< // 	protoc-gen-go v1.28.1
---
> // 	protoc-gen-go v1.28.1-devel
24,26c24,26
< 	state         protoimpl.MessageState
< 	sizeCache     protoimpl.SizeCache
< 	unknownFields protoimpl.UnknownFields
---
> 	state         struct{}
> 	sizeCache     struct{}
> 	unknownFields struct{}
95,97c95,97
< 	state         protoimpl.MessageState
< 	sizeCache     protoimpl.SizeCache
< 	unknownFields protoimpl.UnknownFields
---
> 	state         struct{}
> 	sizeCache     struct{}
> 	unknownFields struct{}

```