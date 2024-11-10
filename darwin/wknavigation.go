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

// IWKNavigation Root Interface
type IWKNavigation interface {
	IObject
	Data() WKNavigation                  // function
	EffectiveContentMode() WKContentMode // function
	// Release
	//  Release the current object and Data pointer
	Release() // procedure
}

// TWKNavigation Root Object
type TWKNavigation struct {
	TObject
}

func NewWKNavigation(aData WKNavigation) IWKNavigation {
	r1 := wKNavigationImportAPI().SysCallN(0, uintptr(aData))
	return AsWKNavigation(r1)
}

func (m *TWKNavigation) Data() WKNavigation {
	r1 := wKNavigationImportAPI().SysCallN(1, m.Instance())
	return WKNavigation(r1)
}

func (m *TWKNavigation) EffectiveContentMode() WKContentMode {
	r1 := wKNavigationImportAPI().SysCallN(2, m.Instance())
	return WKContentMode(r1)
}

func (m *TWKNavigation) Release() {
	wKNavigationImportAPI().SysCallN(3, m.Instance())
}

var (
	wKNavigationImport       *imports.Imports = nil
	wKNavigationImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKNavigation_Create", 0),
		/*1*/ imports.NewTable("WKNavigation_Data", 0),
		/*2*/ imports.NewTable("WKNavigation_EffectiveContentMode", 0),
		/*3*/ imports.NewTable("WKNavigation_Release", 0),
	}
)

func wKNavigationImportAPI() *imports.Imports {
	if wKNavigationImport == nil {
		wKNavigationImport = NewDefaultImports()
		wKNavigationImport.SetImportTable(wKNavigationImportTables)
		wKNavigationImportTables = nil
	}
	return wKNavigationImport
}
