package gorequest

import "errors"

// WithHeader is an alias for Set
func (s *SuperAgent) WithHeader(key, value string) *SuperAgent {
	s.Set(key, value)
	return s
}

// WithContentType is a short cut for Set("Content-Type",value)
func (s *SuperAgent) WithContentType(value string) *SuperAgent {
	s.Set("Content-Type", value)
	return s
}

// WithShortType is a short cut for Type(typeStr string)
func (s *SuperAgent) WithShortType(shortType string) *SuperAgent {
	if contentType, ok := Types[shortType]; ok {
		s.WithContentType(contentType)
	} else {
		s.Errors = append(s.Errors, errors.New("WithType func: incorrect type \""+shortType+"\""))
	}
	return s
}

// Accept is a short cut for Set("Accept",value)
func (s *SuperAgent) Accept(value string) *SuperAgent {
	s.Set("Accept", value)
	return s
}
