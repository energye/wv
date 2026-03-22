//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkUserContentController Parent: IObject
type IWkUserContentController interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKUserContentController // function
	// UserScripts
	//  The user scripts associated with the user content controller.
	UserScripts() lcl.IStrings // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// AddUserScript
	//  Injects the specified script into the webpage’s content.
	AddUserScript(userScript wvTypes.WKUserScript) // procedure
	// RemoveAllUserScripts
	//  Removes all user scripts from the web view.
	RemoveAllUserScripts() // procedure
	// AddScriptMessageHandlerName
	//  Installs a message handler that you can call from your JavaScript code.
	AddScriptMessageHandlerName(scriptMessageHandler IWkScriptMessageHandler, name string) // procedure
	// RemoveScriptMessageHandlerForName
	//  Uninstalls the custom message handler with the specified name from your JavaScript code.
	RemoveScriptMessageHandlerForName(name string) // procedure
}

type TWkUserContentController struct {
	TObject
}

func (m *TWkUserContentController) Data() wvTypes.WKUserContentController {
	if !m.IsValid() {
		return 0
	}
	r := wkUserContentControllerAPI().SysCallN(1, m.Instance())
	return wvTypes.WKUserContentController(r)
}

func (m *TWkUserContentController) UserScripts() lcl.IStrings {
	if !m.IsValid() {
		return nil
	}
	r := wkUserContentControllerAPI().SysCallN(3, m.Instance())
	return lcl.AsStrings(r)
}

func (m *TWkUserContentController) Release() {
	if !m.IsValid() {
		return
	}
	wkUserContentControllerAPI().SysCallN(4, m.Instance())
}

func (m *TWkUserContentController) AddUserScript(userScript wvTypes.WKUserScript) {
	if !m.IsValid() {
		return
	}
	wkUserContentControllerAPI().SysCallN(5, m.Instance(), uintptr(userScript))
}

func (m *TWkUserContentController) RemoveAllUserScripts() {
	if !m.IsValid() {
		return
	}
	wkUserContentControllerAPI().SysCallN(6, m.Instance())
}

func (m *TWkUserContentController) AddScriptMessageHandlerName(scriptMessageHandler IWkScriptMessageHandler, name string) {
	if !m.IsValid() {
		return
	}
	wkUserContentControllerAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(scriptMessageHandler), api.PasStr(name))
}

func (m *TWkUserContentController) RemoveScriptMessageHandlerForName(name string) {
	if !m.IsValid() {
		return
	}
	wkUserContentControllerAPI().SysCallN(8, m.Instance(), api.PasStr(name))
}

// UserContentController  is static instance
var UserContentController _UserContentControllerClass

// _UserContentControllerClass is class type defined by TWkUserContentController
type _UserContentControllerClass uintptr

// New
//
//	Creates and returns an WKUserContentController object.
func (_UserContentControllerClass) New() IWkUserContentController {
	r := wkUserContentControllerAPI().SysCallN(2)
	return AsWkUserContentController(r)
}

// NewUserContentController class constructor
func NewUserContentController(data wvTypes.WKUserContentController) IWkUserContentController {
	r := wkUserContentControllerAPI().SysCallN(0, uintptr(data))
	return AsWkUserContentController(r)
}

var (
	wkUserContentControllerOnce   base.Once
	wkUserContentControllerImport *imports.Imports = nil
)

func wkUserContentControllerAPI() *imports.Imports {
	wkUserContentControllerOnce.Do(func() {
		wkUserContentControllerImport = api.NewDefaultImports()
		wkUserContentControllerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkUserContentController_Create", 0), // constructor NewUserContentController
			/* 1 */ imports.NewTable("TWkUserContentController_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkUserContentController_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkUserContentController_UserScripts", 0), // function UserScripts
			/* 4 */ imports.NewTable("TWkUserContentController_Release", 0), // procedure Release
			/* 5 */ imports.NewTable("TWkUserContentController_AddUserScript", 0), // procedure AddUserScript
			/* 6 */ imports.NewTable("TWkUserContentController_RemoveAllUserScripts", 0), // procedure RemoveAllUserScripts
			/* 7 */ imports.NewTable("TWkUserContentController_AddScriptMessageHandlerName", 0), // procedure AddScriptMessageHandlerName
			/* 8 */ imports.NewTable("TWkUserContentController_RemoveScriptMessageHandlerForName", 0), // procedure RemoveScriptMessageHandlerForName
		}
	})
	return wkUserContentControllerImport
}
