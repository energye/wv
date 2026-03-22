//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkWebContext Parent: IObject
type IWkWebContext interface {
	IObject
	Data() wvTypes.WebKitWebContext                       // function
	GetCookieManager() wvTypes.WebKitCookieManager        // function
	RegisterURIScheme(scheme string, delegate IWkWebview) // procedure
	SetCacheModel(cacheModel wvTypes.WebKitCacheModel)    // procedure
	DownloadURI(uri string)                               // procedure
	SetWetPreferredLanguages(languages lcl.IStrings)      // procedure
}

type TWkWebContext struct {
	TObject
}

func (m *TWkWebContext) Data() wvTypes.WebKitWebContext {
	if !m.IsValid() {
		return 0
	}
	r := wkWebContextAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitWebContext(r)
}

func (m *TWkWebContext) GetCookieManager() wvTypes.WebKitCookieManager {
	if !m.IsValid() {
		return 0
	}
	r := wkWebContextAPI().SysCallN(2, m.Instance())
	return wvTypes.WebKitCookieManager(r)
}

func (m *TWkWebContext) RegisterURIScheme(scheme string, delegate IWkWebview) {
	if !m.IsValid() {
		return
	}
	wkWebContextAPI().SysCallN(5, m.Instance(), api.PasStr(scheme), base.GetObjectUintptr(delegate))
}

func (m *TWkWebContext) SetCacheModel(cacheModel wvTypes.WebKitCacheModel) {
	if !m.IsValid() {
		return
	}
	wkWebContextAPI().SysCallN(6, m.Instance(), uintptr(cacheModel))
}

func (m *TWkWebContext) DownloadURI(uri string) {
	if !m.IsValid() {
		return
	}
	wkWebContextAPI().SysCallN(7, m.Instance(), api.PasStr(uri))
}

func (m *TWkWebContext) SetWetPreferredLanguages(languages lcl.IStrings) {
	if !m.IsValid() {
		return
	}
	wkWebContextAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(languages))
}

// WebContext  is static instance
var WebContext _WebContextClass

// _WebContextClass is class type defined by TWkWebContext
type _WebContextClass uintptr

func (_WebContextClass) Default() IWkWebContext {
	r := wkWebContextAPI().SysCallN(3)
	return AsWkWebContext(r)
}

func (_WebContextClass) New() IWkWebContext {
	r := wkWebContextAPI().SysCallN(4)
	return AsWkWebContext(r)
}

// NewWebContext class constructor
func NewWebContext(webContext wvTypes.WebKitWebContext) IWkWebContext {
	r := wkWebContextAPI().SysCallN(0, uintptr(webContext))
	return AsWkWebContext(r)
}

var (
	wkWebContextOnce   base.Once
	wkWebContextImport *imports.Imports = nil
)

func wkWebContextAPI() *imports.Imports {
	wkWebContextOnce.Do(func() {
		wkWebContextImport = api.NewDefaultImports()
		wkWebContextImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkWebContext_Create", 0), // constructor NewWebContext
			/* 1 */ imports.NewTable("TWkWebContext_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkWebContext_GetCookieManager", 0), // function GetCookieManager
			/* 3 */ imports.NewTable("TWkWebContext_Default", 0), // static function Default
			/* 4 */ imports.NewTable("TWkWebContext_New", 0), // static function New
			/* 5 */ imports.NewTable("TWkWebContext_RegisterURIScheme", 0), // procedure RegisterURIScheme
			/* 6 */ imports.NewTable("TWkWebContext_SetCacheModel", 0), // procedure SetCacheModel
			/* 7 */ imports.NewTable("TWkWebContext_DownloadURI", 0), // procedure DownloadURI
			/* 8 */ imports.NewTable("TWkWebContext_SetWetPreferredLanguages", 0), // procedure SetWetPreferredLanguages
		}
	})
	return wkWebContextImport
}
