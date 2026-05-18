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

// IWkLoader Parent: IComponent
type IWkLoader interface {
	IComponent
	StartWebKit2() bool // function
	// Webkit2Version
	//  Webkit2Version 4.0 or 4.1
	Webkit2Version() wvTypes.TWebkit2Version         // property Webkit2Version Getter
	SetWebkit2Version(value wvTypes.TWebkit2Version) // property Webkit2Version Setter
	LoaderWebKit2DllPath() string                    // property LoaderWebKit2DllPath Getter
	SetLoaderWebKit2DllPath(value string)            // property LoaderWebKit2DllPath Setter
	LoaderJavascriptCoreDllPath() string             // property LoaderJavascriptCoreDllPath Getter
	SetLoaderJavascriptCoreDllPath(value string)     // property LoaderJavascriptCoreDllPath Setter
	LoaderSoupDllPath() string                       // property LoaderSoupDllPath Getter
	SetLoaderSoupDllPath(value string)               // property LoaderSoupDllPath Setter
}

type TWkLoader struct {
	TComponent
}

func (m *TWkLoader) StartWebKit2() bool {
	if !m.IsValid() {
		return false
	}
	r := wkLoaderAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TWkLoader) Webkit2Version() wvTypes.TWebkit2Version {
	if !m.IsValid() {
		return 0
	}
	r := wkLoaderAPI().SysCallN(2, 0, m.Instance())
	return wvTypes.TWebkit2Version(r)
}

func (m *TWkLoader) SetWebkit2Version(value wvTypes.TWebkit2Version) {
	if !m.IsValid() {
		return
	}
	wkLoaderAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TWkLoader) LoaderWebKit2DllPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkLoaderAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkLoader) SetLoaderWebKit2DllPath(value string) {
	if !m.IsValid() {
		return
	}
	wkLoaderAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TWkLoader) LoaderJavascriptCoreDllPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkLoaderAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkLoader) SetLoaderJavascriptCoreDllPath(value string) {
	if !m.IsValid() {
		return
	}
	wkLoaderAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TWkLoader) LoaderSoupDllPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkLoaderAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkLoader) SetLoaderSoupDllPath(value string) {
	if !m.IsValid() {
		return
	}
	wkLoaderAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
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
			/* 2 */ imports.NewTable("TWkLoader_Webkit2Version", 0), // property Webkit2Version
			/* 3 */ imports.NewTable("TWkLoader_LoaderWebKit2DllPath", 0), // property LoaderWebKit2DllPath
			/* 4 */ imports.NewTable("TWkLoader_LoaderJavascriptCoreDllPath", 0), // property LoaderJavascriptCoreDllPath
			/* 5 */ imports.NewTable("TWkLoader_LoaderSoupDllPath", 0), // property LoaderSoupDllPath
		}
	})
	return wkLoaderImport
}
