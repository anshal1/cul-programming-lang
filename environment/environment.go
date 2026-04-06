package environment

import "github.com/anshal1/custom-language/utils"

type Value struct {
	// type of the variable
	Type string
	// value of the variable
	Value string
	// the actual token
	Token utils.Token
}
type Environment struct {
	// the key here is the variable name
	store map[string]Value
}

func NewEnv() *Environment {
	return &Environment{store: make(map[string]Value)}
}

func (e *Environment) Set(name string, value Value) {
	e.store[name] = value
}

func (e *Environment) Get(name string) (Value, bool) {
	val, ok := e.store[name]
	return val, ok
}
func (e *Environment) Print() map[string]Value {
	return e.store
}
