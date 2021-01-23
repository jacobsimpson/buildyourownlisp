package parser

import (
	"buildyourownlisp/ast"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		program string
		wantErr error
		wantAst interface{}
	}{
		{
			program: "#t",
			wantErr: nil,
			wantAst: []interface{}{
				true,
			},
		},
		{
			program: `"hello world"`,
			wantErr: nil,
			wantAst: []interface{}{
				"hello world",
			},
		},
		{
			program: `98765`,
			wantErr: nil,
			wantAst: []interface{}{
				98765,
			},
		},
		{
			program: "var-name",
			wantErr: nil,
			wantAst: []interface{}{
				"var-name",
			},
		},
		{
			program: "()",
			wantErr: nil,
			wantAst: []interface{}{
				&ast.Cell{},
			},
		},
		{
			program: "))",
			wantErr: errList([]error{
				&parserError{
					Inner:  fmt.Errorf("no match found"),
					pos:    position{line: 1, col: 1, offset: 0},
					prefix: "parser-test:1:1 (0)",
				},
			}),
		},
		{
			program: "(87)",
			wantErr: nil,
			wantAst: []interface{}{
				&ast.Cell{
					Left: 87,
				},
			},
		},
		{
			program: `(    8
			7)`,
			wantErr: nil,
			wantAst: []interface{}{
				&ast.Cell{
					Left:  8,
					Right: &ast.Cell{Left: 7},
				},
			},
		},
		{
			program: ` (    "abc" 7 #t  #f  )    `,
			wantErr: nil,
			wantAst: []interface{}{
				&ast.Cell{
					Left: "abc",
					Right: &ast.Cell{
						Left: 7,
						Right: &ast.Cell{
							Left: true,
							Right: &ast.Cell{
								Left: false,
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.program, func(t *testing.T) {
			assert := assert.New(t)

			got, err := Parse("parser-test", []byte(test.program), Debug(false))

			assert.Equal(test.wantErr, err)
			assert.Equal(test.wantAst, got)
		})
	}
}
