package asserts

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

type tester interface {
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Error(args ...any)
	Fatal(args ...any)
	Helper()
}

func Equal[T any](t tester, expected T, actual T, opts ...cmp.Option) {
	if !cmp.Equal(expected, actual, opts...) {
		diff := cmp.Diff(expected, actual)
		t.Helper()
		t.Fatalf("Expected values to be equal (-want +got):\n%s", diff)
	}
}

func NotEqual[T any](t tester, expected T, actual T, opts ...cmp.Option) {
	if cmp.Equal(expected, actual, opts...) {
		t.Helper()
		t.Fatal("Expected values to not be equal")
	}
}

func Nil(t tester, value interface{}) {
	if value != nil {
		if reflect.ValueOf(value).CanAddr() && !reflect.ValueOf(value).Addr().IsNil() {
			t.Helper()
			t.Fatalf("Not Nil:\n%v", value)
		}
	}
}

func True[T bool](t tester, value T) {
	if value != true {
		t.Helper()
		t.Fatalf("Not true")
	}
}

func False[T bool](t tester, value T) {
	if value != false {
		t.Helper()
		t.Fatalf("Not false")
	}
}
