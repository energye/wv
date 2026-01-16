//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/darwin"
)

type TWKCreateWebView func(sender lcl.IObject, configuration wvTypes.WKWebViewConfiguration, navigationAction wvTypes.WKNavigationAction, windowFeatures wvTypes.WKWindowFeatures) wvTypes.WKWebView
type TWKDecidePolicyForNavigationActionPreferences func(sender lcl.IObject, navigationAction wvTypes.WKNavigationAction, actionPolicy *wvTypes.WKNavigationActionPolicy, preferences *wvTypes.WKWebpagePreferences)
type TWKDecidePolicyForNavigationResponse func(sender lcl.IObject, navigationResponse wvTypes.WKNavigationResponse, responsePolicy *wvTypes.WKNavigationResponsePolicy)
type TWKDownloadCancelCompletionHandler func(sender lcl.IObject, download wvTypes.WKDownload, data uintptr, length int32)
type TWKDownloadDecideDestinationUsingResponseSuggestedFilename func(sender lcl.IObject, download wvTypes.WKDownload, response NSURLResponse, suggestedFilename string)
type TWKDownloadFailWithError func(sender lcl.IObject, download wvTypes.WKDownload, error_ string, data uintptr, length int32)
type TWKDownloadFinish func(sender lcl.IObject, download wvTypes.WKDownload)
type TWKDownloadWillPerformHTTPRedirectionNewRequest func(sender lcl.IObject, download wvTypes.WKDownload, response NSHTTPURLResponse, request NSURLRequest)
type TWKEvaluateJavaScriptCallback func(data uintptr, error_ string)
type TWKRunJavaScriptAlert func(sender lcl.IObject, message string, frame wvTypes.WKFrameInfo)
type TWKRunJavaScriptConfirmCompletion func(sender lcl.IObject, message string, frame wvTypes.WKFrameInfo) bool
type TWKRunJavaScriptTextInputCompletion func(sender lcl.IObject, prompt string, defaultText string, frame wvTypes.WKFrameInfo) string
type TWKStartURLSchemeTask func(sender lcl.IObject, urlSchemeTask wvTypes.WKURLSchemeTask)
type TWKStopURLSchemeTask func(sender lcl.IObject, urlSchemeTask wvTypes.WKURLSchemeTask)
type TWKWebViewDidClose func(sender lcl.IObject)
type TWkCommitNavigation func(sender lcl.IObject, navigation wvTypes.WKNavigation)
type TWkFailNavigationWithError func(sender lcl.IObject, navigation wvTypes.WKNavigation, error_ string)
type TWkFailProvisionalNavigationWithError func(sender lcl.IObject, navigation wvTypes.WKNavigation, error_ string)
type TWkFinishNavigation func(sender lcl.IObject, navigation wvTypes.WKNavigation)
type TWkMouseEvent func(sender lcl.IObject, event TWKButtonEvent)
type TWkNavigationActionDidBecomeDownload func(sender lcl.IObject, navigationAction wvTypes.WKNavigationAction, download wvTypes.WKDownload)
type TWkNavigationResponseDidBecomeDownload func(sender lcl.IObject, navigationResponse wvTypes.WKNavigationResponse, download wvTypes.WKDownload)
type TWkProcessMessageEvent func(sender lcl.IObject, userContentController wvTypes.WKUserContentController, name string, data string)
type TWkReceiveServerRedirectForProvisionalNavigation func(sender lcl.IObject, navigation wvTypes.WKNavigation)
type TWkStartProvisionalNavigation func(sender lcl.IObject, navigation wvTypes.WKNavigation)
type TWkWebContentProcessDidTerminate func(sender lcl.IObject)
