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

// IWKNavigationDelegate Root Interface
//
//	https://developer.apple.com/documentation/webkit/wknavigationdelegate?language=objc
//	Methods for accepting or rejecting navigation changes, and for tracking the progress of navigation requests.
type IWKNavigationDelegate interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKNavigationDelegateProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TWKNavigationDelegate Root Object
//
//	https://developer.apple.com/documentation/webkit/wknavigationdelegate?language=objc
//	Methods for accepting or rejecting navigation changes, and for tracking the progress of navigation requests.
type TWKNavigationDelegate struct {
	TObject
}

func NewWKNavigationDelegate(event IWKNavigationDelegateEvent) IWKNavigationDelegate {
	r1 := wKNavigationDelegateImportAPI().SysCallN(0, GetObjectUintptr(event))
	return AsWKNavigationDelegate(r1)
}

func (m *TWKNavigationDelegate) Data() WKNavigationDelegateProtocol {
	r1 := wKNavigationDelegateImportAPI().SysCallN(1, m.Instance())
	return WKNavigationDelegateProtocol(r1)
}

func (m *TWKNavigationDelegate) Release() {
	wKNavigationDelegateImportAPI().SysCallN(2, m.Instance())
}

var (
	wKNavigationDelegateImport       *imports.Imports = nil
	wKNavigationDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKNavigationDelegate_Create", 0),
		/*1*/ imports.NewTable("WKNavigationDelegate_Data", 0),
		/*2*/ imports.NewTable("WKNavigationDelegate_Release", 0),
	}
)

func wKNavigationDelegateImportAPI() *imports.Imports {
	if wKNavigationDelegateImport == nil {
		wKNavigationDelegateImport = NewDefaultImports()
		wKNavigationDelegateImport.SetImportTable(wKNavigationDelegateImportTables)
		wKNavigationDelegateImportTables = nil
	}
	return wKNavigationDelegateImport
}
