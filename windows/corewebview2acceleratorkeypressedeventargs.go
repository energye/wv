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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2AcceleratorKeyPressedEventArgs Parent: IObject
type ICoreWebView2AcceleratorKeyPressedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2AcceleratorKeyPressedEventArgs // property BaseIntf Getter
	// KeyEventKind
	//  The key event type that caused the event to run.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_keyeventkind">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	KeyEventKind() wvTypes.TWVKeyEventKind // property KeyEventKind Getter
	// VirtualKey
	//  The Win32 virtual key code of the key that was pressed or released. It
	//  is one of the Win32 virtual key constants such as `VK_RETURN` or an
	//  (uppercase) ASCII value such as `A`. Verify whether Ctrl or Alt
	//  are pressed by running `GetKeyState(VK_CONTROL)` or
	//  `GetKeyState(VK_MENU)`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_virtualkey">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	VirtualKey() uint32 // property VirtualKey Getter
	// KeyEventLParam
	//  The `LPARAM` value that accompanied the window message. For more
	//  information, navigate to [WM_KEYDOWN](/windows/win32/inputdev/wm-keydown)
	//  and [WM_KEYUP](/windows/win32/inputdev/wm-keyup).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_keyeventlparam">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	KeyEventLParam() int32 // property KeyEventLParam Getter
	// RepeatCount
	//  Specifies the repeat count for the current message.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</see>
	RepeatCount() uint32 // property RepeatCount Getter
	// ScanCode
	//  Specifies the scan code.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</see>
	ScanCode() uint32 // property ScanCode Getter
	// IsExtendedKey
	//  Indicates that the key is an extended key.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</see>
	IsExtendedKey() bool // property IsExtendedKey Getter
	// IsMenuKeyDown
	//  Indicates that a menu key is held down (context code).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</see>
	IsMenuKeyDown() bool // property IsMenuKeyDown Getter
	// WasKeyDown
	//  Indicates that the key was held down.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</see>
	WasKeyDown() bool // property WasKeyDown Getter
	// IsKeyReleased
	//  Indicates that the key was released.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</see>
	IsKeyReleased() bool // property IsKeyReleased Getter
	// Handled
	//  During `AcceleratorKeyPressedEvent` handler invocation the WebView is
	//  blocked waiting for the decision of if the accelerator is handled by the
	//  host (or not). If the `Handled` property is set to `TRUE` then this
	//  prevents the WebView from performing the default action for this
	//  accelerator key. Otherwise the WebView performs the default action for
	//  the accelerator key.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_handled">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</see>
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
	// IsBrowserAcceleratorKeyEnabled
	//  This property allows developers to enable or disable the browser from handling a specific
	//  browser accelerator key such as Ctrl+P or F3, etc.
	//  Browser accelerator keys are the keys/key combinations that access features specific to
	//  a web browser, including but not limited to:
	//  <code>
	//  - Ctrl-F and F3 for Find on Page
	//  - Ctrl-P for Print
	//  - Ctrl-R and F5 for Reload
	//  - Ctrl-Plus and Ctrl-Minus for zooming
	//  - Ctrl-Shift-C and F12 for DevTools
	//  - Special keys for browser functions, such as Back, Forward, and Search
	//  </code>
	//  This property does not disable accelerator keys related to movement and text editing,
	//  such as:
	//  <code>
	//  - Home, End, Page Up, and Page Down
	//  - Ctrl-X, Ctrl-C, Ctrl-V
	//  - Ctrl-A for Select All
	//  - Ctrl-Z for Undo
	//  </code>
	//  The `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` API is a convenient setting
	//  for developers to disable all the browser accelerator keys together, and sets the default
	//  value for the `IsBrowserAcceleratorKeyEnabled` property.
	//  By default, `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` is `TRUE` and
	//  `IsBrowserAcceleratorKeyEnabled` is `TRUE`.
	//  When developers change `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` setting to `FALSE`,
	//  this will change default value for `IsBrowserAcceleratorKeyEnabled` to `FALSE`.
	//  If developers want specific keys to be handled by the browser after changing the
	//  `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` setting to `FALSE`, they need to enable
	//  these keys by setting `IsBrowserAcceleratorKeyEnabled` to `TRUE`.
	//  This API will give the event arg higher priority over the
	//  `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` setting when we handle the keys.
	//  For browser accelerator keys, when an accelerator key is pressed, the propagation and
	//  processing order is:
	//  <code>
	//  1. A ICoreWebView2Controller.AcceleratorKeyPressed event is raised
	//  2. WebView2 browser feature accelerator key handling
	//  3. Web Content Handling: If the key combination isn't reserved for browser actions,
	//  the key event propagates to the web content, where JavaScript event listeners can
	//  capture and respond to it.
	//  </code>
	//  `ICoreWebView2AcceleratorKeyPressedEventArgs` has a `Handled` property, that developers
	//  can use to mark a key as handled. When the key is marked as handled anywhere along
	//  the path, the event propagation stops, and web content will not receive the key.
	//  With `IsBrowserAcceleratorKeyEnabled` property, if developers mark
	//  `IsBrowserAcceleratorKeyEnabled` as `FALSE`, the browser will skip the WebView2
	//  browser feature accelerator key handling process, but the event propagation
	//  continues, and web content will receive the key combination.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs2#get_isbrowseracceleratorkeyenabled">See the ICoreWebView2AcceleratorKeyPressedEventArgs2 article.</see>
	IsBrowserAcceleratorKeyEnabled() bool         // property IsBrowserAcceleratorKeyEnabled Getter
	SetIsBrowserAcceleratorKeyEnabled(value bool) // property IsBrowserAcceleratorKeyEnabled Setter
}

type TCoreWebView2AcceleratorKeyPressedEventArgs struct {
	TObject
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) BaseIntf() (result ICoreWebView2AcceleratorKeyPressedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2AcceleratorKeyPressedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) KeyEventKind() wvTypes.TWVKeyEventKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVKeyEventKind(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) VirtualKey() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(4, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) KeyEventLParam() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) RepeatCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(6, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) ScanCode() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(7, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) IsExtendedKey() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) IsMenuKeyDown() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) WasKeyDown() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) IsKeyReleased() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) IsBrowserAcceleratorKeyEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) SetIsBrowserAcceleratorKeyEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

// NewCoreWebView2AcceleratorKeyPressedEventArgs class constructor
func NewCoreWebView2AcceleratorKeyPressedEventArgs(args ICoreWebView2AcceleratorKeyPressedEventArgs) ICoreWebView2AcceleratorKeyPressedEventArgs {
	r := coreWebView2AcceleratorKeyPressedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2AcceleratorKeyPressedEventArgs(r)
}

var (
	coreWebView2AcceleratorKeyPressedEventArgsOnce   base.Once
	coreWebView2AcceleratorKeyPressedEventArgsImport *imports.Imports = nil
)

func coreWebView2AcceleratorKeyPressedEventArgsAPI() *imports.Imports {
	coreWebView2AcceleratorKeyPressedEventArgsOnce.Do(func() {
		coreWebView2AcceleratorKeyPressedEventArgsImport = api.NewDefaultImports()
		coreWebView2AcceleratorKeyPressedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_Create", 0), // constructor NewCoreWebView2AcceleratorKeyPressedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_KeyEventKind", 0), // property KeyEventKind
			/* 4 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_VirtualKey", 0), // property VirtualKey
			/* 5 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_KeyEventLParam", 0), // property KeyEventLParam
			/* 6 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_RepeatCount", 0), // property RepeatCount
			/* 7 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_ScanCode", 0), // property ScanCode
			/* 8 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_IsExtendedKey", 0), // property IsExtendedKey
			/* 9 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_IsMenuKeyDown", 0), // property IsMenuKeyDown
			/* 10 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_WasKeyDown", 0), // property WasKeyDown
			/* 11 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_IsKeyReleased", 0), // property IsKeyReleased
			/* 12 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_Handled", 0), // property Handled
			/* 13 */ imports.NewTable("TCoreWebView2AcceleratorKeyPressedEventArgs_IsBrowserAcceleratorKeyEnabled", 0), // property IsBrowserAcceleratorKeyEnabled
		}
	})
	return coreWebView2AcceleratorKeyPressedEventArgsImport
}
