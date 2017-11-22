// +build go1.9

// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol/network"
)

// NavigateReply represents the return values for Navigate in the Page domain.
type NavigateReply struct {
	FrameID   FrameID           `json:"frameId"`             // Frame id that has navigated (or failed to navigate)
	LoaderID  *network.LoaderID `json:"loaderId,omitempty"`  // Loader identifier.
	ErrorText *string           `json:"errorText,omitempty"` // User friendly error message, present if and only if navigation has failed.
}

// GetResourceContentArgs represents the arguments for GetResourceContent in the Page domain.
type GetResourceContentArgs struct {
	FrameID FrameID `json:"frameId"` // Frame id to get resource for.
	URL     string  `json:"url"`     // URL of the resource to get content for.
}

// NewGetResourceContentArgs initializes GetResourceContentArgs with the required arguments.
func NewGetResourceContentArgs(frameID FrameID, url string) *GetResourceContentArgs {
	args := new(GetResourceContentArgs)
	args.FrameID = frameID
	args.URL = url
	return args
}

// SearchInResourceArgs represents the arguments for SearchInResource in the Page domain.
type SearchInResourceArgs struct {
	FrameID       FrameID `json:"frameId"`                 // Frame id for resource to search in.
	URL           string  `json:"url"`                     // URL of the resource to search in.
	Query         string  `json:"query"`                   // String to search for.
	CaseSensitive *bool   `json:"caseSensitive,omitempty"` // If true, search is case sensitive.
	IsRegex       *bool   `json:"isRegex,omitempty"`       // If true, treats string parameter as regex.
}

// NewSearchInResourceArgs initializes SearchInResourceArgs with the required arguments.
func NewSearchInResourceArgs(frameID FrameID, url string, query string) *SearchInResourceArgs {
	args := new(SearchInResourceArgs)
	args.FrameID = frameID
	args.URL = url
	args.Query = query
	return args
}

// SetCaseSensitive sets the CaseSensitive optional argument. If true, search is case sensitive.
func (a *SearchInResourceArgs) SetCaseSensitive(caseSensitive bool) *SearchInResourceArgs {
	a.CaseSensitive = &caseSensitive
	return a
}

// SetIsRegex sets the IsRegex optional argument. If true, treats string parameter as regex.
func (a *SearchInResourceArgs) SetIsRegex(isRegex bool) *SearchInResourceArgs {
	a.IsRegex = &isRegex
	return a
}

// SetDocumentContentArgs represents the arguments for SetDocumentContent in the Page domain.
type SetDocumentContentArgs struct {
	FrameID FrameID `json:"frameId"` // Frame id to set HTML for.
	HTML    string  `json:"html"`    // HTML content to set.
}

// NewSetDocumentContentArgs initializes SetDocumentContentArgs with the required arguments.
func NewSetDocumentContentArgs(frameID FrameID, html string) *SetDocumentContentArgs {
	args := new(SetDocumentContentArgs)
	args.FrameID = frameID
	args.HTML = html
	return args
}

// CreateIsolatedWorldArgs represents the arguments for CreateIsolatedWorld in the Page domain.
type CreateIsolatedWorldArgs struct {
	FrameID             FrameID `json:"frameId"`                       // Id of the frame in which the isolated world should be created.
	WorldName           *string `json:"worldName,omitempty"`           // An optional name which is reported in the Execution Context.
	GrantUniveralAccess *bool   `json:"grantUniveralAccess,omitempty"` // Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
}

// NewCreateIsolatedWorldArgs initializes CreateIsolatedWorldArgs with the required arguments.
func NewCreateIsolatedWorldArgs(frameID FrameID) *CreateIsolatedWorldArgs {
	args := new(CreateIsolatedWorldArgs)
	args.FrameID = frameID
	return args
}

// SetWorldName sets the WorldName optional argument. An optional name which is reported in the Execution Context.
func (a *CreateIsolatedWorldArgs) SetWorldName(worldName string) *CreateIsolatedWorldArgs {
	a.WorldName = &worldName
	return a
}

// SetGrantUniveralAccess sets the GrantUniveralAccess optional argument. Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
func (a *CreateIsolatedWorldArgs) SetGrantUniveralAccess(grantUniveralAccess bool) *CreateIsolatedWorldArgs {
	a.GrantUniveralAccess = &grantUniveralAccess
	return a
}
