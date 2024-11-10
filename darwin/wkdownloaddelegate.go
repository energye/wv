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

// IWKDownloadDelegate Root Interface
//
//	A protocol you implement to track download progress and handle redirects, authentication challenges, and failures.
//	https://developer.apple.com/documentation/webkit/wkdownloaddelegate?language=objc
type IWKDownloadDelegate interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKDownloadDelegateProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TWKDownloadDelegate Root Object
//
//	A protocol you implement to track download progress and handle redirects, authentication challenges, and failures.
//	https://developer.apple.com/documentation/webkit/wkdownloaddelegate?language=objc
type TWKDownloadDelegate struct {
	TObject
}

func NewWKDownloadDelegate(event IWKDownloadDelegateEvent) IWKDownloadDelegate {
	r1 := wKDownloadDelegateImportAPI().SysCallN(0, GetObjectUintptr(event))
	return AsWKDownloadDelegate(r1)
}

func (m *TWKDownloadDelegate) Data() WKDownloadDelegateProtocol {
	r1 := wKDownloadDelegateImportAPI().SysCallN(1, m.Instance())
	return WKDownloadDelegateProtocol(r1)
}

func (m *TWKDownloadDelegate) Release() {
	wKDownloadDelegateImportAPI().SysCallN(2, m.Instance())
}

var (
	wKDownloadDelegateImport       *imports.Imports = nil
	wKDownloadDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKDownloadDelegate_Create", 0),
		/*1*/ imports.NewTable("WKDownloadDelegate_Data", 0),
		/*2*/ imports.NewTable("WKDownloadDelegate_Release", 0),
	}
)

func wKDownloadDelegateImportAPI() *imports.Imports {
	if wKDownloadDelegateImport == nil {
		wKDownloadDelegateImport = NewDefaultImports()
		wKDownloadDelegateImport.SetImportTable(wKDownloadDelegateImportTables)
		wKDownloadDelegateImportTables = nil
	}
	return wKDownloadDelegateImport
}
