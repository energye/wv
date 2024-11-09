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

// IWKDownload Root Interface
type IWKDownload interface {
	IObject
	Data() WKDownload                                 // function
	OriginalRequest() NSURLRequest                    // function
	Progress() NSProgress                             // function
	Release()                                         // procedure
	SetDelegate(downloadDelegate IWKDownloadDelegate) // procedure
	Cancel()                                          // procedure
}

// TWKDownload Root Object
type TWKDownload struct {
	TObject
}

func NewWKDownload(aData WKDownload, aWebViewDelegate IWKDownloadDelegateEvent) IWKDownload {
	r1 := wKDownloadImportAPI().SysCallN(1, uintptr(aData), GetObjectUintptr(aWebViewDelegate))
	return AsWKDownload(r1)
}

func (m *TWKDownload) Data() WKDownload {
	r1 := wKDownloadImportAPI().SysCallN(2, m.Instance())
	return WKDownload(r1)
}

func (m *TWKDownload) OriginalRequest() NSURLRequest {
	r1 := wKDownloadImportAPI().SysCallN(3, m.Instance())
	return NSURLRequest(r1)
}

func (m *TWKDownload) Progress() NSProgress {
	r1 := wKDownloadImportAPI().SysCallN(4, m.Instance())
	return NSProgress(r1)
}

func (m *TWKDownload) Release() {
	wKDownloadImportAPI().SysCallN(5, m.Instance())
}

func (m *TWKDownload) SetDelegate(downloadDelegate IWKDownloadDelegate) {
	wKDownloadImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(downloadDelegate))
}

func (m *TWKDownload) Cancel() {
	wKDownloadImportAPI().SysCallN(0, m.Instance())
}

var (
	wKDownloadImport       *imports.Imports = nil
	wKDownloadImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKDownload_Cancel", 0),
		/*1*/ imports.NewTable("WKDownload_Create", 0),
		/*2*/ imports.NewTable("WKDownload_Data", 0),
		/*3*/ imports.NewTable("WKDownload_OriginalRequest", 0),
		/*4*/ imports.NewTable("WKDownload_Progress", 0),
		/*5*/ imports.NewTable("WKDownload_Release", 0),
		/*6*/ imports.NewTable("WKDownload_SetDelegate", 0),
	}
)

func wKDownloadImportAPI() *imports.Imports {
	if wKDownloadImport == nil {
		wKDownloadImport = NewDefaultImports()
		wKDownloadImport.SetImportTable(wKDownloadImportTables)
		wKDownloadImportTables = nil
	}
	return wKDownloadImport
}
