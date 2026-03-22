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
	"github.com/energye/lcl/types"
)

// ICoreWebView2Certificate Parent: IObject
type ICoreWebView2Certificate interface {
	IObject
	// ToPemEncoding
	//  PEM encoded data for the certificate.
	//  Returns Base64 encoding of DER encoded certificate.
	//  Read more about PEM at [RFC 1421 Privacy Enhanced Mail]
	//  (https://tools.ietf.org/html/rfc1421).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#topemencoding">See the ICoreWebView2Certificate article.</see>
	ToPemEncoding() string // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Certificate         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2Certificate) // property BaseIntf Setter
	// Subject
	//  Subject of the certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_subject">See the ICoreWebView2Certificate article.</see>
	Subject() string // property Subject Getter
	// Issuer
	//  Name of the certificate authority that issued the certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_issuer">See the ICoreWebView2Certificate article.</see>
	Issuer() string // property Issuer Getter
	// ValidFrom
	//  The valid start date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_validfrom">See the ICoreWebView2Certificate article.</see>
	ValidFrom() types.TDateTime // property ValidFrom Getter
	// ValidTo
	//  The valid expiration date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_validto">See the ICoreWebView2Certificate article.</see>
	ValidTo() types.TDateTime // property ValidTo Getter
	// DerEncodedSerialNumber
	//  Base64 encoding of DER encoded serial number of the certificate.
	//  Read more about DER at [RFC 7468 DER]
	//  (https://tools.ietf.org/html/rfc7468#appendix-B).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_derencodedserialnumber">See the ICoreWebView2Certificate article.</see>
	DerEncodedSerialNumber() string // property DerEncodedSerialNumber Getter
	// DisplayName
	//  Display name for a certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_displayname">See the ICoreWebView2Certificate article.</see>
	DisplayName() string // property DisplayName Getter
	// PemEncodedIssuerCertificateChain
	//  Collection of PEM encoded certificate issuer chain.
	//  In this collection first element is the current certificate followed by
	//  intermediate1, intermediate2...intermediateN-1. Root certificate is the
	//  last element in collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2certificate#get_pemencodedissuercertificatechain">See the ICoreWebView2Certificate article.</see>
	PemEncodedIssuerCertificateChain() ICoreWebView2StringCollection // property PemEncodedIssuerCertificateChain Getter
}

type TCoreWebView2Certificate struct {
	TObject
}

func (m *TCoreWebView2Certificate) ToPemEncoding() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2CertificateAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Certificate) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CertificateAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Certificate) BaseIntf() (result ICoreWebView2Certificate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CertificateAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Certificate(resultPtr)
	return
}

func (m *TCoreWebView2Certificate) SetBaseIntf(value ICoreWebView2Certificate) {
	if !m.IsValid() {
		return
	}
	coreWebView2CertificateAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2Certificate) Subject() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2CertificateAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Certificate) Issuer() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2CertificateAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Certificate) ValidFrom() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	coreWebView2CertificateAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2Certificate) ValidTo() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	coreWebView2CertificateAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2Certificate) DerEncodedSerialNumber() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2CertificateAPI().SysCallN(8, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Certificate) DisplayName() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2CertificateAPI().SysCallN(9, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Certificate) PemEncodedIssuerCertificateChain() (result ICoreWebView2StringCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CertificateAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2StringCollection(resultPtr)
	return
}

// NewCoreWebView2Certificate class constructor
func NewCoreWebView2Certificate(baseIntf ICoreWebView2Certificate) ICoreWebView2Certificate {
	r := coreWebView2CertificateAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Certificate(r)
}

var (
	coreWebView2CertificateOnce   base.Once
	coreWebView2CertificateImport *imports.Imports = nil
)

func coreWebView2CertificateAPI() *imports.Imports {
	coreWebView2CertificateOnce.Do(func() {
		coreWebView2CertificateImport = api.NewDefaultImports()
		coreWebView2CertificateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Certificate_Create", 0), // constructor NewCoreWebView2Certificate
			/* 1 */ imports.NewTable("TCoreWebView2Certificate_ToPemEncoding", 0), // function ToPemEncoding
			/* 2 */ imports.NewTable("TCoreWebView2Certificate_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2Certificate_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2Certificate_Subject", 0), // property Subject
			/* 5 */ imports.NewTable("TCoreWebView2Certificate_Issuer", 0), // property Issuer
			/* 6 */ imports.NewTable("TCoreWebView2Certificate_ValidFrom", 0), // property ValidFrom
			/* 7 */ imports.NewTable("TCoreWebView2Certificate_ValidTo", 0), // property ValidTo
			/* 8 */ imports.NewTable("TCoreWebView2Certificate_DerEncodedSerialNumber", 0), // property DerEncodedSerialNumber
			/* 9 */ imports.NewTable("TCoreWebView2Certificate_DisplayName", 0), // property DisplayName
			/* 10 */ imports.NewTable("TCoreWebView2Certificate_PemEncodedIssuerCertificateChain", 0), // property PemEncodedIssuerCertificateChain
		}
	})
	return coreWebView2CertificateImport
}
