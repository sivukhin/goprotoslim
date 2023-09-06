package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"os"
	"slices"
	"strings"
)

type SlimTypes map[string]struct{}

var (
	protoslimName      = "goprotoslim"
	supportedSlimTypes = []string{
		"protoimpl.MessageState",
		"protoimpl.SizeCache",
		"protoimpl.UnknownFields",
	}
)

func parseGlobalSlimTypes(typesString string) (SlimTypes, error) {
	slimTypes := make(map[string]struct{})
	if typesString != "" {
		typeNames := strings.Split(typesString, "+")
		for _, typeName := range typeNames {
			if !slices.Contains(supportedSlimTypes, typeName) {
				return nil, fmt.Errorf("unsupported slim type name '%v', supported types %v", typeName, supportedSlimTypes)
			}
			slimTypes[typeName] = struct{}{}
		}
	}
	if len(slimTypes) == 0 {
		for _, slimType := range supportedSlimTypes {
			slimTypes[slimType] = struct{}{}
		}
	}
	return slimTypes, nil
}

func replaceSlimTypes(source string, slimTypes SlimTypes) string {
	replaced := source
	for slimType := range slimTypes {
		replaced = strings.ReplaceAll(replaced, slimType, "struct{}")
	}
	return replaced
}

func gen(slimTypesArg *string, options protogen.Options) error {
	if len(os.Args) > 1 {
		return fmt.Errorf("unknown argument %q (this program should be run by protoc, not directly)", os.Args[1])
	}
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %w", err)
	}
	request := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(stdin, request); err != nil {
		return err
	}

	plugin, err := options.New(request)
	if err != nil {
		return fmt.Errorf("unable to parse options: %w", err)
	}

	for _, file := range plugin.Files {
		if file.Generate {
			internal_gengo.GenerateFile(plugin, file)
		}
	}

	// todo (sivukhin, 2023-09-07): add file-level option like "goprotoslim.types"?
	globalSlimTypes, err := parseGlobalSlimTypes(*slimTypesArg)
	if err != nil {
		return fmt.Errorf("invalid %v options: %w", protoslimName, err)
	}

	response := plugin.Response()
	for _, generated := range response.File {
		replaced := replaceSlimTypes(*generated.Content, globalSlimTypes)
		generated.Content = &replaced
	}
	out, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(out)
	return err
}

func main() {
	var flags flag.FlagSet
	var slimTypes string
	flags.StringVar(&slimTypes, "types", "", "slim types to replace")
	err := gen(&slimTypes, protogen.Options{ParamFunc: flags.Set})
	if err != nil {
		panic(err)
	}
}
