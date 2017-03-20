package cdptype

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Map returns the headers decoded into a map.
func (n NetworkHeaders) Map() (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal(n, &m)
	return m, err
}

// MustMap panics if the headers cannot be decoded into a map.
func (n NetworkHeaders) MustMap() map[string]string {
	m := make(map[string]string)
	if err := json.Unmarshal(n, &m); err != nil {
		panic(err)
	}
	return m
}

// Error implements error for RuntimeExceptionDetails.
func (r RuntimeExceptionDetails) Error() string {
	var desc string
	if r.Exception.Description != nil {
		desc = ": " + *r.Exception.Description
	}
	return fmt.Sprintf("cdptype.RuntimeExceptionDetails: %s exception at %d:%d%s", r.Text, r.LineNumber, r.ColumnNumber, desc)
}

var (
	_ error = (*RuntimeExceptionDetails)(nil)
)

// String returns a human readable string of a runtime object.
func (r RuntimeRemoteObject) String() string {
	switch r.Type {
	case "undefined":
		return "undefined"
	case "object":
		switch {
		case r.Preview != nil:
			return r.Preview.String()
		case r.UnserializableValue.Valid():
			return r.UnserializableValue.String()
		}
	}

	if len(r.Value) == 0 && r.Description != nil {
		return *r.Description
	}
	if len(r.Value) == 0 {
		return r.Type + "?"
	}

	return string(r.Value)
}

// String returns a human readable string of the object preview.
func (r RuntimeObjectPreview) String() string {
	var stype, desc string
	if r.Subtype != nil {
		stype = *r.Subtype
	}
	if r.Description != nil {
		desc = *r.Description
	}

	var b bytes.Buffer
	switch r.Type {
	case "object":
		switch stype {
		case "null":
			return "null"
		case "array":
			b.WriteByte('[')
			for _, prop := range r.Properties {
				b.WriteString(prop.string(false))
				b.WriteString(", ")
			}
			if b.Len() >= 2 && len(r.Properties) > 0 {
				b.Truncate(b.Len() - 2)
			}
			b.WriteByte(']')
			return b.String()
		case "date", "map", "regexp", "set", "typedarray":
			stype = ""
		default:
			if val, ok := primitiveValue(r.Properties); ok {
				fmt.Fprintf(&b, "%s(%s)", desc, val)
				return b.String()
			}
			if desc == "Object" {
				if len(r.Properties) == 0 {
					return "{}"
				}
				desc = ""
			}
		}
	case "string":
		fmt.Fprintf(&b, "%q", desc)
		return b.String()
	default:
		return desc
	}

	typeAndDesc := stype != "" && desc != ""
	b.WriteString(stype)
	if typeAndDesc {
		b.WriteByte('(')
	}
	b.WriteString(desc)
	if typeAndDesc {
		b.WriteByte(')')
	}

	if len(r.Properties) == 0 && len(r.Entries) == 0 {
		return b.String()
	}

	b.WriteByte('{')
	for _, prop := range r.Properties {
		b.WriteString(prop.String())
		b.WriteString(", ")
	}
	for _, entry := range r.Entries {
		b.WriteString(entry.String())
		b.WriteString(", ")
	}

	if r.Overflow {
		b.WriteString("...")
	} else if b.Len() >= 2 {
		b.Truncate(b.Len() - 2)
	}

	b.WriteByte('}')
	return b.String()
}

// String returns a human readable string of the property.
func (r RuntimePropertyPreview) String() string {
	return r.string(true)
}

func (r RuntimePropertyPreview) string(showName bool) string {
	var b bytes.Buffer
	if showName {
		b.WriteString(r.Name)
		b.WriteString(": ")
	}
	if r.Value != nil {
		if r.Type == "string" {
			fmt.Fprintf(&b, "%q", *r.Value)
		} else {
			b.WriteString(*r.Value)
		}
	}
	if r.ValuePreview != nil {
		b.WriteString(r.ValuePreview.String())
	}
	return b.String()
}

// String returns a human readable string of the entry preview.
func (r RuntimeEntryPreview) String() string {
	var b bytes.Buffer
	if r.Key != nil {
		b.WriteString(r.Key.String())
		b.WriteString(": ")
	}
	b.WriteString(r.Value.String())
	return b.String()
}

const primitiveValueKey = "[[PrimitiveValue]]"

func primitiveValue(props []RuntimePropertyPreview) (string, bool) {
	for _, prop := range props {
		if prop.Name == primitiveValueKey && prop.Value != nil {
			val := *prop.Value
			if prop.Type == "string" {
				val = fmt.Sprintf("%q", val)
			}
			return val, true
		}
	}
	return "", false
}
