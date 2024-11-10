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
//
//	An object for managing interactions between JavaScript code and your web view, and for filtering content in your web view.
//	https://developer.apple.com/documentation/webkit/wkusercontentcontroller?language=objc
type IWKUserContentController interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKUserContentController // function
	// UserScripts
	//  The user scripts associated with the user content controller.
	UserScripts() IStrings // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// AddUserScript
	//  Injects the specified script into the webpage’s content.
	AddUserScript(userScript WKUserScript) // procedure
	// RemoveAllUserScripts
	//  Removes all user scripts from the web view.
	RemoveAllUserScripts() // procedure
	// AddScriptMessageHandlerName
	//  Installs a message handler that you can call from your JavaScript code.
	AddScriptMessageHandlerName(scriptMessageHandler IWKScriptMessageHandler, name string) // procedure
	// RemoveScriptMessageHandlerForName
	//  Uninstalls the custom message handler with the specified name from your JavaScript code.
	RemoveScriptMessageHandlerForName(name string) // procedure
}

// TWKUserContentController Root Object
//
//	An object for managing interactions between JavaScript code and your web view, and for filtering content in your web view.
//	https://developer.apple.com/documentation/webkit/wkusercontentcontroller?language=objc
type TWKUserContentController struct {
	TObject
}

func NewWKUserContentController(aData WKUserContentController) IWKUserContentController {
	r1 := wKUserContentControllerImportAPI().SysCallN(2, uintptr(aData))
	return AsWKUserContentController(r1)
}

// WKUserContentControllerRef -> IWKUserContentController
var WKUserContentControllerRef wKUserContentController

// wKUserContentController TWKUserContentController Ref
type wKUserContentController uintptr

// New
//
//	Creates and returns an WKUserContentController object.
func (m *wKUserContentController) New() IWKUserContentController {
	r1 := wKUserContentControllerImportAPI().SysCallN(4)
	return AsWKUserContentController(r1)
}

func (m *TWKUserContentController) Data() WKUserContentController {
	r1 := wKUserContentControllerImportAPI().SysCallN(3, m.Instance())
	return WKUserContentController(r1)
}

func (m *TWKUserContentController) UserScripts() IStrings {
	var resultStrings uintptr
	wKUserContentControllerImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&resultStrings)))
	return AsStrings(resultStrings)
}

func (m *TWKUserContentController) Release() {
	wKUserContentControllerImportAPI().SysCallN(5, m.Instance())
}

func (m *TWKUserContentController) AddUserScript(userScript WKUserScript) {
	wKUserContentControllerImportAPI().SysCallN(1, m.Instance(), uintptr(userScript))
}

func (m *TWKUserContentController) RemoveAllUserScripts() {
	wKUserContentControllerImportAPI().SysCallN(6, m.Instance())
}

func (m *TWKUserContentController) AddScriptMessageHandlerName(scriptMessageHandler IWKScriptMessageHandler, name string) {
	wKUserContentControllerImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(scriptMessageHandler), PascalStr(name))
}

func (m *TWKUserContentController) RemoveScriptMessageHandlerForName(name string) {
	wKUserContentControllerImportAPI().SysCallN(7, m.Instance(), PascalStr(name))
}

var (
	wKUserContentControllerImport       *imports.Imports = nil
	wKUserContentControllerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKUserContentController_AddScriptMessageHandlerName", 0),
		/*1*/ imports.NewTable("WKUserContentController_AddUserScript", 0),
		/*2*/ imports.NewTable("WKUserContentController_Create", 0),
		/*3*/ imports.NewTable("WKUserContentController_Data", 0),
		/*4*/ imports.NewTable("WKUserContentController_New", 0),
		/*5*/ imports.NewTable("WKUserContentController_Release", 0),
		/*6*/ imports.NewTable("WKUserContentController_RemoveAllUserScripts", 0),
		/*7*/ imports.NewTable("WKUserContentController_RemoveScriptMessageHandlerForName", 0),
		/*8*/ imports.NewTable("WKUserContentController_UserScripts", 0),
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
