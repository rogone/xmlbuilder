package xmlbuilder

import (
	"bytes"
	"fmt"
)

func NewElement(tagName string, ops ...Option) *Element {
	return newElement(tagName, ops...)
}

type Element struct {
	TagName           string
	IsSingleClosedTag bool
	HasText           bool

	Attrs      []Attr_
	Text       any
	SubElement []*Element
}

func (e *Element) IntentWrite(intent int, buf *bytes.Buffer) (err error) {
	if intent > 0 {
		err = e.writeNewLine(buf)
		if err != nil {
			return err
		}

		err = e.writeIntent(intent, buf)
		if err != nil {
			return err
		}
	}

	err = buf.WriteByte('<')
	if err != nil {
		return err
	}

	_, err = buf.WriteString(e.TagName)
	if err != nil {
		return err
	}

	if len(e.Attrs) > 0 {
		err = buf.WriteByte(' ')
		if err != nil {
			return err
		}

		err = e.writeAttrs(buf)
		if err != nil {
			return err
		}
	}

	isSingleTag := e.IsSingleClosedTag && !e.HasText && len(e.SubElement) == 0
	if isSingleTag {
		_, err = buf.Write([]byte{'/', '>'})
		return err
	}

	err = buf.WriteByte('>')
	if err != nil {
		return err
	}

	if e.HasText {
		_, err = buf.WriteString(fmt.Sprintf("%v", e.Text))
		if err != nil {
			return err
		}
	}

	subIntent := intent
	if intent >= 0 {
		subIntent += 1
	}

	for _, it := range e.SubElement {
		err = it.IntentWrite(subIntent, buf)
		if err != nil {
			return err
		}
	}

	return e.writeCloseTag(intent, buf)

}

func (e *Element) writeCloseTag(intent int, buf *bytes.Buffer) (err error) {
	if intent >= 0 && len(e.SubElement) > 0 {
		err = e.writeNewLine(buf)
		if err != nil {
			return err
		}

		if intent > 0 {
			err = e.writeIntent(intent, buf)
			if err != nil {
				return err
			}
		}
	}

	_, err = buf.Write([]byte{'<', '/'})
	if err != nil {
		return err
	}

	_, err = buf.WriteString(e.TagName)
	if err != nil {
		return err
	}

	err = buf.WriteByte('>')
	if err != nil {
		return err
	}
	return nil
}

func (e *Element) writeAttrs(buf *bytes.Buffer) (err error) {
	count := len(e.Attrs)
	for i := range e.Attrs {
		it := &e.Attrs[i]
		err = it.Write(buf)
		if err != nil {
			return err
		}

		if i+1 < count {
			err = buf.WriteByte(' ')
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *Element) writeNewLine(buf *bytes.Buffer) (err error) {
	_, err = buf.WriteString(NewLine)
	return err
}

func (e *Element) writeIntent(intent int, buf *bytes.Buffer) (err error) {
	_, err = buf.Write(bytes.Repeat([]byte{XmlIntentChar}, intent*XmlIntentCount))
	return
}

func newElement(tagName string, ops ...Option) *Element {
	e := &Element{TagName: tagName}
	for _, f := range ops {
		f(e)
	}
	return e
}
