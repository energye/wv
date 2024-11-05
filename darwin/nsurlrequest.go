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
type INSURLRequest interface {
	IObject
	Data() NSURLRequest                                 // function
	URL() NSURL                                         // function
	CachePolicy() NSURLRequestCachePolicy               // function
	TimeoutInterval() (resultFloat64 float64)           // function
	MainDocumentURL() NSURL                             // function
	NetworkServiceType() NSURLRequestNetworkServiceType // function
	AllowsCellularAccess() bool                         // function
	HTTPMethod() string                                 // function
	AllHTTPHeaderFields() string                        // function
	ValueForHTTPHeaderField(aField string) string       // function
}

// TNSURLRequest Root Object
type TNSURLRequest struct {
	TObject
}

func NewNSURLRequest() INSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(3)
	return AsNSURLRequest(r1)
}

func NewNSURLRequest1(aData NSURLRequest) INSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(4, uintptr(aData))
	return AsNSURLRequest(r1)
}

// NSURLRequestRef -> INSURLRequest
var NSURLRequestRef nSURLRequest

// nSURLRequest TNSURLRequest Ref
type nSURLRequest uintptr

func (m *nSURLRequest) RequestWithURL(uRL NSURL) INSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(9, uintptr(uRL))
	return AsNSURLRequest(r1)
}

func (m *nSURLRequest) SupportsSecureCoding() bool {
	r1 := nSURLRequestImportAPI().SysCallN(11)
	return GoBool(r1)
}

func (m *nSURLRequest) RequestWithURLCachePolicyTimeoutInterval(uRL NSURL, cachePolicy NSURLRequestCachePolicy, timeoutInterval NSTimeInterval) INSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(10, uintptr(uRL), uintptr(cachePolicy), uintptr(unsafePointer(&timeoutInterval)))
	return AsNSURLRequest(r1)
}

func (m *TNSURLRequest) Data() NSURLRequest {
	r1 := nSURLRequestImportAPI().SysCallN(5, m.Instance())
	return NSURLRequest(r1)
}

func (m *TNSURLRequest) URL() NSURL {
	r1 := nSURLRequestImportAPI().SysCallN(13, m.Instance())
	return NSURL(r1)
}

func (m *TNSURLRequest) CachePolicy() NSURLRequestCachePolicy {
	r1 := nSURLRequestImportAPI().SysCallN(2, m.Instance())
	return NSURLRequestCachePolicy(r1)
}

func (m *TNSURLRequest) TimeoutInterval() (resultFloat64 float64) {
	nSURLRequestImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TNSURLRequest) MainDocumentURL() NSURL {
	r1 := nSURLRequestImportAPI().SysCallN(7, m.Instance())
	return NSURL(r1)
}

func (m *TNSURLRequest) NetworkServiceType() NSURLRequestNetworkServiceType {
	r1 := nSURLRequestImportAPI().SysCallN(8, m.Instance())
	return NSURLRequestNetworkServiceType(r1)
}

func (m *TNSURLRequest) AllowsCellularAccess() bool {
	r1 := nSURLRequestImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TNSURLRequest) HTTPMethod() string {
	r1 := nSURLRequestImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TNSURLRequest) AllHTTPHeaderFields() string {
	r1 := nSURLRequestImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TNSURLRequest) ValueForHTTPHeaderField(aField string) string {
	r1 := nSURLRequestImportAPI().SysCallN(14, m.Instance(), PascalStr(aField))
	return GoStr(r1)
}

var (
	nSURLRequestImport       *imports.Imports = nil
	nSURLRequestImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSURLRequest_AllHTTPHeaderFields", 0),
		/*1*/ imports.NewTable("NSURLRequest_AllowsCellularAccess", 0),
		/*2*/ imports.NewTable("NSURLRequest_CachePolicy", 0),
		/*3*/ imports.NewTable("NSURLRequest_Create", 0),
		/*4*/ imports.NewTable("NSURLRequest_Create1", 0),
		/*5*/ imports.NewTable("NSURLRequest_Data", 0),
		/*6*/ imports.NewTable("NSURLRequest_HTTPMethod", 0),
		/*7*/ imports.NewTable("NSURLRequest_MainDocumentURL", 0),
		/*8*/ imports.NewTable("NSURLRequest_NetworkServiceType", 0),
		/*9*/ imports.NewTable("NSURLRequest_RequestWithURL", 0),
		/*10*/ imports.NewTable("NSURLRequest_RequestWithURLCachePolicyTimeoutInterval", 0),
		/*11*/ imports.NewTable("NSURLRequest_SupportsSecureCoding", 0),
		/*12*/ imports.NewTable("NSURLRequest_TimeoutInterval", 0),
		/*13*/ imports.NewTable("NSURLRequest_URL", 0),
		/*14*/ imports.NewTable("NSURLRequest_ValueForHTTPHeaderField", 0),
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
