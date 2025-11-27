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
	"github.com/energye/lcl/lcl"
)

// ICoreWebView2WebMessageReceivedEventArgs Parent: lcl.IObject
type ICoreWebView2WebMessageReceivedEventArgs interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebMessageReceivedEventArgs // property BaseIntf Getter
	// Source
	//  The URI of the document that sent this web message.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs#get_source">See the ICoreWebView2WebMessageReceivedEventArgs article.</see>
	Source() string // property Source Getter
	// WebMessageAsJson
	//  The message posted from the WebView content to the host converted to a
	//  JSON string. Run this operation to communicate using JavaScript objects.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs#get_webmessageasjson">See the ICoreWebView2WebMessageReceivedEventArgs article.</see>
	WebMessageAsJson() string // property WebMessageAsJson Getter
	// WebMessageAsString
	//  If the message posted from the WebView content to the host is a string
	//  type, this method returns the value of that string. If the message
	//  posted is some other kind of JavaScript type this method fails with the
	//  following error.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs#trygetwebmessageasstring">See the ICoreWebView2WebMessageReceivedEventArgs article.</see>
	WebMessageAsString() string // property WebMessageAsString Getter
	// AdditionalObjects
	//  Additional received WebMessage objects. To pass `additionalObjects` via
	//  WebMessage to the app, use the
	//  `chrome.webview.postMessageWithAdditionalObjects` content API.
	//  Any DOM object type that can be natively representable that has been
	//  passed in to `additionalObjects` parameter will be accessible here.
	//  Currently a WebMessage object can be the `ICoreWebView2File` type.
	//  Entries in the collection can be `nullptr` if `null` or `undefined` was
	//  passed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs2#get_additionalobjects">See the ICoreWebView2WebMessageReceivedEventArgs2 article.</see>
	AdditionalObjects() ICoreWebView2ObjectCollectionView // property AdditionalObjects Getter
}

type TCoreWebView2WebMessageReceivedEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WebMessageReceivedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) BaseIntf() (result ICoreWebView2WebMessageReceivedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebMessageReceivedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebMessageReceivedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) Source() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2WebMessageReceivedEventArgsAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) WebMessageAsJson() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2WebMessageReceivedEventArgsAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) WebMessageAsString() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2WebMessageReceivedEventArgsAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2WebMessageReceivedEventArgs) AdditionalObjects() (result ICoreWebView2ObjectCollectionView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebMessageReceivedEventArgsAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ObjectCollectionView(resultPtr)
	return
}

// NewCoreWebView2WebMessageReceivedEventArgs class constructor
func NewCoreWebView2WebMessageReceivedEventArgs(args ICoreWebView2WebMessageReceivedEventArgs) ICoreWebView2WebMessageReceivedEventArgs {
	r := coreWebView2WebMessageReceivedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2WebMessageReceivedEventArgs(r)
}

var (
	coreWebView2WebMessageReceivedEventArgsOnce   base.Once
	coreWebView2WebMessageReceivedEventArgsImport *imports.Imports = nil
)

func coreWebView2WebMessageReceivedEventArgsAPI() *imports.Imports {
	coreWebView2WebMessageReceivedEventArgsOnce.Do(func() {
		coreWebView2WebMessageReceivedEventArgsImport = api.NewDefaultImports()
		coreWebView2WebMessageReceivedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WebMessageReceivedEventArgs_Create", 0), // constructor NewCoreWebView2WebMessageReceivedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2WebMessageReceivedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2WebMessageReceivedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2WebMessageReceivedEventArgs_Source", 0), // property Source
			/* 4 */ imports.NewTable("TCoreWebView2WebMessageReceivedEventArgs_WebMessageAsJson", 0), // property WebMessageAsJson
			/* 5 */ imports.NewTable("TCoreWebView2WebMessageReceivedEventArgs_WebMessageAsString", 0), // property WebMessageAsString
			/* 6 */ imports.NewTable("TCoreWebView2WebMessageReceivedEventArgs_AdditionalObjects", 0), // property AdditionalObjects
		}
	})
	return coreWebView2WebMessageReceivedEventArgsImport
}
