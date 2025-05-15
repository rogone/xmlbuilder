package xmlbuilder

import (
	"bytes"
	"fmt"
)

type Attr struct {
	Name  string
	Value any //string or interger
}

func (a Attr) Write(buf *bytes.Buffer) error {
	_, err := buf.WriteString(a.Name)
	if err != nil {
		return err
	}
	err = buf.WriteByte('=')
	if err != nil {
		return err
	}
	err = buf.WriteByte('"')
	if err != nil {
		return err
	}

	_, err = buf.WriteString(fmt.Sprintf("%v", a.Value))

	err = buf.WriteByte('"')
	if err != nil {
		return err
	}
	return nil
}
