package snippets

func foo(i int, b bool) {} //@item(snipFoo, "foo(i int, b bool)", "", "func")
func bar(fn func()) func()    {} //@item(snipBar, "bar(fn func())", "", "func")

type Foo struct {
	Bar int //@item(snipFieldBar, "Bar", "int", "field")
}

func (Foo) Baz() func() {} //@item(snipMethodBaz, "Baz()", "func()", "method")
func (Foo) BazBar() func() {} //@item(snipMethodBazBar, "BazBar()", "func()", "method")

func _() {
	f //@snippet(" //", snipFoo, "foo(${1})", "foo(${1:i int}, ${2:b bool})")

	bar //@snippet(" //", snipBar, "bar(${1})", "bar(${1:fn func()})")

	bar(nil) //@snippet("(", snipBar, "", "")
	bar(ba) //@snippet(")", snipBar, "bar(${1})", "bar(${1:fn func()})")
	var f Foo
	bar(f.Ba) //@snippet(")", snipMethodBaz, "Baz()", "Baz()")
	(bar)(nil) //@snippet(")", snipBar, "bar(${1})", "bar(${1:fn func()})")
	(f.Ba)() //@snippet(")", snipMethodBaz, "Baz()", "Baz()")

	Foo{
		B //@snippet(" //", snipFieldBar, "Bar: ${1},", "Bar: ${1:int},")
	}

	Foo{B} //@snippet("}", snipFieldBar, "Bar: ${1}", "Bar: ${1:int}")
	Foo{} //@snippet("}", snipFieldBar, "Bar: ${1}", "Bar: ${1:int}")

	Foo{Foo{}.B} //@snippet("} ", snipFieldBar, "Bar", "Bar")

	var err error
	err.Error() //@snippet("E", Error, "", "")
	f.Baz()     //@snippet("B", snipMethodBaz, "", "")

	// TODO(rstambler): Handle this case correctly.
	f.Baz()     //@fails("(", snipMethodBazBar, "Bar", "Bar")
}
