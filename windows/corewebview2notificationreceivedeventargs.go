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

// ICoreWebView2NotificationReceivedEventArgs Parent: IObject
type ICoreWebView2NotificationReceivedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2NotificationReceivedEventArgs // property BaseIntf Getter
	// SenderOrigin
	//  The origin of the web content that sends the notification, such as
	//  `https://example.com/` or `https://www.example.com/`.
	SenderOrigin() string // property SenderOrigin Getter
	// Notification
	//  The notification that was received. You can access the
	//  properties on the Notification object to show your own notification.
	Notification() ICoreWebView2Notification // property Notification Getter
	// Handled
	//  Sets whether the `NotificationReceived` event is handled by the host after
	//  the event handler completes or if there is a deferral then after the
	//  deferral is completed.
	//
	//  If `Handled` is set to TRUE then WebView will not display the notification
	//  with the default UI, and the host will be responsible for handling the
	//  notification and for letting the web content know that the notification
	//  has been displayed, clicked, or closed. You must set `Handled` to `TRUE`
	//  before you call `ReportShown`, `ReportClicked`,
	//  `ReportClickedWithActionIndex` and `ReportClosed`, otherwise they will
	//  fail with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`. If after the event
	//  handler or deferral completes `Handled` is set to FALSE then WebView will
	//  display the default notification UI. Note that you cannot un-handle this
	//  event once you have set `Handled` to be `TRUE`. The initial value is
	//  FALSE.
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to complete
	//  the event at a later time.
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2NotificationReceivedEventArgs struct {
	TObject
}

func (m *TCoreWebView2NotificationReceivedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationReceivedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NotificationReceivedEventArgs) BaseIntf() (result ICoreWebView2NotificationReceivedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NotificationReceivedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2NotificationReceivedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2NotificationReceivedEventArgs) SenderOrigin() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NotificationReceivedEventArgsAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2NotificationReceivedEventArgs) Notification() (result ICoreWebView2Notification) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NotificationReceivedEventArgsAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Notification(resultPtr)
	return
}

func (m *TCoreWebView2NotificationReceivedEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationReceivedEventArgsAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NotificationReceivedEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2NotificationReceivedEventArgsAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2NotificationReceivedEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NotificationReceivedEventArgsAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2NotificationReceivedEventArgs class constructor
func NewCoreWebView2NotificationReceivedEventArgs(args ICoreWebView2NotificationReceivedEventArgs) ICoreWebView2NotificationReceivedEventArgs {
	r := coreWebView2NotificationReceivedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2NotificationReceivedEventArgs(r)
}

var (
	coreWebView2NotificationReceivedEventArgsOnce   base.Once
	coreWebView2NotificationReceivedEventArgsImport *imports.Imports = nil
)

func coreWebView2NotificationReceivedEventArgsAPI() *imports.Imports {
	coreWebView2NotificationReceivedEventArgsOnce.Do(func() {
		coreWebView2NotificationReceivedEventArgsImport = api.NewDefaultImports()
		coreWebView2NotificationReceivedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2NotificationReceivedEventArgs_Create", 0), // constructor NewCoreWebView2NotificationReceivedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2NotificationReceivedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2NotificationReceivedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2NotificationReceivedEventArgs_SenderOrigin", 0), // property SenderOrigin
			/* 4 */ imports.NewTable("TCoreWebView2NotificationReceivedEventArgs_Notification", 0), // property Notification
			/* 5 */ imports.NewTable("TCoreWebView2NotificationReceivedEventArgs_Handled", 0), // property Handled
			/* 6 */ imports.NewTable("TCoreWebView2NotificationReceivedEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2NotificationReceivedEventArgsImport
}
