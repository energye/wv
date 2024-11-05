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

// IWKPreferences Root Interface
type IWKPreferences interface {
	IObject
	Data() WKPreferences                                    // function
	MinimumFontSize() (resultFloat64 float64)               // function
	JavaScriptEnabled() bool                                // function
	JavaScriptCanOpenWindowsAutomatically() bool            // function
	JavaEnabled() bool                                      // function
	PlugInsEnabled() bool                                   // function
	SetMinimumFontSize(newValue float64)                    // procedure
	SetJavaScriptEnabled(newValue bool)                     // procedure
	SetJavaScriptCanOpenWindowsAutomatically(newValue bool) // procedure
	SetJavaEnabled(newValue bool)                           // procedure
	SetPlugInsEnabled(newValue bool)                        // procedure
	SetTabFocusesLinks(newValue bool)                       // procedure
	SetFraudulentWebsiteWarningEnabled(newValue bool)       // procedure
	// SetValueForKey
	//  procedure SetTextInteractionEnabled(newValue: LongBool); // 11.3
	//  procedure SetElementFullscreenEnabled(newValue: LongBool); // 12.3
	SetValueForKey(value bool, key string) // procedure
}

// TWKPreferences Root Object
type TWKPreferences struct {
	TObject
}

func NewWKPreferences() IWKPreferences {
	r1 := wKPreferencesImportAPI().SysCallN(0)
	return AsWKPreferences(r1)
}

func NewWKPreferences1(aData WKPreferences) IWKPreferences {
	r1 := wKPreferencesImportAPI().SysCallN(1, uintptr(aData))
	return AsWKPreferences(r1)
}

func (m *TWKPreferences) Data() WKPreferences {
	r1 := wKPreferencesImportAPI().SysCallN(2, m.Instance())
	return WKPreferences(r1)
}

func (m *TWKPreferences) MinimumFontSize() (resultFloat64 float64) {
	wKPreferencesImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TWKPreferences) JavaScriptEnabled() bool {
	r1 := wKPreferencesImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TWKPreferences) JavaScriptCanOpenWindowsAutomatically() bool {
	r1 := wKPreferencesImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TWKPreferences) JavaEnabled() bool {
	r1 := wKPreferencesImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TWKPreferences) PlugInsEnabled() bool {
	r1 := wKPreferencesImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TWKPreferences) SetMinimumFontSize(newValue float64) {
	wKPreferencesImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&newValue)))
}

func (m *TWKPreferences) SetJavaScriptEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(11, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetJavaScriptCanOpenWindowsAutomatically(newValue bool) {
	wKPreferencesImportAPI().SysCallN(10, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetJavaEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(9, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetPlugInsEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(13, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetTabFocusesLinks(newValue bool) {
	wKPreferencesImportAPI().SysCallN(14, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetFraudulentWebsiteWarningEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(8, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetValueForKey(value bool, key string) {
	wKPreferencesImportAPI().SysCallN(15, m.Instance(), PascalBool(value), PascalStr(key))
}

var (
	wKPreferencesImport       *imports.Imports = nil
	wKPreferencesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKPreferences_Create", 0),
		/*1*/ imports.NewTable("WKPreferences_Create1", 0),
		/*2*/ imports.NewTable("WKPreferences_Data", 0),
		/*3*/ imports.NewTable("WKPreferences_JavaEnabled", 0),
		/*4*/ imports.NewTable("WKPreferences_JavaScriptCanOpenWindowsAutomatically", 0),
		/*5*/ imports.NewTable("WKPreferences_JavaScriptEnabled", 0),
		/*6*/ imports.NewTable("WKPreferences_MinimumFontSize", 0),
		/*7*/ imports.NewTable("WKPreferences_PlugInsEnabled", 0),
		/*8*/ imports.NewTable("WKPreferences_SetFraudulentWebsiteWarningEnabled", 0),
		/*9*/ imports.NewTable("WKPreferences_SetJavaEnabled", 0),
		/*10*/ imports.NewTable("WKPreferences_SetJavaScriptCanOpenWindowsAutomatically", 0),
		/*11*/ imports.NewTable("WKPreferences_SetJavaScriptEnabled", 0),
		/*12*/ imports.NewTable("WKPreferences_SetMinimumFontSize", 0),
		/*13*/ imports.NewTable("WKPreferences_SetPlugInsEnabled", 0),
		/*14*/ imports.NewTable("WKPreferences_SetTabFocusesLinks", 0),
		/*15*/ imports.NewTable("WKPreferences_SetValueForKey", 0),
	}
)

func wKPreferencesImportAPI() *imports.Imports {
	if wKPreferencesImport == nil {
		wKPreferencesImport = NewDefaultImports()
		wKPreferencesImport.SetImportTable(wKPreferencesImportTables)
		wKPreferencesImportTables = nil
	}
	return wKPreferencesImport
}
