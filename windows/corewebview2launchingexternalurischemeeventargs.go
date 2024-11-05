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

// ICoreWebView2LaunchingExternalUriSchemeEventArgs Parent: IObject
//
//	Event args for LaunchingExternalUriScheme event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</a>
type ICoreWebView2LaunchingExternalUriSchemeEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2LaunchingExternalUriSchemeEventArgs // property
	// Uri
	//  The URI with the external URI scheme to be launched.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#get_uri">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</a>
	Uri() string // property
	// InitiatingOrigin
	//  The origin initiating the external URI scheme launch.
	//  The origin will be an empty string if the request is initiated by calling
	//  `CoreWebView2.Navigate` on the external URI scheme. If a script initiates
	//  the navigation, the `InitiatingOrigin` will be the top-level document's
	//  `Source`, for example, if `window.location` is set to `"calculator://", the
	//  `InitiatingOrigin` will be set to `calculator://`. If the request is initiated
	//  from a child frame, the `InitiatingOrigin` will be the source of that child frame.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#get_initiatingorigin">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</a>
	InitiatingOrigin() string // property
	// IsUserInitiated
	//  `TRUE` when the external URI scheme request was initiated through a user gesture.
	//  NOTE: Being initiated through a user gesture does not mean that user intended
	//  to access the associated resource.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#get_isuserinitiated">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</a>
	IsUserInitiated() bool // property
	// Cancel
	//  The event handler may set this property to `TRUE` to cancel the external URI scheme
	//  launch. If set to `TRUE`, the external URI scheme will not be launched, and the default
	//  dialog is not displayed. This property can be used to replace the normal
	//  handling of launching an external URI scheme.
	//  The initial value of the `Cancel` property is `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#get_cancel">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</a>
	Cancel() bool // property
	// SetCancel Set Cancel
	SetCancel(AValue bool) // property
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs#getdeferral">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
}

// TCoreWebView2LaunchingExternalUriSchemeEventArgs Parent: TObject
//
//	Event args for LaunchingExternalUriScheme event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs">See the ICoreWebView2LaunchingExternalUriSchemeEventArgs article.</a>
type TCoreWebView2LaunchingExternalUriSchemeEventArgs struct {
	TObject
}

func NewCoreWebView2LaunchingExternalUriSchemeEventArgs(aArgs ICoreWebView2LaunchingExternalUriSchemeEventArgs) ICoreWebView2LaunchingExternalUriSchemeEventArgs {
	r1 := coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(2, GetObjectUintptr(aArgs))
	return AsCoreWebView2LaunchingExternalUriSchemeEventArgs(r1)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) Initialized() bool {
	r1 := coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) BaseIntf() ICoreWebView2LaunchingExternalUriSchemeEventArgs {
	var resultCoreWebView2LaunchingExternalUriSchemeEventArgs uintptr
	coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2LaunchingExternalUriSchemeEventArgs)))
	return AsCoreWebView2LaunchingExternalUriSchemeEventArgs(resultCoreWebView2LaunchingExternalUriSchemeEventArgs)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) Uri() string {
	r1 := coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) InitiatingOrigin() string {
	r1 := coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) IsUserInitiated() bool {
	r1 := coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) Cancel() bool {
	r1 := coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) SetCancel(AValue bool) {
	coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2LaunchingExternalUriSchemeEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

var (
	coreWebView2LaunchingExternalUriSchemeEventArgsImport       *imports.Imports = nil
	coreWebView2LaunchingExternalUriSchemeEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2LaunchingExternalUriSchemeEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2LaunchingExternalUriSchemeEventArgs_Cancel", 0),
		/*2*/ imports.NewTable("CoreWebView2LaunchingExternalUriSchemeEventArgs_Create", 0),
		/*3*/ imports.NewTable("CoreWebView2LaunchingExternalUriSchemeEventArgs_Deferral", 0),
		/*4*/ imports.NewTable("CoreWebView2LaunchingExternalUriSchemeEventArgs_Initialized", 0),
		/*5*/ imports.NewTable("CoreWebView2LaunchingExternalUriSchemeEventArgs_InitiatingOrigin", 0),
		/*6*/ imports.NewTable("CoreWebView2LaunchingExternalUriSchemeEventArgs_IsUserInitiated", 0),
		/*7*/ imports.NewTable("CoreWebView2LaunchingExternalUriSchemeEventArgs_Uri", 0),
	}
)

func coreWebView2LaunchingExternalUriSchemeEventArgsImportAPI() *imports.Imports {
	if coreWebView2LaunchingExternalUriSchemeEventArgsImport == nil {
		coreWebView2LaunchingExternalUriSchemeEventArgsImport = NewDefaultImports()
		coreWebView2LaunchingExternalUriSchemeEventArgsImport.SetImportTable(coreWebView2LaunchingExternalUriSchemeEventArgsImportTables)
		coreWebView2LaunchingExternalUriSchemeEventArgsImportTables = nil
	}
	return coreWebView2LaunchingExternalUriSchemeEventArgsImport
}
