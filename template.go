package quark

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/hyperbits/quark/models"
)

func (a *App) Template(templateName, typeName string, data map[string]interface{}) (content string, err error) {
	tmp, err := a.LoadTemplate(templateName, typeName)
	if err != nil {
		return "", err
	}

	funcs := template.FuncMap{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
		"isNotNull": func(t interface{}) bool {
			return t != nil
		},
	}

	var tpl bytes.Buffer
	t, err := template.New(templateName).Funcs(funcs).Parse(tmp)
	if err != nil {
		return "", err
	}
	err = t.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func (a *App) LoadTemplate(name, t string) (text string, err error) {
	var template models.Template
	if result := a.DB.Where("name = ? AND type = ?", name, t).Find(&template); result.Error != nil {
		return "", result.Error
	}
	return template.Content, nil
}

func LoadTemplate(template string) (text string, err error) {
	path := fmt.Sprintf("templates/%s", template)

	data, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		return "", err
	}
	return string(data), nil
}
