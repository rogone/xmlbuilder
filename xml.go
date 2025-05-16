package xmlbuilder

const (
	WinNewLine    = "\r\n"
	LinuxNewLine  = "\n"
	XmlNoIntent   = -1
	XmlWithIntent = 0
)

var (
	XmlHead             = `<?xml version="1.0" encoding="UTF-8"?>`
	XmlIntentCount      = 2
	XmlIntentChar  byte = ' '
	NewLine             = LinuxNewLine
)
