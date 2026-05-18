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

// ICoreWebView2DevToolsProtocolEventReceivedEventArgs Parent: IObject
type ICoreWebView2DevToolsProtocolEventReceivedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DevToolsProtocolEventReceivedEventArgs // property BaseIntf Getter
	// ParameterObjectAsJson
	//  The parameter object of the corresponding `DevToolsProtocol` event
	//  represented as a JSON string.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs#get_parameterobjectasjson">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs article.</see>
	ParameterObjectAsJson() string // property ParameterObjectAsJson Getter
	// SessionId
	//  The sessionId of the target where the event originates from.
	//  Empty string is returned as sessionId if the event comes from the default
	//  session for the top page.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceivedeventargs2#get_sessionid">See the ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 article.</see>
	SessionId() string // property SessionId Getter
}

type TCoreWebView2DevToolsProtocolEventReceivedEventArgs struct {
	TObject
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DevToolsProtocolEventReceivedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) BaseIntf() (result ICoreWebView2DevToolsProtocolEventReceivedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2DevToolsProtocolEventReceivedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) ParameterObjectAsJson() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2DevToolsProtocolEventReceivedEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2DevToolsProtocolEventReceivedEventArgs) SessionId() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2DevToolsProtocolEventReceivedEventArgsAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

// NewCoreWebView2DevToolsProtocolEventReceivedEventArgs class constructor
func NewCoreWebView2DevToolsProtocolEventReceivedEventArgs(args ICoreWebView2DevToolsProtocolEventReceivedEventArgs) ICoreWebView2DevToolsProtocolEventReceivedEventArgs {
	r := coreWebView2DevToolsProtocolEventReceivedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(r)
}

var (
	coreWebView2DevToolsProtocolEventReceivedEventArgsOnce   base.Once
	coreWebView2DevToolsProtocolEventReceivedEventArgsImport *imports.Imports = nil
)

func coreWebView2DevToolsProtocolEventReceivedEventArgsAPI() *imports.Imports {
	coreWebView2DevToolsProtocolEventReceivedEventArgsOnce.Do(func() {
		coreWebView2DevToolsProtocolEventReceivedEventArgsImport = api.NewDefaultImports()
		coreWebView2DevToolsProtocolEventReceivedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2DevToolsProtocolEventReceivedEventArgs_Create", 0), // constructor NewCoreWebView2DevToolsProtocolEventReceivedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2DevToolsProtocolEventReceivedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2DevToolsProtocolEventReceivedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2DevToolsProtocolEventReceivedEventArgs_ParameterObjectAsJson", 0), // property ParameterObjectAsJson
			/* 4 */ imports.NewTable("TCoreWebView2DevToolsProtocolEventReceivedEventArgs_SessionId", 0), // property SessionId
		}
	})
	return coreWebView2DevToolsProtocolEventReceivedEventArgsImport
}
