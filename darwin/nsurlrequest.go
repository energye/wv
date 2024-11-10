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

// INSURLRequest Root Interface
//
//	A URL load request that is independent of protocol or URL scheme.
//	https://developer.apple.com/documentation/foundation/nsurlrequest?language=objc
type INSURLRequest interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSURLRequest // function
	// URL
	//  The URL being requested.
	URL() NSURL // function
	// CachePolicy
	//  The request’s cache policy.
	CachePolicy() NSURLRequestCachePolicy // function
	// TimeoutInterval
	//  The request’s timeout interval, in seconds.
	TimeoutInterval() (resultFloat64 float64) // function
	// MainDocumentURL
	//  The main document URL associated with the request.
	MainDocumentURL() NSURL // function
	// NetworkServiceType
	//  The network service type of the request.
	NetworkServiceType() NSURLRequestNetworkServiceType // function
	// AllowsCellularAccess
	//  A Boolean value that indicates whether the request is allowed to use the cellular radio(if present).
	AllowsCellularAccess() bool // function
	// HTTPMethod
	//  The HTTP request method.
	HTTPMethod() string // function
	// AllHTTPHeaderFields
	//  A dictionary containing all of the HTTP header fields for a request. JSON string.
	AllHTTPHeaderFields() string // function
	// ValueForHTTPHeaderField
	//  Returns the value of the specified HTTP header field.
	ValueForHTTPHeaderField(aField string) string // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TNSURLRequest Root Object
//
//	A URL load request that is independent of protocol or URL scheme.
//	https://developer.apple.com/documentation/foundation/nsurlrequest?language=objc
type TNSURLRequest struct {
	TObject
}

func NewNSURLRequest(aData NSURLRequest) INSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(3, uintptr(aData))
	return AsNSURLRequest(r1)
}

// NSURLRequestRef -> INSURLRequest
var NSURLRequestRef nSURLRequest

// nSURLRequest TNSURLRequest Ref
type nSURLRequest uintptr

// New
//
//	Creates and returns an NSURLRequest object.
func (m *nSURLRequest) New() INSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(8)
	return AsNSURLRequest(r1)
}

// RequestWithURL
//
//	Creates and returns a URL request for a specified URL.
func (m *nSURLRequest) RequestWithURL(uRL NSURL) INSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(10, uintptr(uRL))
	return AsNSURLRequest(r1)
}

// SupportsSecureCoding
//
//	A Boolean value indicating whether the NSURLRequest implements the NSSecureCoding protocol.
func (m *nSURLRequest) SupportsSecureCoding() bool {
	r1 := nSURLRequestImportAPI().SysCallN(12)
	return GoBool(r1)
}

// RequestWithURLCachePolicyTimeoutInterval
//
//	Creates and returns an initialized URL request with specified URL, cache policy, and timeout values.
func (m *nSURLRequest) RequestWithURLCachePolicyTimeoutInterval(uRL NSURL, cachePolicy NSURLRequestCachePolicy, timeoutInterval NSTimeInterval) INSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(11, uintptr(uRL), uintptr(cachePolicy), uintptr(unsafePointer(&timeoutInterval)))
	return AsNSURLRequest(r1)
}

func (m *TNSURLRequest) Data() NSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(4, m.Instance())
	return NSURLRequest(r1)
}

func (m *TNSURLRequest) URL() NSURL {
	r1 := nSURLRequestImportAPI().SysCallN(14, m.Instance())
	return NSURL(r1)
}

func (m *TNSURLRequest) CachePolicy() NSURLRequestCachePolicy {
	r1 := nSURLRequestImportAPI().SysCallN(2, m.Instance())
	return NSURLRequestCachePolicy(r1)
}

func (m *TNSURLRequest) TimeoutInterval() (resultFloat64 float64) {
	nSURLRequestImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TNSURLRequest) MainDocumentURL() NSURL {
	r1 := nSURLRequestImportAPI().SysCallN(6, m.Instance())
	return NSURL(r1)
}

func (m *TNSURLRequest) NetworkServiceType() NSURLRequestNetworkServiceType {
	r1 := nSURLRequestImportAPI().SysCallN(7, m.Instance())
	return NSURLRequestNetworkServiceType(r1)
}

func (m *TNSURLRequest) AllowsCellularAccess() bool {
	r1 := nSURLRequestImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TNSURLRequest) HTTPMethod() string {
	r1 := nSURLRequestImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TNSURLRequest) AllHTTPHeaderFields() string {
	r1 := nSURLRequestImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TNSURLRequest) ValueForHTTPHeaderField(aField string) string {
	r1 := nSURLRequestImportAPI().SysCallN(15, m.Instance(), PascalStr(aField))
	return GoStr(r1)
}

func (m *TNSURLRequest) Release() {
	nSURLRequestImportAPI().SysCallN(9, m.Instance())
}

var (
	nSURLRequestImport       *imports.Imports = nil
	nSURLRequestImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSURLRequest_AllHTTPHeaderFields", 0),
		/*1*/ imports.NewTable("NSURLRequest_AllowsCellularAccess", 0),
		/*2*/ imports.NewTable("NSURLRequest_CachePolicy", 0),
		/*3*/ imports.NewTable("NSURLRequest_Create", 0),
		/*4*/ imports.NewTable("NSURLRequest_Data", 0),
		/*5*/ imports.NewTable("NSURLRequest_HTTPMethod", 0),
		/*6*/ imports.NewTable("NSURLRequest_MainDocumentURL", 0),
		/*7*/ imports.NewTable("NSURLRequest_NetworkServiceType", 0),
		/*8*/ imports.NewTable("NSURLRequest_New", 0),
		/*9*/ imports.NewTable("NSURLRequest_Release", 0),
		/*10*/ imports.NewTable("NSURLRequest_RequestWithURL", 0),
		/*11*/ imports.NewTable("NSURLRequest_RequestWithURLCachePolicyTimeoutInterval", 0),
		/*12*/ imports.NewTable("NSURLRequest_SupportsSecureCoding", 0),
		/*13*/ imports.NewTable("NSURLRequest_TimeoutInterval", 0),
		/*14*/ imports.NewTable("NSURLRequest_URL", 0),
		/*15*/ imports.NewTable("NSURLRequest_ValueForHTTPHeaderField", 0),
	}
)

func nSURLRequestImportAPI() *imports.Imports {
	if nSURLRequestImport == nil {
		nSURLRequestImport = NewDefaultImports()
		nSURLRequestImport.SetImportTable(nSURLRequestImportTables)
		nSURLRequestImportTables = nil
	}
	return nSURLRequestImport
}
