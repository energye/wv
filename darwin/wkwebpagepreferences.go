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

// IWKWebpagePreferences Root Interface
type IWKWebpagePreferences interface {
	IObject
	Data() WKWebpagePreferences                   // function
	PreferredContentMode() WKContentMode          // function
	AllowsContentJavaScript() bool                // function
	SetPreferredContentMode(aValue WKContentMode) // procedure
	SetAllowsContentJavaScript(aValue bool)       // procedure
}

// TWKWebpagePreferences Root Object
type TWKWebpagePreferences struct {
	TObject
}

func NewWKWebpagePreferences(aData WKWebpagePreferences) IWKWebpagePreferences {
	r1 := wKWebpagePreferencesImportAPI().SysCallN(1, uintptr(aData))
	return AsWKWebpagePreferences(r1)
}

// WKWebpagePreferencesRef -> IWKWebpagePreferences
var WKWebpagePreferencesRef wKWebpagePreferences

// wKWebpagePreferences TWKWebpagePreferences Ref
type wKWebpagePreferences uintptr

func (m *wKWebpagePreferences) New() IWKWebpagePreferences {
	r1 := wKWebpagePreferencesImportAPI().SysCallN(3)
	return AsWKWebpagePreferences(r1)
}

func (m *TWKWebpagePreferences) Data() WKWebpagePreferences {
	r1 := wKWebpagePreferencesImportAPI().SysCallN(2, m.Instance())
	return WKWebpagePreferences(r1)
}

func (m *TWKWebpagePreferences) PreferredContentMode() WKContentMode {
	r1 := wKWebpagePreferencesImportAPI().SysCallN(4, m.Instance())
	return WKContentMode(r1)
}

func (m *TWKWebpagePreferences) AllowsContentJavaScript() bool {
	r1 := wKWebpagePreferencesImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TWKWebpagePreferences) SetPreferredContentMode(aValue WKContentMode) {
	wKWebpagePreferencesImportAPI().SysCallN(6, m.Instance(), uintptr(aValue))
}

func (m *TWKWebpagePreferences) SetAllowsContentJavaScript(aValue bool) {
	wKWebpagePreferencesImportAPI().SysCallN(5, m.Instance(), PascalBool(aValue))
}

var (
	wKWebpagePreferencesImport       *imports.Imports = nil
	wKWebpagePreferencesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKWebpagePreferences_AllowsContentJavaScript", 0),
		/*1*/ imports.NewTable("WKWebpagePreferences_Create", 0),
		/*2*/ imports.NewTable("WKWebpagePreferences_Data", 0),
		/*3*/ imports.NewTable("WKWebpagePreferences_New", 0),
		/*4*/ imports.NewTable("WKWebpagePreferences_PreferredContentMode", 0),
		/*5*/ imports.NewTable("WKWebpagePreferences_SetAllowsContentJavaScript", 0),
		/*6*/ imports.NewTable("WKWebpagePreferences_SetPreferredContentMode", 0),
	}
)

func wKWebpagePreferencesImportAPI() *imports.Imports {
	if wKWebpagePreferencesImport == nil {
		wKWebpagePreferencesImport = NewDefaultImports()
		wKWebpagePreferencesImport.SetImportTable(wKWebpagePreferencesImportTables)
		wKWebpagePreferencesImportTables = nil
	}
	return wKWebpagePreferencesImport
}
