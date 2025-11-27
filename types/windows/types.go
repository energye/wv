//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package windows

// TWV2DebugLog ENUM
//
//	Debug log values used by TWVLoader.DebugLog
type TWV2DebugLog = int32

const (
	DlDisabled TWV2DebugLog = iota
	DlEnabled
	DlEnabledStdOut
	DlEnabledStdErr
)

// TWV2DebugLogLevel ENUM
//
//	Debug log level used when the logging is enabled
type TWV2DebugLogLevel = int32

const (
	DllDefault TWV2DebugLogLevel = iota
	DllInfo
	DllWarning
	DllError
	DllFatal
)

// TWV2EditingCommand ENUM
//
//	Blink editing commands used by the "Input.dispatchKeyEvent" DevTools method.
//	<see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
//	<see href="https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h">See the Chromium sources.</see>
type TWV2EditingCommand = int32

const (
	EcAlignCenter TWV2EditingCommand = iota
	EcAlignJustified
	EcAlignLeft
	EcAlignRight
	EcBackColor
	EcBackwardDelete
	EcBold
	EcCopy
	EcCreateLink
	EcCut
	EcDefaultParagraphSeparator
	EcDelete
	EcDeleteBackward
	EcDeleteBackwardByDecomposingPreviousCharacter
	EcDeleteForward
	EcDeleteToBeginningOfLine
	EcDeleteToBeginningOfParagraph
	EcDeleteToEndOfLine
	EcDeleteToEndOfParagraph
	EcDeleteToMark
	EcDeleteWordBackward
	EcDeleteWordForward
	EcFindString
	EcFontName
	EcFontSize
	EcFontSizeDelta
	EcForeColor
	EcFormatBlock
	EcForwardDelete
	EcHiliteColor
	EcIgnoreSpelling
	EcIndent
	EcInsertBacktab
	EcInsertHorizontalRule
	EcInsertHTML
	EcInsertImage
	EcInsertLineBreak
	EcInsertNewline
	EcInsertNewlineInQuotedContent
	EcInsertOrderedList
	EcInsertParagraph
	EcInsertTab
	EcInsertText
	EcInsertUnorderedList
	EcItalic
	EcJustifyCenter
	EcJustifyFull
	EcJustifyLeft
	EcJustifyNone
	EcJustifyRight
	EcMakeTextWritingDirectionLeftToRight
	EcMakeTextWritingDirectionNatural
	EcMakeTextWritingDirectionRightToLeft
	EcMoveBackward
	EcMoveBackwardAndModifySelection
	EcMoveDown
	EcMoveDownAndModifySelection
	EcMoveForward
	EcMoveForwardAndModifySelection
	EcMoveLeft
	EcMoveLeftAndModifySelection
	EcMovePageDown
	EcMovePageDownAndModifySelection
	EcMovePageUp
	EcMovePageUpAndModifySelection
	EcMoveParagraphBackward
	EcMoveParagraphBackwardAndModifySelection
	EcMoveParagraphForward
	EcMoveParagraphForwardAndModifySelection
	EcMoveRight
	EcMoveRightAndModifySelection
	EcMoveToBeginningOfDocument
	EcMoveToBeginningOfDocumentAndModifySelection
	EcMoveToBeginningOfLine
	EcMoveToBeginningOfLineAndModifySelection
	EcMoveToBeginningOfParagraph
	EcMoveToBeginningOfParagraphAndModifySelection
	EcMoveToBeginningOfSentence
	EcMoveToBeginningOfSentenceAndModifySelection
	EcMoveToEndOfDocument
	EcMoveToEndOfDocumentAndModifySelection
	EcMoveToEndOfLine
	EcMoveToEndOfLineAndModifySelection
	EcMoveToEndOfParagraph
	EcMoveToEndOfParagraphAndModifySelection
	EcMoveToEndOfSentence
	EcMoveToEndOfSentenceAndModifySelection
	EcMoveToLeftEndOfLine
	EcMoveToLeftEndOfLineAndModifySelection
	EcMoveToRightEndOfLine
	EcMoveToRightEndOfLineAndModifySelection
	EcMoveUp
	EcMoveUpAndModifySelection
	EcMoveWordBackward
	EcMoveWordBackwardAndModifySelection
	EcMoveWordForward
	EcMoveWordForwardAndModifySelection
	EcMoveWordLeft
	EcMoveWordLeftAndModifySelection
	EcMoveWordRight
	EcMoveWordRightAndModifySelection
	EcOutdent
	EcOverWrite
	EcPaste
	EcPasteAndMatchStyle
	EcPasteGlobalSelection
	EcPrint
	EcRedo
	EcRemoveFormat
	EcScrollLineDown
	EcScrollLineUp
	EcScrollPageBackward
	EcScrollPageForward
	EcScrollToBeginningOfDocument
	EcScrollToEndOfDocument
	EcSelectAll
	EcSelectLine
	EcSelectParagraph
	EcSelectSentence
	EcSelectToMark
	EcSelectWord
	EcSetMark
	EcStrikethrough
	EcStyleWithCSS
	EcSubscript
	EcSuperscript
	EcSwapWithMark
	EcToggleBold
	EcToggleItalic
	EcToggleUnderline
	EcTranspose
	EcUnderline
	EcUndo
	EcUnlink
	EcUnscript
	EcUnselect
	EcUseCSS
	EcYank
	EcYankAndSelect
)

// TWV2KeyEventType ENUM
//
//	Event type used by TWVBrowserBase.SimulateKeyEvent
type TWV2KeyEventType = int32

const (
	KetKeyDown TWV2KeyEventType = iota
	KetKeyUp
	KetRawKeyDown
	KetChar
)

// TWV2LoaderStatus ENUM
//
//	TWVLoader status values
type TWV2LoaderStatus = int32

const (
	WvlsCreated TWV2LoaderStatus = iota
	WvlsLoading
	WvlsLoaded
	WvlsImported
	WvlsInitialized
	WvlsError
	WvlsUnloaded
)

// TWVAutoplayPolicy ENUM
//
//	Autoplay policy types used by TWVLoader.AutoplayPolicy. See the --autoplay-policy switch.
type TWVAutoplayPolicy = int32

const (
	AppDefault TWVAutoplayPolicy = iota
	AppDocumentUserActivationRequired
	AppNoUserGestureRequired
	AppUserGestureRequired
)

// TWVClearDataStorageTypes ENUM
//
//	Used by TWVBrowserBase.ClearDataForOrigin to clear the storage
type TWVClearDataStorageTypes = int32

const (
	CdstAppCache TWVClearDataStorageTypes = iota
	CdstCookies
	CdstFileSystems
	CdstIndexeddb
	CdstLocalStorage
	CdstShaderCache
	CdstWebsql
	CdstServiceWorkers
	CdstCacheStorage
	CdstAll
)

// TWVState ENUM
//
//	Represents the state of a setting.
type TWVState = int32

const (
	STATE_DEFAULT TWVState = iota
	STATE_ENABLED
	STATE_DISABLED
)
