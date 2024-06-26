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

// ICoreWebView2BrowserExtension Parent: IObject
//
//	Provides a set of properties for managing an Extension, which includes
//	an ID, name, and whether it is enabled or not, and the ability to Remove
//	the Extension, and enable or disable it.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextension">See the ICoreWebView2BrowserExtension article.</a>
type ICoreWebView2BrowserExtension interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BrowserExtension // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2BrowserExtension) // property
	// ID
	//  This is the browser extension's ID. This is the same browser extension ID returned by
	//  the browser extension API [`chrome.runtime.id`](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/runtime/id).
	//  Please see that documentation for more details on how the ID is generated.
	//  After an extension is removed, calling `Id` will return the id of the extension that is removed.
	ID() string // property
	// Name
	//  This is the browser extension's name. This value is defined in this browser extension's
	//  manifest.json file. If manifest.json define extension's localized name, this value will
	//  be the localized version of the name.
	//  Please see [Manifest.json name](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/manifest.json/name)
	//  for more details.
	//  After an extension is removed, calling `Name` will return the name of the extension that is removed.
	Name() string // property
	// IsEnabled
	//  If `isEnabled` is true then the Extension is enabled and running in WebView instances.
	//  If it is false then the Extension is disabled and not running in WebView instances.
	//  When a Extension is first installed, `IsEnable` are default to be `TRUE`.
	//  `isEnabled` is persisted per profile.
	//  After an extension is removed, calling `isEnabled` will return the value at the time it was removed.
	IsEnabled() bool // property
	// Remove
	//  Removes this browser extension from its WebView2 Profile. The browser extension is removed
	//  immediately including from all currently running HTML documents associated with this
	//  WebView2 Profile. The removal is persisted and future uses of this profile will not have this
	//  extension installed. After an extension is removed, calling `Remove` again will cause an exception.
	//  The TWVBrowserBase.OnBrowserExtensionRemoveCompleted event is triggered when this function finishes.
	Remove(aBrowserComponent IComponent) bool // function
	// Enable
	//  Sets whether this browser extension is enabled or disabled. This change applies immediately
	//  to the extension in all HTML documents in all WebView2s associated with this profile.
	//  After an extension is removed, calling `Enable` will not change the value of `IsEnabled`.
	//  The TWVBrowserBase.OnBrowserExtensionEnableCompleted event is triggered when this function finishes.
	Enable(aIsEnabled bool, aBrowserComponent IComponent) bool // function
}

// TCoreWebView2BrowserExtension Parent: TObject
//
//	Provides a set of properties for managing an Extension, which includes
//	an ID, name, and whether it is enabled or not, and the ability to Remove
//	the Extension, and enable or disable it.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextension">See the ICoreWebView2BrowserExtension article.</a>
type TCoreWebView2BrowserExtension struct {
	TObject
}

func NewCoreWebView2BrowserExtension(aBaseIntf ICoreWebView2BrowserExtension) ICoreWebView2BrowserExtension {
	r1 := coreWebView2BrowserExtensionImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2BrowserExtension(r1)
}

func (m *TCoreWebView2BrowserExtension) Initialized() bool {
	r1 := coreWebView2BrowserExtensionImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2BrowserExtension) BaseIntf() ICoreWebView2BrowserExtension {
	var resultCoreWebView2BrowserExtension uintptr
	coreWebView2BrowserExtensionImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2BrowserExtension)))
	return AsCoreWebView2BrowserExtension(resultCoreWebView2BrowserExtension)
}

func (m *TCoreWebView2BrowserExtension) SetBaseIntf(AValue ICoreWebView2BrowserExtension) {
	coreWebView2BrowserExtensionImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2BrowserExtension) ID() string {
	r1 := coreWebView2BrowserExtensionImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2BrowserExtension) Name() string {
	r1 := coreWebView2BrowserExtensionImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2BrowserExtension) IsEnabled() bool {
	r1 := coreWebView2BrowserExtensionImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2BrowserExtension) Remove(aBrowserComponent IComponent) bool {
	r1 := coreWebView2BrowserExtensionImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2BrowserExtension) Enable(aIsEnabled bool, aBrowserComponent IComponent) bool {
	r1 := coreWebView2BrowserExtensionImportAPI().SysCallN(2, m.Instance(), PascalBool(aIsEnabled), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

var (
	coreWebView2BrowserExtensionImport       *imports.Imports = nil
	coreWebView2BrowserExtensionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2BrowserExtension_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2BrowserExtension_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2BrowserExtension_Enable", 0),
		/*3*/ imports.NewTable("CoreWebView2BrowserExtension_ID", 0),
		/*4*/ imports.NewTable("CoreWebView2BrowserExtension_Initialized", 0),
		/*5*/ imports.NewTable("CoreWebView2BrowserExtension_IsEnabled", 0),
		/*6*/ imports.NewTable("CoreWebView2BrowserExtension_Name", 0),
		/*7*/ imports.NewTable("CoreWebView2BrowserExtension_Remove", 0),
	}
)

func coreWebView2BrowserExtensionImportAPI() *imports.Imports {
	if coreWebView2BrowserExtensionImport == nil {
		coreWebView2BrowserExtensionImport = NewDefaultImports()
		coreWebView2BrowserExtensionImport.SetImportTable(coreWebView2BrowserExtensionImportTables)
		coreWebView2BrowserExtensionImportTables = nil
	}
	return coreWebView2BrowserExtensionImport
}
