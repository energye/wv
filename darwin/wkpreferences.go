//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkPreferences Parent: lcl.IObject
type IWkPreferences interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKPreferences // function
	// MinimumFontSize
	//  Returns The minimum font size, in points.
	MinimumFontSize() float64 // function
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

type TWkPreferences struct {
	lcl.TObject
}

func (m *TWkPreferences) Data() wvTypes.WKPreferences {
	if !m.IsValid() {
		return 0
	}
	r := wkPreferencesAPI().SysCallN(1, m.Instance())
	return wvTypes.WKPreferences(r)
}

func (m *TWkPreferences) MinimumFontSize() (result float64) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkPreferences) JavaScriptEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wkPreferencesAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TWkPreferences) JavaScriptCanOpenWindowsAutomatically() bool {
	if !m.IsValid() {
		return false
	}
	r := wkPreferencesAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TWkPreferences) JavaEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wkPreferencesAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TWkPreferences) PlugInsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wkPreferencesAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TWkPreferences) Release() {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(8, m.Instance())
}

func (m *TWkPreferences) SetMinimumFontSize(newValue float64) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&newValue)))
}

func (m *TWkPreferences) SetJavaScriptEnabled(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(10, m.Instance(), api.PasBool(newValue))
}

func (m *TWkPreferences) SetJavaScriptCanOpenWindowsAutomatically(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(11, m.Instance(), api.PasBool(newValue))
}

func (m *TWkPreferences) SetJavaEnabled(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(12, m.Instance(), api.PasBool(newValue))
}

func (m *TWkPreferences) SetPlugInsEnabled(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(13, m.Instance(), api.PasBool(newValue))
}

func (m *TWkPreferences) SetTabFocusesLinks(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(14, m.Instance(), api.PasBool(newValue))
}

func (m *TWkPreferences) SetFraudulentWebsiteWarningEnabled(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(15, m.Instance(), api.PasBool(newValue))
}

func (m *TWkPreferences) SetTextInteractionEnabled(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(16, m.Instance(), api.PasBool(newValue))
}

func (m *TWkPreferences) SetElementFullscreenEnabled(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(17, m.Instance(), api.PasBool(newValue))
}

func (m *TWkPreferences) SetValueForKey(value bool, key string) {
	if !m.IsValid() {
		return
	}
	wkPreferencesAPI().SysCallN(18, m.Instance(), api.PasBool(value), api.PasStr(key))
}

// Preferences  is static instance
var Preferences _PreferencesClass

// _PreferencesClass is class type defined by TWkPreferences
type _PreferencesClass uintptr

// New
//
//	Creates and returns an WKPreferences object.
func (_PreferencesClass) New() IWkPreferences {
	r := wkPreferencesAPI().SysCallN(2)
	return AsWkPreferences(r)
}

// NewPreferences class constructor
func NewPreferences(data wvTypes.WKPreferences) IWkPreferences {
	r := wkPreferencesAPI().SysCallN(0, uintptr(data))
	return AsWkPreferences(r)
}

var (
	wkPreferencesOnce   base.Once
	wkPreferencesImport *imports.Imports = nil
)

func wkPreferencesAPI() *imports.Imports {
	wkPreferencesOnce.Do(func() {
		wkPreferencesImport = api.NewDefaultImports()
		wkPreferencesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkPreferences_Create", 0), // constructor NewPreferences
			/* 1 */ imports.NewTable("TWkPreferences_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkPreferences_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkPreferences_MinimumFontSize", 0), // function MinimumFontSize
			/* 4 */ imports.NewTable("TWkPreferences_JavaScriptEnabled", 0), // function JavaScriptEnabled
			/* 5 */ imports.NewTable("TWkPreferences_JavaScriptCanOpenWindowsAutomatically", 0), // function JavaScriptCanOpenWindowsAutomatically
			/* 6 */ imports.NewTable("TWkPreferences_JavaEnabled", 0), // function JavaEnabled
			/* 7 */ imports.NewTable("TWkPreferences_PlugInsEnabled", 0), // function PlugInsEnabled
			/* 8 */ imports.NewTable("TWkPreferences_Release", 0), // procedure Release
			/* 9 */ imports.NewTable("TWkPreferences_SetMinimumFontSize", 0), // procedure SetMinimumFontSize
			/* 10 */ imports.NewTable("TWkPreferences_SetJavaScriptEnabled", 0), // procedure SetJavaScriptEnabled
			/* 11 */ imports.NewTable("TWkPreferences_SetJavaScriptCanOpenWindowsAutomatically", 0), // procedure SetJavaScriptCanOpenWindowsAutomatically
			/* 12 */ imports.NewTable("TWkPreferences_SetJavaEnabled", 0), // procedure SetJavaEnabled
			/* 13 */ imports.NewTable("TWkPreferences_SetPlugInsEnabled", 0), // procedure SetPlugInsEnabled
			/* 14 */ imports.NewTable("TWkPreferences_SetTabFocusesLinks", 0), // procedure SetTabFocusesLinks
			/* 15 */ imports.NewTable("TWkPreferences_SetFraudulentWebsiteWarningEnabled", 0), // procedure SetFraudulentWebsiteWarningEnabled
			/* 16 */ imports.NewTable("TWkPreferences_SetTextInteractionEnabled", 0), // procedure SetTextInteractionEnabled
			/* 17 */ imports.NewTable("TWkPreferences_SetElementFullscreenEnabled", 0), // procedure SetElementFullscreenEnabled
			/* 18 */ imports.NewTable("TWkPreferences_SetValueForKey", 0), // procedure SetValueForKey
		}
	})
	return wkPreferencesImport
}
