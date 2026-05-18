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

// ICoreWebView2ServerCertificateErrorDetectedEventArgs Parent: IObject
type ICoreWebView2ServerCertificateErrorDetectedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ServerCertificateErrorDetectedEventArgs // property BaseIntf Getter
	// ErrorStatus
	//  The TLS error code for the invalid certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#get_errorstatus">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</see>
	ErrorStatus() wvTypes.TWVWebErrorStatus // property ErrorStatus Getter
	// RequestUri
	//  URI associated with the request for the invalid certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#get_requesturi">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</see>
	RequestUri() string // property RequestUri Getter
	// ServerCertificate
	//  Returns the server certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#get_servercertificate">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</see>
	ServerCertificate() ICoreWebView2Certificate // property ServerCertificate Getter
	// Action
	//  The action of the server certificate error detection.
	//  The default value is `COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_DEFAULT`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#get_action">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</see>
	Action() wvTypes.TWVServerCertificateErrorAction         // property Action Getter
	SetAction(value wvTypes.TWVServerCertificateErrorAction) // property Action Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2servercertificateerrordetectedeventargs#getdeferral">See the ICoreWebView2ServerCertificateErrorDetectedEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2ServerCertificateErrorDetectedEventArgs struct {
	TObject
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) BaseIntf() (result ICoreWebView2ServerCertificateErrorDetectedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ServerCertificateErrorDetectedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) ErrorStatus() wvTypes.TWVWebErrorStatus {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVWebErrorStatus(r)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) RequestUri() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) ServerCertificate() (result ICoreWebView2Certificate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Certificate(resultPtr)
	return
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) Action() wvTypes.TWVServerCertificateErrorAction {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(6, 0, m.Instance())
	return wvTypes.TWVServerCertificateErrorAction(r)
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) SetAction(value wvTypes.TWVServerCertificateErrorAction) {
	if !m.IsValid() {
		return
	}
	coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2ServerCertificateErrorDetectedEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2ServerCertificateErrorDetectedEventArgs class constructor
func NewCoreWebView2ServerCertificateErrorDetectedEventArgs(args ICoreWebView2ServerCertificateErrorDetectedEventArgs) ICoreWebView2ServerCertificateErrorDetectedEventArgs {
	r := coreWebView2ServerCertificateErrorDetectedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2ServerCertificateErrorDetectedEventArgs(r)
}

var (
	coreWebView2ServerCertificateErrorDetectedEventArgsOnce   base.Once
	coreWebView2ServerCertificateErrorDetectedEventArgsImport *imports.Imports = nil
)

func coreWebView2ServerCertificateErrorDetectedEventArgsAPI() *imports.Imports {
	coreWebView2ServerCertificateErrorDetectedEventArgsOnce.Do(func() {
		coreWebView2ServerCertificateErrorDetectedEventArgsImport = api.NewDefaultImports()
		coreWebView2ServerCertificateErrorDetectedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ServerCertificateErrorDetectedEventArgs_Create", 0), // constructor NewCoreWebView2ServerCertificateErrorDetectedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2ServerCertificateErrorDetectedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ServerCertificateErrorDetectedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ServerCertificateErrorDetectedEventArgs_ErrorStatus", 0), // property ErrorStatus
			/* 4 */ imports.NewTable("TCoreWebView2ServerCertificateErrorDetectedEventArgs_RequestUri", 0), // property RequestUri
			/* 5 */ imports.NewTable("TCoreWebView2ServerCertificateErrorDetectedEventArgs_ServerCertificate", 0), // property ServerCertificate
			/* 6 */ imports.NewTable("TCoreWebView2ServerCertificateErrorDetectedEventArgs_Action", 0), // property Action
			/* 7 */ imports.NewTable("TCoreWebView2ServerCertificateErrorDetectedEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2ServerCertificateErrorDetectedEventArgsImport
}
