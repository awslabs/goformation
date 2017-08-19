package main

import (
	"fmt"
	"net/url"
)

// ResourceGenerator takes AWS CloudFormation Resource Specification
// documents, and generates Go structs and a JSON Schema from them.
type ResourceGenerator struct {
	filenames []string
	urls      []string
}

// NewResourceGenerator an array of AWS CloudFormation Resource Specification
// documents, and generates Go structs and a JSON Schema from them.
// The input can be a mix of URLs (https://) or files (file://).
func NewResourceGenerator(urls []string) (*ResourceGenerator, error) {

	rg := &ResourceGenerator{}

	for _, found := range urls {

		uri, err := url.Parse(found)
		if err != nil {
			return nil, err
		}

		switch uri.Scheme {
		case "https":
			rg.urls = append(rg.urls, uri.String())
		case "http":
			rg.urls = append(rg.urls, uri.String())
		case "file":
			rg.filenames = append(rg.filenames, uri.String())
		default:
			return nil, fmt.Errorf("invalid URL scheme %s", uri.Scheme)
		}

	}

	return rg, nil

}
