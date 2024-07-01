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

// ICoreWebView2ClientCertificate Parent: IObject
//
//	Provides access to the client certificate metadata.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate">See the ICoreWebView2ClientCertificate article.</a>
type ICoreWebView2ClientCertificate interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ClientCertificate // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2ClientCertificate) // property
	// Subject
	//  Subject of the certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_subject">See the ICoreWebView2ClientCertificate article.</a>
	Subject() string // property
	// Issuer
	//  Name of the certificate authority that issued the certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_issuer">See the ICoreWebView2ClientCertificate article.</a>
	Issuer() string // property
	// ValidFrom
	//  The valid start date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_validfrom">See the ICoreWebView2ClientCertificate article.</a>
	ValidFrom() (resultDateTime float64) // property
	// ValidTo
	//  The valid expiration date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_validto">See the ICoreWebView2ClientCertificate article.</a>
	ValidTo() (resultDateTime float64) // property
	// DerEncodedSerialNumber
	//  Base64 encoding of DER encoded serial number of the certificate.
	//  Read more about DER at [RFC 7468 DER]
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_derencodedserialnumber">See the ICoreWebView2ClientCertificate article.</a>
	DerEncodedSerialNumber() string // property
	// DisplayName
	//  Display name for a certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_displayname">See the ICoreWebView2ClientCertificate article.</a>
	DisplayName() string // property
	// PemEncodedIssuerCertificateChain
	//  Collection of PEM encoded client certificate issuer chain.
	//  In this collection first element is the current certificate followed by
	//  intermediate1, intermediate2...intermediateN-1. Root certificate is the
	//  last element in collection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_pemencodedissuercertificatechain">See the ICoreWebView2ClientCertificate article.</a>
	PemEncodedIssuerCertificateChain() ICoreWebView2StringCollection // property
	// Kind
	//  Kind of a certificate(eg., smart card, pin, other).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_kind">See the ICoreWebView2ClientCertificate article.</a>
	Kind() TWVClientCertificateKind // property
	// ToPemEncoding
	//  PEM encoded data for the certificate.
	//  Returns Base64 encoding of DER encoded certificate.
	//  Read more about PEM at [RFC 1421 Privacy Enhanced Mail]
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#topemencoding">See the ICoreWebView2ClientCertificate article.</a>
	ToPemEncoding() string // function
}

// TCoreWebView2ClientCertificate Parent: TObject
//
//	Provides access to the client certificate metadata.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate">See the ICoreWebView2ClientCertificate article.</a>
type TCoreWebView2ClientCertificate struct {
	TObject
}

func NewCoreWebView2ClientCertificate(aBaseIntf ICoreWebView2ClientCertificate) ICoreWebView2ClientCertificate {
	r1 := coreWebView2ClientCertificateImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ClientCertificate(r1)
}

func (m *TCoreWebView2ClientCertificate) Initialized() bool {
	r1 := coreWebView2ClientCertificateImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificate) BaseIntf() ICoreWebView2ClientCertificate {
	var resultCoreWebView2ClientCertificate uintptr
	coreWebView2ClientCertificateImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ClientCertificate)))
	return AsCoreWebView2ClientCertificate(resultCoreWebView2ClientCertificate)
}

func (m *TCoreWebView2ClientCertificate) SetBaseIntf(AValue ICoreWebView2ClientCertificate) {
	coreWebView2ClientCertificateImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ClientCertificate) Subject() string {
	r1 := coreWebView2ClientCertificateImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificate) Issuer() string {
	r1 := coreWebView2ClientCertificateImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificate) ValidFrom() (resultDateTime float64) {
	coreWebView2ClientCertificateImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCoreWebView2ClientCertificate) ValidTo() (resultDateTime float64) {
	coreWebView2ClientCertificateImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCoreWebView2ClientCertificate) DerEncodedSerialNumber() string {
	r1 := coreWebView2ClientCertificateImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificate) DisplayName() string {
	r1 := coreWebView2ClientCertificateImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ClientCertificate) PemEncodedIssuerCertificateChain() ICoreWebView2StringCollection {
	var resultCoreWebView2StringCollection uintptr
	coreWebView2ClientCertificateImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2StringCollection)))
	return AsCoreWebView2StringCollection(resultCoreWebView2StringCollection)
}

func (m *TCoreWebView2ClientCertificate) Kind() TWVClientCertificateKind {
	r1 := coreWebView2ClientCertificateImportAPI().SysCallN(6, m.Instance())
	return TWVClientCertificateKind(r1)
}

func (m *TCoreWebView2ClientCertificate) ToPemEncoding() string {
	r1 := coreWebView2ClientCertificateImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

var (
	coreWebView2ClientCertificateImport       *imports.Imports = nil
	coreWebView2ClientCertificateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ClientCertificate_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2ClientCertificate_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2ClientCertificate_DerEncodedSerialNumber", 0),
		/*3*/ imports.NewTable("CoreWebView2ClientCertificate_DisplayName", 0),
		/*4*/ imports.NewTable("CoreWebView2ClientCertificate_Initialized", 0),
		/*5*/ imports.NewTable("CoreWebView2ClientCertificate_Issuer", 0),
		/*6*/ imports.NewTable("CoreWebView2ClientCertificate_Kind", 0),
		/*7*/ imports.NewTable("CoreWebView2ClientCertificate_PemEncodedIssuerCertificateChain", 0),
		/*8*/ imports.NewTable("CoreWebView2ClientCertificate_Subject", 0),
		/*9*/ imports.NewTable("CoreWebView2ClientCertificate_ToPemEncoding", 0),
		/*10*/ imports.NewTable("CoreWebView2ClientCertificate_ValidFrom", 0),
		/*11*/ imports.NewTable("CoreWebView2ClientCertificate_ValidTo", 0),
	}
)

func coreWebView2ClientCertificateImportAPI() *imports.Imports {
	if coreWebView2ClientCertificateImport == nil {
		coreWebView2ClientCertificateImport = NewDefaultImports()
		coreWebView2ClientCertificateImport.SetImportTable(coreWebView2ClientCertificateImportTables)
		coreWebView2ClientCertificateImportTables = nil
	}
	return coreWebView2ClientCertificateImport
}
