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

// ICoreWebView2ControllerOptions Parent: IObject
type ICoreWebView2ControllerOptions interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ControllerOptions         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2ControllerOptions) // property BaseIntf Setter
	// ProfileName
	//  `ProfileName` property is to specify a profile name, which is only allowed to contain
	//  the following ASCII characters. It has a maximum length of 64 characters excluding the null-terminator.
	//  It is ASCII case insensitive.
	//  * alphabet characters: a-z and A-Z
	//  * digit characters: 0-9
	//  * and '#', '@', '$', '(', ')', '+', '-', '_', '~', '.', ' ' (space).
	//  Note: the text must not end with a period '.' or ' ' (space). And, although upper-case letters are
	//  allowed, they're treated just as lower-case counterparts because the profile name will be mapped to
	//  the real profile directory path on disk and Windows file system handles path names in a case-insensitive way.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions#get_profilename">See the ICoreWebView2ControllerOptions article.</see>
	ProfileName() string         // property ProfileName Getter
	SetProfileName(value string) // property ProfileName Setter
	// IsInPrivateModeEnabled
	//  `IsInPrivateModeEnabled` property is to enable/disable InPrivate mode.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions#get_isinprivatemodeenabled">See the ICoreWebView2ControllerOptions article.</see>
	IsInPrivateModeEnabled() bool         // property IsInPrivateModeEnabled Getter
	SetIsInPrivateModeEnabled(value bool) // property IsInPrivateModeEnabled Setter
	// ScriptLocale
	//  The default locale for the WebView2. It sets the default locale for all
	//  Intl JavaScript APIs and other JavaScript APIs that depend on it, namely
	//  `Intl.DateTimeFormat()` which affects string formatting like
	//  in the time/date formats. Example: `Intl.DateTimeFormat().format(new Date())`
	//  The intended locale value is in the format of
	//  BCP 47 Language Tags. More information can be found from
	//  [IETF BCP47](https://www.ietf.org/rfc/bcp/bcp47.html).
	//  This property sets the locale for a CoreWebView2Environment used to create the
	//  WebView2ControllerOptions object, which is passed as a parameter in
	//  `CreateCoreWebView2ControllerWithOptions`.
	//  Changes to the ScriptLocale property apply to renderer processes created after
	//  the change. Any existing renderer processes will continue to use the previous
	//  ScriptLocale value. To ensure changes are applied to all renderer process,
	//  close and restart the CoreWebView2Environment and all associated WebView2 objects.
	//  The default value for ScriptLocale will depend on the WebView2 language
	//  and OS region. If the language portions of the WebView2 language and OS region
	//  match, then it will use the OS region. Otherwise, it will use the WebView2
	//  language.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2#get_scriptlocale">See the ICoreWebView2ControllerOptions2 article.</see>
	ScriptLocale() string         // property ScriptLocale Getter
	SetScriptLocale(value string) // property ScriptLocale Setter
}

type TCoreWebView2ControllerOptions struct {
	TObject
}

func (m *TCoreWebView2ControllerOptions) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerOptionsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ControllerOptions) BaseIntf() (result ICoreWebView2ControllerOptions) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ControllerOptionsAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ControllerOptions(resultPtr)
	return
}

func (m *TCoreWebView2ControllerOptions) SetBaseIntf(value ICoreWebView2ControllerOptions) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerOptionsAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2ControllerOptions) ProfileName() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ControllerOptionsAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ControllerOptions) SetProfileName(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerOptionsAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2ControllerOptions) IsInPrivateModeEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerOptionsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ControllerOptions) SetIsInPrivateModeEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerOptionsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2ControllerOptions) ScriptLocale() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ControllerOptionsAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ControllerOptions) SetScriptLocale(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerOptionsAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

// NewCoreWebView2ControllerOptions class constructor
func NewCoreWebView2ControllerOptions(baseIntf ICoreWebView2ControllerOptions) ICoreWebView2ControllerOptions {
	r := coreWebView2ControllerOptionsAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ControllerOptions(r)
}

var (
	coreWebView2ControllerOptionsOnce   base.Once
	coreWebView2ControllerOptionsImport *imports.Imports = nil
)

func coreWebView2ControllerOptionsAPI() *imports.Imports {
	coreWebView2ControllerOptionsOnce.Do(func() {
		coreWebView2ControllerOptionsImport = api.NewDefaultImports()
		coreWebView2ControllerOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ControllerOptions_Create", 0), // constructor NewCoreWebView2ControllerOptions
			/* 1 */ imports.NewTable("TCoreWebView2ControllerOptions_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ControllerOptions_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ControllerOptions_ProfileName", 0), // property ProfileName
			/* 4 */ imports.NewTable("TCoreWebView2ControllerOptions_IsInPrivateModeEnabled", 0), // property IsInPrivateModeEnabled
			/* 5 */ imports.NewTable("TCoreWebView2ControllerOptions_ScriptLocale", 0), // property ScriptLocale
		}
	})
	return coreWebView2ControllerOptionsImport
}
