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

	wvTypes "github.com/energye/wv/types/windows"
)

// MiscFunc  is static instance
var MiscFunc _MiscFuncClass

// _MiscFuncClass is class type defined by MiscFunc
type _MiscFuncClass uintptr

func (_MiscFuncClass) LowestChromiumVersion() string {
	r := uWVMiscFunctionsAPI().SysCallN(0)
	return api.GoStr(r)
}

func (_MiscFuncClass) LowestLoaderDLLVersion() string {
	r := uWVMiscFunctionsAPI().SysCallN(1)
	return api.GoStr(r)
}

func (_MiscFuncClass) EnvironmentCreationErrorToString(errorCode types.HRESULT) string {
	r := uWVMiscFunctionsAPI().SysCallN(2, uintptr(errorCode))
	return api.GoStr(r)
}

func (_MiscFuncClass) ControllerCreationErrorToString(errorCode types.HRESULT) string {
	r := uWVMiscFunctionsAPI().SysCallN(3, uintptr(errorCode))
	return api.GoStr(r)
}

func (_MiscFuncClass) ControllerOptionsCreationErrorToString(errorCode types.HRESULT) string {
	r := uWVMiscFunctionsAPI().SysCallN(4, uintptr(errorCode))
	return api.GoStr(r)
}

func (_MiscFuncClass) CompositionControllerCreationErrorToString(errorCode types.HRESULT) string {
	r := uWVMiscFunctionsAPI().SysCallN(5, uintptr(errorCode))
	return api.GoStr(r)
}

func (_MiscFuncClass) GetScreenDPI() int32 {
	r := uWVMiscFunctionsAPI().SysCallN(6)
	return int32(r)
}

func (_MiscFuncClass) GetDeviceScaleFactor() (result float32) {
	uWVMiscFunctionsAPI().SysCallN(7, uintptr(base.UnsafePointer(&result)))
	return
}

func (_MiscFuncClass) EditingCommandToString(editingCommand wvTypes.TWV2EditingCommand) string {
	r := uWVMiscFunctionsAPI().SysCallN(8, uintptr(editingCommand))
	return api.GoStr(r)
}

func (_MiscFuncClass) DeleteDirContents(directory string, excludeFiles lcl.IStringList) bool {
	r := uWVMiscFunctionsAPI().SysCallN(9, api.PasStr(directory), base.GetObjectUintptr(excludeFiles))
	return api.GoBool(r)
}

func (_MiscFuncClass) SystemCursorIDToDelphiCursor(systemCursorID uint32) types.TCursor {
	r := uWVMiscFunctionsAPI().SysCallN(10, uintptr(systemCursorID))
	return types.TCursor(r)
}

func (_MiscFuncClass) TryIso8601BasicToDate(str string, outDate *types.TDateTime) bool {
	var datePtr uintptr
	r := uWVMiscFunctionsAPI().SysCallN(11, api.PasStr(str), uintptr(base.UnsafePointer(&datePtr)))
	*outDate = types.TDateTime(datePtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) JSONUnescape(source string) string {
	r := uWVMiscFunctionsAPI().SysCallN(12, api.PasStr(source))
	return api.GoStr(r)
}

func (_MiscFuncClass) JSONEscape(source string) string {
	r := uWVMiscFunctionsAPI().SysCallN(13, api.PasStr(source))
	return api.GoStr(r)
}

func (_MiscFuncClass) CustomURLDecode(encodedStr string) string {
	r := uWVMiscFunctionsAPI().SysCallN(14, api.PasStr(encodedStr))
	return api.GoStr(r)
}

func (_MiscFuncClass) CustomPathIsRelative(path string) bool {
	r := uWVMiscFunctionsAPI().SysCallN(15, api.PasStr(path))
	return api.GoBool(r)
}

func (_MiscFuncClass) CustomPathCanonicalize(originalPath string, canonicalPath *string) bool {
	canonicalPathPtr := api.PasStr(*canonicalPath)
	r := uWVMiscFunctionsAPI().SysCallN(16, api.PasStr(originalPath), uintptr(base.UnsafePointer(&canonicalPathPtr)))
	*canonicalPath = api.GoStr(canonicalPathPtr)
	return api.GoBool(r)
}

func (_MiscFuncClass) CustomAbsolutePath(path string, mustExist bool) string {
	r := uWVMiscFunctionsAPI().SysCallN(17, api.PasStr(path), api.PasBool(mustExist))
	return api.GoStr(r)
}

func (_MiscFuncClass) CustomPathIsURL(path string) bool {
	r := uWVMiscFunctionsAPI().SysCallN(18, api.PasStr(path))
	return api.GoBool(r)
}

func (_MiscFuncClass) CustomPathIsUNC(path string) bool {
	r := uWVMiscFunctionsAPI().SysCallN(19, api.PasStr(path))
	return api.GoBool(r)
}

func (_MiscFuncClass) GetModulePath() string {
	r := uWVMiscFunctionsAPI().SysCallN(20)
	return api.GoStr(r)
}

func (_MiscFuncClass) EscapeCommandLineParameter(parameter string) string {
	r := uWVMiscFunctionsAPI().SysCallN(21, api.PasStr(parameter))
	return api.GoStr(r)
}

func (_MiscFuncClass) OutputDebugMessage(message string) {
	uWVMiscFunctionsAPI().SysCallN(22, api.PasStr(message))
}

func (_MiscFuncClass) LogMouseEvent(eventKind wvTypes.TWVMouseEventKind, virtualKeys wvTypes.TWVMouseEventVirtualKeys, mouseData uint32, point types.TPoint) {
	uWVMiscFunctionsAPI().SysCallN(23, uintptr(eventKind), uintptr(virtualKeys), uintptr(mouseData), uintptr(base.UnsafePointer(&point)))
}

var (
	uWVMiscFunctionsOnce   base.Once
	uWVMiscFunctionsImport *imports.Imports = nil
)

func uWVMiscFunctionsAPI() *imports.Imports {
	uWVMiscFunctionsOnce.Do(func() {
		uWVMiscFunctionsImport = api.NewDefaultImports()
		uWVMiscFunctionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("uWVMiscFunctions_LowestChromiumVersion", 0), // static function LowestChromiumVersion
			/* 1 */ imports.NewTable("uWVMiscFunctions_LowestLoaderDLLVersion", 0), // static function LowestLoaderDLLVersion
			/* 2 */ imports.NewTable("uWVMiscFunctions_EnvironmentCreationErrorToString", 0), // static function EnvironmentCreationErrorToString
			/* 3 */ imports.NewTable("uWVMiscFunctions_ControllerCreationErrorToString", 0), // static function ControllerCreationErrorToString
			/* 4 */ imports.NewTable("uWVMiscFunctions_ControllerOptionsCreationErrorToString", 0), // static function ControllerOptionsCreationErrorToString
			/* 5 */ imports.NewTable("uWVMiscFunctions_CompositionControllerCreationErrorToString", 0), // static function CompositionControllerCreationErrorToString
			/* 6 */ imports.NewTable("uWVMiscFunctions_GetScreenDPI", 0), // static function GetScreenDPI
			/* 7 */ imports.NewTable("uWVMiscFunctions_GetDeviceScaleFactor", 0), // static function GetDeviceScaleFactor
			/* 8 */ imports.NewTable("uWVMiscFunctions_EditingCommandToString", 0), // static function EditingCommandToString
			/* 9 */ imports.NewTable("uWVMiscFunctions_DeleteDirContents", 0), // static function DeleteDirContents
			/* 10 */ imports.NewTable("uWVMiscFunctions_SystemCursorIDToDelphiCursor", 0), // static function SystemCursorIDToDelphiCursor
			/* 11 */ imports.NewTable("uWVMiscFunctions_TryIso8601BasicToDate", 0), // static function TryIso8601BasicToDate
			/* 12 */ imports.NewTable("uWVMiscFunctions_JSONUnescape", 0), // static function JSONUnescape
			/* 13 */ imports.NewTable("uWVMiscFunctions_JSONEscape", 0), // static function JSONEscape
			/* 14 */ imports.NewTable("uWVMiscFunctions_CustomURLDecode", 0), // static function CustomURLDecode
			/* 15 */ imports.NewTable("uWVMiscFunctions_CustomPathIsRelative", 0), // static function CustomPathIsRelative
			/* 16 */ imports.NewTable("uWVMiscFunctions_CustomPathCanonicalize", 0), // static function CustomPathCanonicalize
			/* 17 */ imports.NewTable("uWVMiscFunctions_CustomAbsolutePath", 0), // static function CustomAbsolutePath
			/* 18 */ imports.NewTable("uWVMiscFunctions_CustomPathIsURL", 0), // static function CustomPathIsURL
			/* 19 */ imports.NewTable("uWVMiscFunctions_CustomPathIsUNC", 0), // static function CustomPathIsUNC
			/* 20 */ imports.NewTable("uWVMiscFunctions_GetModulePath", 0), // static function GetModulePath
			/* 21 */ imports.NewTable("uWVMiscFunctions_EscapeCommandLineParameter", 0), // static function EscapeCommandLineParameter
			/* 22 */ imports.NewTable("uWVMiscFunctions_OutputDebugMessage", 0), // static procedure OutputDebugMessage
			/* 23 */ imports.NewTable("uWVMiscFunctions_LogMouseEvent", 0), // static procedure LogMouseEvent
		}
	})
	return uWVMiscFunctionsImport
}
