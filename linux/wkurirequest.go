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

// IWkURIRequest Root Interface
type IWkURIRequest interface {
	IObject
	URI() string                  // property
	SetURI(AValue string)         // property
	Data() WebKitURIRequest       // function
	Method() string               // function
	Headers() PSoupMessageHeaders // function
}

// TWkURIRequest Root Object
type TWkURIRequest struct {
	TObject
}

func NewWkURIRequest(aURIRequest WebKitURIRequest) IWkURIRequest {
	r1 := wkURIRequestImportAPI().SysCallN(0, uintptr(aURIRequest))
	return AsWkURIRequest(r1)
}

// WkURIRequestRef -> IWkURIRequest
var WkURIRequestRef wkURIRequest

// wkURIRequest TWkURIRequest Ref
type wkURIRequest uintptr

func (m *wkURIRequest) NewURIRequest(aUri string) IWkURIRequest {
	r1 := wkURIRequestImportAPI().SysCallN(4, PascalStr(aUri))
	return AsWkURIRequest(r1)
}

func (m *TWkURIRequest) URI() string {
	r1 := wkURIRequestImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWkURIRequest) SetURI(AValue string) {
	wkURIRequestImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWkURIRequest) Data() WebKitURIRequest {
	r1 := wkURIRequestImportAPI().SysCallN(1, m.Instance())
	return WebKitURIRequest(r1)
}

func (m *TWkURIRequest) Method() string {
	r1 := wkURIRequestImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TWkURIRequest) Headers() PSoupMessageHeaders {
	r1 := wkURIRequestImportAPI().SysCallN(2, m.Instance())
	return PSoupMessageHeaders(r1)
}

var (
	wkURIRequestImport       *imports.Imports = nil
	wkURIRequestImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkURIRequest_Create", 0),
		/*1*/ imports.NewTable("WkURIRequest_Data", 0),
		/*2*/ imports.NewTable("WkURIRequest_Headers", 0),
		/*3*/ imports.NewTable("WkURIRequest_Method", 0),
		/*4*/ imports.NewTable("WkURIRequest_NewURIRequest", 0),
		/*5*/ imports.NewTable("WkURIRequest_URI", 0),
	}
)

func wkURIRequestImportAPI() *imports.Imports {
	if wkURIRequestImport == nil {
		wkURIRequestImport = NewDefaultImports()
		wkURIRequestImport.SetImportTable(wkURIRequestImportTables)
		wkURIRequestImportTables = nil
	}
	return wkURIRequestImport
}
