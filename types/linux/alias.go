//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import "github.com/energye/lcl/types"

type WebKitContextMenuItem uintptr
type WebKitContextMenu uintptr
type WebKitJavascriptResult uintptr
type WebKitNavigationAction uintptr
type WebKitPolicyDecision uintptr
type WebKitNavigationPolicyDecision = WebKitPolicyDecision
type WebKitResponsePolicyDecision = WebKitPolicyDecision
type WebKitSettings uintptr
type WebKitURIRequest uintptr
type WebKitURIResponse uintptr
type WebKitURISchemeRequest uintptr
type WebKitWebView uintptr
type WebKitURISchemeResponse uintptr
type WebKitWebsitePolicies uintptr
type WebKitWebContext uintptr
type WebKitCookieManager uintptr
type PWkAction uintptr
type PSoupDate uintptr
type PSoupMessageHeaders uintptr
type PInputStream uintptr
type PSoupCookie uintptr
type PList uintptr

type TScrolledWindow = uintptr

// TGdkModifierType SET
type TGdkModifierType = types.TSet

// JSCCheckSyntaxMode ENUM
type JSCCheckSyntaxMode = int32
