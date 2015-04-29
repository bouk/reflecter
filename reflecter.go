package reflecter

import (
	"github.com/bouk/symme"
	"reflect"
	"unsafe"
)

var types []reflect.Type

type enterface struct{ typ, val uintptr }

func init() {
	// Make sure the right stuff is linked
	reflect.SliceOf(reflect.TypeOf(1))

	t, err := symme.Table()
	if err != nil {
		panic(err)
	}
	rtypeType := reflect.TypeOf(1)
	var typelinks func() []uintptr
	tlLocation := uintptr(t.LookupFunc("reflect.typelinks").Value)
	*(*uintptr)(unsafe.Pointer(&typelinks)) = uintptr(unsafe.Pointer(&tlLocation))

	for _, typ := range typelinks() {
		(*enterface)(unsafe.Pointer(&rtypeType)).val = typ
		types = append(types, rtypeType)
	}
}

// Types returns a slice of all the Types that are in this program
func Types() []reflect.Type {
	return types
}
