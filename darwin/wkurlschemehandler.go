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
type IWKURLSchemeHandler interface {
	IObject
	Data() WKURLSchemeHandlerProtocol // function
}

// TWKURLSchemeHandler Root Object
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

var (
	wKURLSchemeHandlerImport       *imports.Imports = nil
	wKURLSchemeHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKURLSchemeHandler_Create", 0),
		/*1*/ imports.NewTable("WKURLSchemeHandler_Data", 0),
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
