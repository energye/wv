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
)

// IWkLoader Parent: lcl.IComponent
type IWkLoader interface {
	lcl.IComponent
	StartWebKit2() bool                          // function
	LoaderWebKit2DllPath() string                // property LoaderWebKit2DllPath Getter
	SetLoaderWebKit2DllPath(value string)        // property LoaderWebKit2DllPath Setter
	LoaderJavascriptCoreDllPath() string         // property LoaderJavascriptCoreDllPath Getter
	SetLoaderJavascriptCoreDllPath(value string) // property LoaderJavascriptCoreDllPath Setter
	LoaderSoupDllPath() string                   // property LoaderSoupDllPath Getter
	SetLoaderSoupDllPath(value string)           // property LoaderSoupDllPath Setter
}

type TWkLoader struct {
	lcl.TComponent
}

func (m *TWkLoader) StartWebKit2() bool {
	if !m.IsValid() {
		return false
	}
	r := wkLoaderAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TWkLoader) LoaderWebKit2DllPath() string {
	if !m.IsValid() {
		return ""
	}
	r := wkLoaderAPI().SysCallN(2, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWkLoader) SetLoaderWebKit2DllPath(value string) {
	if !m.IsValid() {
		return
	}
	wkLoaderAPI().SysCallN(2, 1, m.Instance(), api.PasStr(value))
}

func (m *TWkLoader) LoaderJavascriptCoreDllPath() string {
	if !m.IsValid() {
		return ""
	}
	r := wkLoaderAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWkLoader) SetLoaderJavascriptCoreDllPath(value string) {
	if !m.IsValid() {
		return
	}
	wkLoaderAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TWkLoader) LoaderSoupDllPath() string {
	if !m.IsValid() {
		return ""
	}
	r := wkLoaderAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWkLoader) SetLoaderSoupDllPath(value string) {
	if !m.IsValid() {
		return
	}
	wkLoaderAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

// NewLoader class constructor
func NewLoader(owner lcl.IComponent) IWkLoader {
	r := wkLoaderAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsWkLoader(r)
}

var (
	wkLoaderOnce   base.Once
	wkLoaderImport *imports.Imports = nil
)

func wkLoaderAPI() *imports.Imports {
	wkLoaderOnce.Do(func() {
		wkLoaderImport = api.NewDefaultImports()
		wkLoaderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkLoader_Create", 0), // constructor NewLoader
			/* 1 */ imports.NewTable("TWkLoader_StartWebKit2", 0), // function StartWebKit2
			/* 2 */ imports.NewTable("TWkLoader_LoaderWebKit2DllPath", 0), // property LoaderWebKit2DllPath
			/* 3 */ imports.NewTable("TWkLoader_LoaderJavascriptCoreDllPath", 0), // property LoaderJavascriptCoreDllPath
			/* 4 */ imports.NewTable("TWkLoader_LoaderSoupDllPath", 0), // property LoaderSoupDllPath
		}
	})
	return wkLoaderImport
}
