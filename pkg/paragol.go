// Package paragol allows you to manipulate text with nested parenthesis,
// as frequently encountered in LISP s-expressions.
// Particular care is taken for it to work well for the Clojure LISP dialect,
// but it should work with any format that employs balanced parenthesis, square brackets
// and curly braces.
//
// In the following description we use paranthesis, but you could replace square brackets or curly braces
// in the examples and it wouldn't affect the semantics.
//
package paragol

import (
	p "github.com/vektah/goparsify"
)

func parse(src string) error {
	var y p.Parser // forward decl for recursive parser
	around := func(start, stop string) p.Parserish { return p.Seq(p.Chars(start), &y, p.Seq(p.Chars(stop))) }

	body := p.NotChars("()[]{}")
	exp := p.Any(around("(", ")"), around("[", "]"), around("{", "}"))
	y = p.Any(exp, body)

	_, err := p.Run(y, src)
	return err
}

// IsBalanced returns true if the parenthesis in expression are balanced
func IsBalanced(src string) bool {
	return parse(src) == nil
}
