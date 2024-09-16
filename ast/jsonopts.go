package ast

// compressedJsonOptions compresses the 12 NodeToggle JSON options, plus the 2
// MarshalOptions, for a total of 14 options, down into a single unint16 for
// memory savings.
//
// NOTE(charles): this is a stopgap measure until upstream Go has JOSN v2 in
// the standard library, which seems to have stalled at time of writing. In
// future, we can probably refactor this to be cleaner once the changes land
// upstream.
type compressedJsonOptions uint16

func (o compressedJsonOptions) getMarshalOptions() MarshalOptions {
	return MarshalOptions{
		IncludeLocation: NodeToggle{
			Term:           o.getTerm(),
			Package:        o.getPackage(),
			Comment:        o.getComment(),
			Import:         o.getImport(),
			Rule:           o.getRule(),
			Head:           o.getHead(),
			Expr:           o.getExpr(),
			SomeDecl:       o.getSomeDecl(),
			Every:          o.getEvery(),
			With:           o.getWith(),
			Annotations:    o.getAnnotations(),
			AnnotationsRef: o.getAnnotationsRef(),
		},
		IncludeLocationText: o.getIncludeLocationText(),
		ExcludeLocationFile: o.getExcludeLocationFile(),
	}
}

func (o compressedJsonOptions) getField(n int) bool {
	return (uint16(o) & (1 << n)) == 0
}

func (o compressedJsonOptions) setField(n int, val bool) compressedJsonOptions {
	x := uint16(0)
	if val {
		x = uint16(1)
	}
	return compressedJsonOptions((uint16(o) & (^(1 << n))) | (x << n))
}

func (o compressedJsonOptions) getTerm() bool {
	return o.getField(0)
}

func (o compressedJsonOptions) setTerm(val bool) compressedJsonOptions {
	return o.setField(0, val)
}

func (o compressedJsonOptions) getPackage() bool {
	return o.getField(1)
}

func (o compressedJsonOptions) setPackage(val bool) compressedJsonOptions {
	return o.setField(1, val)
}

func (o compressedJsonOptions) getComment() bool {
	return o.getField(2)
}

func (o compressedJsonOptions) setComment(val bool) compressedJsonOptions {
	return o.setField(2, val)
}

func (o compressedJsonOptions) getImport() bool {
	return o.getField(3)
}

func (o compressedJsonOptions) setImport(val bool) compressedJsonOptions {
	return o.setField(3, val)
}

func (o compressedJsonOptions) getRule() bool {
	return o.getField(4)
}

func (o compressedJsonOptions) setRule(val bool) compressedJsonOptions {
	return o.setField(4, val)
}

func (o compressedJsonOptions) getHead() bool {
	return o.getField(5)
}

func (o compressedJsonOptions) setHead(val bool) compressedJsonOptions {
	return o.setField(5, val)
}

func (o compressedJsonOptions) getExpr() bool {
	return o.getField(6)
}

func (o compressedJsonOptions) setExpr(val bool) compressedJsonOptions {
	return o.setField(6, val)
}

func (o compressedJsonOptions) getSomeDecl() bool {
	return o.getField(7)
}

func (o compressedJsonOptions) setSomeDecl(val bool) compressedJsonOptions {
	return o.setField(7, val)
}

func (o compressedJsonOptions) getEvery() bool {
	return o.getField(8)
}

func (o compressedJsonOptions) setEvery(val bool) compressedJsonOptions {
	return o.setField(8, val)
}

func (o compressedJsonOptions) getWith() bool {
	return o.getField(9)
}

func (o compressedJsonOptions) setWith(val bool) compressedJsonOptions {
	return o.setField(9, val)
}

func (o compressedJsonOptions) getAnnotations() bool {
	return o.getField(10)
}

func (o compressedJsonOptions) setAnnotations(val bool) compressedJsonOptions {
	return o.setField(10, val)
}

func (o compressedJsonOptions) getAnnotationsRef() bool {
	return o.getField(11)
}

func (o compressedJsonOptions) setAnnotationsRef(val bool) compressedJsonOptions {
	return o.setField(11, val)
}

func (o compressedJsonOptions) getIncludeLocationText() bool {
	return o.getField(12)
}

func (o compressedJsonOptions) setIncludeLocationText(val bool) compressedJsonOptions {
	return o.setField(12, val)
}

func (o compressedJsonOptions) getExcludeLocationFile() bool {
	return o.getField(13)
}

func (o compressedJsonOptions) setExcludeLocationFile(val bool) compressedJsonOptions {
	return o.setField(13, val)
}
