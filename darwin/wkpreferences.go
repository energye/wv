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
//
//	An object that encapsulates the standard behaviors to apply to websites.
//	https://developer.apple.com/documentation/webkit/wkpreferences?language=objc
type IWKPreferences interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKPreferences // function
	// MinimumFontSize
	//  Returns The minimum font size, in points.
	MinimumFontSize() (resultFloat64 float64) // function
	// JavaScriptEnabled
	//  Returns A Boolean value that indicates whether JavaScript is enabled.
	JavaScriptEnabled() bool // function
	// JavaScriptCanOpenWindowsAutomatically
	//  Returns A Boolean value that indicates whether JavaScript can open windows without user interaction.
	JavaScriptCanOpenWindowsAutomatically() bool // function
	// JavaEnabled
	//  Returns A Boolean value that indicates whether Java is enabled.
	JavaEnabled() bool // function
	// PlugInsEnabled
	//  Returns A Boolean value that indicates whether plug-ins are enabled.
	PlugInsEnabled() bool // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// SetMinimumFontSize
	//  Sets The minimum font size, in points.
	SetMinimumFontSize(newValue float64) // procedure
	// SetJavaScriptEnabled
	//  Sets A Boolean value that indicates whether JavaScript is enabled.
	SetJavaScriptEnabled(newValue bool) // procedure
	// SetJavaScriptCanOpenWindowsAutomatically
	//  Sets A Boolean value that indicates whether JavaScript can open windows without user interaction.
	SetJavaScriptCanOpenWindowsAutomatically(newValue bool) // procedure
	// SetJavaEnabled
	//  Sets A Boolean value that indicates whether Java is enabled.
	SetJavaEnabled(newValue bool) // procedure
	// SetPlugInsEnabled
	//  Sets A Boolean value that indicates whether plug-ins are enabled.
	SetPlugInsEnabled(newValue bool) // procedure
	// SetTabFocusesLinks
	//  Sets A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
	SetTabFocusesLinks(newValue bool) // procedure
	// SetFraudulentWebsiteWarningEnabled
	//  Sets A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
	SetFraudulentWebsiteWarningEnabled(newValue bool) // procedure
	// SetTextInteractionEnabled
	//  Sets A Boolean value that indicates whether to allow people to select or otherwise interact with text. 11.3
	SetTextInteractionEnabled(newValue bool) // procedure
	// SetElementFullscreenEnabled
	//  Sets A Boolean value that indicates whether a web view can display content full screen. 12.3
	SetElementFullscreenEnabled(newValue bool) // procedure
	// SetValueForKey
	//  Writes a custom value to a field
	SetValueForKey(value bool, key string) // procedure
}

// TWKPreferences Root Object
//
//	An object that encapsulates the standard behaviors to apply to websites.
//	https://developer.apple.com/documentation/webkit/wkpreferences?language=objc
type TWKPreferences struct {
	TObject
}

func NewWKPreferences(aData WKPreferences) IWKPreferences {
	r1 := wKPreferencesImportAPI().SysCallN(0, uintptr(aData))
	return AsWKPreferences(r1)
}

// WKPreferencesRef -> IWKPreferences
var WKPreferencesRef wKPreferences

// wKPreferences TWKPreferences Ref
type wKPreferences uintptr

// New
//
//	Creates and returns an WKPreferences object.
func (m *wKPreferences) New() IWKPreferences {
	r1 := wKPreferencesImportAPI().SysCallN(6)
	return AsWKPreferences(r1)
}

func (m *TWKPreferences) Data() WKPreferences {
	r1 := wKPreferencesImportAPI().SysCallN(1, m.Instance())
	return WKPreferences(r1)
}

func (m *TWKPreferences) MinimumFontSize() (resultFloat64 float64) {
	wKPreferencesImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TWKPreferences) JavaScriptEnabled() bool {
	r1 := wKPreferencesImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TWKPreferences) JavaScriptCanOpenWindowsAutomatically() bool {
	r1 := wKPreferencesImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TWKPreferences) JavaEnabled() bool {
	r1 := wKPreferencesImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TWKPreferences) PlugInsEnabled() bool {
	r1 := wKPreferencesImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TWKPreferences) Release() {
	wKPreferencesImportAPI().SysCallN(8, m.Instance())
}

func (m *TWKPreferences) SetMinimumFontSize(newValue float64) {
	wKPreferencesImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(&newValue)))
}

func (m *TWKPreferences) SetJavaScriptEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(13, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetJavaScriptCanOpenWindowsAutomatically(newValue bool) {
	wKPreferencesImportAPI().SysCallN(12, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetJavaEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(11, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetPlugInsEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(15, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetTabFocusesLinks(newValue bool) {
	wKPreferencesImportAPI().SysCallN(16, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetFraudulentWebsiteWarningEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(10, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetTextInteractionEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(17, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetElementFullscreenEnabled(newValue bool) {
	wKPreferencesImportAPI().SysCallN(9, m.Instance(), PascalBool(newValue))
}

func (m *TWKPreferences) SetValueForKey(value bool, key string) {
	wKPreferencesImportAPI().SysCallN(18, m.Instance(), PascalBool(value), PascalStr(key))
}

var (
	wKPreferencesImport       *imports.Imports = nil
	wKPreferencesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKPreferences_Create", 0),
		/*1*/ imports.NewTable("WKPreferences_Data", 0),
		/*2*/ imports.NewTable("WKPreferences_JavaEnabled", 0),
		/*3*/ imports.NewTable("WKPreferences_JavaScriptCanOpenWindowsAutomatically", 0),
		/*4*/ imports.NewTable("WKPreferences_JavaScriptEnabled", 0),
		/*5*/ imports.NewTable("WKPreferences_MinimumFontSize", 0),
		/*6*/ imports.NewTable("WKPreferences_New", 0),
		/*7*/ imports.NewTable("WKPreferences_PlugInsEnabled", 0),
		/*8*/ imports.NewTable("WKPreferences_Release", 0),
		/*9*/ imports.NewTable("WKPreferences_SetElementFullscreenEnabled", 0),
		/*10*/ imports.NewTable("WKPreferences_SetFraudulentWebsiteWarningEnabled", 0),
		/*11*/ imports.NewTable("WKPreferences_SetJavaEnabled", 0),
		/*12*/ imports.NewTable("WKPreferences_SetJavaScriptCanOpenWindowsAutomatically", 0),
		/*13*/ imports.NewTable("WKPreferences_SetJavaScriptEnabled", 0),
		/*14*/ imports.NewTable("WKPreferences_SetMinimumFontSize", 0),
		/*15*/ imports.NewTable("WKPreferences_SetPlugInsEnabled", 0),
		/*16*/ imports.NewTable("WKPreferences_SetTabFocusesLinks", 0),
		/*17*/ imports.NewTable("WKPreferences_SetTextInteractionEnabled", 0),
		/*18*/ imports.NewTable("WKPreferences_SetValueForKey", 0),
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
