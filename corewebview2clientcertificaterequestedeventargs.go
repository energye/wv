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

// ICoreWebView2ClientCertificateRequestedEventArgs Parent: IObject
//
//	Event args for the ClientCertificateRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
type ICoreWebView2ClientCertificateRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ClientCertificateRequestedEventArgs // property
	// Host
	//  Host name of the server that requested client certificate authentication.
	//  Normalization rules applied to the hostname are:
	//  * Convert to lowercase characters for ascii characters.
	//  * Punycode is used for representing non ascii characters.
	//  * Strip square brackets for IPV6 address.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_host">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Host() string // property
	// Port
	//  Port of the server that requested client certificate authentication.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_port">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Port() int32 // property
	// IsProxy
	//  Returns true if the server that issued this request is an http proxy.
	//  Returns false if the server is the origin server.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_isproxy">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	IsProxy() bool // property
	// AllowedCertificateAuthorities
	//  Returns the `ICoreWebView2StringCollection`.
	//  The collection contains Base64 encoding of DER encoded distinguished names of
	//  certificate authorities allowed by the server.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_allowedcertificateauthorities">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	AllowedCertificateAuthorities() ICoreWebView2StringCollection // property
	// MutuallyTrustedCertificates
	//  Returns the `ICoreWebView2ClientCertificateCollection` when client
	//  certificate authentication is requested. The collection contains mutually
	//  trusted CA certificates.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_mutuallytrustedcertificates">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	MutuallyTrustedCertificates() ICoreWebView2ClientCertificateCollection // property
	// SelectedCertificate
	//  Returns the selected certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_selectedcertificate">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	SelectedCertificate() ICoreWebView2ClientCertificate // property
	// SetSelectedCertificate Set SelectedCertificate
	SetSelectedCertificate(AValue ICoreWebView2ClientCertificate) // property
	// Cancel
	//  You may set this flag to cancel the certificate selection. If canceled,
	//  the request is aborted regardless of the `Handled` property. By default the
	//  value is `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_cancel">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Cancel() bool // property
	// SetCancel Set Cancel
	SetCancel(AValue bool) // property
	// Handled
	//  You may set this flag to `TRUE` to respond to the server with or
	//  without a certificate. If this flag is `TRUE` with a `SelectedCertificate`
	//  it responds to the server with the selected certificate otherwise respond to the
	//  server without a certificate. By default the value of `Handled` and `Cancel` are `FALSE`
	//  and display default client certificate selection dialog prompt to allow the user to
	//  choose a certificate. The `SelectedCertificate` is ignored unless `Handled` is set `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_handled">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Handled() bool // property
	// SetHandled Set Handled
	SetHandled(AValue bool) // property
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#getdeferral">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
}

// TCoreWebView2ClientCertificateRequestedEventArgs Parent: TObject
//
//	Event args for the ClientCertificateRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</a>
type TCoreWebView2ClientCertificateRequestedEventArgs struct {
	TObject
}

func NewCoreWebView2ClientCertificateRequestedEventArgs(aArgs ICoreWebView2ClientCertificateRequestedEventArgs) ICoreWebView2ClientCertificateRequestedEventArgs {
	r1 := coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(3, GetObjectUintptr(aArgs))
	return AsCoreWebView2ClientCertificateRequestedEventArgs(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Initialized() bool {
	r1 := coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) BaseIntf() ICoreWebView2ClientCertificateRequestedEventArgs {
	var resultCoreWebView2ClientCertificateRequestedEventArgs uintptr
	coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ClientCertificateRequestedEventArgs)))
	return AsCoreWebView2ClientCertificateRequestedEventArgs(resultCoreWebView2ClientCertificateRequestedEventArgs)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Host() string {
	r1 := coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Port() int32 {
	r1 := coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) IsProxy() bool {
	r1 := coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) AllowedCertificateAuthorities() ICoreWebView2StringCollection {
	var resultCoreWebView2StringCollection uintptr
	coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2StringCollection)))
	return AsCoreWebView2StringCollection(resultCoreWebView2StringCollection)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) MutuallyTrustedCertificates() ICoreWebView2ClientCertificateCollection {
	var resultCoreWebView2ClientCertificateCollection uintptr
	coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ClientCertificateCollection)))
	return AsCoreWebView2ClientCertificateCollection(resultCoreWebView2ClientCertificateCollection)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SelectedCertificate() ICoreWebView2ClientCertificate {
	var resultCoreWebView2ClientCertificate uintptr
	coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ClientCertificate)))
	return AsCoreWebView2ClientCertificate(resultCoreWebView2ClientCertificate)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetSelectedCertificate(AValue ICoreWebView2ClientCertificate) {
	coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(11, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Cancel() bool {
	r1 := coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetCancel(AValue bool) {
	coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Handled() bool {
	r1 := coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetHandled(AValue bool) {
	coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	coreWebView2ClientCertificateRequestedEventArgsImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

var (
	coreWebView2ClientCertificateRequestedEventArgsImport       *imports.Imports = nil
	coreWebView2ClientCertificateRequestedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_AllowedCertificateAuthorities", 0),
		/*1*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_BaseIntf", 0),
		/*2*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_Cancel", 0),
		/*3*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_Create", 0),
		/*4*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_Deferral", 0),
		/*5*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_Handled", 0),
		/*6*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_Host", 0),
		/*7*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_Initialized", 0),
		/*8*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_IsProxy", 0),
		/*9*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_MutuallyTrustedCertificates", 0),
		/*10*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_Port", 0),
		/*11*/ imports.NewTable("CoreWebView2ClientCertificateRequestedEventArgs_SelectedCertificate", 0),
	}
)

func coreWebView2ClientCertificateRequestedEventArgsImportAPI() *imports.Imports {
	if coreWebView2ClientCertificateRequestedEventArgsImport == nil {
		coreWebView2ClientCertificateRequestedEventArgsImport = NewDefaultImports()
		coreWebView2ClientCertificateRequestedEventArgsImport.SetImportTable(coreWebView2ClientCertificateRequestedEventArgsImportTables)
		coreWebView2ClientCertificateRequestedEventArgsImportTables = nil
	}
	return coreWebView2ClientCertificateRequestedEventArgsImport
}
