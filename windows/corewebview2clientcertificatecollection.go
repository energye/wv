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

// ICoreWebView2ClientCertificateCollection Parent: IObject
type ICoreWebView2ClientCertificateCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ClientCertificateCollection // property BaseIntf Getter
	// Count
	//  The number of elements contained in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificatecollection#get_count">See the ICoreWebView2ClientCertificateCollection article.</see>
	Count() uint32 // property Count Getter
	// Items
	//  Gets the element at the given index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificatecollection#getvalueatindex">See the ICoreWebView2ClientCertificateCollection article.</see>
	Items(idx uint32) ICoreWebView2ClientCertificate // property Items Getter
}

type TCoreWebView2ClientCertificateCollection struct {
	TObject
}

func (m *TCoreWebView2ClientCertificateCollection) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ClientCertificateCollectionAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ClientCertificateCollection) BaseIntf() (result ICoreWebView2ClientCertificateCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateCollectionAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ClientCertificateCollection(resultPtr)
	return
}

func (m *TCoreWebView2ClientCertificateCollection) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ClientCertificateCollectionAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2ClientCertificateCollection) Items(idx uint32) (result ICoreWebView2ClientCertificate) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ClientCertificateCollectionAPI().SysCallN(4, m.Instance(), uintptr(idx), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ClientCertificate(resultPtr)
	return
}

// NewCoreWebView2ClientCertificateCollection class constructor
func NewCoreWebView2ClientCertificateCollection(baseIntf ICoreWebView2ClientCertificateCollection) ICoreWebView2ClientCertificateCollection {
	r := coreWebView2ClientCertificateCollectionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ClientCertificateCollection(r)
}

var (
	coreWebView2ClientCertificateCollectionOnce   base.Once
	coreWebView2ClientCertificateCollectionImport *imports.Imports = nil
)

func coreWebView2ClientCertificateCollectionAPI() *imports.Imports {
	coreWebView2ClientCertificateCollectionOnce.Do(func() {
		coreWebView2ClientCertificateCollectionImport = api.NewDefaultImports()
		coreWebView2ClientCertificateCollectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ClientCertificateCollection_Create", 0), // constructor NewCoreWebView2ClientCertificateCollection
			/* 1 */ imports.NewTable("TCoreWebView2ClientCertificateCollection_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ClientCertificateCollection_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ClientCertificateCollection_Count", 0), // property Count
			/* 4 */ imports.NewTable("TCoreWebView2ClientCertificateCollection_Items", 0), // property Items
		}
	})
	return coreWebView2ClientCertificateCollectionImport
}
