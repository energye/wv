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

// ICoreWebView2LaunchingExternalUriSchemeEventArgs Parent: IObject
type ICoreWebView2LaunchingExternalUriSchemeEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2LaunchingExternalUriSchemeEventArgs // property BaseIntf Getter
	// Uri
	//  The URI with the external URI scheme to be launched.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#get_uri">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</see>
	Uri() string // property Uri Getter
	// InitiatingOrigin
	//  The origin initiating the external URI scheme launch.
	//  The origin will be an empty string if the request is initiated by calling
	//  `CoreWebView2.Navigate` on the external URI scheme. If a script initiates
	//  the navigation, the `InitiatingOrigin` will be the top-level document's
	//  `Source`, for example, if `window.location` is set to `"calculator://"`, the
	//  `InitiatingOrigin` will be set to `calculator://`. If the request is initiated
	//  from a child frame, the `InitiatingOrigin` will be the source of that child frame.
	//  If the `InitiatingOrigin` is
	//  [opaque](https://html.spec.whatwg.org/multipage/origin.html#concept-origin-opaque),
	//  the `InitiatingOrigin` reported in the event args will be its precursor origin.
	//  The precursor origin is the origin that created the opaque origin. For example, if
	//  a frame on example.com opens a subframe with a different opaque origin, the subframe's
	//  precursor origin is example.com.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#get_initiatingorigin">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</see>
	InitiatingOrigin() string // property InitiatingOrigin Getter
	// IsUserInitiated
	//  `TRUE` when the external URI scheme request was initiated through a user gesture.
	//  \>NOTE: Being initiated through a user gesture does not mean that user intended
	//  to access the associated resource.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#get_isuserinitiated">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</see>
	IsUserInitiated() bool // property IsUserInitiated Getter
	// Cancel
	//  The event handler may set this property to `TRUE` to cancel the external URI scheme
	//  launch. If set to `TRUE`, the external URI scheme will not be launched, and the default
	//  dialog is not displayed. This property can be used to replace the normal
	//  handling of launching an external URI scheme.
	//  The initial value of the `Cancel` property is `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#get_cancel">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</see>
	Cancel() bool         // property Cancel Getter
	SetCancel(value bool) // property Cancel Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#getdeferral">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2LaunchingExternalUriSchemeEventArgs struct {
	TObject
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) BaseIntf() (result ICoreWebView2LaunchingExternalUriSchemeEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2LaunchingExternalUriSchemeEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) Uri() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) InitiatingOrigin() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) IsUserInitiated() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) SetCancel(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2LaunchingExternalUriSchemeEventArgs class constructor
func NewCoreWebView2LaunchingExternalUriSchemeEventArgs(args ICoreWebView2LaunchingExternalUriSchemeEventArgs) ICoreWebView2LaunchingExternalUriSchemeEventArgs {
	r := coreWebView2LaunchingExternalUriSchemeEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2LaunchingExternalUriSchemeEventArgs(r)
}

var (
	coreWebView2LaunchingExternalUriSchemeEventArgsOnce   base.Once
	coreWebView2LaunchingExternalUriSchemeEventArgsImport *imports.Imports = nil
)

func coreWebView2LaunchingExternalUriSchemeEventArgsAPI() *imports.Imports {
	coreWebView2LaunchingExternalUriSchemeEventArgsOnce.Do(func() {
		coreWebView2LaunchingExternalUriSchemeEventArgsImport = api.NewDefaultImports()
		coreWebView2LaunchingExternalUriSchemeEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2LaunchingExternalUriSchemeEventArgs_Create", 0), // constructor NewCoreWebView2LaunchingExternalUriSchemeEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2LaunchingExternalUriSchemeEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2LaunchingExternalUriSchemeEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2LaunchingExternalUriSchemeEventArgs_Uri", 0), // property Uri
			/* 4 */ imports.NewTable("TCoreWebView2LaunchingExternalUriSchemeEventArgs_InitiatingOrigin", 0), // property InitiatingOrigin
			/* 5 */ imports.NewTable("TCoreWebView2LaunchingExternalUriSchemeEventArgs_IsUserInitiated", 0), // property IsUserInitiated
			/* 6 */ imports.NewTable("TCoreWebView2LaunchingExternalUriSchemeEventArgs_Cancel", 0), // property Cancel
			/* 7 */ imports.NewTable("TCoreWebView2LaunchingExternalUriSchemeEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2LaunchingExternalUriSchemeEventArgsImport
}
