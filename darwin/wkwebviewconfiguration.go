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

// IWKWebViewConfiguration Root Interface
type IWKWebViewConfiguration interface {
	IObject
	Data() WKWebViewConfiguration                                                          // function
	Preferences() WKPreferences                                                            // function
	UserContentController() WKUserContentController                                        // function
	SuppressesIncrementalRendering() bool                                                  // function
	ApplicationNameForUserAgent() string                                                   // function
	SetPreferences(preferences WKPreferences)                                              // procedure
	SetUserContentController(userContentController WKUserContentController)                // procedure
	SetSuppressesIncrementalRendering(newValue bool)                                       // procedure
	SetApplicationNameForUserAgent(newValue string)                                        // procedure
	SetURLSchemeHandlerForURLScheme(urlSchemeHandler WKURLSchemeHandler, urlScheme string) // procedure
}

// TWKWebViewConfiguration Root Object
type TWKWebViewConfiguration struct {
	TObject
}

func NewWKWebViewConfiguration(aData WKWebViewConfiguration) IWKWebViewConfiguration {
	r1 := wKWebViewConfigurationImportAPI().SysCallN(1, uintptr(aData))
	return AsWKWebViewConfiguration(r1)
}

// WKWebViewConfigurationRef -> IWKWebViewConfiguration
var WKWebViewConfigurationRef wKWebViewConfiguration

// wKWebViewConfiguration TWKWebViewConfiguration Ref
type wKWebViewConfiguration uintptr

func (m *wKWebViewConfiguration) New() IWKWebViewConfiguration {
	r1 := wKWebViewConfigurationImportAPI().SysCallN(3)
	return AsWKWebViewConfiguration(r1)
}

func (m *TWKWebViewConfiguration) Data() WKWebViewConfiguration {
	r1 := wKWebViewConfigurationImportAPI().SysCallN(2, m.Instance())
	return WKWebViewConfiguration(r1)
}

func (m *TWKWebViewConfiguration) Preferences() WKPreferences {
	r1 := wKWebViewConfigurationImportAPI().SysCallN(4, m.Instance())
	return WKPreferences(r1)
}

func (m *TWKWebViewConfiguration) UserContentController() WKUserContentController {
	r1 := wKWebViewConfigurationImportAPI().SysCallN(11, m.Instance())
	return WKUserContentController(r1)
}

func (m *TWKWebViewConfiguration) SuppressesIncrementalRendering() bool {
	r1 := wKWebViewConfigurationImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TWKWebViewConfiguration) ApplicationNameForUserAgent() string {
	r1 := wKWebViewConfigurationImportAPI().SysCallN(0, m.Instance())
	return GoStr(r1)
}

func (m *TWKWebViewConfiguration) SetPreferences(preferences WKPreferences) {
	wKWebViewConfigurationImportAPI().SysCallN(6, m.Instance(), uintptr(preferences))
}

func (m *TWKWebViewConfiguration) SetUserContentController(userContentController WKUserContentController) {
	wKWebViewConfigurationImportAPI().SysCallN(9, m.Instance(), uintptr(userContentController))
}

func (m *TWKWebViewConfiguration) SetSuppressesIncrementalRendering(newValue bool) {
	wKWebViewConfigurationImportAPI().SysCallN(7, m.Instance(), PascalBool(newValue))
}

func (m *TWKWebViewConfiguration) SetApplicationNameForUserAgent(newValue string) {
	wKWebViewConfigurationImportAPI().SysCallN(5, m.Instance(), PascalStr(newValue))
}

func (m *TWKWebViewConfiguration) SetURLSchemeHandlerForURLScheme(urlSchemeHandler WKURLSchemeHandler, urlScheme string) {
	wKWebViewConfigurationImportAPI().SysCallN(8, m.Instance(), uintptr(urlSchemeHandler), PascalStr(urlScheme))
}

var (
	wKWebViewConfigurationImport       *imports.Imports = nil
	wKWebViewConfigurationImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKWebViewConfiguration_ApplicationNameForUserAgent", 0),
		/*1*/ imports.NewTable("WKWebViewConfiguration_Create", 0),
		/*2*/ imports.NewTable("WKWebViewConfiguration_Data", 0),
		/*3*/ imports.NewTable("WKWebViewConfiguration_New", 0),
		/*4*/ imports.NewTable("WKWebViewConfiguration_Preferences", 0),
		/*5*/ imports.NewTable("WKWebViewConfiguration_SetApplicationNameForUserAgent", 0),
		/*6*/ imports.NewTable("WKWebViewConfiguration_SetPreferences", 0),
		/*7*/ imports.NewTable("WKWebViewConfiguration_SetSuppressesIncrementalRendering", 0),
		/*8*/ imports.NewTable("WKWebViewConfiguration_SetURLSchemeHandlerForURLScheme", 0),
		/*9*/ imports.NewTable("WKWebViewConfiguration_SetUserContentController", 0),
		/*10*/ imports.NewTable("WKWebViewConfiguration_SuppressesIncrementalRendering", 0),
		/*11*/ imports.NewTable("WKWebViewConfiguration_UserContentController", 0),
	}
)

func wKWebViewConfigurationImportAPI() *imports.Imports {
	if wKWebViewConfigurationImport == nil {
		wKWebViewConfigurationImport = NewDefaultImports()
		wKWebViewConfigurationImport.SetImportTable(wKWebViewConfigurationImportTables)
		wKWebViewConfigurationImportTables = nil
	}
	return wKWebViewConfigurationImport
}
