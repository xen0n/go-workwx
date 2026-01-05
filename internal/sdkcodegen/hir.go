//go:build sdkcodegen

package main

// High-level IR for the API spec being described.
type hir struct {
	topics []topic
}

// An API topic being described.
type topic struct {
	models []apiModel
	calls  []apiCall

	// TODO: retain source order
	// map[languageTag][]snippet
	inlineCodeSections map[string][]string
}

type visibility int

const (
	visibilityPrivate visibility = 1
	visibilityPublic  visibility = 2
)

// A model used by the APIs.
type apiModel struct {
	ident  string
	doc    string
	vis    visibility
	fields []apiModelField

	// TODO: retain source order
	// map[languageTag][]snippet
	inlineCodeSections map[string][]string
}

type apiModelField struct {
	ident string
	doc   string
	typ   string
	vis   visibility
	tags  map[string]string
}

type apiMethod int

const (
	apiMethodUnknown   apiMethod = 0
	apiMethodGET       apiMethod = 1
	apiMethodGETBinary apiMethod = 2
	apiMethodPOSTJSON  apiMethod = 3
	apiMethodPOSTMedia apiMethod = 4
)

// An API call.
type apiCall struct {
	ident string
	doc   string
	vis   visibility

	reqType  string
	respType string

	needsAccessToken bool

	method  apiMethod
	httpURI string
}
