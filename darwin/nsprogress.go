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

// INSProgress Root Interface
//
//	The NSProgress class provides a self-contained mechanism for progress reporting.
//	It makes it easy for code that performs work to report the progress of that work,
//	and for user interface code to observe that progress for presentation to the user.
//	Specifically, you can use a progress object to show the user a progress bar and
//	explanatory text that update as you do work. It also allows the user to cancel or pause work.
//	https://developer.apple.com/documentation/foundation/nsprogress?language=objc
type INSProgress interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSProgress // function
	// TotalUnitCount
	//  Returns The total number of tracked units of work for the current progress.
	TotalUnitCount() (resultInt64 int64) // function
	// CompletedUnitCount
	//  Returns The number of completed units of work for the current job.
	CompletedUnitCount() (resultInt64 int64) // function
	// LocalizedDescription
	//  Returns A localized description of tracked progress for the receiver.
	LocalizedDescription() string // function
	// LocalizedAdditionalDescription
	//  Returns A more specific localized description of tracked progress for the receiver.
	LocalizedAdditionalDescription() string // function
	// IsCancellable
	//  Returns A Boolean value that indicates whether the receiver is tracking work that you can cancel.
	IsCancellable() bool // function
	// IsPausable
	//  Returns A Boolean value that indicates whether the receiver is tracking work that you can pause.
	IsPausable() bool // function
	// IsCancelled
	//  A Boolean value that Indicates whether the receiver is tracking canceled work.
	IsCancelled() bool // function
	// IsPaused
	//  A Boolean value that indicates whether the receiver is tracking paused work.
	IsPaused() bool // function
	// CancellationHandler
	//  Returns The block to invoke when canceling progress.
	CancellationHandler() uintptr // function
	// PausingHandler
	//  Returns The block to invoke when pausing progress.
	PausingHandler() uintptr // function
	// IsIndeterminate
	//  A Boolean value that indicates whether the tracked progress is indeterminate.
	IsIndeterminate() bool // function
	// FractionCompleted
	//  The fraction of the overall work that the progress object completes, including work from its suboperations.
	FractionCompleted() (resultFloat64 float64) // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// SetTotalUnitCount
	//  Sets The total number of tracked units of work for the current progress.
	SetTotalUnitCount(aValue int64) // procedure
	// SetCompletedUnitCount
	//  Sets The number of completed units of work for the current job.
	SetCompletedUnitCount(aValue int64) // procedure
	// SetLocalizedDescription
	//  Sets A localized description of tracked progress for the receiver.
	SetLocalizedDescription(aValue string) // procedure
	// SetLocalizedAdditionalDescription
	//  Sets A more specific localized description of tracked progress for the receiver.
	SetLocalizedAdditionalDescription(aValue string) // procedure
	// SetCancellable
	//  Sets A Boolean value that indicates whether the receiver is tracking work that you can cancel.
	SetCancellable(aValue bool) // procedure
	// SetPausable
	//  Sets A Boolean value that indicates whether the receiver is tracking work that you can pause.
	SetPausable(aValue bool) // procedure
	// SetCancellationHandler
	//  Sets The block to invoke when canceling progress.
	SetCancellationHandler(aValue uintptr) // procedure
	// SetPausingHandler
	//  Sets The block to invoke when pausing progress.
	SetPausingHandler(aValue uintptr) // procedure
	// Cancel
	//  Cancels progress tracking.
	Cancel() // procedure
	// Pause
	//  Pauses progress tracking.
	Pause() // procedure
	// Publish
	//  Publishes the progress object for other processes to observe it.
	Publish() // procedure
	// Unpublish
	//  Removes a progress object from publication, making it unobservable by other processes.
	Unpublish() // procedure
}

// TNSProgress Root Object
//
//	The NSProgress class provides a self-contained mechanism for progress reporting.
//	It makes it easy for code that performs work to report the progress of that work,
//	and for user interface code to observe that progress for presentation to the user.
//	Specifically, you can use a progress object to show the user a progress bar and
//	explanatory text that update as you do work. It also allows the user to cancel or pause work.
//	https://developer.apple.com/documentation/foundation/nsprogress?language=objc
type TNSProgress struct {
	TObject
}

func NewNSProgress(aData NSProgress) INSProgress {
	r1 := nSProgressImportAPI().SysCallN(3, uintptr(aData))
	return AsNSProgress(r1)
}

// NSProgressRef -> INSProgress
var NSProgressRef nSProgress

// nSProgress TNSProgress Ref
type nSProgress uintptr

// CurrentProgress
//
//	Returns the progress instance, if any.
func (m *nSProgress) CurrentProgress() INSProgress {
	r1 := nSProgressImportAPI().SysCallN(4)
	return AsNSProgress(r1)
}

// ProgressWithTotalUnitCount
//
//	Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform.
func (m *nSProgress) ProgressWithTotalUnitCount(aCount int64) INSProgress {
	r1 := nSProgressImportAPI().SysCallN(16, uintptr(unsafePointer(&aCount)))
	return AsNSProgress(r1)
}

func (m *TNSProgress) Data() NSProgress {
	r1 := nSProgressImportAPI().SysCallN(5, m.Instance())
	return NSProgress(r1)
}

func (m *TNSProgress) TotalUnitCount() (resultInt64 int64) {
	nSProgressImportAPI().SysCallN(27, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TNSProgress) CompletedUnitCount() (resultInt64 int64) {
	nSProgressImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TNSProgress) LocalizedDescription() string {
	r1 := nSProgressImportAPI().SysCallN(13, m.Instance())
	return GoStr(r1)
}

func (m *TNSProgress) LocalizedAdditionalDescription() string {
	r1 := nSProgressImportAPI().SysCallN(12, m.Instance())
	return GoStr(r1)
}

func (m *TNSProgress) IsCancellable() bool {
	r1 := nSProgressImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TNSProgress) IsPausable() bool {
	r1 := nSProgressImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TNSProgress) IsCancelled() bool {
	r1 := nSProgressImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TNSProgress) IsPaused() bool {
	r1 := nSProgressImportAPI().SysCallN(11, m.Instance())
	return GoBool(r1)
}

func (m *TNSProgress) CancellationHandler() uintptr {
	r1 := nSProgressImportAPI().SysCallN(1, m.Instance())
	return uintptr(r1)
}

func (m *TNSProgress) PausingHandler() uintptr {
	r1 := nSProgressImportAPI().SysCallN(15, m.Instance())
	return uintptr(r1)
}

func (m *TNSProgress) IsIndeterminate() bool {
	r1 := nSProgressImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TNSProgress) FractionCompleted() (resultFloat64 float64) {
	nSProgressImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TNSProgress) Release() {
	nSProgressImportAPI().SysCallN(18, m.Instance())
}

func (m *TNSProgress) SetTotalUnitCount(aValue int64) {
	nSProgressImportAPI().SysCallN(26, m.Instance(), uintptr(unsafePointer(&aValue)))
}

func (m *TNSProgress) SetCompletedUnitCount(aValue int64) {
	nSProgressImportAPI().SysCallN(21, m.Instance(), uintptr(unsafePointer(&aValue)))
}

func (m *TNSProgress) SetLocalizedDescription(aValue string) {
	nSProgressImportAPI().SysCallN(23, m.Instance(), PascalStr(aValue))
}

func (m *TNSProgress) SetLocalizedAdditionalDescription(aValue string) {
	nSProgressImportAPI().SysCallN(22, m.Instance(), PascalStr(aValue))
}

func (m *TNSProgress) SetCancellable(aValue bool) {
	nSProgressImportAPI().SysCallN(19, m.Instance(), PascalBool(aValue))
}

func (m *TNSProgress) SetPausable(aValue bool) {
	nSProgressImportAPI().SysCallN(24, m.Instance(), PascalBool(aValue))
}

func (m *TNSProgress) SetCancellationHandler(aValue uintptr) {
	nSProgressImportAPI().SysCallN(20, m.Instance(), uintptr(aValue))
}

func (m *TNSProgress) SetPausingHandler(aValue uintptr) {
	nSProgressImportAPI().SysCallN(25, m.Instance(), uintptr(aValue))
}

func (m *TNSProgress) Cancel() {
	nSProgressImportAPI().SysCallN(0, m.Instance())
}

func (m *TNSProgress) Pause() {
	nSProgressImportAPI().SysCallN(14, m.Instance())
}

func (m *TNSProgress) Publish() {
	nSProgressImportAPI().SysCallN(17, m.Instance())
}

func (m *TNSProgress) Unpublish() {
	nSProgressImportAPI().SysCallN(28, m.Instance())
}

var (
	nSProgressImport       *imports.Imports = nil
	nSProgressImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSProgress_Cancel", 0),
		/*1*/ imports.NewTable("NSProgress_CancellationHandler", 0),
		/*2*/ imports.NewTable("NSProgress_CompletedUnitCount", 0),
		/*3*/ imports.NewTable("NSProgress_Create", 0),
		/*4*/ imports.NewTable("NSProgress_CurrentProgress", 0),
		/*5*/ imports.NewTable("NSProgress_Data", 0),
		/*6*/ imports.NewTable("NSProgress_FractionCompleted", 0),
		/*7*/ imports.NewTable("NSProgress_IsCancellable", 0),
		/*8*/ imports.NewTable("NSProgress_IsCancelled", 0),
		/*9*/ imports.NewTable("NSProgress_IsIndeterminate", 0),
		/*10*/ imports.NewTable("NSProgress_IsPausable", 0),
		/*11*/ imports.NewTable("NSProgress_IsPaused", 0),
		/*12*/ imports.NewTable("NSProgress_LocalizedAdditionalDescription", 0),
		/*13*/ imports.NewTable("NSProgress_LocalizedDescription", 0),
		/*14*/ imports.NewTable("NSProgress_Pause", 0),
		/*15*/ imports.NewTable("NSProgress_PausingHandler", 0),
		/*16*/ imports.NewTable("NSProgress_ProgressWithTotalUnitCount", 0),
		/*17*/ imports.NewTable("NSProgress_Publish", 0),
		/*18*/ imports.NewTable("NSProgress_Release", 0),
		/*19*/ imports.NewTable("NSProgress_SetCancellable", 0),
		/*20*/ imports.NewTable("NSProgress_SetCancellationHandler", 0),
		/*21*/ imports.NewTable("NSProgress_SetCompletedUnitCount", 0),
		/*22*/ imports.NewTable("NSProgress_SetLocalizedAdditionalDescription", 0),
		/*23*/ imports.NewTable("NSProgress_SetLocalizedDescription", 0),
		/*24*/ imports.NewTable("NSProgress_SetPausable", 0),
		/*25*/ imports.NewTable("NSProgress_SetPausingHandler", 0),
		/*26*/ imports.NewTable("NSProgress_SetTotalUnitCount", 0),
		/*27*/ imports.NewTable("NSProgress_TotalUnitCount", 0),
		/*28*/ imports.NewTable("NSProgress_Unpublish", 0),
	}
)

func nSProgressImportAPI() *imports.Imports {
	if nSProgressImport == nil {
		nSProgressImport = NewDefaultImports()
		nSProgressImport.SetImportTable(nSProgressImportTables)
		nSProgressImportTables = nil
	}
	return nSProgressImport
}
