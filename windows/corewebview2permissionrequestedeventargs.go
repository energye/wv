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

// ICoreWebView2PermissionRequestedEventArgs Parent: IObject
type ICoreWebView2PermissionRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PermissionRequestedEventArgs // property BaseIntf Getter
	// URI
	//  The origin of the web content that requests the permission.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#get_uri">See the ICoreWebView2PermissionRequestedEventArgs article.</see>
	URI() string // property URI Getter
	// State
	//  The status of a permission request, (for example is the request is granted).
	//  The default value is `COREWEBVIEW2_PERMISSION_STATE_DEFAULT`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#get_state">See the ICoreWebView2PermissionRequestedEventArgs article.</see>
	State() wvTypes.TWVPermissionState         // property State Getter
	SetState(value wvTypes.TWVPermissionState) // property State Setter
	// PermissionKind
	//  The type of the permission that is requested.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#get_permissionkind">See the ICoreWebView2PermissionRequestedEventArgs article.</see>
	PermissionKind() wvTypes.TWVPermissionKind // property PermissionKind Getter
	// IsUserInitiated
	//  `TRUE` when the permission request was initiated through a user gesture.
	//  NOTE: Being initiated through a user gesture does not mean that user intended
	//  to access the associated resource.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#get_isuserinitiated">See the ICoreWebView2PermissionRequestedEventArgs article.</see>
	IsUserInitiated() bool // property IsUserInitiated Getter
	// Deferral
	//  Gets an `ICoreWebView2Deferral` object. Use the deferral object to make
	//  the permission decision at a later time. The deferral only applies to the
	//  current request, and does not prevent the `OnPermissionRequested` event from
	//  getting raised for new requests. However, for some permission kinds the
	//  WebView will avoid creating a new request if there is a pending request of
	//  the same kind.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs#getdeferral">See the ICoreWebView2PermissionRequestedEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
	// Handled
	//  By default, both the `OnPermissionRequested` event handlers on the
	//  `CoreWebView2Frame` and the `CoreWebView2` will be invoked, with the
	//  `CoreWebView2Frame` event handlers invoked first. The host may
	//  set this flag to `TRUE` within the `CoreWebView2Frame` event handlers
	//  to prevent the remaining `CoreWebView2` event handlers from being invoked.
	//  If a deferral is taken on the event args, then you must synchronously
	//  set `Handled` to TRUE prior to taking your deferral to prevent the
	//  `CoreWebView2`s event handlers from being invoked.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs2#get_handled">See the ICoreWebView2PermissionRequestedEventArgs2 article.</see>
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
	// SavesInProfile
	//  The permission state set from the `PermissionRequested` event is saved in
	//  the profile by default; it persists across sessions and becomes the new
	//  default behavior for future `PermissionRequested` events. Browser
	//  heuristics can affect whether the event continues to be raised when the
	//  state is saved in the profile. Set the `SavesInProfile` property to
	//  `FALSE` to not persist the state beyond the current request, and to
	//  continue to receive `PermissionRequested`
	//  events for this origin and permission kind.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs3#get_savesinprofile">See the ICoreWebView2PermissionRequestedEventArgs3 article.</see>
	SavesInProfile() bool         // property SavesInProfile Getter
	SetSavesInProfile(value bool) // property SavesInProfile Setter
}

type TCoreWebView2PermissionRequestedEventArgs struct {
	TObject
}

func (m *TCoreWebView2PermissionRequestedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) BaseIntf() (result ICoreWebView2PermissionRequestedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2PermissionRequestedEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2PermissionRequestedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2PermissionRequestedEventArgs) URI() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) State() wvTypes.TWVPermissionState {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(5, 0, m.Instance())
	return wvTypes.TWVPermissionState(r)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) SetState(value wvTypes.TWVPermissionState) {
	if !m.IsValid() {
		return
	}
	coreWebView2PermissionRequestedEventArgsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PermissionRequestedEventArgs) PermissionKind() wvTypes.TWVPermissionKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(6, m.Instance())
	return wvTypes.TWVPermissionKind(r)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) IsUserInitiated() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2PermissionRequestedEventArgsAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

func (m *TCoreWebView2PermissionRequestedEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2PermissionRequestedEventArgsAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2PermissionRequestedEventArgs) SavesInProfile() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PermissionRequestedEventArgs) SetSavesInProfile(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2PermissionRequestedEventArgsAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

// NewCoreWebView2PermissionRequestedEventArgs class constructor
func NewCoreWebView2PermissionRequestedEventArgs(args ICoreWebView2PermissionRequestedEventArgs) ICoreWebView2PermissionRequestedEventArgs {
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2PermissionRequestedEventArgs(r)
}

// NewCoreWebView2PermissionRequestedEventArgsWithPermissionRequestedEventArgs2 class constructor
func NewCoreWebView2PermissionRequestedEventArgsWithPermissionRequestedEventArgs2(args ICoreWebView2PermissionRequestedEventArgs2) ICoreWebView2PermissionRequestedEventArgs {
	r := coreWebView2PermissionRequestedEventArgsAPI().SysCallN(1, uintptr(args))
	return AsCoreWebView2PermissionRequestedEventArgs(r)
}

var (
	coreWebView2PermissionRequestedEventArgsOnce   base.Once
	coreWebView2PermissionRequestedEventArgsImport *imports.Imports = nil
)

func coreWebView2PermissionRequestedEventArgsAPI() *imports.Imports {
	coreWebView2PermissionRequestedEventArgsOnce.Do(func() {
		coreWebView2PermissionRequestedEventArgsImport = api.NewDefaultImports()
		coreWebView2PermissionRequestedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_Create", 0), // constructor NewCoreWebView2PermissionRequestedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_CreateWithPermissionRequestedEventArgs2", 0), // constructor NewCoreWebView2PermissionRequestedEventArgsWithPermissionRequestedEventArgs2
			/* 2 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_URI", 0), // property URI
			/* 5 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_State", 0), // property State
			/* 6 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_PermissionKind", 0), // property PermissionKind
			/* 7 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_IsUserInitiated", 0), // property IsUserInitiated
			/* 8 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_Deferral", 0), // property Deferral
			/* 9 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_Handled", 0), // property Handled
			/* 10 */ imports.NewTable("TCoreWebView2PermissionRequestedEventArgs_SavesInProfile", 0), // property SavesInProfile
		}
	})
	return coreWebView2PermissionRequestedEventArgsImport
}
