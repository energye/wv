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

// INSURLResponse Root Interface
type INSURLResponse interface {
	IObject
	Data() NSURLResponse                                                                                                           // function
	InitWithURLMIMETypeExpectedContentLengthTextEncodingName(uRL NSURL, mIMEType string, length int32, name string) INSURLResponse // function
	URL() NSURL                                                                                                                    // function
	MIMEType() string                                                                                                              // function
	ExpectedContentLength() (resultInt64 int64)                                                                                    // function
	TextEncodingName() string                                                                                                      // function
	SuggestedFilename() string                                                                                                     // function
}

// TNSURLResponse Root Object
type TNSURLResponse struct {
	TObject
}

func NewNSURLResponse(aData NSURLResponse) INSURLResponse {
	r1 := nSURLResponseImportAPI().SysCallN(0, uintptr(aData))
	return AsNSURLResponse(r1)
}

// NSURLResponseRef -> INSURLResponse
var NSURLResponseRef nSURLResponse

// nSURLResponse TNSURLResponse Ref
type nSURLResponse uintptr

func (m *nSURLResponse) New() INSURLResponse {
	r1 := nSURLResponseImportAPI().SysCallN(5)
	return AsNSURLResponse(r1)
}

func (m *TNSURLResponse) Data() NSURLResponse {
	r1 := nSURLResponseImportAPI().SysCallN(1, m.Instance())
	return NSURLResponse(r1)
}

func (m *TNSURLResponse) InitWithURLMIMETypeExpectedContentLengthTextEncodingName(uRL NSURL, mIMEType string, length int32, name string) INSURLResponse {
	r1 := nSURLResponseImportAPI().SysCallN(3, m.Instance(), uintptr(uRL), PascalStr(mIMEType), uintptr(length), PascalStr(name))
	return AsNSURLResponse(r1)
}

func (m *TNSURLResponse) URL() NSURL {
	r1 := nSURLResponseImportAPI().SysCallN(8, m.Instance())
	return NSURL(r1)
}

func (m *TNSURLResponse) MIMEType() string {
	r1 := nSURLResponseImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

func (m *TNSURLResponse) ExpectedContentLength() (resultInt64 int64) {
	nSURLResponseImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TNSURLResponse) TextEncodingName() string {
	r1 := nSURLResponseImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TNSURLResponse) SuggestedFilename() string {
	r1 := nSURLResponseImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

var (
	nSURLResponseImport       *imports.Imports = nil
	nSURLResponseImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSURLResponse_Create", 0),
		/*1*/ imports.NewTable("NSURLResponse_Data", 0),
		/*2*/ imports.NewTable("NSURLResponse_ExpectedContentLength", 0),
		/*3*/ imports.NewTable("NSURLResponse_InitWithURLMIMETypeExpectedContentLengthTextEncodingName", 0),
		/*4*/ imports.NewTable("NSURLResponse_MIMEType", 0),
		/*5*/ imports.NewTable("NSURLResponse_New", 0),
		/*6*/ imports.NewTable("NSURLResponse_SuggestedFilename", 0),
		/*7*/ imports.NewTable("NSURLResponse_TextEncodingName", 0),
		/*8*/ imports.NewTable("NSURLResponse_URL", 0),
	}
)

func nSURLResponseImportAPI() *imports.Imports {
	if nSURLResponseImport == nil {
		nSURLResponseImport = NewDefaultImports()
		nSURLResponseImport.SetImportTable(nSURLResponseImportTables)
		nSURLResponseImportTables = nil
	}
	return nSURLResponseImport
}
