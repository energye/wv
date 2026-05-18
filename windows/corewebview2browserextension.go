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
)

// ICoreWebView2BrowserExtension Parent: IObject
type ICoreWebView2BrowserExtension interface {
	IObject
	// Remove
	//  Removes this browser extension from its WebView2 Profile. The browser extension is removed
	//  immediately including from all currently running HTML documents associated with this
	//  WebView2 Profile. The removal is persisted and future uses of this profile will not have this
	//  extension installed. After an extension is removed, calling `Remove` again will cause an exception.
	//  The TWVBrowserBase.OnBrowserExtensionRemoveCompleted event is triggered when this function finishes.
	Remove(browserComponent lcl.IComponent) bool // function
	// Enable
	//  Sets whether this browser extension is enabled or disabled. This change applies immediately
	//  to the extension in all HTML documents in all WebView2s associated with this profile.
	//  After an extension is removed, calling `Enable` will not change the value of `IsEnabled`.
	//  The TWVBrowserBase.OnBrowserExtensionEnableCompleted event is triggered when this function finishes.
	Enable(isEnabled bool, browserComponent lcl.IComponent) bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BrowserExtension         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2BrowserExtension) // property BaseIntf Setter
	// ID
	//  This is the browser extension's ID. This is the same browser extension ID returned by
	//  the browser extension API [`chrome.runtime.id`](https://developer.mozilla.org/docs/Mozilla/Add-ons/WebExtensions/API/runtime/id).
	//  Please see that documentation for more details on how the ID is generated.
	//  After an extension is removed, calling `Id` will return the id of the extension that is removed.
	ID() string // property ID Getter
	// Name
	//  This is the browser extension's name. This value is defined in this browser extension's
	//  manifest.json file. If manifest.json define extension's localized name, this value will
	//  be the localized version of the name.
	//  Please see [Manifest.json name](https://developer.mozilla.org/docs/Mozilla/Add-ons/WebExtensions/manifest.json/name)
	//  for more details.
	//  After an extension is removed, calling `Name` will return the name of the extension that is removed.
	Name() string // property Name Getter
	// IsEnabled
	//  If `isEnabled` is true then the Extension is enabled and running in WebView instances.
	//  If it is false then the Extension is disabled and not running in WebView instances.
	//  When a Extension is first installed, `IsEnable` are default to be `TRUE`.
	//  `isEnabled` is persisted per profile.
	//  After an extension is removed, calling `isEnabled` will return the value at the time it was removed.
	IsEnabled() bool // property IsEnabled Getter
}

type TCoreWebView2BrowserExtension struct {
	TObject
}

func (m *TCoreWebView2BrowserExtension) Remove(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BrowserExtensionAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2BrowserExtension) Enable(isEnabled bool, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BrowserExtensionAPI().SysCallN(2, m.Instance(), api.PasBool(isEnabled), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2BrowserExtension) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BrowserExtensionAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2BrowserExtension) BaseIntf() (result ICoreWebView2BrowserExtension) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2BrowserExtensionAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2BrowserExtension(resultPtr)
	return
}

func (m *TCoreWebView2BrowserExtension) SetBaseIntf(value ICoreWebView2BrowserExtension) {
	if !m.IsValid() {
		return
	}
	coreWebView2BrowserExtensionAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2BrowserExtension) ID() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2BrowserExtensionAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2BrowserExtension) Name() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2BrowserExtensionAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2BrowserExtension) IsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BrowserExtensionAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

// NewCoreWebView2BrowserExtension class constructor
func NewCoreWebView2BrowserExtension(baseIntf ICoreWebView2BrowserExtension) ICoreWebView2BrowserExtension {
	r := coreWebView2BrowserExtensionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2BrowserExtension(r)
}

var (
	coreWebView2BrowserExtensionOnce   base.Once
	coreWebView2BrowserExtensionImport *imports.Imports = nil
)

func coreWebView2BrowserExtensionAPI() *imports.Imports {
	coreWebView2BrowserExtensionOnce.Do(func() {
		coreWebView2BrowserExtensionImport = api.NewDefaultImports()
		coreWebView2BrowserExtensionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2BrowserExtension_Create", 0), // constructor NewCoreWebView2BrowserExtension
			/* 1 */ imports.NewTable("TCoreWebView2BrowserExtension_Remove", 0), // function Remove
			/* 2 */ imports.NewTable("TCoreWebView2BrowserExtension_Enable", 0), // function Enable
			/* 3 */ imports.NewTable("TCoreWebView2BrowserExtension_Initialized", 0), // property Initialized
			/* 4 */ imports.NewTable("TCoreWebView2BrowserExtension_BaseIntf", 0), // property BaseIntf
			/* 5 */ imports.NewTable("TCoreWebView2BrowserExtension_ID", 0), // property ID
			/* 6 */ imports.NewTable("TCoreWebView2BrowserExtension_Name", 0), // property Name
			/* 7 */ imports.NewTable("TCoreWebView2BrowserExtension_IsEnabled", 0), // property IsEnabled
		}
	})
	return coreWebView2BrowserExtensionImport
}
