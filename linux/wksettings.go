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

// IWkSettings Root Interface
type IWkSettings interface {
	IObject
	EnableJavascript() bool                                                        // property
	SetEnableJavascript(AValue bool)                                               // property
	AutoLoadImages() bool                                                          // property
	SetAutoLoadImages(AValue bool)                                                 // property
	EnableDeveloperExtras() bool                                                   // property
	SetEnableDeveloperExtras(AValue bool)                                          // property
	Data() WebKitSettings                                                          // function
	SetUserAgentWithApplicationDetails(applicationname, applicationversion string) // procedure
	SetHardwareAccelerationPolicy(policy WebKitHardwareAccelerationPolicy)         // procedure
}

// TWkSettings Root Object
type TWkSettings struct {
	TObject
}

func NewWkSettings() IWkSettings {
	r1 := wkSettingsImportAPI().SysCallN(1)
	return AsWkSettings(r1)
}

func NewWkSettings1(aSetting WebKitSettings) IWkSettings {
	r1 := wkSettingsImportAPI().SysCallN(2, uintptr(aSetting))
	return AsWkSettings(r1)
}

func (m *TWkSettings) EnableJavascript() bool {
	r1 := wkSettingsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWkSettings) SetEnableJavascript(AValue bool) {
	wkSettingsImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWkSettings) AutoLoadImages() bool {
	r1 := wkSettingsImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWkSettings) SetAutoLoadImages(AValue bool) {
	wkSettingsImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWkSettings) EnableDeveloperExtras() bool {
	r1 := wkSettingsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWkSettings) SetEnableDeveloperExtras(AValue bool) {
	wkSettingsImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWkSettings) Data() WebKitSettings {
	r1 := wkSettingsImportAPI().SysCallN(3, m.Instance())
	return WebKitSettings(r1)
}

func (m *TWkSettings) SetUserAgentWithApplicationDetails(applicationname, applicationversion string) {
	wkSettingsImportAPI().SysCallN(7, m.Instance(), PascalStr(applicationname), PascalStr(applicationversion))
}

func (m *TWkSettings) SetHardwareAccelerationPolicy(policy WebKitHardwareAccelerationPolicy) {
	wkSettingsImportAPI().SysCallN(6, m.Instance(), uintptr(policy))
}

var (
	wkSettingsImport       *imports.Imports = nil
	wkSettingsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkSettings_AutoLoadImages", 0),
		/*1*/ imports.NewTable("WkSettings_Create", 0),
		/*2*/ imports.NewTable("WkSettings_Create1", 0),
		/*3*/ imports.NewTable("WkSettings_Data", 0),
		/*4*/ imports.NewTable("WkSettings_EnableDeveloperExtras", 0),
		/*5*/ imports.NewTable("WkSettings_EnableJavascript", 0),
		/*6*/ imports.NewTable("WkSettings_SetHardwareAccelerationPolicy", 0),
		/*7*/ imports.NewTable("WkSettings_SetUserAgentWithApplicationDetails", 0),
	}
)

func wkSettingsImportAPI() *imports.Imports {
	if wkSettingsImport == nil {
		wkSettingsImport = NewDefaultImports()
		wkSettingsImport.SetImportTable(wkSettingsImportTables)
		wkSettingsImportTables = nil
	}
	return wkSettingsImport
}
