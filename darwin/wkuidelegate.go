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

// IWKUIDelegate Root Interface
//
//	The methods for presenting native user interface elements on behalf of a webpage.
//	https://developer.apple.com/documentation/webkit/wkuidelegate?language=objc
type IWKUIDelegate interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKUIDelegateProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TWKUIDelegate Root Object
//
//	The methods for presenting native user interface elements on behalf of a webpage.
//	https://developer.apple.com/documentation/webkit/wkuidelegate?language=objc
type TWKUIDelegate struct {
	TObject
}

func NewWKUIDelegate(event IWKUIDelegateEvent) IWKUIDelegate {
	r1 := wKUIDelegateImportAPI().SysCallN(0, GetObjectUintptr(event))
	return AsWKUIDelegate(r1)
}

func (m *TWKUIDelegate) Data() WKUIDelegateProtocol {
	r1 := wKUIDelegateImportAPI().SysCallN(1, m.Instance())
	return WKUIDelegateProtocol(r1)
}

func (m *TWKUIDelegate) Release() {
	wKUIDelegateImportAPI().SysCallN(2, m.Instance())
}

var (
	wKUIDelegateImport       *imports.Imports = nil
	wKUIDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKUIDelegate_Create", 0),
		/*1*/ imports.NewTable("WKUIDelegate_Data", 0),
		/*2*/ imports.NewTable("WKUIDelegate_Release", 0),
	}
)

func wKUIDelegateImportAPI() *imports.Imports {
	if wKUIDelegateImport == nil {
		wKUIDelegateImport = NewDefaultImports()
		wKUIDelegateImport.SetImportTable(wKUIDelegateImportTables)
		wKUIDelegateImportTables = nil
	}
	return wKUIDelegateImport
}
