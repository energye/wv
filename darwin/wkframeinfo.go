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

// IWKFrameInfo Root Interface
type IWKFrameInfo interface {
	IObject
	Data() WKFrameInfo     // function
	IsMainFrame() bool     // function
	Request() NSURLRequest // function
	Release()              // procedure
}

// TWKFrameInfo Root Object
type TWKFrameInfo struct {
	TObject
}

func NewWKFrameInfo(aData WKFrameInfo) IWKFrameInfo {
	r1 := wKFrameInfoImportAPI().SysCallN(0, uintptr(aData))
	return AsWKFrameInfo(r1)
}

func (m *TWKFrameInfo) Data() WKFrameInfo {
	r1 := wKFrameInfoImportAPI().SysCallN(1, m.Instance())
	return WKFrameInfo(r1)
}

func (m *TWKFrameInfo) IsMainFrame() bool {
	r1 := wKFrameInfoImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TWKFrameInfo) Request() NSURLRequest {
	r1 := wKFrameInfoImportAPI().SysCallN(4, m.Instance())
	return NSURLRequest(r1)
}

func (m *TWKFrameInfo) Release() {
	wKFrameInfoImportAPI().SysCallN(3, m.Instance())
}

var (
	wKFrameInfoImport       *imports.Imports = nil
	wKFrameInfoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKFrameInfo_Create", 0),
		/*1*/ imports.NewTable("WKFrameInfo_Data", 0),
		/*2*/ imports.NewTable("WKFrameInfo_IsMainFrame", 0),
		/*3*/ imports.NewTable("WKFrameInfo_Release", 0),
		/*4*/ imports.NewTable("WKFrameInfo_Request", 0),
	}
)

func wKFrameInfoImportAPI() *imports.Imports {
	if wKFrameInfoImport == nil {
		wKFrameInfoImport = NewDefaultImports()
		wKFrameInfoImport.SetImportTable(wKFrameInfoImportTables)
		wKFrameInfoImportTables = nil
	}
	return wKFrameInfoImport
}
