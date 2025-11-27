//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
)

// IWkNSURLRequest Parent: lcl.IObject
type IWkNSURLRequest interface {
	lcl.IObject
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
	TimeoutInterval() float64 // function
	// MainDocumentURL
	//  The main document URL associated with the request.
	MainDocumentURL() NSURL // function
	// NetworkServiceType
	//  The network service type of the request.
	NetworkServiceType() NSURLRequestNetworkServiceType // function
	// AllowsCellularAccess
	//  A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
	AllowsCellularAccess() bool // function
	// HTTPMethod
	//  The HTTP request method.
	HTTPMethod() string // function
	// AllHTTPHeaderFields
	//  A dictionary containing all of the HTTP header fields for a request. JSON string.
	AllHTTPHeaderFields() string // function
	// ValueForHTTPHeaderField
	//  Returns the value of the specified HTTP header field.
	ValueForHTTPHeaderField(field string) string // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNSURLRequest struct {
	lcl.TObject
}

func (m *TWkNSURLRequest) Data() NSURLRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLRequestAPI().SysCallN(1, m.Instance())
	return NSURLRequest(r)
}

func (m *TWkNSURLRequest) URL() NSURL {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLRequestAPI().SysCallN(6, m.Instance())
	return NSURL(r)
}

func (m *TWkNSURLRequest) CachePolicy() NSURLRequestCachePolicy {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLRequestAPI().SysCallN(7, m.Instance())
	return NSURLRequestCachePolicy(r)
}

func (m *TWkNSURLRequest) TimeoutInterval() (result float64) {
	if !m.IsValid() {
		return
	}
	wkNSURLRequestAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkNSURLRequest) MainDocumentURL() NSURL {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLRequestAPI().SysCallN(9, m.Instance())
	return NSURL(r)
}

func (m *TWkNSURLRequest) NetworkServiceType() NSURLRequestNetworkServiceType {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLRequestAPI().SysCallN(10, m.Instance())
	return NSURLRequestNetworkServiceType(r)
}

func (m *TWkNSURLRequest) AllowsCellularAccess() bool {
	if !m.IsValid() {
		return false
	}
	r := wkNSURLRequestAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TWkNSURLRequest) HTTPMethod() string {
	if !m.IsValid() {
		return ""
	}
	r := wkNSURLRequestAPI().SysCallN(12, m.Instance())
	return api.GoStr(r)
}

func (m *TWkNSURLRequest) AllHTTPHeaderFields() string {
	if !m.IsValid() {
		return ""
	}
	r := wkNSURLRequestAPI().SysCallN(13, m.Instance())
	return api.GoStr(r)
}

func (m *TWkNSURLRequest) ValueForHTTPHeaderField(field string) string {
	if !m.IsValid() {
		return ""
	}
	r := wkNSURLRequestAPI().SysCallN(14, m.Instance(), api.PasStr(field))
	return api.GoStr(r)
}

func (m *TWkNSURLRequest) Release() {
	if !m.IsValid() {
		return
	}
	wkNSURLRequestAPI().SysCallN(15, m.Instance())
}

// URLRequest  is static instance
var URLRequest _URLRequestClass

// _URLRequestClass is class type defined by TWkNSURLRequest
type _URLRequestClass uintptr

// New
//
//	Creates and returns an NSURLRequest object.
func (_URLRequestClass) New() IWkNSURLRequest {
	r := wkNSURLRequestAPI().SysCallN(2)
	return AsWkNSURLRequest(r)
}

// RequestWithURL
//
//	Creates and returns a URL request for a specified URL.
func (_URLRequestClass) RequestWithURL(uRL NSURL) IWkNSURLRequest {
	r := wkNSURLRequestAPI().SysCallN(3, uintptr(uRL))
	return AsWkNSURLRequest(r)
}

// SupportsSecureCoding
//
//	A Boolean value indicating whether the NSURLRequest implements the NSSecureCoding protocol.
func (_URLRequestClass) SupportsSecureCoding() bool {
	r := wkNSURLRequestAPI().SysCallN(4)
	return api.GoBool(r)
}

// RequestWithURLCachePolicyTimeoutInterval
//
//	Creates and returns an initialized URL request with specified URL, cache policy, and timeout values.
func (_URLRequestClass) RequestWithURLCachePolicyTimeoutInterval(uRL NSURL, cachePolicy NSURLRequestCachePolicy, timeoutInterval NSTimeInterval) IWkNSURLRequest {
	r := wkNSURLRequestAPI().SysCallN(5, uintptr(uRL), uintptr(cachePolicy), uintptr(base.UnsafePointer(&timeoutInterval)))
	return AsWkNSURLRequest(r)
}

// NewURLRequest class constructor
func NewURLRequest(data NSURLRequest) IWkNSURLRequest {
	r := wkNSURLRequestAPI().SysCallN(0, uintptr(data))
	return AsWkNSURLRequest(r)
}

var (
	wkNSURLRequestOnce   base.Once
	wkNSURLRequestImport *imports.Imports = nil
)

func wkNSURLRequestAPI() *imports.Imports {
	wkNSURLRequestOnce.Do(func() {
		wkNSURLRequestImport = api.NewDefaultImports()
		wkNSURLRequestImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNSURLRequest_Create", 0), // constructor NewURLRequest
			/* 1 */ imports.NewTable("TWkNSURLRequest_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNSURLRequest_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkNSURLRequest_RequestWithURL", 0), // static function RequestWithURL
			/* 4 */ imports.NewTable("TWkNSURLRequest_SupportsSecureCoding", 0), // static function SupportsSecureCoding
			/* 5 */ imports.NewTable("TWkNSURLRequest_RequestWithURLCachePolicyTimeoutInterval", 0), // static function RequestWithURLCachePolicyTimeoutInterval
			/* 6 */ imports.NewTable("TWkNSURLRequest_URL", 0), // function URL
			/* 7 */ imports.NewTable("TWkNSURLRequest_CachePolicy", 0), // function CachePolicy
			/* 8 */ imports.NewTable("TWkNSURLRequest_TimeoutInterval", 0), // function TimeoutInterval
			/* 9 */ imports.NewTable("TWkNSURLRequest_MainDocumentURL", 0), // function MainDocumentURL
			/* 10 */ imports.NewTable("TWkNSURLRequest_NetworkServiceType", 0), // function NetworkServiceType
			/* 11 */ imports.NewTable("TWkNSURLRequest_AllowsCellularAccess", 0), // function AllowsCellularAccess
			/* 12 */ imports.NewTable("TWkNSURLRequest_HTTPMethod", 0), // function HTTPMethod
			/* 13 */ imports.NewTable("TWkNSURLRequest_AllHTTPHeaderFields", 0), // function AllHTTPHeaderFields
			/* 14 */ imports.NewTable("TWkNSURLRequest_ValueForHTTPHeaderField", 0), // function ValueForHTTPHeaderField
			/* 15 */ imports.NewTable("TWkNSURLRequest_Release", 0), // procedure Release
		}
	})
	return wkNSURLRequestImport
}
