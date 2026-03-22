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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2PrintSettings Parent: IObject
type ICoreWebView2PrintSettings interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PrintSettings // property BaseIntf Getter
	// Orientation
	//  The orientation can be portrait or landscape. The default orientation is
	//  portrait. See `COREWEBVIEW2_PRINT_ORIENTATION`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_orientation">See the ICoreWebView2PrintSettings article.</see>
	Orientation() wvTypes.TWVPrintOrientation         // property Orientation Getter
	SetOrientation(value wvTypes.TWVPrintOrientation) // property Orientation Setter
	// ScaleFactor
	//  The scale factor is a value between 0.1 and 2.0. The default is 1.0.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_scalefactor">See the ICoreWebView2PrintSettings article.</see>
	ScaleFactor() float64         // property ScaleFactor Getter
	SetScaleFactor(value float64) // property ScaleFactor Setter
	// PageWidth
	//  The page width in inches. The default width is 8.5 inches.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_pagewidth">See the ICoreWebView2PrintSettings article.</see>
	PageWidth() float64         // property PageWidth Getter
	SetPageWidth(value float64) // property PageWidth Setter
	// PageHeight
	//  The page height in inches. The default height is 11 inches.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_pageheight">See the ICoreWebView2PrintSettings article.</see>
	PageHeight() float64         // property PageHeight Getter
	SetPageHeight(value float64) // property PageHeight Setter
	// MarginTop
	//  The top margin in inches. The default is 1 cm, or ~0.4 inches.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_margintop">See the ICoreWebView2PrintSettings article.</see>
	MarginTop() float64         // property MarginTop Getter
	SetMarginTop(value float64) // property MarginTop Setter
	// MarginBottom
	//  The bottom margin in inches. The default is 1 cm, or ~0.4 inches.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_marginbottom">See the ICoreWebView2PrintSettings article.</see>
	MarginBottom() float64         // property MarginBottom Getter
	SetMarginBottom(value float64) // property MarginBottom Setter
	// MarginLeft
	//  The left margin in inches. The default is 1 cm, or ~0.4 inches.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_marginleft">See the ICoreWebView2PrintSettings article.</see>
	MarginLeft() float64         // property MarginLeft Getter
	SetMarginLeft(value float64) // property MarginLeft Setter
	// MarginRight
	//  The right margin in inches. The default is 1 cm, or ~0.4 inches.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_marginright">See the ICoreWebView2PrintSettings article.</see>
	MarginRight() float64         // property MarginRight Getter
	SetMarginRight(value float64) // property MarginRight Setter
	// ShouldPrintBackgrounds
	//  `TRUE` if background colors and images should be printed. The default value
	//  is `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_shouldprintbackgrounds">See the ICoreWebView2PrintSettings article.</see>
	ShouldPrintBackgrounds() bool         // property ShouldPrintBackgrounds Getter
	SetShouldPrintBackgrounds(value bool) // property ShouldPrintBackgrounds Setter
	// ShouldPrintSelectionOnly
	//  `TRUE` if only the current end user's selection of HTML in the document
	//  should be printed. The default value is `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_shouldprintselectiononly">See the ICoreWebView2PrintSettings article.</see>
	ShouldPrintSelectionOnly() bool         // property ShouldPrintSelectionOnly Getter
	SetShouldPrintSelectionOnly(value bool) // property ShouldPrintSelectionOnly Setter
	// ShouldPrintHeaderAndFooter
	//  `TRUE` if header and footer should be printed. The default value is `FALSE`.
	//  The header consists of the date and time of printing, and the title of the
	//  page. The footer consists of the URI and page number. The height of the
	//  header and footer is 0.5 cm, or ~0.2 inches.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_shouldprintheaderandfooter">See the ICoreWebView2PrintSettings article.</see>
	ShouldPrintHeaderAndFooter() bool         // property ShouldPrintHeaderAndFooter Getter
	SetShouldPrintHeaderAndFooter(value bool) // property ShouldPrintHeaderAndFooter Setter
	// HeaderTitle
	//  The title in the header if `ShouldPrintHeaderAndFooter` is `TRUE`. The
	//  default value is the title of the current document.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_headertitle">See the ICoreWebView2PrintSettings article.</see>
	HeaderTitle() string         // property HeaderTitle Getter
	SetHeaderTitle(value string) // property HeaderTitle Setter
	// FooterUri
	//  The URI in the footer if `ShouldPrintHeaderAndFooter` is `TRUE`. The
	//  default value is the current URI.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings#get_footeruri">See the ICoreWebView2PrintSettings article.</see>
	FooterUri() string         // property FooterUri Getter
	SetFooterUri(value string) // property FooterUri Setter
	// PageRanges
	//  Page range to print.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_pageranges">See the ICoreWebView2PrintSettings2 article.</see>
	PageRanges() string         // property PageRanges Getter
	SetPageRanges(value string) // property PageRanges Setter
	// PagesPerSide
	//  Prints multiple pages of a document on a single piece of paper. Choose from 1, 2, 4, 6, 9 or 16.
	//  The default value is 1.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_pagesperside">See the ICoreWebView2PrintSettings2 article.</see>
	PagesPerSide() int32         // property PagesPerSide Getter
	SetPagesPerSide(value int32) // property PagesPerSide Setter
	// Copies
	//  Number of copies to print. Minimum value is `1` and the maximum copies count is `999`.
	//  The default value is 1.
	//  This value is ignored in PrintToPdfStream method.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_copies">See the ICoreWebView2PrintSettings2 article.</see>
	Copies() int32         // property Copies Getter
	SetCopies(value int32) // property Copies Setter
	// Collation
	//  Printer collation. See `COREWEBVIEW2_PRINT_COLLATION` for descriptions of
	//  collation. The default value is `COREWEBVIEW2_PRINT_COLLATION_DEFAULT`.
	//  Printing uses default value of printer's collation if an invalid value is provided
	//  for the specific printer. This value is ignored in PrintToPdfStream method.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_collation">See the ICoreWebView2PrintSettings2 article.</see>
	Collation() wvTypes.TWVPrintCollation         // property Collation Getter
	SetCollation(value wvTypes.TWVPrintCollation) // property Collation Setter
	// ColorMode
	//  Printer color mode. See `COREWEBVIEW2_PRINT_COLOR_MODE` for descriptions
	//  of color modes. The default value is `COREWEBVIEW2_PRINT_COLOR_MODE_DEFAULT`.
	//  Printing uses default value of printer supported color if an invalid value is provided
	//  for the specific printer.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_colormode">See the ICoreWebView2PrintSettings2 article.</see>
	ColorMode() wvTypes.TWVPrintColorMode         // property ColorMode Getter
	SetColorMode(value wvTypes.TWVPrintColorMode) // property ColorMode Setter
	// Duplex
	//  Printer duplex settings. See `COREWEBVIEW2_PRINT_DUPLEX` for descriptions of duplex.
	//  The default value is `COREWEBVIEW2_PRINT_DUPLEX_DEFAULT`.
	//  Printing uses default value of printer's duplex if an invalid value is provided
	//  for the specific printer. This value is ignored in PrintToPdfStream method.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_duplex">See the ICoreWebView2PrintSettings2 article.</see>
	Duplex() wvTypes.TWVPrintDuplex         // property Duplex Getter
	SetDuplex(value wvTypes.TWVPrintDuplex) // property Duplex Setter
	// MediaSize
	//  Printer media size. See `COREWEBVIEW2_PRINT_MEDIA_SIZE` for descriptions of media size.
	//  The default value is `COREWEBVIEW2_PRINT_MEDIA_SIZE_DEFAULT`.
	//  If media size is `COREWEBVIEW2_PRINT_MEDIA_SIZE_CUSTOM`, you should set the `PageWidth`
	//  and `PageHeight`.
	//  Printing uses default value of printer supported media size if an invalid value is provided
	//  for the specific printer. This value is ignored in PrintToPdfStream method.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_mediasize">See the ICoreWebView2PrintSettings2 article.</see>
	MediaSize() wvTypes.TWVPrintMediaSize         // property MediaSize Getter
	SetMediaSize(value wvTypes.TWVPrintMediaSize) // property MediaSize Setter
	// PrinterName
	//  The name of the printer to use. Defaults to empty string.
	//  If the printer name is empty string or null, then it prints to the default
	//  printer on the user OS. This value is ignored in PrintToPdfStream method.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2#get_printername">See the ICoreWebView2PrintSettings2 article.</see>
	PrinterName() string         // property PrinterName Getter
	SetPrinterName(value string) // property PrinterName Setter
}

type TCoreWebView2PrintSettings struct {
	TObject
}

func (m *TCoreWebView2PrintSettings) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PrintSettings) BaseIntf() (result ICoreWebView2PrintSettings) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2PrintSettingsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2PrintSettings(resultPtr)
	return
}

func (m *TCoreWebView2PrintSettings) Orientation() wvTypes.TWVPrintOrientation {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(3, 0, m.Instance())
	return wvTypes.TWVPrintOrientation(r)
}

func (m *TCoreWebView2PrintSettings) SetOrientation(value wvTypes.TWVPrintOrientation) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PrintSettings) ScaleFactor() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PrintSettings) SetScaleFactor(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(4, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PrintSettings) PageWidth() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PrintSettings) SetPageWidth(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(5, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PrintSettings) PageHeight() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PrintSettings) SetPageHeight(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(6, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PrintSettings) MarginTop() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PrintSettings) SetMarginTop(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(7, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PrintSettings) MarginBottom() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PrintSettings) SetMarginBottom(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PrintSettings) MarginLeft() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PrintSettings) SetMarginLeft(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(9, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PrintSettings) MarginRight() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PrintSettings) SetMarginRight(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(10, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PrintSettings) ShouldPrintBackgrounds() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PrintSettings) SetShouldPrintBackgrounds(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2PrintSettings) ShouldPrintSelectionOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PrintSettings) SetShouldPrintSelectionOnly(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2PrintSettings) ShouldPrintHeaderAndFooter() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PrintSettings) SetShouldPrintHeaderAndFooter(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2PrintSettings) HeaderTitle() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(14, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2PrintSettings) SetHeaderTitle(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(14, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2PrintSettings) FooterUri() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(15, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2PrintSettings) SetFooterUri(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(15, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2PrintSettings) PageRanges() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(16, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2PrintSettings) SetPageRanges(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(16, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2PrintSettings) PagesPerSide() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2PrintSettings) SetPagesPerSide(value int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PrintSettings) Copies() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2PrintSettings) SetCopies(value int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PrintSettings) Collation() wvTypes.TWVPrintCollation {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(19, 0, m.Instance())
	return wvTypes.TWVPrintCollation(r)
}

func (m *TCoreWebView2PrintSettings) SetCollation(value wvTypes.TWVPrintCollation) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PrintSettings) ColorMode() wvTypes.TWVPrintColorMode {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(20, 0, m.Instance())
	return wvTypes.TWVPrintColorMode(r)
}

func (m *TCoreWebView2PrintSettings) SetColorMode(value wvTypes.TWVPrintColorMode) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PrintSettings) Duplex() wvTypes.TWVPrintDuplex {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(21, 0, m.Instance())
	return wvTypes.TWVPrintDuplex(r)
}

func (m *TCoreWebView2PrintSettings) SetDuplex(value wvTypes.TWVPrintDuplex) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PrintSettings) MediaSize() wvTypes.TWVPrintMediaSize {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(22, 0, m.Instance())
	return wvTypes.TWVPrintMediaSize(r)
}

func (m *TCoreWebView2PrintSettings) SetMediaSize(value wvTypes.TWVPrintMediaSize) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PrintSettings) PrinterName() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2PrintSettingsAPI().SysCallN(23, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2PrintSettings) SetPrinterName(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2PrintSettingsAPI().SysCallN(23, 1, m.Instance(), api.PasStr(value))
}

// NewCoreWebView2PrintSettings class constructor
func NewCoreWebView2PrintSettings(baseIntf ICoreWebView2PrintSettings) ICoreWebView2PrintSettings {
	r := coreWebView2PrintSettingsAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2PrintSettings(r)
}

var (
	coreWebView2PrintSettingsOnce   base.Once
	coreWebView2PrintSettingsImport *imports.Imports = nil
)

func coreWebView2PrintSettingsAPI() *imports.Imports {
	coreWebView2PrintSettingsOnce.Do(func() {
		coreWebView2PrintSettingsImport = api.NewDefaultImports()
		coreWebView2PrintSettingsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2PrintSettings_Create", 0), // constructor NewCoreWebView2PrintSettings
			/* 1 */ imports.NewTable("TCoreWebView2PrintSettings_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2PrintSettings_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2PrintSettings_Orientation", 0), // property Orientation
			/* 4 */ imports.NewTable("TCoreWebView2PrintSettings_ScaleFactor", 0), // property ScaleFactor
			/* 5 */ imports.NewTable("TCoreWebView2PrintSettings_PageWidth", 0), // property PageWidth
			/* 6 */ imports.NewTable("TCoreWebView2PrintSettings_PageHeight", 0), // property PageHeight
			/* 7 */ imports.NewTable("TCoreWebView2PrintSettings_MarginTop", 0), // property MarginTop
			/* 8 */ imports.NewTable("TCoreWebView2PrintSettings_MarginBottom", 0), // property MarginBottom
			/* 9 */ imports.NewTable("TCoreWebView2PrintSettings_MarginLeft", 0), // property MarginLeft
			/* 10 */ imports.NewTable("TCoreWebView2PrintSettings_MarginRight", 0), // property MarginRight
			/* 11 */ imports.NewTable("TCoreWebView2PrintSettings_ShouldPrintBackgrounds", 0), // property ShouldPrintBackgrounds
			/* 12 */ imports.NewTable("TCoreWebView2PrintSettings_ShouldPrintSelectionOnly", 0), // property ShouldPrintSelectionOnly
			/* 13 */ imports.NewTable("TCoreWebView2PrintSettings_ShouldPrintHeaderAndFooter", 0), // property ShouldPrintHeaderAndFooter
			/* 14 */ imports.NewTable("TCoreWebView2PrintSettings_HeaderTitle", 0), // property HeaderTitle
			/* 15 */ imports.NewTable("TCoreWebView2PrintSettings_FooterUri", 0), // property FooterUri
			/* 16 */ imports.NewTable("TCoreWebView2PrintSettings_PageRanges", 0), // property PageRanges
			/* 17 */ imports.NewTable("TCoreWebView2PrintSettings_PagesPerSide", 0), // property PagesPerSide
			/* 18 */ imports.NewTable("TCoreWebView2PrintSettings_Copies", 0), // property Copies
			/* 19 */ imports.NewTable("TCoreWebView2PrintSettings_Collation", 0), // property Collation
			/* 20 */ imports.NewTable("TCoreWebView2PrintSettings_ColorMode", 0), // property ColorMode
			/* 21 */ imports.NewTable("TCoreWebView2PrintSettings_Duplex", 0), // property Duplex
			/* 22 */ imports.NewTable("TCoreWebView2PrintSettings_MediaSize", 0), // property MediaSize
			/* 23 */ imports.NewTable("TCoreWebView2PrintSettings_PrinterName", 0), // property PrinterName
		}
	})
	return coreWebView2PrintSettingsImport
}
