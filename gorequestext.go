package gorequest

// WithHeader is an alias for Set
func (s *SuperAgent) WithHeader(key, value string) *SuperAgent {
	s.Set(key, value)
	return s
}

// WithContent is a short cut for Set("Content-Type",value)
func (s *SuperAgent) WithContent(value string) *SuperAgent {
	s.Set("Content-Type", value)
	return s
}

// Accept is a short cut for Set("Accept",value)
func (s *SuperAgent) Accept(value string) *SuperAgent {
	s.Set("Accept", value)
	return s
}
