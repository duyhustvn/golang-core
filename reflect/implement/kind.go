package implement

import (
	"fmt"
	"strconv"
	"unsafe"
)

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint

const (
	//       1 << 5 - 1 = 31 = 0001.1111
	// There are 27 Type. 27 = 0001.1011
	kindMask = (1 << 5) - 1
)

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

type tflag uint8
type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype
type textOff int32 // offset from top of text section

// rtype must be kept in sync with ../runtime/type.go:/^type._type
type rtype struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      tflag   // extra type information flags
	align      uint8   // aligment of variable for this type
	fieldAlign uint8   // aligment of struct field with this type

	kind uint8

	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> == ?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte   // garbage collection data
	str       nameOff // string form
	ptrToThis typeOff // type for pointer to this type, maybe zero
}

func (t *rtype) Kind() Kind {
	//       1 << 5 - 1 = 31 = 0001.1111
	// There are 27 Type. 27 = 0001.1011
	return Kind(t.kind & kindMask)
}

func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}
	return "kind" + strconv.Itoa(int(k))
}

var kindNames = []string{
	Invalid:       "invalid",
	Bool:          "bool",
	Int:           "int",
	Int8:          "int8",
	Int16:         "int16",
	Int32:         "int32",
	Int64:         "int64",
	Uint:          "uint",
	Uint8:         "uint8",
	Uint16:        "uint16",
	Uint32:        "uint32",
	Uint64:        "uint64",
	Uintptr:       "uintptr",
	Float32:       "float32",
	Float64:       "float64",
	Complex64:     "complex64",
	Complex128:    "complex128",
	Array:         "array",
	Chan:          "chan",
	Func:          "func",
	Interface:     "interface",
	Map:           "map",
	Pointer:       "ptr",
	Slice:         "slice",
	String:        "string",
	Struct:        "struct",
	UnsafePointer: "unsafe.Pointer",
}

func ImplementKind(r rtype) {
	// r := rtype{kind: 2}

	fmt.Println(r.Kind())
}
