package presenter

import (
	"strings"
)

const (
	unknownFormat   format = "unknown"
	jsonFormat      format = "json"
	tableFormat     format = "table"
	cycloneDXFormat format = "cyclonedx"
	sarifFormat     format = "sarif"
	templateFormat  format = "template"
	embeddedVEXJSON format = "embedded-cyclonedx-vex-json"
	embeddedVEXXML  format = "embedded-cyclonedx-vex-xml"
)

// format is a dedicated type to represent a specific kind of presenter output format.
type format string

func (f format) String() string {
	return string(f)
}

// parse returns the presenter.format specified by the given user input.
func parse(userInput string) format {
	switch strings.ToLower(userInput) {
	case "":
		return tableFormat
	case strings.ToLower(jsonFormat.String()):
		return jsonFormat
	case strings.ToLower(tableFormat.String()):
		return tableFormat
	case strings.ToLower(cycloneDXFormat.String()):
		return cycloneDXFormat
	case strings.ToLower(sarifFormat.String()):
		return sarifFormat
	case strings.ToLower(templateFormat.String()):
		return templateFormat
	case strings.ToLower(embeddedVEXJSON.String()):
		return embeddedVEXJSON
	case strings.ToLower(embeddedVEXXML.String()):
		return embeddedVEXXML
	default:
		return unknownFormat
	}
}

// AvailableFormats is a list of presenter format options available to users.
var AvailableFormats = []format{
	jsonFormat,
	tableFormat,
	cycloneDXFormat,
	sarifFormat,
	templateFormat,
	embeddedVEXJSON,
	embeddedVEXXML,
}
