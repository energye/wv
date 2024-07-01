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

// ICoreWebView2Certificate Parent: IObject
//
//	Provides access to the certificate metadata.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate">See the ICoreWebView2Certificate article.</a>
type ICoreWebView2Certificate interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Certificate // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2Certificate) // property
	// Subject
	//  Subject of the certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_subject">See the ICoreWebView2Certificate article.</a>
	Subject() string // property
	// Issuer
	//  Name of the certificate authority that issued the certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_issuer">See the ICoreWebView2Certificate article.</a>
	Issuer() string // property
	// ValidFrom
	//  The valid start date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_validfrom">See the ICoreWebView2Certificate article.</a>
	ValidFrom() (resultDateTime float64) // property
	// ValidTo
	//  The valid expiration date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_validto">See the ICoreWebView2Certificate article.</a>
	ValidTo() (resultDateTime float64) // property
	// DerEncodedSerialNumber
	//  Base64 encoding of DER encoded serial number of the certificate.
	//  Read more about DER at [RFC 7468 DER]
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_derencodedserialnumber">See the ICoreWebView2Certificate article.</a>
	DerEncodedSerialNumber() string // property
	// DisplayName
	//  Display name for a certificate.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_displayname">See the ICoreWebView2Certificate article.</a>
	DisplayName() string // property
	// PemEncodedIssuerCertificateChain
	//  Collection of PEM encoded certificate issuer chain.
	//  In this collection first element is the current certificate followed by
	//  intermediate1, intermediate2...intermediateN-1. Root certificate is the
	//  last element in collection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_pemencodedissuercertificatechain">See the ICoreWebView2Certificate article.</a>
	PemEncodedIssuerCertificateChain() ICoreWebView2StringCollection // property
	// ToPemEncoding
	//  PEM encoded data for the certificate.
	//  Returns Base64 encoding of DER encoded certificate.
	//  Read more about PEM at [RFC 1421 Privacy Enhanced Mail]
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#topemencoding">See the ICoreWebView2Certificate article.</a>
	ToPemEncoding() string // function
}

// TCoreWebView2Certificate Parent: TObject
//
//	Provides access to the certificate metadata.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate">See the ICoreWebView2Certificate article.</a>
type TCoreWebView2Certificate struct {
	TObject
}

func NewCoreWebView2Certificate(aBaseIntf ICoreWebView2Certificate) ICoreWebView2Certificate {
	r1 := coreWebView2CertificateImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Certificate(r1)
}

func (m *TCoreWebView2Certificate) Initialized() bool {
	r1 := coreWebView2CertificateImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Certificate) BaseIntf() ICoreWebView2Certificate {
	var resultCoreWebView2Certificate uintptr
	coreWebView2CertificateImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2Certificate)))
	return AsCoreWebView2Certificate(resultCoreWebView2Certificate)
}

func (m *TCoreWebView2Certificate) SetBaseIntf(AValue ICoreWebView2Certificate) {
	coreWebView2CertificateImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2Certificate) Subject() string {
	r1 := coreWebView2CertificateImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Certificate) Issuer() string {
	r1 := coreWebView2CertificateImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Certificate) ValidFrom() (resultDateTime float64) {
	coreWebView2CertificateImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCoreWebView2Certificate) ValidTo() (resultDateTime float64) {
	coreWebView2CertificateImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCoreWebView2Certificate) DerEncodedSerialNumber() string {
	r1 := coreWebView2CertificateImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Certificate) DisplayName() string {
	r1 := coreWebView2CertificateImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Certificate) PemEncodedIssuerCertificateChain() ICoreWebView2StringCollection {
	var resultCoreWebView2StringCollection uintptr
	coreWebView2CertificateImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2StringCollection)))
	return AsCoreWebView2StringCollection(resultCoreWebView2StringCollection)
}

func (m *TCoreWebView2Certificate) ToPemEncoding() string {
	r1 := coreWebView2CertificateImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

var (
	coreWebView2CertificateImport       *imports.Imports = nil
	coreWebView2CertificateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2Certificate_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2Certificate_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2Certificate_DerEncodedSerialNumber", 0),
		/*3*/ imports.NewTable("CoreWebView2Certificate_DisplayName", 0),
		/*4*/ imports.NewTable("CoreWebView2Certificate_Initialized", 0),
		/*5*/ imports.NewTable("CoreWebView2Certificate_Issuer", 0),
		/*6*/ imports.NewTable("CoreWebView2Certificate_PemEncodedIssuerCertificateChain", 0),
		/*7*/ imports.NewTable("CoreWebView2Certificate_Subject", 0),
		/*8*/ imports.NewTable("CoreWebView2Certificate_ToPemEncoding", 0),
		/*9*/ imports.NewTable("CoreWebView2Certificate_ValidFrom", 0),
		/*10*/ imports.NewTable("CoreWebView2Certificate_ValidTo", 0),
	}
)

func coreWebView2CertificateImportAPI() *imports.Imports {
	if coreWebView2CertificateImport == nil {
		coreWebView2CertificateImport = NewDefaultImports()
		coreWebView2CertificateImport.SetImportTable(coreWebView2CertificateImportTables)
		coreWebView2CertificateImportTables = nil
	}
	return coreWebView2CertificateImport
}
