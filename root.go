package xmlbuilder

import (
	"bytes"
)

func NewRoot(tagName string, ops ...Option) *Root {
	root := &Root{Element{TagName: tagName}}
	for _, f := range ops {
		f(&root.Element)
	}
	return root
}

type Root struct {
	Element
}

func (r *Root) IntentWrite(intent int, buf *bytes.Buffer) (err error) {
	_, err = buf.WriteString(XmlHead)
	if err != nil {
		return err
	}

	if intent >= 0 {
		_, err = buf.WriteString(NewLine)
		if err != nil {
			return err
		}
	}

	return r.Element.IntentWrite(intent, buf)
}
