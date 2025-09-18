package mediatype_test

import (
	"fmt"

	"github.com/aereal/mediatype"
)

func ExampleParse() {
	mt, err := mediatype.Parse("text/plain;charset=utf-8")
	if err != nil {
		panic(err)
	}
	fmt.Printf("type=%s sub=%s params=%#v\n", mt.Type, mt.SubType, mt.Parameters)
	// Output: type=text sub=plain params=map[string]string{"charset":"utf-8"}
}

func ExampleMediaType_Equal() {
	a, err := mediatype.Parse("text/plain")
	if err != nil {
		panic(err)
	}
	b, err := mediatype.Parse("text/plain;charset=utf-8")
	if err != nil {
		panic(err)
	}
	c, err := mediatype.Parse("text/xml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("a==b: %v\n", a.Equal(b))
	fmt.Printf("b==c: %v\n", b.Equal(c))
	fmt.Printf("a==c: %v\n", a.Equal(c))
	// Output:
	// a==b: true
	// b==c: false
	// a==c: false
}

func ExampleSubType_Base() {
	fmt.Printf("plain: %s\n", mediatype.SubType("plain").Base())
	fmt.Printf("svg+xml: %s\n", mediatype.SubType("svg+xml").Base())
	// Output:
	// plain: plain
	// svg+xml: svg
}

func ExampleSubType_Suffix() {
	{
		suffix, hasSuffix := mediatype.SubType("plain").Suffix()
		fmt.Printf("plain: %q %v\n", suffix, hasSuffix)
	}
	{
		suffix, hasSuffix := mediatype.SubType("svg+xml").Suffix()
		fmt.Printf("svg+xml: %q %v\n", suffix, hasSuffix)
	}
	{
		suffix, hasSuffix := mediatype.SubType("svg+").Suffix()
		fmt.Printf("svg+: %q %v\n", suffix, hasSuffix)
	}
	// Output:
	// plain: "" false
	// svg+xml: "xml" true
	// svg+: "" false
}
