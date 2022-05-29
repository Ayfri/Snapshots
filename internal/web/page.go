package web

import (
	"io/fs"
)

type Page map[string]any

func (s *Site) openPage(file string) (*Page, error) {
	var body []byte
	var err error
	if file == "." {
		file = BaseTemplate
	}

	stat, err := fs.Stat(s.fs, file)
	if err == nil {
		body, err = s.readFile(".", file)
	}
	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		return nil, nil
	}

	page := make(Page)
	page["body"] = string(body)
	page["File"] = file

	return &page, nil
}
