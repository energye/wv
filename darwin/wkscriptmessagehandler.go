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
//
//	An interface for receiving messages from JavaScript code running in a webpage.
//	https://developer.apple.com/documentation/webkit/wkscriptmessagehandler?language=objc
type IWKScriptMessageHandler interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKScriptMessageHandlerProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TWKScriptMessageHandler Root Object
//
//	An interface for receiving messages from JavaScript code running in a webpage.
//	https://developer.apple.com/documentation/webkit/wkscriptmessagehandler?language=objc
type TWKScriptMessageHandler struct {
	TObject
}

func NewWKScriptMessageHandler(event IReceiveScriptMessageDelegateEvent) IWKScriptMessageHandler {
	r1 := wKScriptMessageHandlerImportAPI().SysCallN(0, GetObjectUintptr(event))
	return AsWKScriptMessageHandler(r1)
}

func (m *TWKScriptMessageHandler) Data() WKScriptMessageHandlerProtocol {
	r1 := wKScriptMessageHandlerImportAPI().SysCallN(1, m.Instance())
	return WKScriptMessageHandlerProtocol(r1)
}

func (m *TWKScriptMessageHandler) Release() {
	wKScriptMessageHandlerImportAPI().SysCallN(2, m.Instance())
}

var (
	wKScriptMessageHandlerImport       *imports.Imports = nil
	wKScriptMessageHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKScriptMessageHandler_Create", 0),
		/*1*/ imports.NewTable("WKScriptMessageHandler_Data", 0),
		/*2*/ imports.NewTable("WKScriptMessageHandler_Release", 0),
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
