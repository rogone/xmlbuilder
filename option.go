package xmlbuilder

type Option func(*Element)

func CloseSingleTag(b bool) Option {
	return func(e *Element) {
		e.IsSingleClosedTag = b
	}
}

func Text(text any) Option {
	return func(e *Element) {
		e.HasText = true
		e.Text = text
	}
}

func AddAttr(name string, value any) Option {
	return func(e *Element) {
		e.Attrs = append(e.Attrs, Attr{
			Name:  name,
			Value: value,
		})
	}
}

func SubElement(tagName string, ops ...Option) Option {
	return func(e *Element) {
		sub := newElement(tagName, ops...)
		e.SubElement = append(e.SubElement, sub)
	}
}
