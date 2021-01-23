package ast

import "fmt"

type Symbol string

type Cell struct {
	Left  interface{}
	Right interface{}
}

func (c *Cell) String() string {
	return fmt.Sprintf("(%v . %v)", c.Left, c.Right)
}
