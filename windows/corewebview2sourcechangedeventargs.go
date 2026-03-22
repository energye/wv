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
)

// ICoreWebView2SourceChangedEventArgs Parent: IObject
type ICoreWebView2SourceChangedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2SourceChangedEventArgs // property BaseIntf Getter
	// IsNewDocument
	//  `TRUE` if the page being navigated to is a new document.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventargs#get_isnewdocument">See the ICoreWebView2SourceChangedEventArgs article.</see>
	IsNewDocument() bool // property IsNewDocument Getter
}

type TCoreWebView2SourceChangedEventArgs struct {
	TObject
}

func (m *TCoreWebView2SourceChangedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SourceChangedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SourceChangedEventArgs) BaseIntf() (result ICoreWebView2SourceChangedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2SourceChangedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2SourceChangedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2SourceChangedEventArgs) IsNewDocument() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SourceChangedEventArgsAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

// NewCoreWebView2SourceChangedEventArgs class constructor
func NewCoreWebView2SourceChangedEventArgs(args ICoreWebView2SourceChangedEventArgs) ICoreWebView2SourceChangedEventArgs {
	r := coreWebView2SourceChangedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2SourceChangedEventArgs(r)
}

var (
	coreWebView2SourceChangedEventArgsOnce   base.Once
	coreWebView2SourceChangedEventArgsImport *imports.Imports = nil
)

func coreWebView2SourceChangedEventArgsAPI() *imports.Imports {
	coreWebView2SourceChangedEventArgsOnce.Do(func() {
		coreWebView2SourceChangedEventArgsImport = api.NewDefaultImports()
		coreWebView2SourceChangedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2SourceChangedEventArgs_Create", 0), // constructor NewCoreWebView2SourceChangedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2SourceChangedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2SourceChangedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2SourceChangedEventArgs_IsNewDocument", 0), // property IsNewDocument
		}
	})
	return coreWebView2SourceChangedEventArgsImport
}
