//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// ICoreWebView2PrintSettings Parent: IObject
//
//	Settings used by the PrintToPdf method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings">See the ICoreWebView2PrintSettings article.</a>
type ICoreWebView2PrintSettings interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PrintSettings // property
	// Orientation
	//  The orientation can be portrait or landscape. The default orientation is
	//  portrait. See `COREWEBVIEW2_PRINT_ORIENTATION`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_orientation">See the ICoreWebView2PrintSettings article.</a>
	Orientation() TWVPrintOrientation // property
	// SetOrientation Set Orientation
	SetOrientation(AValue TWVPrintOrientation) // property
	// ScaleFactor
	//  The scale factor is a value between 0.1 and 2.0. The default is 1.0.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_scalefactor">See the ICoreWebView2PrintSettings article.</a>
	ScaleFactor() (resultDouble float64) // property
	// SetScaleFactor Set ScaleFactor
	SetScaleFactor(AValue float64) // property
	// PageWidth
	//  The page width in inches. The default width is 8.5 inches.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_pagewidth">See the ICoreWebView2PrintSettings article.</a>
	PageWidth() (resultDouble float64) // property
	// SetPageWidth Set PageWidth
	SetPageWidth(AValue float64) // property
	// PageHeight
	//  The page height in inches. The default height is 11 inches.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_pageheight">See the ICoreWebView2PrintSettings article.</a>
	PageHeight() (resultDouble float64) // property
	// SetPageHeight Set PageHeight
	SetPageHeight(AValue float64) // property
	// MarginTop
	//  The top margin in inches. The default is 1 cm, or ~0.4 inches.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_margintop">See the ICoreWebView2PrintSettings article.</a>
	MarginTop() (resultDouble float64) // property
	// SetMarginTop Set MarginTop
	SetMarginTop(AValue float64) // property
	// MarginBottom
	//  The bottom margin in inches. The default is 1 cm, or ~0.4 inches.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_marginbottom">See the ICoreWebView2PrintSettings article.</a>
	MarginBottom() (resultDouble float64) // property
	// SetMarginBottom Set MarginBottom
	SetMarginBottom(AValue float64) // property
	// MarginLeft
	//  The left margin in inches. The default is 1 cm, or ~0.4 inches.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_marginleft">See the ICoreWebView2PrintSettings article.</a>
	MarginLeft() (resultDouble float64) // property
	// SetMarginLeft Set MarginLeft
	SetMarginLeft(AValue float64) // property
	// MarginRight
	//  The right margin in inches. The default is 1 cm, or ~0.4 inches.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_marginright">See the ICoreWebView2PrintSettings article.</a>
	MarginRight() (resultDouble float64) // property
	// SetMarginRight Set MarginRight
	SetMarginRight(AValue float64) // property
	// ShouldPrintBackgrounds
	//  `TRUE` if background colors and images should be printed. The default value
	//  is `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_shouldprintbackgrounds">See the ICoreWebView2PrintSettings article.</a>
	ShouldPrintBackgrounds() bool // property
	// SetShouldPrintBackgrounds Set ShouldPrintBackgrounds
	SetShouldPrintBackgrounds(AValue bool) // property
	// ShouldPrintSelectionOnly
	//  `TRUE` if only the current end user's selection of HTML in the document
	//  should be printed. The default value is `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_shouldprintselectiononly">See the ICoreWebView2PrintSettings article.</a>
	ShouldPrintSelectionOnly() bool // property
	// SetShouldPrintSelectionOnly Set ShouldPrintSelectionOnly
	SetShouldPrintSelectionOnly(AValue bool) // property
	// ShouldPrintHeaderAndFooter
	//  `TRUE` if header and footer should be printed. The default value is `FALSE`.
	//  The header consists of the date and time of printing, and the title of the
	//  page. The footer consists of the URI and page number. The height of the
	//  header and footer is 0.5 cm, or ~0.2 inches.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_shouldprintheaderandfooter">See the ICoreWebView2PrintSettings article.</a>
	ShouldPrintHeaderAndFooter() bool // property
	// SetShouldPrintHeaderAndFooter Set ShouldPrintHeaderAndFooter
	SetShouldPrintHeaderAndFooter(AValue bool) // property
	// HeaderTitle
	//  The title in the header if `ShouldPrintHeaderAndFooter` is `TRUE`. The
	//  default value is the title of the current document.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_headertitle">See the ICoreWebView2PrintSettings article.</a>
	HeaderTitle() string // property
	// SetHeaderTitle Set HeaderTitle
	SetHeaderTitle(AValue string) // property
	// FooterUri
	//  The URI in the footer if `ShouldPrintHeaderAndFooter` is `TRUE`. The
	//  default value is the current URI.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_footeruri">See the ICoreWebView2PrintSettings article.</a>
	FooterUri() string // property
	// SetFooterUri Set FooterUri
	SetFooterUri(AValue string) // property
	// PageRanges
	//  Page range to print.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_pageranges">See the ICoreWebView2PrintSettings2 article.</a>
	PageRanges() string // property
	// SetPageRanges Set PageRanges
	SetPageRanges(AValue string) // property
	// PagesPerSide
	//  Prints multiple pages of a document on a single piece of paper. Choose from 1, 2, 4, 6, 9 or 16.
	//  The default value is 1.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_pagesperside">See the ICoreWebView2PrintSettings2 article.</a>
	PagesPerSide() int32 // property
	// SetPagesPerSide Set PagesPerSide
	SetPagesPerSide(AValue int32) // property
	// Copies
	//  Number of copies to print. Minimum value is `1` and the maximum copies count is `999`.
	//  The default value is 1.
	//  This value is ignored in PrintToPdfStream method.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_copies">See the ICoreWebView2PrintSettings2 article.</a>
	Copies() int32 // property
	// SetCopies Set Copies
	SetCopies(AValue int32) // property
	// Collation
	//  Printer collation. See `COREWEBVIEW2_PRINT_COLLATION` for descriptions of
	//  collation. The default value is `COREWEBVIEW2_PRINT_COLLATION_DEFAULT`.
	//  Printing uses default value of printer's collation if an invalid value is provided
	//  for the specific printer. This value is ignored in PrintToPdfStream method.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_collation">See the ICoreWebView2PrintSettings2 article.</a>
	Collation() TWVPrintCollation // property
	// SetCollation Set Collation
	SetCollation(AValue TWVPrintCollation) // property
	// ColorMode
	//  Printer color mode. See `COREWEBVIEW2_PRINT_COLOR_MODE` for descriptions
	//  of color modes. The default value is `COREWEBVIEW2_PRINT_COLOR_MODE_DEFAULT`.
	//  Printing uses default value of printer supported color if an invalid value is provided
	//  for the specific printer.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_colormode">See the ICoreWebView2PrintSettings2 article.</a>
	ColorMode() TWVPrintColorMode // property
	// SetColorMode Set ColorMode
	SetColorMode(AValue TWVPrintColorMode) // property
	// Duplex
	//  Printer duplex settings. See `COREWEBVIEW2_PRINT_DUPLEX` for descriptions of duplex.
	//  The default value is `COREWEBVIEW2_PRINT_DUPLEX_DEFAULT`.
	//  Printing uses default value of printer's duplex if an invalid value is provided
	//  for the specific printer. This value is ignored in PrintToPdfStream method.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_duplex">See the ICoreWebView2PrintSettings2 article.</a>
	Duplex() TWVPrintDuplex // property
	// SetDuplex Set Duplex
	SetDuplex(AValue TWVPrintDuplex) // property
	// MediaSize
	//  Printer media size. See `COREWEBVIEW2_PRINT_MEDIA_SIZE` for descriptions of media size.
	//  The default value is `COREWEBVIEW2_PRINT_MEDIA_SIZE_DEFAULT`.
	//  If media size is `COREWEBVIEW2_PRINT_MEDIA_SIZE_CUSTOM`, you should set the `PageWidth`
	//  and `PageHeight`.
	//  Printing uses default value of printer supported media size if an invalid value is provided
	//  for the specific printer. This value is ignored in PrintToPdfStream method.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_mediasize">See the ICoreWebView2PrintSettings2 article.</a>
	MediaSize() TWVPrintMediaSize // property
	// SetMediaSize Set MediaSize
	SetMediaSize(AValue TWVPrintMediaSize) // property
	// PrinterName
	//  The name of the printer to use. Defaults to empty string.
	//  If the printer name is empty string or null, then it prints to the default
	//  printer on the user OS. This value is ignored in PrintToPdfStream method.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_printername">See the ICoreWebView2PrintSettings2 article.</a>
	PrinterName() string // property
	// SetPrinterName Set PrinterName
	SetPrinterName(AValue string) // property
}

// TCoreWebView2PrintSettings Parent: TObject
//
//	Settings used by the PrintToPdf method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings">See the ICoreWebView2PrintSettings article.</a>
type TCoreWebView2PrintSettings struct {
	TObject
}

func NewCoreWebView2PrintSettings(aBaseIntf ICoreWebView2PrintSettings) ICoreWebView2PrintSettings {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(4, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2PrintSettings(r1)
}

func (m *TCoreWebView2PrintSettings) Initialized() bool {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2PrintSettings) BaseIntf() ICoreWebView2PrintSettings {
	var resultCoreWebView2PrintSettings uintptr
	coreWebView2PrintSettingsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2PrintSettings)))
	return AsCoreWebView2PrintSettings(resultCoreWebView2PrintSettings)
}

func (m *TCoreWebView2PrintSettings) Orientation() TWVPrintOrientation {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TWVPrintOrientation(r1)
}

func (m *TCoreWebView2PrintSettings) SetOrientation(AValue TWVPrintOrientation) {
	coreWebView2PrintSettingsImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PrintSettings) ScaleFactor() (resultDouble float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(20, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2PrintSettings) SetScaleFactor(AValue float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(20, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2PrintSettings) PageWidth() (resultDouble float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(17, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2PrintSettings) SetPageWidth(AValue float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(17, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2PrintSettings) PageHeight() (resultDouble float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(15, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2PrintSettings) SetPageHeight(AValue float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(15, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2PrintSettings) MarginTop() (resultDouble float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(12, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2PrintSettings) SetMarginTop(AValue float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(12, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2PrintSettings) MarginBottom() (resultDouble float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(9, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2PrintSettings) SetMarginBottom(AValue float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(9, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2PrintSettings) MarginLeft() (resultDouble float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(10, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2PrintSettings) SetMarginLeft(AValue float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(10, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2PrintSettings) MarginRight() (resultDouble float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(11, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2PrintSettings) SetMarginRight(AValue float64) {
	coreWebView2PrintSettingsImportAPI().SysCallN(11, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2PrintSettings) ShouldPrintBackgrounds() bool {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2PrintSettings) SetShouldPrintBackgrounds(AValue bool) {
	coreWebView2PrintSettingsImportAPI().SysCallN(21, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2PrintSettings) ShouldPrintSelectionOnly() bool {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2PrintSettings) SetShouldPrintSelectionOnly(AValue bool) {
	coreWebView2PrintSettingsImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2PrintSettings) ShouldPrintHeaderAndFooter() bool {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2PrintSettings) SetShouldPrintHeaderAndFooter(AValue bool) {
	coreWebView2PrintSettingsImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2PrintSettings) HeaderTitle() string {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2PrintSettings) SetHeaderTitle(AValue string) {
	coreWebView2PrintSettingsImportAPI().SysCallN(7, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2PrintSettings) FooterUri() string {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2PrintSettings) SetFooterUri(AValue string) {
	coreWebView2PrintSettingsImportAPI().SysCallN(6, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2PrintSettings) PageRanges() string {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2PrintSettings) SetPageRanges(AValue string) {
	coreWebView2PrintSettingsImportAPI().SysCallN(16, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2PrintSettings) PagesPerSide() int32 {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PrintSettings) SetPagesPerSide(AValue int32) {
	coreWebView2PrintSettingsImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PrintSettings) Copies() int32 {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PrintSettings) SetCopies(AValue int32) {
	coreWebView2PrintSettingsImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PrintSettings) Collation() TWVPrintCollation {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TWVPrintCollation(r1)
}

func (m *TCoreWebView2PrintSettings) SetCollation(AValue TWVPrintCollation) {
	coreWebView2PrintSettingsImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PrintSettings) ColorMode() TWVPrintColorMode {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TWVPrintColorMode(r1)
}

func (m *TCoreWebView2PrintSettings) SetColorMode(AValue TWVPrintColorMode) {
	coreWebView2PrintSettingsImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PrintSettings) Duplex() TWVPrintDuplex {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TWVPrintDuplex(r1)
}

func (m *TCoreWebView2PrintSettings) SetDuplex(AValue TWVPrintDuplex) {
	coreWebView2PrintSettingsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PrintSettings) MediaSize() TWVPrintMediaSize {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TWVPrintMediaSize(r1)
}

func (m *TCoreWebView2PrintSettings) SetMediaSize(AValue TWVPrintMediaSize) {
	coreWebView2PrintSettingsImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PrintSettings) PrinterName() string {
	r1 := coreWebView2PrintSettingsImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2PrintSettings) SetPrinterName(AValue string) {
	coreWebView2PrintSettingsImportAPI().SysCallN(19, 1, m.Instance(), PascalStr(AValue))
}

var (
	coreWebView2PrintSettingsImport       *imports.Imports = nil
	coreWebView2PrintSettingsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2PrintSettings_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2PrintSettings_Collation", 0),
		/*2*/ imports.NewTable("CoreWebView2PrintSettings_ColorMode", 0),
		/*3*/ imports.NewTable("CoreWebView2PrintSettings_Copies", 0),
		/*4*/ imports.NewTable("CoreWebView2PrintSettings_Create", 0),
		/*5*/ imports.NewTable("CoreWebView2PrintSettings_Duplex", 0),
		/*6*/ imports.NewTable("CoreWebView2PrintSettings_FooterUri", 0),
		/*7*/ imports.NewTable("CoreWebView2PrintSettings_HeaderTitle", 0),
		/*8*/ imports.NewTable("CoreWebView2PrintSettings_Initialized", 0),
		/*9*/ imports.NewTable("CoreWebView2PrintSettings_MarginBottom", 0),
		/*10*/ imports.NewTable("CoreWebView2PrintSettings_MarginLeft", 0),
		/*11*/ imports.NewTable("CoreWebView2PrintSettings_MarginRight", 0),
		/*12*/ imports.NewTable("CoreWebView2PrintSettings_MarginTop", 0),
		/*13*/ imports.NewTable("CoreWebView2PrintSettings_MediaSize", 0),
		/*14*/ imports.NewTable("CoreWebView2PrintSettings_Orientation", 0),
		/*15*/ imports.NewTable("CoreWebView2PrintSettings_PageHeight", 0),
		/*16*/ imports.NewTable("CoreWebView2PrintSettings_PageRanges", 0),
		/*17*/ imports.NewTable("CoreWebView2PrintSettings_PageWidth", 0),
		/*18*/ imports.NewTable("CoreWebView2PrintSettings_PagesPerSide", 0),
		/*19*/ imports.NewTable("CoreWebView2PrintSettings_PrinterName", 0),
		/*20*/ imports.NewTable("CoreWebView2PrintSettings_ScaleFactor", 0),
		/*21*/ imports.NewTable("CoreWebView2PrintSettings_ShouldPrintBackgrounds", 0),
		/*22*/ imports.NewTable("CoreWebView2PrintSettings_ShouldPrintHeaderAndFooter", 0),
		/*23*/ imports.NewTable("CoreWebView2PrintSettings_ShouldPrintSelectionOnly", 0),
	}
)

func coreWebView2PrintSettingsImportAPI() *imports.Imports {
	if coreWebView2PrintSettingsImport == nil {
		coreWebView2PrintSettingsImport = NewDefaultImports()
		coreWebView2PrintSettingsImport.SetImportTable(coreWebView2PrintSettingsImportTables)
		coreWebView2PrintSettingsImportTables = nil
	}
	return coreWebView2PrintSettingsImport
}
