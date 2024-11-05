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

// IWKScriptMessageHandler Root Interface
type IWKScriptMessageHandler interface {
	IObject
	Data() WKScriptMessageHandlerProtocol // function
}

// TWKScriptMessageHandler Root Object
type TWKScriptMessageHandler struct {
	TObject
}

func NewWKScriptMessageHandler(event IReceiveScriptMessageEvent) IWKScriptMessageHandler {
	r1 := wKScriptMessageHandlerImportAPI().SysCallN(0, GetObjectUintptr(event))
	return AsWKScriptMessageHandler(r1)
}

func (m *TWKScriptMessageHandler) Data() WKScriptMessageHandlerProtocol {
	r1 := wKScriptMessageHandlerImportAPI().SysCallN(1, m.Instance())
	return WKScriptMessageHandlerProtocol(r1)
}

var (
	wKScriptMessageHandlerImport       *imports.Imports = nil
	wKScriptMessageHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKScriptMessageHandler_Create", 0),
		/*1*/ imports.NewTable("WKScriptMessageHandler_Data", 0),
	}
)

func wKScriptMessageHandlerImportAPI() *imports.Imports {
	if wKScriptMessageHandlerImport == nil {
		wKScriptMessageHandlerImport = NewDefaultImports()
		wKScriptMessageHandlerImport.SetImportTable(wKScriptMessageHandlerImportTables)
		wKScriptMessageHandlerImportTables = nil
	}
	return wKScriptMessageHandlerImport
}
