package main

import (
	"os"
	"strings"
	"text/template"
)

var scheme_file = "file://"

func generate(from, exec string) error {
	logger.Infof("parsing %s", from)
	var t *template.Template
	var err error
	if strings.HasPrefix(from, scheme_file) {
		path := strings.TrimPrefix(from, scheme_file)
		t, err = template.ParseFiles(path)
	}
	if err != nil {
		return err
	}
	if exec != "" {
		t.ExecuteTemplate(os.Stdout, exec, nil)
	} else {
		t.Execute(os.Stdout, nil)
	}
	return err
}
