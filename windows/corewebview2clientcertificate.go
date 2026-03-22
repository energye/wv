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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2ClientCertificate Parent: IObject
type ICoreWebView2ClientCertificate interface {
	IObject
	// ToPemEncoding
	//  PEM encoded data for the certificate.
	//  Returns Base64 encoding of DER encoded certificate.
	//  Read more about PEM at [RFC 1421 Privacy Enhanced Mail]
	//  (https://tools.ietf.org/html/rfc1421).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#topemencoding">See the ICoreWebView2ClientCertificate article.</see>
	ToPemEncoding() string // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ClientCertificate         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2ClientCertificate) // property BaseIntf Setter
	// Subject
	//  Subject of the certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_subject">See the ICoreWebView2ClientCertificate article.</see>
	Subject() string // property Subject Getter
	// Issuer
	//  Name of the certificate authority that issued the certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_issuer">See the ICoreWebView2ClientCertificate article.</see>
	Issuer() string // property Issuer Getter
	// ValidFrom
	//  The valid start date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_validfrom">See the ICoreWebView2ClientCertificate article.</see>
	ValidFrom() types.TDateTime // property ValidFrom Getter
	// ValidTo
	//  The valid expiration date and time for the certificate as the number of seconds since
	//  the UNIX epoch.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_validto">See the ICoreWebView2ClientCertificate article.</see>
	ValidTo() types.TDateTime // property ValidTo Getter
	// DerEncodedSerialNumber
	//  Base64 encoding of DER encoded serial number of the certificate.
	//  Read more about DER at [RFC 7468 DER]
	//  (https://tools.ietf.org/html/rfc7468#appendix-B).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_derencodedserialnumber">See the ICoreWebView2ClientCertificate article.</see>
	DerEncodedSerialNumber() string // property DerEncodedSerialNumber Getter
	// DisplayName
	//  Display name for a certificate.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_displayname">See the ICoreWebView2ClientCertificate article.</see>
	DisplayName() string // property DisplayName Getter
	// PemEncodedIssuerCertificateChain
	//  Collection of PEM encoded client certificate issuer chain.
	//  In this collection first element is the current certificate followed by
	//  intermediate1, intermediate2...intermediateN-1. Root certificate is the
	//  last element in collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_pemencodedissuercertificatechain">See the ICoreWebView2ClientCertificate article.</see>
	PemEncodedIssuerCertificateChain() ICoreWebView2StringCollection // property PemEncodedIssuerCertificateChain Getter
	// Kind
	//  Kind of a certificate (eg., smart card, pin, other).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate#get_kind">See the ICoreWebView2ClientCertificate article.</see>
	Kind() wvTypes.TWVClientCertificateKind // property Kind Getter
}

type TCoreWebView2ClientCertificate struct {
	TObject
}

func (m *TCoreWebView2ClientCertificate) ToPemEncoding() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ClientCertificateAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ClientCertificate) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ClientCertificateAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ClientCertificate) BaseIntf() (result ICoreWebView2ClientCertificate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ClientCertificate(resultPtr)
	return
}

func (m *TCoreWebView2ClientCertificate) SetBaseIntf(value ICoreWebView2ClientCertificate) {
	if !m.IsValid() {
		return
	}
	coreWebView2ClientCertificateAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2ClientCertificate) Subject() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ClientCertificateAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ClientCertificate) Issuer() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ClientCertificateAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ClientCertificate) ValidFrom() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	coreWebView2ClientCertificateAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2ClientCertificate) ValidTo() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	coreWebView2ClientCertificateAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2ClientCertificate) DerEncodedSerialNumber() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ClientCertificateAPI().SysCallN(8, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ClientCertificate) DisplayName() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ClientCertificateAPI().SysCallN(9, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ClientCertificate) PemEncodedIssuerCertificateChain() (result ICoreWebView2StringCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2StringCollection(resultPtr)
	return
}

func (m *TCoreWebView2ClientCertificate) Kind() wvTypes.TWVClientCertificateKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ClientCertificateAPI().SysCallN(11, m.Instance())
	return wvTypes.TWVClientCertificateKind(r)
}

// NewCoreWebView2ClientCertificate class constructor
func NewCoreWebView2ClientCertificate(baseIntf ICoreWebView2ClientCertificate) ICoreWebView2ClientCertificate {
	r := coreWebView2ClientCertificateAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ClientCertificate(r)
}

var (
	coreWebView2ClientCertificateOnce   base.Once
	coreWebView2ClientCertificateImport *imports.Imports = nil
)

func coreWebView2ClientCertificateAPI() *imports.Imports {
	coreWebView2ClientCertificateOnce.Do(func() {
		coreWebView2ClientCertificateImport = api.NewDefaultImports()
		coreWebView2ClientCertificateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ClientCertificate_Create", 0), // constructor NewCoreWebView2ClientCertificate
			/* 1 */ imports.NewTable("TCoreWebView2ClientCertificate_ToPemEncoding", 0), // function ToPemEncoding
			/* 2 */ imports.NewTable("TCoreWebView2ClientCertificate_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2ClientCertificate_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2ClientCertificate_Subject", 0), // property Subject
			/* 5 */ imports.NewTable("TCoreWebView2ClientCertificate_Issuer", 0), // property Issuer
			/* 6 */ imports.NewTable("TCoreWebView2ClientCertificate_ValidFrom", 0), // property ValidFrom
			/* 7 */ imports.NewTable("TCoreWebView2ClientCertificate_ValidTo", 0), // property ValidTo
			/* 8 */ imports.NewTable("TCoreWebView2ClientCertificate_DerEncodedSerialNumber", 0), // property DerEncodedSerialNumber
			/* 9 */ imports.NewTable("TCoreWebView2ClientCertificate_DisplayName", 0), // property DisplayName
			/* 10 */ imports.NewTable("TCoreWebView2ClientCertificate_PemEncodedIssuerCertificateChain", 0), // property PemEncodedIssuerCertificateChain
			/* 11 */ imports.NewTable("TCoreWebView2ClientCertificate_Kind", 0), // property Kind
		}
	})
	return coreWebView2ClientCertificateImport
}
