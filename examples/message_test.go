package examples

import (
	"github.com/sivukhin/goprotoslim/examples/def"
	"github.com/sivukhin/goprotoslim/examples/slim"
	"testing"
	"unsafe"
)

func TestMessageSize(t *testing.T) {
	slimSize := unsafe.Sizeof(slim.Message{})
	defSize := unsafe.Sizeof(def.Message{})
	t.Logf("slimSize: %v, defSize: %v", slimSize, defSize)
	if slimSize >= defSize {
		t.Errorf("slimSize >= defSize: %v >= %v", slimSize, defSize)
	}
}

func TestAddressSize(t *testing.T) {
	slimSize := unsafe.Sizeof(slim.Address{})
	defSize := unsafe.Sizeof(def.Address{})
	t.Logf("slimSize: %v, defSize: %v", slimSize, defSize)
	if slimSize >= defSize {
		t.Errorf("slimSize >= defSize: %v >= %v", slimSize, defSize)
	}
}
