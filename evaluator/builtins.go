package evaluator

import (
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": object.GetBuiltinByName("len"),
	"first": object.GetBuiltinByName("first"),
	"rest": object.GetBuiltinByName("rest"),
	"last": object.GetBuiltinByName("last"),
	"push": object.GetBuiltinByName("push"),
	"puts": object.GetBuiltinByName("puts"),
}
