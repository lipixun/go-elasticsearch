// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// DeleteScript - the scripting module enables you to use scripts to evaluate custom expressions. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-scripting.html for more info.
//
// lang: script language.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) DeleteScript(lang string, options ...*Option) (*DeleteScriptResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "DELETE",
	}
	methodOptions := supportedOptions["DeleteScript"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &DeleteScriptResponse{resp}, err
}

// DeleteScriptResponse is the response for DeleteScript
type DeleteScriptResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
