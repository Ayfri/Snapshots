package web

import (
	"bytes"
	"html/template"
	"net/http"
)

func (s *Site) renderHTML(p Page, tmpl string, r *http.Request) ([]byte, error) {
	p2 := make(Page)
	for k, v := range p {
		p2[k] = v
	}

	url, ok := p2["URL"].(string)
	if !ok || url == "" {
		p2["URL"] = r.URL.Path
	}
	p2["request"] = r

	base, err := s.readFile(".", tmpl)
	if err != nil {
		return nil, err
	}

	t := template.New(BaseTemplate).Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int { return a / b },
		"cache": func(name string) interface{} {
			if value, exists := s.cache.Get(name); exists {
				return value
			}
			return nil
		},
		"raw":      raw,
		"toString": toString,
	})
	t.Funcs(s.funcs)

	if _, err := t.Parse(string(base)); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, p2); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
