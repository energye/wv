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

// IWkScriptMessageHandler Parent: lcl.IObject
type IWkScriptMessageHandler interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKScriptMessageHandlerProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkScriptMessageHandler struct {
	lcl.TObject
}

func (m *TWkScriptMessageHandler) Data() wvTypes.WKScriptMessageHandlerProtocol {
	if !m.IsValid() {
		return 0
	}
	r := wkScriptMessageHandlerAPI().SysCallN(1, m.Instance())
	return wvTypes.WKScriptMessageHandlerProtocol(r)
}

func (m *TWkScriptMessageHandler) Release() {
	if !m.IsValid() {
		return
	}
	wkScriptMessageHandlerAPI().SysCallN(2, m.Instance())
}

// NewScriptMessageHandler class constructor
func NewScriptMessageHandler(event IWkWebview) IWkScriptMessageHandler {
	r := wkScriptMessageHandlerAPI().SysCallN(0, base.GetObjectUintptr(event))
	return AsWkScriptMessageHandler(r)
}

var (
	wkScriptMessageHandlerOnce   base.Once
	wkScriptMessageHandlerImport *imports.Imports = nil
)

func wkScriptMessageHandlerAPI() *imports.Imports {
	wkScriptMessageHandlerOnce.Do(func() {
		wkScriptMessageHandlerImport = api.NewDefaultImports()
		wkScriptMessageHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkScriptMessageHandler_Create", 0), // constructor NewScriptMessageHandler
			/* 1 */ imports.NewTable("TWkScriptMessageHandler_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkScriptMessageHandler_Release", 0), // procedure Release
		}
	})
	return wkScriptMessageHandlerImport
}
