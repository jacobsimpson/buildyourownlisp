package evaluator

import (
	"buildyourownlisp/ast"
	"fmt"
)

func Evaluate(statement interface{}) (interface{}, error) {
	v := statement.([]interface{})
	result := []interface{}{}
	for _, s := range v {
		e, err := evaluate(s)
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}

func evaluate(statement interface{}) (interface{}, error) {
	switch v := statement.(type) {
	case int, string, bool:
		return v, nil
	case *ast.Cell:
		return evaluateList(v)
	}
	return nil, fmt.Errorf("unable to evaluate %v", statement)
}

func evaluateList(head *ast.Cell) (interface{}, error) {
	switch h := head.Left.(type) {
	case ast.Symbol:
		return evaluateFunction(h, head.Right)
	}
	return nil, fmt.Errorf("Wrong type to apply: %v", head.Left)
}

func evaluateFunction(s ast.Symbol, args interface{}) (interface{}, error) {
	fn := functions[s]
	if fn == nil {
		return nil, fmt.Errorf("Unbound variable: %v", s)
	}

	a, err := evaluateArgs(args)
	if err != nil {
		return nil, err
	}
	return fn(a)
}

func evaluateArgs(args interface{}) ([]interface{}, error) {
	result := []interface{}{}
	var a *ast.Cell
	if args != nil {
		a = args.(*ast.Cell)
	}
	for a != nil {
		switch v := a.Left.(type) {
		case *ast.Cell:
			o, err := evaluateList(v)
			if err != nil {
				return nil, err
			}
			result = append(result, o)
		default:
			result = append(result, v)
		}
		if a.Right == nil {
			a = nil
		} else {
			a = a.Right.(*ast.Cell)
		}
	}
	return result, nil
}

var functions map[ast.Symbol]func(args []interface{}) (interface{}, error)

func init() {
	functions = map[ast.Symbol]func(args []interface{}) (interface{}, error){
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}
}

func add(args []interface{}) (interface{}, error) {
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

func subtract(args []interface{}) (interface{}, error) {
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

func multiply(args []interface{}) (interface{}, error) {
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

func divide(args []interface{}) (interface{}, error) {
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
