package cell

import (
	"testing"
)

type Foo struct {
	bar string
}

func (f *Foo) Fill(k string, v interface{}) {
	if k == "bar" {
		f.bar = v.(string)
	}
}

func TestSpilt(t *testing.T) {
	Regist("foo", func() interface{} { return &Foo{"bar"} })
	foo := Spilt("foo", nil).(*Foo)
	if foo.bar != "bar" {
		t.Errorf("it should be 'bar', but it is %s", foo.bar)
	}
	foo = Spilt("foo", &SplitOptions{[]string{"bar"}, []interface{}{"foo"}}).(*Foo)
	if foo.bar != "foo" {
		t.Errorf("it should be 'foo', but it is %s", foo.bar)
	}
	foo = Spilt("foo", nil).(*Foo)
	if foo.bar != "bar" {
		t.Errorf("it should be 'bar', but it is %s", foo.bar)
	}
}

func BenchmarkSpilt(b *testing.B) {
	Regist("foo", func() interface{} { return &Foo{} })
	for i := 0; i < b.N; i++ {
		_ = Spilt("foo", nil).(*Foo)
	}
}

func BenchmarkSpiltWithFill(b *testing.B) {
	Regist("foo", func() interface{} { return &Foo{} })
	o := &SplitOptions{[]string{"bar"}, []interface{}{"foo"}}
	for i := 0; i < b.N; i++ {
		_ = Spilt("foo", o).(*Foo)
	}
}
