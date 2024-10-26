package evaluator

import "monkey/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		}},

	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				elements := arg.Elements
				if len(elements) == 0 {
					return NULL
				}
				return elements[0]

			default:
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				elements := arg.Elements
				if len(elements) == 0 {
					return NULL
				}
				return elements[len(elements)-1]

			default:
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				elements := arg.Elements
				length := len(elements)
				if len(elements) == 0 {
					return NULL
				}
				newElements := make([]object.Object, length-1)
				copy(newElements, elements[1:])
				return &object.Array{Elements: newElements}

			default:
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}
		},
	},

	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				elements := arg.Elements
				length := len(elements)
				newElements := make([]object.Object, length+1)
				copy(newElements, elements)
				newElements[length] = args[1]
				return &object.Array{Elements: newElements}

			default:
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
		},
	},

	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				println(arg.Inspect())
			}
			return NULL
		},
	},
}
