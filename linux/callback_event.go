//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
	wvTypes "github.com/energye/wv/types/linux"
)

type callback struct {
	name string
	cb   func(getVal func(i int) uintptr)
}

func getPtr(val uintptr) base.UnsafePointer {
	return base.UnsafePointer(val)
}

func makeTWkAddCookieFinishEvent(cb TWkAddCookieFinishEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkAddCookieFinishEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Result: boolean; error: string);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			error_ := api.GoStr(getVal(2))
			cb(sender, result, error_)
		},
	}
}

func makeTWkContextMenuCommandEvent(cb TWkContextMenuCommandEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkContextMenuCommandEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; menuID: integer);
			sender := lcl.AsObject(getVal(0))
			menuID := int32(getVal(1))
			cb(sender, menuID)
		},
	}
}

func makeTWkContextMenuDismissedEvent(cb TWkContextMenuDismissedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkContextMenuDismissedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(Sender: TObject);
			sender := lcl.AsObject(getVal(0))
			cb(sender)
		},
	}
}

func makeTWkContextMenuEvent(cb TWkContextMenuEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkContextMenuEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(Sender: TObject; contextMenu: WebKitContextMenu; defaultAction: PWkAction): boolean;
			sender := lcl.AsObject(getVal(0))
			contextMenu := wvTypes.WebKitContextMenu(getVal(1))
			defaultAction := wvTypes.PWkAction(getVal(2))
			ret := cb(sender, contextMenu, defaultAction)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTWkDecidePolicyEvent(cb TWkDecidePolicyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkDecidePolicyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(Sender: TObject; decision: WebKitPolicyDecision; type_: WebKitPolicyDecisionType): boolean;
			sender := lcl.AsObject(getVal(0))
			decision := wvTypes.WebKitPolicyDecision(getVal(1))
			type_ := wvTypes.WebKitPolicyDecisionType(getVal(2))
			ret := cb(sender, decision, type_)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTWkDeleteCookieFinishEvent(cb TWkDeleteCookieFinishEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkDeleteCookieFinishEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Result: boolean; error: string);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			error_ := api.GoStr(getVal(2))
			cb(sender, result, error_)
		},
	}
}

func makeTWkExecuteScriptFinishedEvent(cb TWkExecuteScriptFinishedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkExecuteScriptFinishedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const jsValue: TWkJSValue);
			sender := lcl.AsObject(getVal(0))
			jsValue := AsWkJSValue(getVal(1))
			cb(sender, jsValue)
		},
	}
}

func makeTWkGetAcceptPolicyFinishEvent(cb TWkGetAcceptPolicyFinishEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkGetAcceptPolicyFinishEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; policy: WebKitCookieAcceptPolicy; error: string);
			sender := lcl.AsObject(getVal(0))
			policy := wvTypes.WebKitCookieAcceptPolicy(getVal(1))
			error_ := api.GoStr(getVal(2))
			cb(sender, policy, error_)
		},
	}
}

func makeTWkGetCookiesFinishEvent(cb TWkGetCookiesFinishEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkGetCookiesFinishEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; cookieList: PList; error: string);
			sender := lcl.AsObject(getVal(0))
			cookieList := wvTypes.PList(getVal(1))
			error_ := api.GoStr(getVal(2))
			cb(sender, cookieList, error_)
		},
	}
}

func makeTWkLoadChangeEvent(cb TWkLoadChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkLoadChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; loadEvent: WebKitLoadEvent);
			sender := lcl.AsObject(getVal(0))
			loadEvent := wvTypes.WebKitLoadEvent(getVal(1))
			cb(sender, loadEvent)
		},
	}
}

func makeTWkLoadFailedEvent(cb TWkLoadFailedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkLoadFailedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(Sender: TObject; loadEvent: WebKitLoadEvent; failingUri: string; error: string): boolean;
			sender := lcl.AsObject(getVal(0))
			loadEvent := wvTypes.WebKitLoadEvent(getVal(1))
			failingUri := api.GoStr(getVal(2))
			error_ := api.GoStr(getVal(3))
			ret := cb(sender, loadEvent, failingUri, error_)
			*(*bool)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTWkMouseEvent(cb TWkMouseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkMouseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(Sender: TObject; event: PWkButtonEvent): boolean;
			sender := lcl.AsObject(getVal(0))
			event := *(*TWkButtonEvent)(getPtr(getVal(1)))
			ret := cb(sender, event)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTWkProcessMessageEvent(cb TWkProcessMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkProcessMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const jsValue: TWkJSValue; processId: TWkProcessId);
			sender := lcl.AsObject(getVal(0))
			jsValue := AsWkJSValue(getVal(1))
			processId := wvTypes.TWkProcessId(getVal(2))
			cb(sender, jsValue, processId)
		},
	}
}

func makeTWkURISchemeRequestEvent(cb TWkURISchemeRequestEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkURISchemeRequestEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const uriSchemeRequest: WebKitURISchemeRequest);
			sender := lcl.AsObject(getVal(0))
			uriSchemeRequest := wvTypes.WebKitURISchemeRequest(getVal(1))
			cb(sender, uriSchemeRequest)
		},
	}
}

func makeTWkWebProcessTerminatedEvent(cb TWkWebProcessTerminatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkWebProcessTerminatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; reason: WebKitWebProcessTerminationReason);
			sender := lcl.AsObject(getVal(0))
			reason := wvTypes.WebKitWebProcessTerminationReason(getVal(1))
			cb(sender, reason)
		},
	}
}
