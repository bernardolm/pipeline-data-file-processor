package processor

import (
	"fmt"
	"reflect"

	"github.com/k0kubun/pp"
)

var typeRegistry = make(map[string]reflect.Type)

func typeRegister(T interface{}) {
	t := reflect.TypeOf(T)
	fmt.Printf("registering %s...", t.Name())
	// typeRegistry["MyString"] = reflect.TypeOf(MyString{})
	// typeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
	// typeRegistry[reflect.TypeOf(T).Name()] = reflect.TypeOf(v)
	typeRegistry[t.Name()] = t
}

func RegisterSource(S ISource) {
	typeRegister(S)
}

func RegisterOutput(S IOutput) {
	typeRegister(S)
}

func MakeInstance(name string) (interface{}, error) {
	if _, ok := typeRegistry[name]; !ok {
		return nil, fmt.Errorf("type %s isn't registered", name)
	}

	v := reflect.New(typeRegistry[name])

	return v.Elem().Interface(), nil
}

func ListRegistredTypes() {
	fmt.Printf("\nRegistred types\n")
	pp.Print(typeRegistry)
}
