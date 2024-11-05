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
type INSURL interface {
	IObject
	Data() NSURL             // function
	AbsoluteString() string  // function
	RelativeString() string  // function
	BaseURL() INSURL         // function
	AbsoluteURL() INSURL     // function
	Scheme() string          // function
	Host() string            // function
	Port() int32             // function
	User() string            // function
	Password() string        // function
	Path() string            // function
	Fragment() string        // function
	ParameterString() string // function
	Query() string           // function
}

// TNSURL Root Object
type TNSURL struct {
	TObject
}

func NewNSURL() INSURL {
	r1 := nSURLImportAPI().SysCallN(3)
	return AsNSURL(r1)
}

func NewNSURL1(aData NSURL) INSURL {
	r1 := nSURLImportAPI().SysCallN(4, uintptr(aData))
	return AsNSURL(r1)
}

// NSURLRef -> INSURL
var NSURLRef nSURL

// nSURL TNSURL Ref
type nSURL uintptr

func (m *nSURL) URLWithString(uRLString string) INSURL {
	r1 := nSURLImportAPI().SysCallN(15, PascalStr(uRLString))
	return AsNSURL(r1)
}

func (m *nSURL) URLWithStringRelativeToURL(uRLString string, baseURL NSURL) INSURL {
	r1 := nSURLImportAPI().SysCallN(16, PascalStr(uRLString), uintptr(baseURL))
	return AsNSURL(r1)
}

func (m *TNSURL) Data() NSURL {
	r1 := nSURLImportAPI().SysCallN(5, m.Instance())
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
	r1 := nSURLImportAPI().SysCallN(14, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Host() string {
	r1 := nSURLImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TNSURL) Port() int32 {
	r1 := nSURLImportAPI().SysCallN(11, m.Instance())
	return int32(r1)
}

func (m *TNSURL) User() string {
	r1 := nSURLImportAPI().SysCallN(17, m.Instance())
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
	r1 := nSURLImportAPI().SysCallN(6, m.Instance())
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

var (
	nSURLImport       *imports.Imports = nil
	nSURLImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSURL_AbsoluteString", 0),
		/*1*/ imports.NewTable("NSURL_AbsoluteURL", 0),
		/*2*/ imports.NewTable("NSURL_BaseURL", 0),
		/*3*/ imports.NewTable("NSURL_Create", 0),
		/*4*/ imports.NewTable("NSURL_Create1", 0),
		/*5*/ imports.NewTable("NSURL_Data", 0),
		/*6*/ imports.NewTable("NSURL_Fragment", 0),
		/*7*/ imports.NewTable("NSURL_Host", 0),
		/*8*/ imports.NewTable("NSURL_ParameterString", 0),
		/*9*/ imports.NewTable("NSURL_Password", 0),
		/*10*/ imports.NewTable("NSURL_Path", 0),
		/*11*/ imports.NewTable("NSURL_Port", 0),
		/*12*/ imports.NewTable("NSURL_Query", 0),
		/*13*/ imports.NewTable("NSURL_RelativeString", 0),
		/*14*/ imports.NewTable("NSURL_Scheme", 0),
		/*15*/ imports.NewTable("NSURL_URLWithString", 0),
		/*16*/ imports.NewTable("NSURL_URLWithStringRelativeToURL", 0),
		/*17*/ imports.NewTable("NSURL_User", 0),
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
