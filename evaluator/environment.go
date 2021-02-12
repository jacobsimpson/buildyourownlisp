package evaluator

import (
	"buildyourownlisp/ast"
	"fmt"
)

type Environment struct {
	parent   *Environment
	bindings map[ast.Symbol]interface{}
}

type InterpreterFn func(environment *Environment, args []interface{}) (interface{}, error)
type SpecialFormFn func(environment *Environment, args interface{}) (interface{}, error)

func NewEnvironment(p *Environment) *Environment {
	return &Environment{
		parent:   p,
		bindings: map[ast.Symbol]interface{}{},
	}
}

func NewRootEnvironment() *Environment {
	e := &Environment{
		bindings: map[ast.Symbol]interface{}{
			ast.Symbol("+"):  InterpreterFn(add),
			ast.Symbol("-"):  InterpreterFn(subtract),
			ast.Symbol("*"):  InterpreterFn(multiply),
			ast.Symbol("/"):  InterpreterFn(divide),
			ast.Symbol("if"): SpecialFormFn(ifFn),
		},
	}
	return e
}

func add(environment *Environment, args []interface{}) (interface{}, error) {
	result := 0
	for i, c := range args {
		switch v := c.(type) {
		case int:
			result += v
		default:
			return nil, fmt.Errorf("In procedure +: Wrong type argument in position %d: %v", i+1, v)
		}
	}
	return result, nil
}

func subtract(environment *Environment, args []interface{}) (interface{}, error) {
	result := 0
	firstElement := true
	for i, c := range args {
		switch v := c.(type) {
		case int:
			if firstElement {
				firstElement = false
			}
			result -= v
		default:
			return nil, fmt.Errorf("In procedure -: Wrong type argument in position %d: %v", i+1, v)
		}
	}
	if firstElement {
		return nil, fmt.Errorf("Wrong number of arguments to -")
	}
	return result, nil
}

func multiply(environment *Environment, args []interface{}) (interface{}, error) {
	result := 1
	for i, c := range args {
		switch v := c.(type) {
		case int:
			result *= v
		default:
			return nil, fmt.Errorf("In procedure *: Wrong type argument in position %d: %v", i+1, v)
		}
	}
	return result, nil
}

func divide(environment *Environment, args []interface{}) (interface{}, error) {
	result := 1
	firstElement := true
	for i, c := range args {
		switch v := c.(type) {
		case int:
			if firstElement {
				firstElement = false
			}
			if v == 0 {
				return nil, fmt.Errorf("In procedure /: division by 0")
			}
			result /= v
		default:
			fmt.Println("5")
			return nil, fmt.Errorf("In procedure /: Wrong type argument in position %d: %v", i+1, v)
		}
	}
	if firstElement {
		return nil, fmt.Errorf("Wrong number of arguments to /")
	}
	return result, nil
}

func ifFn(environment *Environment, args interface{}) (interface{}, error) {
	fmt.Printf("ifFn: %+v\n", args)
	if args == nil {
		return nil, fmt.Errorf("Missing parameters")
	}

	c, ok := args.(*ast.Cell)
	if !ok {
		return nil, fmt.Errorf("Wrong parameters")
	}

	if c.Right == nil {
		return nil, fmt.Errorf("Missing branch in if statement")
	}

	v, err := evaluate(environment, c.Left)
	if err != nil {
		return nil, err
	}
	fmt.Printf("ifFn, predicate evaluation = %+v\n", v)
	if isTruthy(v) {
		_, err := evaluate(environment, c.Right)
		if err != nil {
			return nil, err
		}
	}
	//predicate := true
	//branch := true
	//alternate := true
	//func ifFn(environment *Environment, args interface{}) (interface{}, error) {

	return nil, fmt.Errorf("Unimplemented")
	//result := 1
	//firstElement := true
	//for i, c := range args {
	//	switch v := c.(type) {
	//	case int:
	//		if firstElement {
	//			firstElement = false
	//		}
	//		if v == 0 {
	//			return nil, fmt.Errorf("In procedure /: division by 0")
	//		}
	//		result /= v
	//	default:
	//		fmt.Println("5")
	//		return nil, fmt.Errorf("In procedure /: Wrong type argument in position %d: %v", i+1, v)
	//	}
	//}
	//if firstElement {
	//	return nil, fmt.Errorf("Wrong number of arguments to /")
	//}
	//return result, nil
}

func isTruthy(arg interface{}) bool {
	switch v := arg.(type) {
	case bool:
		return v
	default:
		return true
	}
}
