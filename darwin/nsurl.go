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

// INSURL Root Interface
//
//	An object that represents the location of a resource, such as an item on a remote server or the path to a local file.
//	https://developer.apple.com/documentation/foundation/nsurl?language=objc
type INSURL interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSURL // function
	// AbsoluteString
	//  The URL string for the receiver as an absolute URL.(read-only)
	AbsoluteString() string // function
	// RelativeString
	//  A string representation of the relative portion of the URL.(read-only)
	RelativeString() string // function
	// BaseURL
	//  The base URL.(read-only)
	BaseURL() INSURL // function
	// AbsoluteURL
	//  An absolute URL that refers to the same resource as the receiver.(read-only)
	AbsoluteURL() INSURL // function
	// Scheme
	//  The scheme.(read-only)
	Scheme() string // function
	// Host
	//  The host, conforming to RFC 1808.(read-only)
	Host() string // function
	// Port
	//  The port, conforming to RFC 1808.
	Port() int32 // function
	// User
	//  The user name, conforming to RFC 1808.
	User() string // function
	// Password
	//  The password conforming to RFC 1808.(read-only)
	Password() string // function
	// Path
	//  The path, conforming to RFC 1808.(read-only)
	Path() string // function
	// Fragment
	//  The fragment identifier, conforming to RFC 1808.(read-only)
	Fragment() string // function
	// ParameterString
	//  The parameter string conforming to RFC 1808.(read-only)
	//  Deprecated
	ParameterString() string // function
	// Query
	//  The query string, conforming to RFC 1808.
	Query() string // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TNSURL Root Object
//
//	An object that represents the location of a resource, such as an item on a remote server or the path to a local file.
//	https://developer.apple.com/documentation/foundation/nsurl?language=objc
type TNSURL struct {
	TObject
}

func NewNSURL(aData NSURL) INSURL {
	r1 := nSURLImportAPI().SysCallN(3, uintptr(aData))
	return AsNSURL(r1)
}

// NSURLRef -> INSURL
var NSURLRef nSURL

// nSURL TNSURL Ref
type nSURL uintptr

// New
//
//	Creates and returns an NSURL object.
func (m *nSURL) New() INSURL {
	r1 := nSURLImportAPI().SysCallN(7)
	return AsNSURL(r1)
}

// URLWithString
//
//	Creates and returns an NSURL object initialized with a provided URL string.
func (m *nSURL) URLWithString(uRLString string) INSURL {
	r1 := nSURLImportAPI().SysCallN(16, PascalStr(uRLString))
	return AsNSURL(r1)
}

// URLWithStringRelativeToURL
//
//	Creates and returns an NSURL object initialized with a base URL and a relative string.
func (m *nSURL) URLWithStringRelativeToURL(uRLString string, baseURL NSURL) INSURL {
	r1 := nSURLImportAPI().SysCallN(17, PascalStr(uRLString), uintptr(baseURL))
	return AsNSURL(r1)
}

func (m *TNSURL) Data() NSURL {
	r1 := nSURLImportAPI().SysCallN(4, m.Instance())
	return NSURL(r1)
}

func (m *TNSURL) AbsoluteString() string {
	r1 := nSURLImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) RelativeString() string {
	r1 := nSURLImportAPI().SysCallN(13, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) BaseURL() INSURL {
	r1 := nSURLImportAPI().SysCallN(2, m.Instance())
	return AsNSURL(r1)
}

func (m *TNSURL) AbsoluteURL() INSURL {
	r1 := nSURLImportAPI().SysCallN(1, m.Instance())
	return AsNSURL(r1)
}

func (m *TNSURL) Scheme() string {
	r1 := nSURLImportAPI().SysCallN(15, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Host() string {
	r1 := nSURLImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Port() int32 {
	r1 := nSURLImportAPI().SysCallN(11, m.Instance())
	return int32(r1)
}

func (m *TNSURL) User() string {
	r1 := nSURLImportAPI().SysCallN(18, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Password() string {
	r1 := nSURLImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Path() string {
	r1 := nSURLImportAPI().SysCallN(10, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Fragment() string {
	r1 := nSURLImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) ParameterString() string {
	r1 := nSURLImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Query() string {
	r1 := nSURLImportAPI().SysCallN(12, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Release() {
	nSURLImportAPI().SysCallN(14, m.Instance())
}

var (
	nSURLImport       *imports.Imports = nil
	nSURLImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSURL_AbsoluteString", 0),
		/*1*/ imports.NewTable("NSURL_AbsoluteURL", 0),
		/*2*/ imports.NewTable("NSURL_BaseURL", 0),
		/*3*/ imports.NewTable("NSURL_Create", 0),
		/*4*/ imports.NewTable("NSURL_Data", 0),
		/*5*/ imports.NewTable("NSURL_Fragment", 0),
		/*6*/ imports.NewTable("NSURL_Host", 0),
		/*7*/ imports.NewTable("NSURL_New", 0),
		/*8*/ imports.NewTable("NSURL_ParameterString", 0),
		/*9*/ imports.NewTable("NSURL_Password", 0),
		/*10*/ imports.NewTable("NSURL_Path", 0),
		/*11*/ imports.NewTable("NSURL_Port", 0),
		/*12*/ imports.NewTable("NSURL_Query", 0),
		/*13*/ imports.NewTable("NSURL_RelativeString", 0),
		/*14*/ imports.NewTable("NSURL_Release", 0),
		/*15*/ imports.NewTable("NSURL_Scheme", 0),
		/*16*/ imports.NewTable("NSURL_URLWithString", 0),
		/*17*/ imports.NewTable("NSURL_URLWithStringRelativeToURL", 0),
		/*18*/ imports.NewTable("NSURL_User", 0),
	}
)

func nSURLImportAPI() *imports.Imports {
	if nSURLImport == nil {
		nSURLImport = NewDefaultImports()
		nSURLImport.SetImportTable(nSURLImportTables)
		nSURLImportTables = nil
	}
	return nSURLImport
}
