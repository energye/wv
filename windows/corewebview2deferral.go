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

// ICoreWebView2Deferral Parent: IObject
type ICoreWebView2Deferral interface {
	IObject
	// Complete
	//  Completes the associated deferred event. Complete should only be run
	//  once for each deferral taken.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2deferral#complete">See the ICoreWebView2Deferral article.</see>
	Complete() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Deferral // property BaseIntf Getter
}

type TCoreWebView2Deferral struct {
	TObject
}

func (m *TCoreWebView2Deferral) Complete() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DeferralAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Deferral) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DeferralAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Deferral) BaseIntf() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2DeferralAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2Deferral class constructor
func NewCoreWebView2Deferral(baseIntf ICoreWebView2Deferral) ICoreWebView2Deferral {
	r := coreWebView2DeferralAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Deferral(r)
}

var (
	coreWebView2DeferralOnce   base.Once
	coreWebView2DeferralImport *imports.Imports = nil
)

func coreWebView2DeferralAPI() *imports.Imports {
	coreWebView2DeferralOnce.Do(func() {
		coreWebView2DeferralImport = api.NewDefaultImports()
		coreWebView2DeferralImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Deferral_Create", 0), // constructor NewCoreWebView2Deferral
			/* 1 */ imports.NewTable("TCoreWebView2Deferral_Complete", 0), // function Complete
			/* 2 */ imports.NewTable("TCoreWebView2Deferral_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2Deferral_BaseIntf", 0), // property BaseIntf
		}
	})
	return coreWebView2DeferralImport
}
