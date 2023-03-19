package main

import "fmt"

var a1 int
var a2 = 1

// a3 := 1

func singleVar() {
	fmt.Println(">>> singleVar")

	var a1 int
	a1 = 1
	fmt.Println("a1 =", a1)

	var a2 = 1
	fmt.Println("a2 =", a2)

	a3 := 1
	fmt.Println("a3 =", a3)
}

func multiVar() {
	fmt.Println(">>> multiVar")

	var a0, a1, a2 int
	a1, a2, a3 := 1, 2, 3
	fmt.Println("a0 =", a0, ", a1 =", a1, ", a2 =", a2, ", a3 =", a3)

	var i1, f1, b1, r1, s1 = 1, 1.1, true, '1', "1"
	fmt.Println("i1 =", i1, ", f1 =", f1, ", b1 =", b1, ", r1 =", r1, ", s1 =", s1)

	i3, f3, b3, r3, s3 := 1, 1.1, true, '1', "1"
	fmt.Println("i3 =", i3, ", f3 =", f3, ", b3 =", b3, ", r3 =", r3, ", s3 =", s3)

	var (
		i2 int
		f2 float64
		b2 bool
		r2 rune
		s2 string
	)
	fmt.Println("i2 =", i2, ", f2 =", f2, ", b2 =", b2, ", r2 =", r2, ", s2 =", s2)

}

func defaultVal() {
	fmt.Println(">>> defaultVal")
	var (
		i            int
		i8           int8
		i16          int16
		i32          int32
		i64          int64
		ui           uint
		ui8          uint8
		ui16         uint16
		ui32         uint32
		ui64         uint64
		f32          float32
		f64          float64
		c64          complex64
		c128         complex128
		b            bool
		r            rune
		s            string
		arrayVar     [0]int
		sliceVar     []int
		mapVar       map[string]int
		pointerVal   *int
		structVal    struct{}
		functionVal  func(int) int
		interfaceVal interface{}
		chanVal      chan int
	)
	fmt.Println("i =", i)
	fmt.Println("i8 =", i8)
	fmt.Println("i16 =", i16)
	fmt.Println("i32 =", i32)
	fmt.Println("i64 =", i64)
	fmt.Println("ui =", ui)
	fmt.Println("ui8 =", ui8)
	fmt.Println("ui16 =", ui16)
	fmt.Println("ui32 =", ui32)
	fmt.Println("ui64 =", ui64)
	fmt.Println("c64 =", c64)
	fmt.Println("c128 =", c128)
	fmt.Println("f32 =", f32)
	fmt.Println("f64 =", f64)
	fmt.Println("b =", b)
	fmt.Println("r =", r)
	fmt.Println("s =", s)
	fmt.Println("arrayVar =", arrayVar)
	fmt.Println("sliceVar =", sliceVar)
	fmt.Println("mapVar =", mapVar)
	fmt.Println("pointerVal =", pointerVal)
	fmt.Println("structVal =", structVal)
	fmt.Println("functionVal =", functionVal)
	fmt.Println("interfaceVal =", interfaceVal)
	fmt.Println("chanVal =", chanVal)

}

func main() {
	singleVar()
	multiVar()
	defaultVal()
}
