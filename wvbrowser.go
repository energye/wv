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

// IWVBrowser Parent: IWVBrowserBase
//
//	VCL and LCL version of TWVBrowserBase that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type IWVBrowser interface {
	IWVBrowserBase
}

// TWVBrowser Parent: TWVBrowserBase
//
//	VCL and LCL version of TWVBrowserBase that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type TWVBrowser struct {
	TWVBrowserBase
}

func NewWVBrowser(AOwner IComponent) IWVBrowser {
	r1 := wVBrowserImportAPI().SysCallN(0, GetObjectUintptr(AOwner))
	return AsWVBrowser(r1)
}

var (
	wVBrowserImport       *imports.Imports = nil
	wVBrowserImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WVBrowser_Create", 0),
	}
)

func wVBrowserImportAPI() *imports.Imports {
	if wVBrowserImport == nil {
		wVBrowserImport = NewDefaultImports()
		wVBrowserImport.SetImportTable(wVBrowserImportTables)
		wVBrowserImportTables = nil
	}
	return wVBrowserImport
}
