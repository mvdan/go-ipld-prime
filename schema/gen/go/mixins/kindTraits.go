package mixins

import (
	"io"

	ipld "github.com/ipld/go-ipld-prime"
)

// kindTraitsGenerator is a embedded in all the other mixins,
// and handles all the method generation which is a pure function of the kind.
//
// OVERRIDE THE METHODS THAT DO APPLY TO YOUR KIND;
// the default method bodies produced by this mixin are those that return errors,
// and that is not what you want for the methods that *are* interesting for your kind.
// The kindTraitsGenerator methods will panic if called for a kind that should've overriden them.
//
// If you're implementing something that can hold "any" kind,
// probably none of these methods apply to you at all.
//
// The other types in this package embed kindTraitsGenerator with a name,
// and only forward the methods to it that don't apply for their kind;
// this means when they're used as an anonymous embed, they grant
// all the appropriate dummy methods to their container,
// while leaving the ones that are still needed entirely absent,
// so the compiler helpfully tells you to finish rather than waiting until
// runtime to panic if a should-have-been-overriden method slips through.
type kindTraitsGenerator struct {
	PkgName    string
	TypeName   string // as will be printed in messages (e.g. can be goosed up a bit, like "Thing.Repr" instead of "_Thing__Repr").
	TypeSymbol string // the identifier in code (sometimes is munged internals like "_Thing__Repr" corresponding to no publicly admitted schema.Type.Name).
	Kind       ipld.ReprKind
}

func (g kindTraitsGenerator) emitNodeMethodLookupString(w io.Writer) {
	if ipld.ReprKindSet_JustMap.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) LookupString(string) (ipld.Node, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.LookupString("")
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodLookup(w io.Writer) {
	if ipld.ReprKindSet_JustMap.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) Lookup(ipld.Node) (ipld.Node, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.Lookup(nil)
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodLookupIndex(w io.Writer) {
	if ipld.ReprKindSet_JustList.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) LookupIndex(idx int) (ipld.Node, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.LookupIndex(0)
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodLookupSegment(w io.Writer) {
	if ipld.ReprKindSet_Recursive.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.LookupSegment(seg)
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodMapIterator(w io.Writer) {
	if ipld.ReprKindSet_JustMap.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) MapIterator() ipld.MapIterator {
			return nil
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodListIterator(w io.Writer) {
	if ipld.ReprKindSet_JustList.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) ListIterator() ipld.ListIterator {
			return nil
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodLength(w io.Writer) {
	if ipld.ReprKindSet_Recursive.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) Length() int {
			return -1
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodIsUndefined(w io.Writer) {
	doTemplate(`
		func ({{ .TypeSymbol }}) IsUndefined() bool {
			return false
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodIsNull(w io.Writer) {
	doTemplate(`
		func ({{ .TypeSymbol }}) IsNull() bool {
			return false
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodAsBool(w io.Writer) {
	if ipld.ReprKindSet_JustBool.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) AsBool() (bool, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.AsBool()
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodAsInt(w io.Writer) {
	if ipld.ReprKindSet_JustInt.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) AsInt() (int, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.AsInt()
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodAsFloat(w io.Writer) {
	if ipld.ReprKindSet_JustFloat.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) AsFloat() (float64, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.AsFloat()
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodAsString(w io.Writer) {
	if ipld.ReprKindSet_JustString.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) AsString() (string, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.AsString()
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodAsBytes(w io.Writer) {
	if ipld.ReprKindSet_JustBytes.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) AsBytes() ([]byte, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.AsBytes()
		}
	`, w, g)
}

func (g kindTraitsGenerator) emitNodeMethodAsLink(w io.Writer) {
	if ipld.ReprKindSet_JustLink.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .TypeSymbol }}) AsLink() (ipld.Link, error) {
			return mixins.{{ .Kind }}{"{{ .PkgName }}.{{ .TypeName }}"}.AsLink()
		}
	`, w, g)
}

// kindAssemblerTraitsGenerator is an awfully lot like kindTraitsGenerator,
// except applying to methods for builders and assemblers.
type kindAssemblerTraitsGenerator struct {
	PkgName       string
	TypeName      string // as will be printed in messages (e.g. can be goosed up a bit, like "Thing.Repr" instead of "_Thing__Repr").
	AppliedPrefix string // the prefix of what to attach methods to... this one is a little wild: should probably be either "_{{ .Type | TypeSymbol }}__" or "_{{ .Type | TypeSymbol }}__Repr", and we'll just add the words "Builder" and "Assembler".
	Kind          ipld.ReprKind
}

// bailed on extracting a common emitNodeBuilderType: too many variations in content and pointer placement to be worth it.
// bailed on extracting a common emitNodeBuilderMethods: same.
// bailed on extracting a common emitNodeAssemblerType: same.
//
// If you try to do these, you'll probably need:
//  - an explicit understanding of if generating representations or not
//  - to still be ready for boatloads of exceptions if the representation isn't directly castable to and from the type-level node.

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodBeginMap(w io.Writer) {
	if ipld.ReprKindSet_JustMap.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.BeginMap(0)
		}
	`, w, g)
}

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodBeginList(w io.Writer) {
	if ipld.ReprKindSet_JustList.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.BeginList(0)
		}
	`, w, g)
}

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodAssignNull(w io.Writer) {
	if ipld.ReprKindSet_JustNull.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) AssignNull() error {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.AssignNull()
		}
	`, w, g)
}

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodAssignBool(w io.Writer) {
	if ipld.ReprKindSet_JustBool.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) AssignBool(bool) error {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.AssignBool(false)
		}
	`, w, g)
}

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodAssignInt(w io.Writer) {
	if ipld.ReprKindSet_JustInt.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) AssignInt(int) error {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.AssignInt(0)
		}
	`, w, g)
}

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodAssignFloat(w io.Writer) {
	if ipld.ReprKindSet_JustFloat.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) AssignFloat(float64) error {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.AssignFloat(0)
		}
	`, w, g)
}

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodAssignString(w io.Writer) {
	if ipld.ReprKindSet_JustString.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) AssignString(string) error {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.AssignString("")
		}
	`, w, g)
}

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodAssignBytes(w io.Writer) {
	if ipld.ReprKindSet_JustBytes.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) AssignBytes([]byte) error {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.AssignBytes(nil)
		}
	`, w, g)
}

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodAssignLink(w io.Writer) {
	if ipld.ReprKindSet_JustLink.Contains(g.Kind) {
		panic("gen internals error: you should've overriden this")
	}
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) AssignLink(ipld.Link) error {
			return mixins.{{ .Kind }}Assembler{"{{ .PkgName }}.{{ .TypeName }}"}.AssignLink(nil)
		}
	`, w, g)
}

// bailed on extracting a common emitNodeAssemblerMethodAssignNode: way too many variations.

func (g kindAssemblerTraitsGenerator) emitNodeAssemblerMethodStyle(w io.Writer) {
	doTemplate(`
		func ({{ .AppliedPrefix }}Assembler) Style() ipld.NodeStyle {
			return {{ .AppliedPrefix }}Style{}
		}
	`, w, g)
}

// bailed on extracting a common emitNodeAssemblerOtherBits: it's just self-evident there's nothing common there.
