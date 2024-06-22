//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	"github.com/energye/lcl/api"
)

func (m *TWVLoader) SetOnEnvironmentCreated(fn TOnLoaderNotifyEvent) {
	predefImportAPI().SysCallN(0, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnInitializationError(fn TOnLoaderNotifyEvent) {
	predefImportAPI().SysCallN(1, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnGetCustomSchemes(fn TOnLoaderGetCustomSchemesEvent) {
	predefImportAPI().SysCallN(2, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnNewBrowserVersionAvailable(fn TOnLoaderNewBrowserVersionAvailableEvent) {
	predefImportAPI().SysCallN(3, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnBrowserProcessExited(fn TOnLoaderBrowserProcessExitedEvent) {
	predefImportAPI().SysCallN(4, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnProcessInfosChanged(fn TOnLoaderProcessInfosChangedEvent) {
	predefImportAPI().SysCallN(5, api.MakeEventDataPtr(fn))
}
