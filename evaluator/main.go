package evaluator

import (
	"buildyourownlisp/ast"
	"fmt"
)

func Evaluate(environment *Environment, statement interface{}) (interface{}, error) {
	v := statement.([]interface{})
	result := []interface{}{}
	for _, s := range v {
		e, err := evaluate(environment, s)
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}

func evaluate(environment *Environment, statement interface{}) (interface{}, error) {
	switch v := statement.(type) {
	case int, string, bool:
		return v, nil
	case *ast.Cell:
		return evaluateList(environment, v)
	}
	return nil, fmt.Errorf("unable to evaluate %v", statement)
}

func evaluateList(environment *Environment, head *ast.Cell) (interface{}, error) {
	switch h := head.Left.(type) {
	case ast.Symbol:
		return evaluateFunction(environment, h, head.Right)
	}
	fmt.Printf("evaluateList %+v\n", head.Left)
	return nil, fmt.Errorf("Wrong type to apply: %v", head.Left)
}

func evaluateFunction(environment *Environment, s ast.Symbol, args interface{}) (interface{}, error) {
	v := environment.bindings[s]
	if v == nil {
		return nil, fmt.Errorf("Unbound variable: %v", s)
	}

	if fn, ok := v.(SpecialFormFn); ok {
		return evaluateSpecialForm(environment, fn, args)
	}

	fn, ok := v.(InterpreterFn)
	if !ok {
		fmt.Printf("evaluateFunction\n")
		return nil, fmt.Errorf("Wrong type to apply: %v", s)
	}

	a, err := evaluateArgs(environment, args)
	if err != nil {
		return nil, err
	}
	return fn(environment, a)
}

func evaluateSpecialForm(environment *Environment, sf SpecialFormFn, args interface{}) (interface{}, error) {
	fmt.Printf("evaluateSpecialForm: %+v\n", args)
	return sf(environment, args)
	//v, err := evaluate(environment, predicate)
	//if err != nil {
	//	return nil, err
	//}
	//v := environment.bindings[s]
	//if v == nil {
	//	return nil, fmt.Errorf("Unbound variable: %v", s)
	//}

	//fn, ok := v.(InterpreterFn)
	//if ok {
	//	return evaluateSpecialForm(environment, args)
	//}

	//fn, ok := v.(InterpreterFn)
	//if !ok {
	//	return nil, fmt.Errorf("Wrong type to apply: %v", s)
	//}

	//a, err := evaluateArgs(environment, args)
	//if err != nil {
	//	return nil, err
	//}
	//return fn(environment, a)
}

func evaluateArgs(environment *Environment, args interface{}) ([]interface{}, error) {
	result := []interface{}{}
	var a *ast.Cell
	if args != nil {
		a = args.(*ast.Cell)
	}
	for a != nil {
		switch v := a.Left.(type) {
		case *ast.Cell:
			o, err := evaluateList(environment, v)
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
