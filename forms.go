package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type errors map[string][]string

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$")

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

type Form struct {
	url.Values
	Errors errors
}

func NewForm(form url.Values) *Form {
	return &Form{
		form,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) *Form {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}

	return f
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) MaxLength(field string, n int) *Form {
	value := f.Get(field)
	if value == "" {
		return f
	}

	if utf8.RuneCountInString(value) > n {
		f.Errors.Add(field, fmt.Sprintf("This field is too long, (maximun %d characteres)", n))
	}

	return f
}

func (f *Form) MinLength(field string, n int) *Form {
	value := f.Get(field)
	if value == "" {
		return f
	}

	if utf8.RuneCountInString(value) < n {
		f.Errors.Add(field, fmt.Sprintf("This field is too short, (minimun %d characteres)", n))
	}

	return f
}

func (f *Form) Matches(field string, pattern *regexp.Regexp) *Form {
	value := f.Get(field)
	if value == "" {
		return f
	}

	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}

	return f
}
