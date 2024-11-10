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

// IWKURLSchemeHandler Root Interface
//
//	A protocol for loading resources with URL schemes that WebKit doesn't handle.
//	https://developer.apple.com/documentation/webkit/wkurlschemehandler?language=objc
type IWKURLSchemeHandler interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKURLSchemeHandlerProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TWKURLSchemeHandler Root Object
//
//	A protocol for loading resources with URL schemes that WebKit doesn't handle.
//	https://developer.apple.com/documentation/webkit/wkurlschemehandler?language=objc
type TWKURLSchemeHandler struct {
	TObject
}

func NewWKURLSchemeHandler(event IWKURLSchemeHandlerDelegateEvent) IWKURLSchemeHandler {
	r1 := wKURLSchemeHandlerImportAPI().SysCallN(0, GetObjectUintptr(event))
	return AsWKURLSchemeHandler(r1)
}

func (m *TWKURLSchemeHandler) Data() WKURLSchemeHandlerProtocol {
	r1 := wKURLSchemeHandlerImportAPI().SysCallN(1, m.Instance())
	return WKURLSchemeHandlerProtocol(r1)
}

func (m *TWKURLSchemeHandler) Release() {
	wKURLSchemeHandlerImportAPI().SysCallN(2, m.Instance())
}

var (
	wKURLSchemeHandlerImport       *imports.Imports = nil
	wKURLSchemeHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKURLSchemeHandler_Create", 0),
		/*1*/ imports.NewTable("WKURLSchemeHandler_Data", 0),
		/*2*/ imports.NewTable("WKURLSchemeHandler_Release", 0),
	}
)

func wKURLSchemeHandlerImportAPI() *imports.Imports {
	if wKURLSchemeHandlerImport == nil {
		wKURLSchemeHandlerImport = NewDefaultImports()
		wKURLSchemeHandlerImport.SetImportTable(wKURLSchemeHandlerImportTables)
		wKURLSchemeHandlerImportTables = nil
	}
	return wKURLSchemeHandlerImport
}
