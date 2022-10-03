// https://go101.org/article/reflection.html
package main

import (
	"fmt"
	"reflect"
)

type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}

func (f F) M() {}

type I interface {
	m(s string) bool
	M()
}

func reflectStruct() {
	var x struct {
		F F
		i I
	}

	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())        // struct
	fmt.Println(tx.NumField())    // number of field in struct 2
	fmt.Println(tx.Field(1).Name) // name of field 1 is i

	// Package path is an intrinsic property of
	// non-exported selectors (fields or methods).
	fmt.Println(tx.Field(0).PkgPath) // no value
	fmt.Println(tx.Field(1).PkgPath) // main

	tf, ti := tx.Field(0).Type, tx.Field(1).Type
	fmt.Println(tf.Kind())       // function
	fmt.Println(ti.Kind())       // interface
	fmt.Println(tf.IsVariadic()) // false
	fmt.Println(tf.NumIn(), tf.NumOut())

	t0, t1, t2 := tf.In(0), tf.In(1), tf.Out(0)
	fmt.Println(t0.Kind(), t1.Kind(), t2.Kind()) // string int bool

	fmt.Println(tf.NumMethod(), ti.NumMethod()) // 1 2
	fmt.Println(tf.Method(0).Name)              // M
	fmt.Println(ti.Method(0).Name)              // M
	fmt.Println(ti.Method(1).Name)              // m

	_, ok1 := tf.MethodByName("m")
	_, ok2 := ti.MethodByName("m")
	fmt.Println(ok1, ok2) // false true

	type T struct {
		X int  `max:"99" min:"0" default:"0"`
		Y bool `optional:"yes"`
		Z bool `optional:"yes,omitempty"`
	}

	t := reflect.TypeOf(T{})
	f1 := t.Field(0).Tag
	f2 := t.Field(1).Tag
	f3 := t.Field(2).Tag
	fmt.Println(f1) // max:"99" min:"0" default:"0"
	fmt.Println(f2) // optional: "yes"
	fmt.Println(f3) // optional: "yes,omitempty"

	fmt.Println(reflect.TypeOf(f1))     // reflect.StructTag
	fmt.Println(f1.Get("max"))          // 99
	fmt.Println(f1.Get("not_found"))    // empty if not found
	fmt.Println(f1.Lookup("max"))       // 99 true
	fmt.Println(f1.Lookup("not_found")) // empty false
	fmt.Println(f2.Lookup("optional"))  // yes true
	fmt.Println(f3.Lookup("optional"))  // yes,omitempty true
}

func reflectType() {
	type A = [16]int16
	var c <-chan map[A][]byte

	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())

	tm := tc.Elem() // channel of map -> tm is map

	ta, tb := tm.Key(), tm.Elem() // tm is map of A (array of int16) to slice of byte -> ta: array tb: slice
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind())

	tx, ty := ta.Elem(), tb.Elem()    // ta is array of int 16 -> ta element is int -> tx: int
	fmt.Println(tx.Kind(), ty.Kind()) // tb is slice of byte -> tb element is byte(uint8) tb: uint8
}

func InvertSlice(args []reflect.Value) []reflect.Value {
	inSlice, n := args[0], args[0].Len()
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, n)

	for i := n - 1; i >= 0; i-- {
		element := inSlice.Index(i)
		outSlice = reflect.Append(outSlice, element)
	}

	return []reflect.Value{outSlice}
}

func Bind(p interface{}, f func([]reflect.Value) []reflect.Value) {
	invert := reflect.ValueOf(p).Elem()
	invert.Set(reflect.MakeFunc(invert.Type(), f))
}

type T1 struct {
	A, b int
}

func (t T1) AddSubThenScale(n int, m int) (int, int, int) {
	return n * (t.A + t.b), n * (t.A - t.b), m
}

func reflectValue() {
	n := 123
	p := &n

	vp := reflect.ValueOf(p)
	fmt.Println(vp.CanSet(), vp.CanAddr()) // false false

	vn := vp.Elem()
	fmt.Println(vn.CanSet(), vn.CanAddr()) // true true

	vn.Set(reflect.ValueOf(789))
	fmt.Println(n) // 789

	var s struct {
		X interface{} // an exported field
		y interface{} // a non-exported field
	}

	vpp := reflect.ValueOf(&s)
	// if vpp represent a pointer the fowlling
	// line is equivalent to "vs := vpp.Element()"
	vs := reflect.Indirect(vpp)
	// vx and vy both represent interface value
	vx, vy := vs.Field(0), vs.Field(1)
	fmt.Println(vx.CanSet(), vy.CanAddr()) // true true
	// vy is addressable but not modifiable
	fmt.Println(vy.CanSet(), vy.CanAddr()) // false true

	vb := reflect.ValueOf(123)
	vx.Set(vb) // okay vx is modifiable
	// vy.Set(vb) // will panic, for vy is unmodifiable
	fmt.Println(s)                      // {123 }
	fmt.Println(vx.IsNil(), vy.IsNil()) // false true

	var invertInts func([]int) []int
	Bind(&invertInts, InvertSlice)
	fmt.Println(invertInts([]int{2, 3, 5})) // 5 3 2

	var invertStrs func([]string) []string
	Bind(&invertStrs, InvertSlice)
	fmt.Println(invertStrs([]string{"Go", "C"})) // C Go

	// If the underlying value of a reflect.Value is a function value, then we can call the Call method of the reflect.Value to call the underlying function.
	t := T1{5, 2}
	vt := reflect.ValueOf(t)
	vm := vt.MethodByName("AddSubThenScale")
	results := vm.Call([]reflect.Value{reflect.ValueOf(3), reflect.ValueOf(5)})
	fmt.Println(results[0].Int(), results[1].Int(), results[2].Int()) // 21 9 5

	neg := func(x int) int {
		return -x
	}
	vf := reflect.ValueOf(neg)
	fmt.Println(vf.Call(results[:1])[0].Int())                 // -21
	fmt.Println(vf.Call([]reflect.Value{vt.FieldByName("A")})[ // change "A" to "b" will cause panic
	0].Int())                                                  // -5
}

func reflectMap() {
	valueOf := reflect.ValueOf
	m := map[string]int{"Unix": 1973, "Windows": 1985}
	v := valueOf(m)

	// A zero second Value argument means to delete an entry
	v.SetMapIndex(valueOf("Windows"), valueOf(1991))
	v.SetMapIndex(valueOf("MacOs"), valueOf(2000))

	for i := v.MapRange(); i.Next(); {
		fmt.Println(i.Key(), "\t:", i.Value())
	}
}

func reflectChannel() {
	c := make(chan string, 2)
	vc := reflect.ValueOf(c)
	fmt.Println(vc.Len(), vc.Cap()) // 0 2

	vc.Send(reflect.ValueOf("C"))
	fmt.Println(vc.Len(), vc.Cap()) // 1 2

	succeeded := vc.TrySend(reflect.ValueOf("Go"))
	fmt.Println(succeeded)          // true
	fmt.Println(vc.Len(), vc.Cap()) // 2 2

	succeeded = vc.TrySend(reflect.ValueOf("C++"))
	fmt.Println(succeeded)          // false
	fmt.Println(vc.Len(), vc.Cap()) // 2 2

	vs, succeeded := vc.TryRecv()
	fmt.Println(vs.String(), succeeded) // C true

	vs, sentBeforeClosed := vc.Recv()
	fmt.Println(vs.String(), sentBeforeClosed) // Go true

	vs, succeeded = vc.TryRecv()
	fmt.Println(vs.String()) //
	fmt.Println(succeeded)   // false
}

func reflectChannelWithSelect() {
	c := make(chan int, 1)
	vc := reflect.ValueOf(c)
	succeeded := vc.TrySend(reflect.ValueOf(123))
	fmt.Println(succeeded, vc.Len(), vc.Cap()) // true 1 1

	vSend, vZero := reflect.ValueOf(789), reflect.Value{}
	branches := []reflect.SelectCase{
		{Dir: reflect.SelectDefault, Chan: vZero, Send: vZero},
		{Dir: reflect.SelectRecv, Chan: vc, Send: vZero},
		{Dir: reflect.SelectSend, Chan: vc, Send: vSend},
	}

	selIndex, vRecv, sentBeforeClosed := reflect.Select(branches)
	fmt.Println(selIndex)         // 1
	fmt.Println(sentBeforeClosed) // true
	fmt.Println(vRecv)            // 123
	vc.Close()

	// Remove the send case branch this time,
	// for it may cause panic
	selIndex, _, sentBeforeClosed = reflect.Select(branches[:2])
	fmt.Println(selIndex, sentBeforeClosed) // 1 false
}

func reflectZeroValue() {
	var z reflect.Value // a zero Value value
	fmt.Println(z)      //

	v := reflect.ValueOf((*int)(nil)).Elem()
	fmt.Println(v) //

	fmt.Println(v == z) // true

	var i = reflect.ValueOf([]interface{}{nil}).Index(0)
	fmt.Println(i)             //
	fmt.Println(i.Elem() == z) // true
	fmt.Println(i.Elem())      //
}

func main() {
	fmt.Println("REFLECT TYPE")
	reflectType()

	fmt.Println("\n\n\nREFLECT STRUCT")
	reflectStruct()

	fmt.Println("\n\n\nREFLECT VALUE")
	reflectValue()

	fmt.Println("\n\n\nREFLECT MAP")
	reflectMap()

	fmt.Println("\n\n\nREFLECT CHANNEL")
	reflectChannel()

	fmt.Println("\n\n\nREFLECT CHANNEL WITH SELECT")
	reflectChannelWithSelect()

	fmt.Println("\n\n\nREFLECT ZERO VALUE")
	reflectZeroValue()
}

// A Kind represents the specific kind of type that a Type represents. The zero Kind is not a valid kind.
// Elem returns a type's element type. It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
// IsVariadic reports whether a function type's final input parameter is a "..." parameter
// NumMethod returns the number of methods accessible using Method (public method)
