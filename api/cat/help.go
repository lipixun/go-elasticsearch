// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"net/url"
)

// Help - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat.html for more info.
//
// options: optional parameters. Supports the following functional options: WithHelp, WithS, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cat) Help(options ...*Option) (*HelpResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Help"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &HelpResponse{resp}, err
}

// HelpResponse is the response for Help
type HelpResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
