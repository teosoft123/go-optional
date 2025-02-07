package optional

import "fmt"

func ExampleOption_IsNone() {
	some := Some[int](1)
	fmt.Printf("%v\n", some.IsNone())
	none := None[int]()
	fmt.Printf("%v\n", none.IsNone())
	// Output:
	// false
	// true
}

func ExampleOption_IsSome() {
	some := Some[int](1)
	fmt.Printf("%v\n", some.IsSome())
	none := None[int]()
	fmt.Printf("%v\n", none.IsSome())
	// Output:
	// true
	// false
}

func ExampleOption_Take() {
	some := Some[int](1)
	v, err := some.Take()
	fmt.Printf("%d\n", v)
	fmt.Printf("%v\n", err == nil)

	none := None[int]()
	_, err = none.Take()
	fmt.Printf("%v\n", err == nil)

	// Output:
	// 1
	// true
	// false
}

func ExampleOption_TakeOr() {
	some := Some[int](1)
	v := some.TakeOr(666)
	fmt.Printf("%d\n", v)

	none := None[int]()
	v = none.TakeOr(666)
	fmt.Printf("%d\n", v)

	// Output:
	// 1
	// 666
}

func ExampleOption_TakeOrElse() {
	some := Some[int](1)
	v := some.TakeOrElse(func () int {
		return 666
	})
	fmt.Printf("%d\n", v)

	none := None[int]()
	v = none.TakeOrElse(func () int {
		return 666
	})
	fmt.Printf("%d\n", v)

	// Output:
	// 1
	// 666
}

func ExampleOption_Filter() {
	isEven := func (v int) bool {
		return v % 2 == 0
	}

	some := Some[int](2)
	opt := some.Filter(isEven)
	fmt.Printf("%d\n", opt.TakeOr(0))

	some = Some[int](1)
	opt = some.Filter(isEven)
	fmt.Printf("%d\n", opt.TakeOr(0))

	none := None[int]()
	opt = none.Filter(isEven)
	fmt.Printf("%d\n", opt.TakeOr(0))

	// Output:
	// 2
	// 0
	// 0
}

func ExampleMap() {
	mapper := func (v int) string {
		return fmt.Sprintf("%d", v)
	}

	some := Some[int](1)
	opt := Map(some, mapper)
	fmt.Printf("%s\n", opt.TakeOr("N/A"))

	none := None[int]()
	opt = Map(none, mapper)
	fmt.Printf("%s\n", opt.TakeOr("N/A"))

	// Output:
	// 1
	// N/A
}

func ExampleMapOr() {
	mapper := func (v int) string {
		return fmt.Sprintf("%d", v)
	}

	some := Some[int](1)
	mapped := MapOr(some, "N/A", mapper)
	fmt.Printf("%s\n", mapped)

	none := None[int]()
	mapped = MapOr(none, "N/A", mapper)
	fmt.Printf("%s\n", mapped)

	// Output:
	// 1
	// N/A
}

func ExampleZip() {
	maybePair := Zip(Some[int](1), Some[string]("foo"))
	pair, err := maybePair.Take()
	fmt.Printf("is none => %v\n", maybePair.IsNone())
	fmt.Printf("err is nil => %v\n", err == nil)
	fmt.Printf("%d %s\n", pair.Value1, pair.Value2)

	maybePair = Zip(Some[int](1), None[string]())
	fmt.Printf("is none => %v\n", maybePair.IsNone())

	maybePair = Zip(None[int](), Some[string]("foo"))
	fmt.Printf("is none => %v\n", maybePair.IsNone())

	maybePair = Zip(None[int](), None[string]())
	fmt.Printf("is none => %v\n", maybePair.IsNone())

	// Output:
	// is none => false
	// err is nil => true
	// 1 foo
	// is none => true
	// is none => true
	// is none => true
}

func ExampleZipWith() {
	type Data struct {
		A int
		B string
	}

	zipper := func (v1 int, v2 string) Data {
		return Data{
			A: v1,
			B: v2,
		}
	}

	maybeData := ZipWith(Some[int](1), Some[string]("foo"), zipper)
	fmt.Printf("is none => %v\n", maybeData.IsNone())
	d, err := maybeData.Take()
	fmt.Printf("err is nil => %v\n", err == nil)
	fmt.Printf("%d %s\n", d.A, d.B)

	maybeData = ZipWith(Some[int](1), None[string](), zipper)
	fmt.Printf("is none => %v\n", maybeData.IsNone())
	maybeData = ZipWith(None[int](), Some[string]("foo"), zipper)
	fmt.Printf("is none => %v\n", maybeData.IsNone())
	maybeData = ZipWith(None[int](), None[string](), zipper)
	fmt.Printf("is none => %v\n", maybeData.IsNone())

	// Output:
	// is none => false
	// err is nil => true
	// 1 foo
	// is none => true
	// is none => true
	// is none => true
}

func ExampleUnzip() {
	// see also ExampleZip()

	pair := Pair[int, string]{
		Value1: 123,
		Value2: "foo",
	}

	o1, o2 := Unzip(Some[Pair[int, string]](pair))
	fmt.Printf("%d\n", o1.TakeOr(0))
	fmt.Printf("%s\n", o2.TakeOr(""))

	o1, o2 = Unzip(None[Pair[int, string]]())
	fmt.Printf("is none => %v\n", o1.IsNone())
	fmt.Printf("is none => %v\n", o2.IsNone())

	// Output:
	// 123
	// foo
	// is none => true
	// is none => true
}

func ExampleUnzipWith() {
	// see also ExampleZipWith()

	type Data struct {
		A int
		B string
	}

	unzipper := func (d Data) (int, string) {
		return d.A, d.B
	}

	o1, o2 := UnzipWith(Some[Data](Data{
		A: 123,
		B: "foo",
	}), unzipper)
	fmt.Printf("%d\n", o1.TakeOr(0))
	fmt.Printf("%s\n", o2.TakeOr(""))

	o1, o2 = UnzipWith(None[Data](), unzipper)
	fmt.Printf("is none => %v\n", o1.IsNone())
	fmt.Printf("is none => %v\n", o2.IsNone())

	// Output:
	// 123
	// foo
	// is none => true
	// is none => true
}
