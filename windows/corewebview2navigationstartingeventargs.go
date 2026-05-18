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

// ICoreWebView2NavigationStartingEventArgs Parent: IObject
type ICoreWebView2NavigationStartingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2NavigationStartingEventArgs // property BaseIntf Getter
	// URI
	//  The uri of the requested navigation.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_uri">See the ICoreWebView2NavigationStartingEventArgs article.</see>
	URI() string // property URI Getter
	// IsUserInitiated
	//  `TRUE` when the navigation was initiated through a user gesture as
	//  opposed to programmatic navigation by page script. Navigations initiated
	//  via WebView2 APIs are considered as user initiated.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_isuserinitiated">See the ICoreWebView2NavigationStartingEventArgs article.</see>
	IsUserInitiated() bool // property IsUserInitiated Getter
	// IsRedirected
	//  `TRUE` when the navigation is redirected.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_isredirected">See the ICoreWebView2NavigationStartingEventArgs article.</see>
	IsRedirected() bool // property IsRedirected Getter
	// Cancel
	//  The host may set this flag to cancel the navigation. If set, the
	//  navigation is not longer present and the content of the current page is
	//  intact. For performance reasons, `GET` HTTP requests may happen, while
	//  the host is responding. You may set cookies and use part of a request
	//  for the navigation. Navigations to about schemes are cancellable, unless
	//  `msWebView2CancellableAboutNavigations` feature flag is disabled.
	//  Cancellation of frame navigation to `srcdoc` is not supported and
	//  wil be ignored. A cancelled navigation will fire a `NavigationCompleted`
	//  event with a `WebErrorStatus` of
	//  `COREWEBVIEW2_WEB_ERROR_STATUS_OPERATION_CANCELED`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_cancel">See the ICoreWebView2NavigationStartingEventArgs article.</see>
	Cancel() bool         // property Cancel Getter
	SetCancel(value bool) // property Cancel Setter
	// NavigationID
	//  The ID of the navigation.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_navigationid">See the ICoreWebView2NavigationStartingEventArgs article.</see>
	NavigationID() uint64 // property NavigationID Getter
	// RequestHeaders
	//  The HTTP request headers for the navigation.
	//  \>NOTE: You are not able to modify the HTTP request headers in a
	//  `NavigationStarting` event.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs#get_requestheaders">See the ICoreWebView2NavigationStartingEventArgs article.</see>
	RequestHeaders() ICoreWebView2HttpRequestHeaders // property RequestHeaders Getter
	// AdditionalAllowedFrameAncestors
	//  Get additional allowed frame ancestors set by the host app.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs2#get_additionalallowedframeancestors">See the ICoreWebView2NavigationStartingEventArgs3 article.</see>
	AdditionalAllowedFrameAncestors() string         // property AdditionalAllowedFrameAncestors Getter
	SetAdditionalAllowedFrameAncestors(value string) // property AdditionalAllowedFrameAncestors Setter
	// NavigationKind
	//  Get the navigation kind of this navigation.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs3#get_navigationkind">See the ICoreWebView2NavigationStartingEventArgs3 article.</see>
	NavigationKind() wvTypes.TWVNavigationKind // property NavigationKind Getter
}

type TCoreWebView2NavigationStartingEventArgs struct {
	TObject
}

func (m *TCoreWebView2NavigationStartingEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NavigationStartingEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NavigationStartingEventArgs) BaseIntf() (result ICoreWebView2NavigationStartingEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NavigationStartingEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2NavigationStartingEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2NavigationStartingEventArgs) URI() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2NavigationStartingEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2NavigationStartingEventArgs) IsUserInitiated() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NavigationStartingEventArgsAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NavigationStartingEventArgs) IsRedirected() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NavigationStartingEventArgsAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NavigationStartingEventArgs) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NavigationStartingEventArgsAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NavigationStartingEventArgs) SetCancel(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2NavigationStartingEventArgsAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2NavigationStartingEventArgs) NavigationID() (result uint64) {
	if !m.IsValid() {
		return
	}
	coreWebView2NavigationStartingEventArgsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2NavigationStartingEventArgs) RequestHeaders() (result ICoreWebView2HttpRequestHeaders) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NavigationStartingEventArgsAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpRequestHeaders(resultPtr)
	return
}

func (m *TCoreWebView2NavigationStartingEventArgs) AdditionalAllowedFrameAncestors() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2NavigationStartingEventArgsAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2NavigationStartingEventArgs) SetAdditionalAllowedFrameAncestors(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2NavigationStartingEventArgsAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2NavigationStartingEventArgs) NavigationKind() wvTypes.TWVNavigationKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2NavigationStartingEventArgsAPI().SysCallN(10, m.Instance())
	return wvTypes.TWVNavigationKind(r)
}

// NewCoreWebView2NavigationStartingEventArgs class constructor
func NewCoreWebView2NavigationStartingEventArgs(args ICoreWebView2NavigationStartingEventArgs) ICoreWebView2NavigationStartingEventArgs {
	r := coreWebView2NavigationStartingEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2NavigationStartingEventArgs(r)
}

var (
	coreWebView2NavigationStartingEventArgsOnce   base.Once
	coreWebView2NavigationStartingEventArgsImport *imports.Imports = nil
)

func coreWebView2NavigationStartingEventArgsAPI() *imports.Imports {
	coreWebView2NavigationStartingEventArgsOnce.Do(func() {
		coreWebView2NavigationStartingEventArgsImport = api.NewDefaultImports()
		coreWebView2NavigationStartingEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_Create", 0), // constructor NewCoreWebView2NavigationStartingEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_URI", 0), // property URI
			/* 4 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_IsUserInitiated", 0), // property IsUserInitiated
			/* 5 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_IsRedirected", 0), // property IsRedirected
			/* 6 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_Cancel", 0), // property Cancel
			/* 7 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_NavigationID", 0), // property NavigationID
			/* 8 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_RequestHeaders", 0), // property RequestHeaders
			/* 9 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_AdditionalAllowedFrameAncestors", 0), // property AdditionalAllowedFrameAncestors
			/* 10 */ imports.NewTable("TCoreWebView2NavigationStartingEventArgs_NavigationKind", 0), // property NavigationKind
		}
	})
	return coreWebView2NavigationStartingEventArgsImport
}
