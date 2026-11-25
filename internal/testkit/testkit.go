// Package testkit has small helpers for the exercise tests.
//
// They let a test check behavior without dictating implementation details:
// methods are looked up by name at runtime, so a test compiles as soon as the
// types it mentions exist, even if later steps of the exercise aren't written
// yet. Constructors may return either T or *T, and methods may use value or
// pointer receivers, unless the behavior being tested depends on it.
package testkit

import (
	"reflect"
	"testing"
)

// Method returns the method called name of recv, as a function of type F.
// It fails the test if the method doesn't exist (yet) or has another signature.
//
// Pass a pointer (&x) as recv: a pointer's method set holds both its value-
// and pointer-receiver methods.
func Method[F any](t testing.TB, recv any, name string) F {
	t.Helper()
	v := reflect.ValueOf(recv)
	m := v.MethodByName(name)
	if !m.IsValid() {
		t.Fatalf("%s has no %s method yet", typeName(v.Type()), name)
	}
	f, ok := m.Interface().(F)
	if !ok {
		t.Fatalf("%s.%s has signature %v, want %v", typeName(v.Type()), name, m.Type(), reflect.TypeFor[F]())
	}
	return f
}

// Value returns v as a T. v may be a T or a non-nil *T.
func Value[T any](t testing.TB, v any) T {
	t.Helper()
	switch x := v.(type) {
	case T:
		return x
	case *T:
		if x == nil {
			t.Fatalf("got a nil *%v", reflect.TypeFor[T]())
		}
		return *x
	}
	t.Fatalf("got a %T, want a %v or a *%v", v, reflect.TypeFor[T](), reflect.TypeFor[T]())
	panic("unreachable")
}

// Ptr returns v if it's a non-nil pointer, or a pointer to a copy of v
// otherwise, so that pointer-receiver methods can be called on it.
func Ptr(t testing.TB, v any) any {
	t.Helper()
	rv := reflect.ValueOf(v)
	switch {
	case !rv.IsValid():
		t.Fatal("got nil")
	case rv.Kind() == reflect.Pointer:
		if rv.IsNil() {
			t.Fatalf("got a nil %v", rv.Type())
		}
		return v
	}
	p := reflect.New(rv.Type())
	p.Elem().Set(rv)
	return p.Interface()
}

// AllFieldsUnexported fails the test unless T is a struct whose fields are
// all unexported, so code in other packages can only use it through its methods.
func AllFieldsUnexported[T any](t testing.TB) {
	t.Helper()
	typ := reflect.TypeFor[T]()
	if typ.Kind() != reflect.Struct {
		t.Fatalf("%s is a %s, want a struct", typ.Name(), typ.Kind())
	}
	if typ.NumField() == 0 {
		t.Fatalf("%s has no fields yet", typ.Name())
	}
	for i := range typ.NumField() {
		if f := typ.Field(i); f.IsExported() {
			t.Errorf("%s.%s is exported: other packages could change it without going through a method", typ.Name(), f.Name)
		}
	}
}

// MapsInitialized fails the test unless v (a struct or a pointer to one) has
// at least one map field, and none of its map fields are nil.
func MapsInitialized(t testing.TB, v any) {
	t.Helper()
	rv := reflect.Indirect(reflect.ValueOf(Ptr(t, v)))
	if rv.Kind() != reflect.Struct {
		t.Fatalf("got a %s, want a struct", rv.Kind())
	}
	found := false
	for i := range rv.NumField() {
		if f := rv.Field(i); f.Kind() == reflect.Map {
			found = true
			if f.IsNil() {
				t.Errorf("%s.%s is a nil map: writing to it would panic", rv.Type().Name(), rv.Type().Field(i).Name)
			}
		}
	}
	if !found {
		t.Fatalf("%s has no map field", rv.Type().Name())
	}
}

func typeName(t reflect.Type) string {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Name() != "" {
		return t.Name()
	}
	return t.String()
}
