//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package windows

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2FileSystemHandle Parent: IObject
type ICoreWebView2FileSystemHandle interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FileSystemHandle // property BaseIntf Getter
	// Kind
	//  The kind of the FileSystemHandle. It can either be a file or a directory.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2filesystemhandle#get_kind">See the ICoreWebView2FileSystemHandle article.</see>
	Kind() wvTypes.TWVFileSystemHandleKind // property Kind Getter
	// Path
	//  The path to the FileSystemHandle.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2filesystemhandle#get_path">See the ICoreWebView2FileSystemHandle article.</see>
	Path() string // property Path Getter
	// Permission
	//  The permissions granted to the FileSystemHandle.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2filesystemhandle#get_permission">See the ICoreWebView2FileSystemHandle article.</see>
	Permission() wvTypes.TWVFileSystemHandlePermission // property Permission Getter
}

type TCoreWebView2FileSystemHandle struct {
	TObject
}

func (m *TCoreWebView2FileSystemHandle) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FileSystemHandleAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2FileSystemHandle) BaseIntf() (result ICoreWebView2FileSystemHandle) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FileSystemHandleAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FileSystemHandle(resultPtr)
	return
}

func (m *TCoreWebView2FileSystemHandle) Kind() wvTypes.TWVFileSystemHandleKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2FileSystemHandleAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVFileSystemHandleKind(r)
}

func (m *TCoreWebView2FileSystemHandle) Path() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2FileSystemHandleAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2FileSystemHandle) Permission() wvTypes.TWVFileSystemHandlePermission {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2FileSystemHandleAPI().SysCallN(5, m.Instance())
	return wvTypes.TWVFileSystemHandlePermission(r)
}

// NewCoreWebView2FileSystemHandle class constructor
func NewCoreWebView2FileSystemHandle(baseIntf ICoreWebView2FileSystemHandle) ICoreWebView2FileSystemHandle {
	r := coreWebView2FileSystemHandleAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2FileSystemHandle(r)
}

var (
	coreWebView2FileSystemHandleOnce   base.Once
	coreWebView2FileSystemHandleImport *imports.Imports = nil
)

func coreWebView2FileSystemHandleAPI() *imports.Imports {
	coreWebView2FileSystemHandleOnce.Do(func() {
		coreWebView2FileSystemHandleImport = api.NewDefaultImports()
		coreWebView2FileSystemHandleImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2FileSystemHandle_Create", 0), // constructor NewCoreWebView2FileSystemHandle
			/* 1 */ imports.NewTable("TCoreWebView2FileSystemHandle_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2FileSystemHandle_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2FileSystemHandle_Kind", 0), // property Kind
			/* 4 */ imports.NewTable("TCoreWebView2FileSystemHandle_Path", 0), // property Path
			/* 5 */ imports.NewTable("TCoreWebView2FileSystemHandle_Permission", 0), // property Permission
		}
	})
	return coreWebView2FileSystemHandleImport
}
