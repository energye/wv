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

// ICoreWebView2DevToolsProtocolEventReceivedEventArgs Parent: IObject
//
//	Event args for the DevToolsProtocolEventReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs article.</a>
type ICoreWebView2DevToolsProtocolEventReceivedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DevToolsProtocolEventReceivedEventArgs // property
	// ParameterObjectAsJson
	//  The parameter object of the corresponding `DevToolsProtocol` event
	//  represented as a JSON string.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs#get_parameterobjectasjson">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs article.</a>
	ParameterObjectAsJson() string // property
	// SessionId
	//  The sessionId of the target where the event originates from.
	//  Empty string is returned as sessionId if the event comes from the default
	//  session for the top page.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs2#get_sessionid">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 article.</a>
	SessionId() string // property
}

// TCoreWebView2DevToolsProtocolEventReceivedEventArgs Parent: TObject
//
//	Event args for the DevToolsProtocolEventReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs article.</a>
type TCoreWebView2DevToolsProtocolEventReceivedEventArgs struct {
	TObject
}

func NewCoreWebView2DevToolsProtocolEventReceivedEventArgs(aArgs ICoreWebView2DevToolsProtocolEventReceivedEventArgs) ICoreWebView2DevToolsProtocolEventReceivedEventArgs {
	r1 := coreWebView2DevToolsProtocolEventReceivedEventArgsImportAPI().SysCallN(1, GetObjectUintptr(aArgs))
	return AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(r1)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) Initialized() bool {
	r1 := coreWebView2DevToolsProtocolEventReceivedEventArgsImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) BaseIntf() ICoreWebView2DevToolsProtocolEventReceivedEventArgs {
	var resultCoreWebView2DevToolsProtocolEventReceivedEventArgs uintptr
	coreWebView2DevToolsProtocolEventReceivedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2DevToolsProtocolEventReceivedEventArgs)))
	return AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(resultCoreWebView2DevToolsProtocolEventReceivedEventArgs)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) ParameterObjectAsJson() string {
	r1 := coreWebView2DevToolsProtocolEventReceivedEventArgsImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) SessionId() string {
	r1 := coreWebView2DevToolsProtocolEventReceivedEventArgsImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

var (
	coreWebView2DevToolsProtocolEventReceivedEventArgsImport       *imports.Imports = nil
	coreWebView2DevToolsProtocolEventReceivedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2DevToolsProtocolEventReceivedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2DevToolsProtocolEventReceivedEventArgs_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2DevToolsProtocolEventReceivedEventArgs_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2DevToolsProtocolEventReceivedEventArgs_ParameterObjectAsJson", 0),
		/*4*/ imports.NewTable("CoreWebView2DevToolsProtocolEventReceivedEventArgs_SessionId", 0),
	}
)

func coreWebView2DevToolsProtocolEventReceivedEventArgsImportAPI() *imports.Imports {
	if coreWebView2DevToolsProtocolEventReceivedEventArgsImport == nil {
		coreWebView2DevToolsProtocolEventReceivedEventArgsImport = NewDefaultImports()
		coreWebView2DevToolsProtocolEventReceivedEventArgsImport.SetImportTable(coreWebView2DevToolsProtocolEventReceivedEventArgsImportTables)
		coreWebView2DevToolsProtocolEventReceivedEventArgsImportTables = nil
	}
	return coreWebView2DevToolsProtocolEventReceivedEventArgsImport
}
