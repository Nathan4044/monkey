package evaluator

import (
	"fmt"
	"monkey/object"
	"strings"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[0]
				} else {
					return NULL
				}
			case *object.String:
				if len(arg.Value) > 0 {
					return &object.String{Value: string(arg.Value[0])}
				} else {
					return NULL
				}
			default:
				return newError("argument to `first` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				length := len(arg.Elements)

				if length > 0 {
					newElements := make([]object.Object, length-1, length-1)
					copy(newElements, arg.Elements[1:length])

					return &object.Array{Elements: newElements}
				}

				return NULL
			case *object.String:
				length := len(arg.Value)

				if length > 0 {
					newString := strings.Clone(arg.Value[1:length])
					return &object.String{Value: newString}
				}

				return NULL
			default:
				return newError("argument to `first` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				length := len(arg.Elements)
				if length > 0 {
					return arg.Elements[length-1]
				} else {
					return NULL
				}
			case *object.String:
				length := len(arg.Value)
				if length > 0 {
					return &object.String{Value: string(arg.Value[length-1])}
				} else {
					return NULL
				}
			default:
				return newError("argument to `first` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newArr := make([]object.Object, length+1, length+1)

			copy(newArr, arr.Elements)
			newArr[length] = args[1]

			return &object.Array{Elements: newArr}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
