//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkURISchemeRequest Parent: IObject
type IWkURISchemeRequest interface {
	IObject
	Data() wvTypes.WebKitURISchemeRequest                                       // function
	Scheme() string                                                             // function
	Uri() string                                                                // function
	Path() string                                                               // function
	WebView() wvTypes.WebKitWebView                                             // function
	Method() string                                                             // function
	Headers() wvTypes.PSoupMessageHeaders                                       // function
	Body() wvTypes.PInputStream                                                 // function
	Finish(stream wvTypes.PInputStream, streamLength int64, contentType string) // procedure
	FinishWithResponse(response wvTypes.WebKitURISchemeResponse)                // procedure
	FinishError(domain int32, code int32, errorMessage string)                  // procedure
}

type TWkURISchemeRequest struct {
	TObject
}

func (m *TWkURISchemeRequest) Data() wvTypes.WebKitURISchemeRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkURISchemeRequestAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitURISchemeRequest(r)
}

func (m *TWkURISchemeRequest) Scheme() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkURISchemeRequestAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkURISchemeRequest) Uri() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkURISchemeRequestAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkURISchemeRequest) Path() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkURISchemeRequestAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkURISchemeRequest) WebView() wvTypes.WebKitWebView {
	if !m.IsValid() {
		return 0
	}
	r := wkURISchemeRequestAPI().SysCallN(5, m.Instance())
	return wvTypes.WebKitWebView(r)
}

func (m *TWkURISchemeRequest) Method() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkURISchemeRequestAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkURISchemeRequest) Headers() wvTypes.PSoupMessageHeaders {
	if !m.IsValid() {
		return 0
	}
	r := wkURISchemeRequestAPI().SysCallN(7, m.Instance())
	return wvTypes.PSoupMessageHeaders(r)
}

func (m *TWkURISchemeRequest) Body() wvTypes.PInputStream {
	if !m.IsValid() {
		return 0
	}
	r := wkURISchemeRequestAPI().SysCallN(8, m.Instance())
	return wvTypes.PInputStream(r)
}

func (m *TWkURISchemeRequest) Finish(stream wvTypes.PInputStream, streamLength int64, contentType string) {
	if !m.IsValid() {
		return
	}
	wkURISchemeRequestAPI().SysCallN(9, m.Instance(), uintptr(stream), uintptr(base.UnsafePointer(&streamLength)), api.PasStr(contentType))
}

func (m *TWkURISchemeRequest) FinishWithResponse(response wvTypes.WebKitURISchemeResponse) {
	if !m.IsValid() {
		return
	}
	wkURISchemeRequestAPI().SysCallN(10, m.Instance(), uintptr(response))
}

func (m *TWkURISchemeRequest) FinishError(domain int32, code int32, errorMessage string) {
	if !m.IsValid() {
		return
	}
	wkURISchemeRequestAPI().SysCallN(11, m.Instance(), uintptr(domain), uintptr(code), api.PasStr(errorMessage))
}

// NewURISchemeRequest class constructor
func NewURISchemeRequest(uRISchemeRequest wvTypes.WebKitURISchemeRequest) IWkURISchemeRequest {
	r := wkURISchemeRequestAPI().SysCallN(0, uintptr(uRISchemeRequest))
	return AsWkURISchemeRequest(r)
}

var (
	wkURISchemeRequestOnce   base.Once
	wkURISchemeRequestImport *imports.Imports = nil
)

func wkURISchemeRequestAPI() *imports.Imports {
	wkURISchemeRequestOnce.Do(func() {
		wkURISchemeRequestImport = api.NewDefaultImports()
		wkURISchemeRequestImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkURISchemeRequest_Create", 0), // constructor NewURISchemeRequest
			/* 1 */ imports.NewTable("TWkURISchemeRequest_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkURISchemeRequest_Scheme", 0), // function Scheme
			/* 3 */ imports.NewTable("TWkURISchemeRequest_Uri", 0), // function Uri
			/* 4 */ imports.NewTable("TWkURISchemeRequest_Path", 0), // function Path
			/* 5 */ imports.NewTable("TWkURISchemeRequest_WebView", 0), // function WebView
			/* 6 */ imports.NewTable("TWkURISchemeRequest_Method", 0), // function Method
			/* 7 */ imports.NewTable("TWkURISchemeRequest_Headers", 0), // function Headers
			/* 8 */ imports.NewTable("TWkURISchemeRequest_Body", 0), // function Body
			/* 9 */ imports.NewTable("TWkURISchemeRequest_Finish", 0), // procedure Finish
			/* 10 */ imports.NewTable("TWkURISchemeRequest_FinishWithResponse", 0), // procedure FinishWithResponse
			/* 11 */ imports.NewTable("TWkURISchemeRequest_FinishError", 0), // procedure FinishError
		}
	})
	return wkURISchemeRequestImport
}
