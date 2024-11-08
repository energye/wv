//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

type IWKDownloadEvent interface {
}

type IWKNavigationDelegateEvent interface {
}

type IReceiveScriptMessageEvent interface {
}

type IWKUIDelegateEvent interface {
}

type IWKURLSchemeHandlerEvent interface {
}

// WKScriptMessageHandlerProtocol

type TWkProcessMessageEvent func(sender IObject, userContentController WKUserContentController, name, data string)

// WKNavigationDelegateProtocol

type TWKDecidePolicyForNavigationActionPreferences func(sender IObject, navigationAction WKNavigationAction, actionPolicy *WKNavigationActionPolicy, preferences *WKWebpagePreferences)
type TWKDecidePolicyForNavigationResponse func(sender IObject, navigationResponse WKNavigationResponse, responsePolicy *WKNavigationResponsePolicy)
type TWkStartProvisionalNavigation func(sender IObject, navigation WKNavigation)
type TWkReceiveServerRedirectForProvisionalNavigation func(sender IObject, navigation WKNavigation)
type TWkFailProvisionalNavigationWithError func(sender IObject, navigation WKNavigation, error string)
type TWkCommitNavigation func(sender IObject, navigation WKNavigation)
type TWkFinishNavigation func(sender IObject, navigation WKNavigation)
type TWkFailNavigationWithError func(sender IObject, navigation WKNavigation, error string)

// TWkReceiveAuthenticationChallenge func(sender IObject, challenge NSURLAuthenticationChallenge, var disposition NSURLSessionAuthChallengeDisposition, var credential INSURLCredential)

type TWkWebContentProcessDidTerminate func(sender IObject)

// TWkAuthenticationChallengeShouldAllowDeprecatedTLS func(sender IObject, challenge NSURLAuthenticationChallenge) Boolean

type TWkNavigationActionDidBecomeDownload func(sender IObject, navigationAction WKNavigationAction, download WKDownload)
type TWkNavigationResponseDidBecomeDownload func(sender IObject, navigationResponse WKNavigationResponse, download WKDownload)

// WKURLSchemeHandlerProtocol

type TWKStartURLSchemeTask func(sender IObject, urlSchemeTask WKURLSchemeTask)
type TWKStopURLSchemeTask func(sender IObject, urlSchemeTask WKURLSchemeTask)

// WKUIDelegateProtocol

type TWKCreateWebView func(sender IObject, configuration WKWebViewConfiguration, navigationAction WKNavigationAction, windowFeatures WKWindowFeatures) WKWebView
type TWKRunJavaScriptAlert func(sender IObject, message string, frame WKFrameInfo)
type TWKRunJavaScriptConfirmCompletion func(sender IObject, message string, frame WKFrameInfo) bool
type TWKRunJavaScriptTextInputCompletion func(sender IObject, prompt string, defaultText string, frame WKFrameInfo) string
type TWKWebViewDidClose func(sender IObject)

// WKDownload

type TWKDownloadCancelCompletionHandler func(sender IObject, download WKDownload, aData uintptr, aLength int32 /*NSData*/)
type TWKDownloadDecideDestinationUsingResponseSuggestedFilename func(sender IObject, download WKDownload, response NSURLResponse, suggestedFilename string)
type TWKDownloadWillPerformHTTPRedirectionNewRequest func(sender IObject, download WKDownload, response NSHTTPURLResponse, request NSURLRequest)

// TWKDownloadReceiveAuthenticationChallenge func(sender IObject, download IWKDownload, challenge NSURLAuthenticationChallenge)

type TWKDownloadFinish func(sender IObject, download WKDownload)
type TWKDownloadFailWithError func(sender IObject, download WKDownload, error string, aData uintptr, aLength int32 /*NSData*/)
