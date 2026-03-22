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

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkWebViewConfiguration Parent: IObject
type IWkWebViewConfiguration interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKWebViewConfiguration // function
	// Preferences
	//  Returns The object that manages the preference-related settings for the web view.
	Preferences() wvTypes.WKPreferences // function
	// UserContentController
	//  Returns The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
	UserContentController() wvTypes.WKUserContentController // function
	// SuppressesIncrementalRendering
	//  Returns A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
	SuppressesIncrementalRendering() bool // function
	// ApplicationNameForUserAgent
	//  Returns The app name that appears in the user agent string.
	ApplicationNameForUserAgent() string // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// SetPreferences
	//  Sets The object that manages the preference-related settings for the web view.
	SetPreferences(preferences wvTypes.WKPreferences) // procedure
	// SetUserContentController
	//  Sets The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
	SetUserContentController(userContentController wvTypes.WKUserContentController) // procedure
	// SetSuppressesIncrementalRendering
	//  Sets A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
	SetSuppressesIncrementalRendering(newValue bool) // procedure
	// SetApplicationNameForUserAgent
	//  Sets The app name that appears in the user agent string.
	SetApplicationNameForUserAgent(newValue string) // procedure
	// SetURLSchemeHandlerForURLScheme
	//  Sets Registers an object to load resources associated with the specified URL scheme.
	SetURLSchemeHandlerForURLScheme(urlSchemeHandler wvTypes.WKURLSchemeHandler, urlScheme string) // procedure
}

type TWkWebViewConfiguration struct {
	TObject
}

func (m *TWkWebViewConfiguration) Data() wvTypes.WKWebViewConfiguration {
	if !m.IsValid() {
		return 0
	}
	r := wkWebViewConfigurationAPI().SysCallN(1, m.Instance())
	return wvTypes.WKWebViewConfiguration(r)
}

func (m *TWkWebViewConfiguration) Preferences() wvTypes.WKPreferences {
	if !m.IsValid() {
		return 0
	}
	r := wkWebViewConfigurationAPI().SysCallN(3, m.Instance())
	return wvTypes.WKPreferences(r)
}

func (m *TWkWebViewConfiguration) UserContentController() wvTypes.WKUserContentController {
	if !m.IsValid() {
		return 0
	}
	r := wkWebViewConfigurationAPI().SysCallN(4, m.Instance())
	return wvTypes.WKUserContentController(r)
}

func (m *TWkWebViewConfiguration) SuppressesIncrementalRendering() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebViewConfigurationAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebViewConfiguration) ApplicationNameForUserAgent() string {
	if !m.IsValid() {
		return ""
	}
	r := wkWebViewConfigurationAPI().SysCallN(6, m.Instance())
	return api.GoStr(r)
}

func (m *TWkWebViewConfiguration) Release() {
	if !m.IsValid() {
		return
	}
	wkWebViewConfigurationAPI().SysCallN(7, m.Instance())
}

func (m *TWkWebViewConfiguration) SetPreferences(preferences wvTypes.WKPreferences) {
	if !m.IsValid() {
		return
	}
	wkWebViewConfigurationAPI().SysCallN(8, m.Instance(), uintptr(preferences))
}

func (m *TWkWebViewConfiguration) SetUserContentController(userContentController wvTypes.WKUserContentController) {
	if !m.IsValid() {
		return
	}
	wkWebViewConfigurationAPI().SysCallN(9, m.Instance(), uintptr(userContentController))
}

func (m *TWkWebViewConfiguration) SetSuppressesIncrementalRendering(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkWebViewConfigurationAPI().SysCallN(10, m.Instance(), api.PasBool(newValue))
}

func (m *TWkWebViewConfiguration) SetApplicationNameForUserAgent(newValue string) {
	if !m.IsValid() {
		return
	}
	wkWebViewConfigurationAPI().SysCallN(11, m.Instance(), api.PasStr(newValue))
}

func (m *TWkWebViewConfiguration) SetURLSchemeHandlerForURLScheme(urlSchemeHandler wvTypes.WKURLSchemeHandler, urlScheme string) {
	if !m.IsValid() {
		return
	}
	wkWebViewConfigurationAPI().SysCallN(12, m.Instance(), uintptr(urlSchemeHandler), api.PasStr(urlScheme))
}

// WebViewConfiguration  is static instance
var WebViewConfiguration _WebViewConfigurationClass

// _WebViewConfigurationClass is class type defined by TWkWebViewConfiguration
type _WebViewConfigurationClass uintptr

// New
//
//	Creates and returns an NSMutableURLRequest object.
func (_WebViewConfigurationClass) New() IWkWebViewConfiguration {
	r := wkWebViewConfigurationAPI().SysCallN(2)
	return AsWkWebViewConfiguration(r)
}

// NewWebViewConfiguration class constructor
func NewWebViewConfiguration(data wvTypes.WKWebViewConfiguration) IWkWebViewConfiguration {
	r := wkWebViewConfigurationAPI().SysCallN(0, uintptr(data))
	return AsWkWebViewConfiguration(r)
}

var (
	wkWebViewConfigurationOnce   base.Once
	wkWebViewConfigurationImport *imports.Imports = nil
)

func wkWebViewConfigurationAPI() *imports.Imports {
	wkWebViewConfigurationOnce.Do(func() {
		wkWebViewConfigurationImport = api.NewDefaultImports()
		wkWebViewConfigurationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkWebViewConfiguration_Create", 0), // constructor NewWebViewConfiguration
			/* 1 */ imports.NewTable("TWkWebViewConfiguration_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkWebViewConfiguration_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkWebViewConfiguration_Preferences", 0), // function Preferences
			/* 4 */ imports.NewTable("TWkWebViewConfiguration_UserContentController", 0), // function UserContentController
			/* 5 */ imports.NewTable("TWkWebViewConfiguration_SuppressesIncrementalRendering", 0), // function SuppressesIncrementalRendering
			/* 6 */ imports.NewTable("TWkWebViewConfiguration_ApplicationNameForUserAgent", 0), // function ApplicationNameForUserAgent
			/* 7 */ imports.NewTable("TWkWebViewConfiguration_Release", 0), // procedure Release
			/* 8 */ imports.NewTable("TWkWebViewConfiguration_SetPreferences", 0), // procedure SetPreferences
			/* 9 */ imports.NewTable("TWkWebViewConfiguration_SetUserContentController", 0), // procedure SetUserContentController
			/* 10 */ imports.NewTable("TWkWebViewConfiguration_SetSuppressesIncrementalRendering", 0), // procedure SetSuppressesIncrementalRendering
			/* 11 */ imports.NewTable("TWkWebViewConfiguration_SetApplicationNameForUserAgent", 0), // procedure SetApplicationNameForUserAgent
			/* 12 */ imports.NewTable("TWkWebViewConfiguration_SetURLSchemeHandlerForURLScheme", 0), // procedure SetURLSchemeHandlerForURLScheme
		}
	})
	return wkWebViewConfigurationImport
}
