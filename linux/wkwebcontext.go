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

// IWkWebContext Root Interface
type IWkWebContext interface {
	IObject
	Data() WebKitWebContext                                                    // function
	GetCookieManager() WebKitCookieManager                                     // function
	RegisterURIScheme(aScheme string, aDelegate IWkSchemeRequestDelegateEvent) // procedure
	SetCacheModel(cachemodel WebKitCacheModel)                                 // procedure
	DownloadURI(uri string)                                                    // procedure
	SetWetPreferredLanguages(languages IStrings)                               // procedure
}

// TWkWebContext Root Object
type TWkWebContext struct {
	TObject
}

func NewWkWebContext(aWebContext WebKitWebContext) IWkWebContext {
	r1 := wkWebContextImportAPI().SysCallN(0, uintptr(aWebContext))
	return AsWkWebContext(r1)
}

// WkWebContextRef -> IWkWebContext
var WkWebContextRef wkWebContext

// wkWebContext TWkWebContext Ref
type wkWebContext uintptr

func (m *wkWebContext) Default() IWkWebContext {
	r1 := wkWebContextImportAPI().SysCallN(2)
	return AsWkWebContext(r1)
}

func (m *wkWebContext) New() IWkWebContext {
	r1 := wkWebContextImportAPI().SysCallN(5)
	return AsWkWebContext(r1)
}

func (m *TWkWebContext) Data() WebKitWebContext {
	r1 := wkWebContextImportAPI().SysCallN(1, m.Instance())
	return WebKitWebContext(r1)
}

func (m *TWkWebContext) GetCookieManager() WebKitCookieManager {
	r1 := wkWebContextImportAPI().SysCallN(4, m.Instance())
	return WebKitCookieManager(r1)
}

func (m *TWkWebContext) RegisterURIScheme(aScheme string, aDelegate IWkSchemeRequestDelegateEvent) {
	wkWebContextImportAPI().SysCallN(6, m.Instance(), PascalStr(aScheme), GetObjectUintptr(aDelegate))
}

func (m *TWkWebContext) SetCacheModel(cachemodel WebKitCacheModel) {
	wkWebContextImportAPI().SysCallN(7, m.Instance(), uintptr(cachemodel))
}

func (m *TWkWebContext) DownloadURI(uri string) {
	wkWebContextImportAPI().SysCallN(3, m.Instance(), PascalStr(uri))
}

func (m *TWkWebContext) SetWetPreferredLanguages(languages IStrings) {
	wkWebContextImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(languages))
}

var (
	wkWebContextImport       *imports.Imports = nil
	wkWebContextImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkWebContext_Create", 0),
		/*1*/ imports.NewTable("WkWebContext_Data", 0),
		/*2*/ imports.NewTable("WkWebContext_Default", 0),
		/*3*/ imports.NewTable("WkWebContext_DownloadURI", 0),
		/*4*/ imports.NewTable("WkWebContext_GetCookieManager", 0),
		/*5*/ imports.NewTable("WkWebContext_New", 0),
		/*6*/ imports.NewTable("WkWebContext_RegisterURIScheme", 0),
		/*7*/ imports.NewTable("WkWebContext_SetCacheModel", 0),
		/*8*/ imports.NewTable("WkWebContext_SetWetPreferredLanguages", 0),
	}
)

func wkWebContextImportAPI() *imports.Imports {
	if wkWebContextImport == nil {
		wkWebContextImport = NewDefaultImports()
		wkWebContextImport.SetImportTable(wkWebContextImportTables)
		wkWebContextImportTables = nil
	}
	return wkWebContextImport
}
