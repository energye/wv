//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/linux"
)

type TWkAddCookieFinishEvent func(sender lcl.IObject, result bool, error_ string)
type TWkContextMenuCommandEvent func(sender lcl.IObject, menuID int32)
type TWkContextMenuDismissedEvent func(sender lcl.IObject)
type TWkContextMenuEvent func(sender lcl.IObject, contextMenu wvTypes.WebKitContextMenu, defaultAction wvTypes.PWkAction) bool
type TWkDecidePolicyEvent func(sender lcl.IObject, decision wvTypes.WebKitPolicyDecision, type_ wvTypes.WebKitPolicyDecisionType) bool
type TWkDeleteCookieFinishEvent func(sender lcl.IObject, result bool, error_ string)
type TWkExecuteScriptFinishedEvent func(sender lcl.IObject, jsValue IWkJSValue)
type TWkGetAcceptPolicyFinishEvent func(sender lcl.IObject, policy wvTypes.WebKitCookieAcceptPolicy, error_ string)
type TWkGetCookiesFinishEvent func(sender lcl.IObject, cookieList wvTypes.PList, error_ string)
type TWkLoadChangeEvent func(sender lcl.IObject, loadEvent wvTypes.WebKitLoadEvent)
type TWkLoadFailedEvent func(sender lcl.IObject, loadEvent wvTypes.WebKitLoadEvent, failingUri string, error_ string) bool
type TWkMouseEvent func(sender lcl.IObject, event TWkButtonEvent) bool
type TWkProcessMessageEvent func(sender lcl.IObject, jsValue IWkJSValue, processId wvTypes.TWkProcessId)
type TWkURISchemeRequestEvent func(sender lcl.IObject, uriSchemeRequest wvTypes.WebKitURISchemeRequest)
type TWkWebProcessTerminatedEvent func(sender lcl.IObject, reason wvTypes.WebKitWebProcessTerminationReason)
