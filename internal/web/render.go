package web

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path"
)

func (s *Site) renderHTML(p Page, tmpl string, r *http.Request) ([]byte, error) {
	p2 := make(Page)
	for k, v := range p {
		p2[k] = v
	}
	p = p2

	url, ok := p["URL"].(string)
	if !ok || url == "" {
		p["URL"] = r.URL.Path
	}
	p["request"] = r

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

	t, err = t.ParseFS(s.fs, "components/*.tmpl")
	if err != nil {
		return nil, err
	}

	if _, err := t.Parse(string(base)); err != nil {
		return nil, err
	}

	layout, _ := p["layout"].(string)
	if layout == "" {
		l, ok := s.findLayout(".", "404")
		if ok {
			layout = l
		} else {
			layout = "none"
		}
	} else {
		l, ok := s.findLayout(".", layout)
		if !ok {
			return nil, fmt.Errorf("layout not found: %s", layout)
		}
		layout = l
	}

	if layout != "none" {
		base, err = s.readFile(".", layout)
		if err != nil {
			return nil, err
		}

		t, err = t.Parse(string(base))
		if err != nil {
			return nil, err
		}
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, p); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *Site) findLayout(dir, name string) (string, bool) {
	name += ".tmpl"
	for {
		abs := path.Join(dir, name)
		if _, err := fs.Stat(s.fs, abs); err == nil {
			return abs, true
		}

		if dir == "." {
			return "", false
		}

		dir = path.Join(dir)
	}
}
