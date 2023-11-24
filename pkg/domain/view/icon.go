package view

import "net/url"

type Icon struct {
	ImageInt    int
	ImageString string
}

type RenderIcon interface {
	Display() url.URL
}

func (i *Icon) Display() url.URL {
	return url.URL{}
}
