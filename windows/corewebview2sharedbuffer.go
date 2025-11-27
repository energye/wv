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
	"github.com/energye/lcl/types"
)

// ICoreWebView2SharedBuffer Parent: lcl.IObject
type ICoreWebView2SharedBuffer interface {
	lcl.IObject
	// Close
	//  Release the backing shared memory. The application should call this API when no
	//  access to the buffer is needed any more, to ensure that the underlying resources
	//  are released timely even if the shared buffer object itself is not released due to
	//  some leaked reference.
	//  After the shared buffer is closed, the buffer address and file mapping handle previously
	//  obtained becomes invalid and cannot be used anymore. Accessing properties of the object
	//  will fail with `RO_E_CLOSED`. Operations like Read or Write on the IStream objects returned
	//  from `OpenStream` will fail with `RO_E_CLOSED`. `PostSharedBufferToScript` will also
	//  fail with `RO_E_CLOSED`.
	//  The script code should call `chrome.webview.releaseBuffer` with
	//  the shared buffer as the parameter to release underlying resources as soon
	//  as it does not need access the shared buffer any more.
	//  When script tries to access the buffer after calling `chrome.webview.releaseBuffer`,
	//  JavaScript `TypeError` exception will be raised complaining about accessing a
	//  detached ArrayBuffer, the same exception when trying to access a transferred ArrayBuffer.
	//  Closing the buffer object on native side doesn't impact access from Script and releasing
	//  the buffer from script doesn't impact access to the buffer from native side.
	//  The underlying shared memory will be released by the OS when both native and script side
	//  release the buffer.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#close">See the ICoreWebView2SharedBuffer article.</see>
	Close() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2SharedBuffer // property BaseIntf Getter
	// Size
	//  The size of the shared buffer in bytes.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#get_size">See the ICoreWebView2SharedBuffer article.</see>
	Size() int64 // property Size Getter
	// Buffer
	//  The memory address of the shared buffer.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#get_buffer">See the ICoreWebView2SharedBuffer article.</see>
	Buffer() types.PByte // property Buffer Getter
	// OpenStream
	//  Get an IStream object that can be used to access the shared buffer.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#openstream">See the ICoreWebView2SharedBuffer article.</see>
	OpenStream() lcl.IStreamAdapter // property OpenStream Getter
	// FileMappingHandle
	//  Returns a handle to the file mapping object that backs this shared buffer.
	//  The returned handle is owned by the shared buffer object. You should not
	//  call CloseHandle on it.
	//  Normal app should use `Buffer` or `OpenStream` to get memory address
	//  or IStream object to access the buffer.
	//  For advanced scenarios, you could use file mapping APIs to obtain other views
	//  or duplicate this handle to another application process and create a view from
	//  the duplicated handle in that process to access the buffer from that separate process.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#get_filemappinghandle">See the ICoreWebView2SharedBuffer article.</see>
	FileMappingHandle() types.HANDLE // property FileMappingHandle Getter
}

type TCoreWebView2SharedBuffer struct {
	lcl.TObject
}

func (m *TCoreWebView2SharedBuffer) Close() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SharedBufferAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SharedBuffer) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SharedBufferAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SharedBuffer) BaseIntf() (result ICoreWebView2SharedBuffer) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2SharedBufferAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2SharedBuffer(resultPtr)
	return
}

func (m *TCoreWebView2SharedBuffer) Size() (result int64) {
	if !m.IsValid() {
		return
	}
	coreWebView2SharedBufferAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2SharedBuffer) Buffer() types.PByte {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2SharedBufferAPI().SysCallN(5, m.Instance())
	return types.PByte(r)
}

func (m *TCoreWebView2SharedBuffer) OpenStream() lcl.IStreamAdapter {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	coreWebView2SharedBufferAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return lcl.AsStreamAdapter(resultPtr)
}

func (m *TCoreWebView2SharedBuffer) FileMappingHandle() types.HANDLE {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2SharedBufferAPI().SysCallN(7, m.Instance())
	return types.HANDLE(r)
}

// NewCoreWebView2SharedBuffer class constructor
func NewCoreWebView2SharedBuffer(baseIntf ICoreWebView2SharedBuffer) ICoreWebView2SharedBuffer {
	r := coreWebView2SharedBufferAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2SharedBuffer(r)
}

var (
	coreWebView2SharedBufferOnce   base.Once
	coreWebView2SharedBufferImport *imports.Imports = nil
)

func coreWebView2SharedBufferAPI() *imports.Imports {
	coreWebView2SharedBufferOnce.Do(func() {
		coreWebView2SharedBufferImport = api.NewDefaultImports()
		coreWebView2SharedBufferImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2SharedBuffer_Create", 0), // constructor NewCoreWebView2SharedBuffer
			/* 1 */ imports.NewTable("TCoreWebView2SharedBuffer_Close", 0), // function Close
			/* 2 */ imports.NewTable("TCoreWebView2SharedBuffer_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2SharedBuffer_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2SharedBuffer_Size", 0), // property Size
			/* 5 */ imports.NewTable("TCoreWebView2SharedBuffer_Buffer", 0), // property Buffer
			/* 6 */ imports.NewTable("TCoreWebView2SharedBuffer_OpenStream", 0), // property OpenStream
			/* 7 */ imports.NewTable("TCoreWebView2SharedBuffer_FileMappingHandle", 0), // property FileMappingHandle
		}
	})
	return coreWebView2SharedBufferImport
}
