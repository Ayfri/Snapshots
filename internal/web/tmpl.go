package web

import (
	"VersionCraft2/internal"
	"fmt"
	"html/template"
)

func toString(v any) string {
	switch v := v.(type) {
	case string:
		return v
	case template.HTML:
		return string(v)
	case nil:
		return ""
	default:
		internal.Fatal(fmt.Errorf("cannot toString %T", v))
	}
	return ""
}

func raw(s interface{}) template.HTML {
	return template.HTML(toString(s))
}
