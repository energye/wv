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
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2NonClientRegionChangedEventArgs Parent: lcl.IObject
type ICoreWebView2NonClientRegionChangedEventArgs interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2NonClientRegionChangedEventArgs // property BaseIntf Getter
	// RegionKind
	//  This property represents the COREWEBVIEW2_NON_CLIENT_REGION_KIND which the
	//  region changed event corresponds to. With this property an app can query
	//  for a collection of rects which have that region kind by using
	//  QueryNonClientRegion on the composition controller.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2nonclientregionchangedeventargs#get_regionkind">See the ICoreWebView2NonClientRegionChangedEventArgs article.</see>
	RegionKind() wvTypes.TWVNonClientRegionKind // property RegionKind Getter
}

type TCoreWebView2NonClientRegionChangedEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2NonClientRegionChangedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NonClientRegionChangedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NonClientRegionChangedEventArgs) BaseIntf() (result ICoreWebView2NonClientRegionChangedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NonClientRegionChangedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2NonClientRegionChangedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2NonClientRegionChangedEventArgs) RegionKind() wvTypes.TWVNonClientRegionKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2NonClientRegionChangedEventArgsAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVNonClientRegionKind(r)
}

// NewCoreWebView2NonClientRegionChangedEventArgs class constructor
func NewCoreWebView2NonClientRegionChangedEventArgs(args ICoreWebView2NonClientRegionChangedEventArgs) ICoreWebView2NonClientRegionChangedEventArgs {
	r := coreWebView2NonClientRegionChangedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2NonClientRegionChangedEventArgs(r)
}

var (
	coreWebView2NonClientRegionChangedEventArgsOnce   base.Once
	coreWebView2NonClientRegionChangedEventArgsImport *imports.Imports = nil
)

func coreWebView2NonClientRegionChangedEventArgsAPI() *imports.Imports {
	coreWebView2NonClientRegionChangedEventArgsOnce.Do(func() {
		coreWebView2NonClientRegionChangedEventArgsImport = api.NewDefaultImports()
		coreWebView2NonClientRegionChangedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2NonClientRegionChangedEventArgs_Create", 0), // constructor NewCoreWebView2NonClientRegionChangedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2NonClientRegionChangedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2NonClientRegionChangedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2NonClientRegionChangedEventArgs_RegionKind", 0), // property RegionKind
		}
	})
	return coreWebView2NonClientRegionChangedEventArgsImport
}
