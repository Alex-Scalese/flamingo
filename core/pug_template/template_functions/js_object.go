package template_functions

import (
	"flamingo/core/pug_template/pugjs"
	"sort"
	"strconv"
)

type (
	JsObject struct{}
	Object   struct{}
)

// Name alias for use in template
func (ol JsObject) Name() string {
	return "Object"
}

// Func as implementation of debug method
func (ol JsObject) Func() interface{} {
	return func() Object {
		return Object{}
	}
}

// NoConvert marker
func (o Object) NoConvert() {}

// Assign all properties from the sources to the target map
func (o Object) Assign(target *pugjs.Map, sources ...*pugjs.Map) pugjs.Object {
	for _, source := range sources {
		if source != nil {
			for k, v := range source.Items {
				target.Items[k] = v
			}
		}
	}

	return target
}

// Keys returns all keys of a map in lexical order
func (o Object) Keys(obj interface{}) *pugjs.Array {
	res := &pugjs.Array{}
	if obj == nil {
		return res
	}
	var tmp []string

	if m, ok := obj.(*pugjs.Map); ok {
		for k := range m.Items {
			tmp = append(tmp, k.String())
		}
	} else if a, ok := obj.(*pugjs.Array); ok {
		for i := 0; i < int(a.Length().(pugjs.Number)); i++ {
			tmp = append(tmp, strconv.Itoa(i))
		}
	}

	sort.Strings(tmp)

	for _, k := range tmp {
		res.Push(pugjs.String(k))
	}

	return res
}