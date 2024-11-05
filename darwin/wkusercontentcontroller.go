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

// IWKUserContentController Root Interface
type IWKUserContentController interface {
	IObject
	Data() WKUserContentController                                                         // function
	UserScripts() IStrings                                                                 // function
	AddUserScript(userScript WKUserScript)                                                 // procedure
	RemoveAllUserScripts()                                                                 // procedure
	AddScriptMessageHandlerName(scriptMessageHandler IWKScriptMessageHandler, name string) // procedure
	RemoveScriptMessageHandlerForName(name string)                                         // procedure
}

// TWKUserContentController Root Object
type TWKUserContentController struct {
	TObject
}

func NewWKUserContentController() IWKUserContentController {
	r1 := wKUserContentControllerImportAPI().SysCallN(2)
	return AsWKUserContentController(r1)
}

func NewWKUserContentController1(aData WKUserContentController) IWKUserContentController {
	r1 := wKUserContentControllerImportAPI().SysCallN(3, uintptr(aData))
	return AsWKUserContentController(r1)
}

func (m *TWKUserContentController) Data() WKUserContentController {
	r1 := wKUserContentControllerImportAPI().SysCallN(4, m.Instance())
	return WKUserContentController(r1)
}

func (m *TWKUserContentController) UserScripts() IStrings {
	var resultStrings uintptr
	wKUserContentControllerImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultStrings)))
	return AsStrings(resultStrings)
}

func (m *TWKUserContentController) AddUserScript(userScript WKUserScript) {
	wKUserContentControllerImportAPI().SysCallN(1, m.Instance(), uintptr(userScript))
}

func (m *TWKUserContentController) RemoveAllUserScripts() {
	wKUserContentControllerImportAPI().SysCallN(5, m.Instance())
}

func (m *TWKUserContentController) AddScriptMessageHandlerName(scriptMessageHandler IWKScriptMessageHandler, name string) {
	wKUserContentControllerImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(scriptMessageHandler), PascalStr(name))
}

func (m *TWKUserContentController) RemoveScriptMessageHandlerForName(name string) {
	wKUserContentControllerImportAPI().SysCallN(6, m.Instance(), PascalStr(name))
}

var (
	wKUserContentControllerImport       *imports.Imports = nil
	wKUserContentControllerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKUserContentController_AddScriptMessageHandlerName", 0),
		/*1*/ imports.NewTable("WKUserContentController_AddUserScript", 0),
		/*2*/ imports.NewTable("WKUserContentController_Create", 0),
		/*3*/ imports.NewTable("WKUserContentController_Create1", 0),
		/*4*/ imports.NewTable("WKUserContentController_Data", 0),
		/*5*/ imports.NewTable("WKUserContentController_RemoveAllUserScripts", 0),
		/*6*/ imports.NewTable("WKUserContentController_RemoveScriptMessageHandlerForName", 0),
		/*7*/ imports.NewTable("WKUserContentController_UserScripts", 0),
	}
)

func wKUserContentControllerImportAPI() *imports.Imports {
	if wKUserContentControllerImport == nil {
		wKUserContentControllerImport = NewDefaultImports()
		wKUserContentControllerImport.SetImportTable(wKUserContentControllerImportTables)
		wKUserContentControllerImportTables = nil
	}
	return wKUserContentControllerImport
}
