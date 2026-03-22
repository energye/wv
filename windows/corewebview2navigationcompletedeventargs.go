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

// ICoreWebView2NavigationCompletedEventArgs Parent: IObject
type ICoreWebView2NavigationCompletedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2NavigationCompletedEventArgs // property BaseIntf Getter
	// IsSuccess
	//  `TRUE` when the navigation is successful. `FALSE` for a navigation that
	//  ended up in an error page (failures due to no network, DNS lookup
	//  failure, HTTP server responds with 4xx), but may also be `FALSE` for
	//  additional scenarios such as `window.stop()` run on navigated page.
	//  Note that WebView2 will report the navigation as 'unsuccessful' if the load
	//  for the navigation did not reach the expected completion for any reason. Such
	//  reasons include potentially catastrophic issues such network and certificate
	//  issues, but can also be the result of intended actions such as the app canceling a navigation or
	//  navigating away before the original navigation completed. Applications should not
	//  just rely on this flag, but also consider the reported WebErrorStatus to
	//  determine whether the failure is indeed catastrophic in their context.
	//  WebErrorStatuses that may indicate a non-catastrophic failure include:
	//   - COREWEBVIEW2_WEB_ERROR_STATUS_OPERATION_CANCELED
	//   - COREWEBVIEW2_WEB_ERROR_STATUS_VALID_AUTHENTICATION_CREDENTIALS_REQUIRED
	//   - COREWEBVIEW2_WEB_ERROR_STATUS_VALID_PROXY_AUTHENTICATION_REQUIRED
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs#get_issuccess">See the ICoreWebView2NavigationCompletedEventArgs article.</see>
	IsSuccess() bool // property IsSuccess Getter
	// WebErrorStatus
	//  The error code if the navigation failed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs#get_weberrorstatus">See the ICoreWebView2NavigationCompletedEventArgs article.</see>
	WebErrorStatus() wvTypes.TWVWebErrorStatus // property WebErrorStatus Getter
	// NavigationID
	//  The ID of the navigation.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs#get_navigationid">See the ICoreWebView2NavigationCompletedEventArgs article.</see>
	NavigationID() uint64 // property NavigationID Getter
	// HttpStatusCode
	//  The HTTP status code of the navigation if it involved an HTTP request.
	//  For instance, this will usually be 200 if the request was successful, 404
	//  if a page was not found, etc. See
	//  https://developer.mozilla.org/docs/Web/HTTP/Status for a list of
	//  common status codes.
	//  The `HttpStatusCode` property will be 0 in the following cases:
	//  * The navigation did not involve an HTTP request. For instance, if it was
	//  a navigation to a file:// URL, or if it was a same-document navigation.
	//  * The navigation failed before a response was received. For instance, if
	//  the hostname was not found, or if there was a network error.
	//  In those cases, you can get more information from the `IsSuccess` and
	//  `WebErrorStatus` properties.
	//  If the navigation receives a successful HTTP response, but the navigated
	//  page calls `window.stop()` before it finishes loading, then
	//  `HttpStatusCode` may contain a success code like 200, but `IsSuccess` will
	//  be FALSE and `WebErrorStatus` will be
	//  `COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_ABORTED`.
	//  Since WebView2 handles HTTP continuations and redirects automatically, it
	//  is unlikely for `HttpStatusCode` to ever be in the 1xx or 3xx ranges.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs2#get_httpstatuscode">See the ICoreWebView2NavigationCompletedEventArgs2 article.</see>
	HttpStatusCode() int32 // property HttpStatusCode Getter
}

type TCoreWebView2NavigationCompletedEventArgs struct {
	TObject
}

func (m *TCoreWebView2NavigationCompletedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NavigationCompletedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) BaseIntf() (result ICoreWebView2NavigationCompletedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NavigationCompletedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2NavigationCompletedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2NavigationCompletedEventArgs) IsSuccess() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NavigationCompletedEventArgsAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) WebErrorStatus() wvTypes.TWVWebErrorStatus {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2NavigationCompletedEventArgsAPI().SysCallN(4, m.Instance())
	return wvTypes.TWVWebErrorStatus(r)
}

func (m *TCoreWebView2NavigationCompletedEventArgs) NavigationID() (result uint64) {
	if !m.IsValid() {
		return
	}
	coreWebView2NavigationCompletedEventArgsAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2NavigationCompletedEventArgs) HttpStatusCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2NavigationCompletedEventArgsAPI().SysCallN(6, m.Instance())
	return int32(r)
}

// NewCoreWebView2NavigationCompletedEventArgs class constructor
func NewCoreWebView2NavigationCompletedEventArgs(args ICoreWebView2NavigationCompletedEventArgs) ICoreWebView2NavigationCompletedEventArgs {
	r := coreWebView2NavigationCompletedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2NavigationCompletedEventArgs(r)
}

var (
	coreWebView2NavigationCompletedEventArgsOnce   base.Once
	coreWebView2NavigationCompletedEventArgsImport *imports.Imports = nil
)

func coreWebView2NavigationCompletedEventArgsAPI() *imports.Imports {
	coreWebView2NavigationCompletedEventArgsOnce.Do(func() {
		coreWebView2NavigationCompletedEventArgsImport = api.NewDefaultImports()
		coreWebView2NavigationCompletedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2NavigationCompletedEventArgs_Create", 0), // constructor NewCoreWebView2NavigationCompletedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2NavigationCompletedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2NavigationCompletedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2NavigationCompletedEventArgs_IsSuccess", 0), // property IsSuccess
			/* 4 */ imports.NewTable("TCoreWebView2NavigationCompletedEventArgs_WebErrorStatus", 0), // property WebErrorStatus
			/* 5 */ imports.NewTable("TCoreWebView2NavigationCompletedEventArgs_NavigationID", 0), // property NavigationID
			/* 6 */ imports.NewTable("TCoreWebView2NavigationCompletedEventArgs_HttpStatusCode", 0), // property HttpStatusCode
		}
	})
	return coreWebView2NavigationCompletedEventArgsImport
}
