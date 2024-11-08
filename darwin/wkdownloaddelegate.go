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
type IWKDownloadDelegate interface {
	IObject
	Data() WKDownloadDelegateProtocol // function
}

// TWKDownloadDelegate Root Object
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

var (
	wKDownloadDelegateImport       *imports.Imports = nil
	wKDownloadDelegateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKDownloadDelegate_Create", 0),
		/*1*/ imports.NewTable("WKDownloadDelegate_Data", 0),
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
