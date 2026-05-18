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

// ICoreWebView2ClientCertificateRequestedEventArgs Parent: IObject
type ICoreWebView2ClientCertificateRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ClientCertificateRequestedEventArgs // property BaseIntf Getter
	// Host
	//  Host name of the server that requested client certificate authentication.
	//  Normalization rules applied to the hostname are:
	//  * Convert to lowercase characters for ascii characters.
	//  * Punycode is used for representing non ascii characters.
	//  * Strip square brackets for IPV6 address.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_host">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	Host() string // property Host Getter
	// Port
	//  Port of the server that requested client certificate authentication.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_port">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	Port() int32 // property Port Getter
	// IsProxy
	//  Returns true if the server that issued this request is an http proxy.
	//  Returns false if the server is the origin server.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_isproxy">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	IsProxy() bool // property IsProxy Getter
	// AllowedCertificateAuthorities
	//  Returns the `ICoreWebView2StringCollection`.
	//  The collection contains Base64 encoding of DER encoded distinguished names of
	//  certificate authorities allowed by the server.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_allowedcertificateauthorities">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	AllowedCertificateAuthorities() ICoreWebView2StringCollection // property AllowedCertificateAuthorities Getter
	// MutuallyTrustedCertificates
	//  Returns the `ICoreWebView2ClientCertificateCollection` when client
	//  certificate authentication is requested. The collection contains mutually
	//  trusted CA certificates.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_mutuallytrustedcertificates">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	MutuallyTrustedCertificates() ICoreWebView2ClientCertificateCollection // property MutuallyTrustedCertificates Getter
	// SelectedCertificate
	//  Returns the selected certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_selectedcertificate">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	SelectedCertificate() ICoreWebView2ClientCertificate         // property SelectedCertificate Getter
	SetSelectedCertificate(value ICoreWebView2ClientCertificate) // property SelectedCertificate Setter
	// Cancel
	//  You may set this flag to cancel the certificate selection. If canceled,
	//  the request is aborted regardless of the `Handled` property. By default the
	//  value is `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_cancel">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	Cancel() bool         // property Cancel Getter
	SetCancel(value bool) // property Cancel Setter
	// Handled
	//  You may set this flag to `TRUE` to respond to the server with or
	//  without a certificate. If this flag is `TRUE` with a `SelectedCertificate`
	//  it responds to the server with the selected certificate otherwise respond to the
	//  server without a certificate. By default the value of `Handled` and `Cancel` are `FALSE`
	//  and display default client certificate selection dialog prompt to allow the user to
	//  choose a certificate. The `SelectedCertificate` is ignored unless `Handled` is set `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#get_handled">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs#getdeferral">See the ICoreWebView2ClientCertificateRequestedEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2ClientCertificateRequestedEventArgs struct {
	TObject
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) BaseIntf() (result ICoreWebView2ClientCertificateRequestedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ClientCertificateRequestedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Host() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Port() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) IsProxy() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) AllowedCertificateAuthorities() (result ICoreWebView2StringCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2StringCollection(resultPtr)
	return
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) MutuallyTrustedCertificates() (result ICoreWebView2ClientCertificateCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ClientCertificateCollection(resultPtr)
	return
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SelectedCertificate() (result ICoreWebView2ClientCertificate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ClientCertificate(resultPtr)
	return
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetSelectedCertificate(value ICoreWebView2ClientCertificate) {
	if !m.IsValid() {
		return
	}
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetCancel(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2ClientCertificateRequestedEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2ClientCertificateRequestedEventArgs class constructor
func NewCoreWebView2ClientCertificateRequestedEventArgs(args ICoreWebView2ClientCertificateRequestedEventArgs) ICoreWebView2ClientCertificateRequestedEventArgs {
	r := coreWebView2ClientCertificateRequestedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2ClientCertificateRequestedEventArgs(r)
}

var (
	coreWebView2ClientCertificateRequestedEventArgsOnce   base.Once
	coreWebView2ClientCertificateRequestedEventArgsImport *imports.Imports = nil
)

func coreWebView2ClientCertificateRequestedEventArgsAPI() *imports.Imports {
	coreWebView2ClientCertificateRequestedEventArgsOnce.Do(func() {
		coreWebView2ClientCertificateRequestedEventArgsImport = api.NewDefaultImports()
		coreWebView2ClientCertificateRequestedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_Create", 0), // constructor NewCoreWebView2ClientCertificateRequestedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_Host", 0), // property Host
			/* 4 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_Port", 0), // property Port
			/* 5 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_IsProxy", 0), // property IsProxy
			/* 6 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_AllowedCertificateAuthorities", 0), // property AllowedCertificateAuthorities
			/* 7 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_MutuallyTrustedCertificates", 0), // property MutuallyTrustedCertificates
			/* 8 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_SelectedCertificate", 0), // property SelectedCertificate
			/* 9 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_Cancel", 0), // property Cancel
			/* 10 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_Handled", 0), // property Handled
			/* 11 */ imports.NewTable("TCoreWebView2ClientCertificateRequestedEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2ClientCertificateRequestedEventArgsImport
}
