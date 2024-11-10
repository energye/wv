//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// INSHTTPURLResponse Parent: INSURLResponse
//
//	The metadata associated with the response to an HTTP protocol URL load request.
//	https://developer.apple.com/documentation/foundation/nshttpurlresponse?language=objc
type INSHTTPURLResponse interface {
	INSURLResponse
	// InitWithURLStatusCodeHTTPVersionHeaderFields
	//  Initializes an HTTP URL response object with a status code, protocol version, and response headers.
	InitWithURLStatusCodeHTTPVersionHeaderFields(url NSURL, statusCode int32, hTTPVersion string, headerFieldsJSONString string) INSHTTPURLResponse // function
	// StatusCode
	//  The response’s HTTP status code.
	StatusCode() int32 // function
	// AllHeaderFields
	//  All HTTP header fields of the response.
	AllHeaderFields() string // function
}

// TNSHTTPURLResponse Parent: TNSURLResponse
//
//	The metadata associated with the response to an HTTP protocol URL load request.
//	https://developer.apple.com/documentation/foundation/nshttpurlresponse?language=objc
type TNSHTTPURLResponse struct {
	TNSURLResponse
}

func NewNSHTTPURLResponse(aData NSHTTPURLResponse) INSHTTPURLResponse {
	r1 := nSHTTPURLResponseImportAPI().SysCallN(1, uintptr(aData))
	return AsNSHTTPURLResponse(r1)
}

// NSHTTPURLResponseRef -> INSHTTPURLResponse
var NSHTTPURLResponseRef nSHTTPURLResponse

// nSHTTPURLResponse TNSHTTPURLResponse Ref
type nSHTTPURLResponse uintptr

// New
//
//	Creates and returns an TNSHTTPURLResponse object.
func (m *nSHTTPURLResponse) New() INSHTTPURLResponse {
	r1 := nSHTTPURLResponseImportAPI().SysCallN(4)
	return AsNSHTTPURLResponse(r1)
}

// LocalizedStringForStatusCode
//
//	Returns a localized string corresponding to a specified HTTP status code.
//	statusCode: https://www.ietf.org/rfc/rfc2616.txt
func (m *nSHTTPURLResponse) LocalizedStringForStatusCode(statusCode int32) string {
	r1 := nSHTTPURLResponseImportAPI().SysCallN(3, uintptr(statusCode))
	return GoStr(r1)
}

func (m *TNSHTTPURLResponse) InitWithURLStatusCodeHTTPVersionHeaderFields(url NSURL, statusCode int32, hTTPVersion string, headerFieldsJSONString string) INSHTTPURLResponse {
	r1 := nSHTTPURLResponseImportAPI().SysCallN(2, m.Instance(), uintptr(url), uintptr(statusCode), PascalStr(hTTPVersion), PascalStr(headerFieldsJSONString))
	return AsNSHTTPURLResponse(r1)
}

func (m *TNSHTTPURLResponse) StatusCode() int32 {
	r1 := nSHTTPURLResponseImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TNSHTTPURLResponse) AllHeaderFields() string {
	r1 := nSHTTPURLResponseImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

var (
	nSHTTPURLResponseImport       *imports.Imports = nil
	nSHTTPURLResponseImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSHTTPURLResponse_AllHeaderFields", 0),
		/*1*/ imports.NewTable("NSHTTPURLResponse_Create", 0),
		/*2*/ imports.NewTable("NSHTTPURLResponse_InitWithURLStatusCodeHTTPVersionHeaderFields", 0),
		/*3*/ imports.NewTable("NSHTTPURLResponse_LocalizedStringForStatusCode", 0),
		/*4*/ imports.NewTable("NSHTTPURLResponse_New", 0),
		/*5*/ imports.NewTable("NSHTTPURLResponse_StatusCode", 0),
	}
)

func nSHTTPURLResponseImportAPI() *imports.Imports {
	if nSHTTPURLResponseImport == nil {
		nSHTTPURLResponseImport = NewDefaultImports()
		nSHTTPURLResponseImport.SetImportTable(nSHTTPURLResponseImportTables)
		nSHTTPURLResponseImportTables = nil
	}
	return nSHTTPURLResponseImport
}
