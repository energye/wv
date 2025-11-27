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

// IWkURLSchemeHandler Parent: lcl.IObject
type IWkURLSchemeHandler interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKURLSchemeHandlerProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkURLSchemeHandler struct {
	lcl.TObject
}

func (m *TWkURLSchemeHandler) Data() wvTypes.WKURLSchemeHandlerProtocol {
	if !m.IsValid() {
		return 0
	}
	r := wkURLSchemeHandlerAPI().SysCallN(1, m.Instance())
	return wvTypes.WKURLSchemeHandlerProtocol(r)
}

func (m *TWkURLSchemeHandler) Release() {
	if !m.IsValid() {
		return
	}
	wkURLSchemeHandlerAPI().SysCallN(2, m.Instance())
}

// NewURLSchemeHandler class constructor
func NewURLSchemeHandler(event IWkWebview) IWkURLSchemeHandler {
	r := wkURLSchemeHandlerAPI().SysCallN(0, base.GetObjectUintptr(event))
	return AsWkURLSchemeHandler(r)
}

var (
	wkURLSchemeHandlerOnce   base.Once
	wkURLSchemeHandlerImport *imports.Imports = nil
)

func wkURLSchemeHandlerAPI() *imports.Imports {
	wkURLSchemeHandlerOnce.Do(func() {
		wkURLSchemeHandlerImport = api.NewDefaultImports()
		wkURLSchemeHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkURLSchemeHandler_Create", 0), // constructor NewURLSchemeHandler
			/* 1 */ imports.NewTable("TWkURLSchemeHandler_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkURLSchemeHandler_Release", 0), // procedure Release
		}
	})
	return wkURLSchemeHandlerImport
}
