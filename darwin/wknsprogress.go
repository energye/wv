//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
)

// IWkNSProgress Parent: lcl.IObject
type IWkNSProgress interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSProgress // function
	// TotalUnitCount
	//  Returns The total number of tracked units of work for the current progress.
	TotalUnitCount() int64 // function
	// CompletedUnitCount
	//  Returns The number of completed units of work for the current job.
	CompletedUnitCount() int64 // function
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
	FractionCompleted() float64 // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// SetTotalUnitCount
	//  Sets The total number of tracked units of work for the current progress.
	SetTotalUnitCount(value int64) // procedure
	// SetCompletedUnitCount
	//  Sets The number of completed units of work for the current job.
	SetCompletedUnitCount(value int64) // procedure
	// SetLocalizedDescription
	//  Sets A localized description of tracked progress for the receiver.
	SetLocalizedDescription(value string) // procedure
	// SetLocalizedAdditionalDescription
	//  Sets A more specific localized description of tracked progress for the receiver.
	SetLocalizedAdditionalDescription(value string) // procedure
	// SetCancellable
	//  Sets A Boolean value that indicates whether the receiver is tracking work that you can cancel.
	SetCancellable(value bool) // procedure
	// SetPausable
	//  Sets A Boolean value that indicates whether the receiver is tracking work that you can pause.
	SetPausable(value bool) // procedure
	// SetCancellationHandler
	//  Sets The block to invoke when canceling progress.
	SetCancellationHandler(value uintptr) // procedure
	// SetPausingHandler
	//  Sets The block to invoke when pausing progress.
	SetPausingHandler(value uintptr) // procedure
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

type TWkNSProgress struct {
	lcl.TObject
}

func (m *TWkNSProgress) Data() NSProgress {
	if !m.IsValid() {
		return 0
	}
	r := wkNSProgressAPI().SysCallN(1, m.Instance())
	return NSProgress(r)
}

func (m *TWkNSProgress) TotalUnitCount() (result int64) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkNSProgress) CompletedUnitCount() (result int64) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkNSProgress) LocalizedDescription() string {
	if !m.IsValid() {
		return ""
	}
	r := wkNSProgressAPI().SysCallN(6, m.Instance())
	return api.GoStr(r)
}

func (m *TWkNSProgress) LocalizedAdditionalDescription() string {
	if !m.IsValid() {
		return ""
	}
	r := wkNSProgressAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
}

func (m *TWkNSProgress) IsCancellable() bool {
	if !m.IsValid() {
		return false
	}
	r := wkNSProgressAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TWkNSProgress) IsPausable() bool {
	if !m.IsValid() {
		return false
	}
	r := wkNSProgressAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TWkNSProgress) IsCancelled() bool {
	if !m.IsValid() {
		return false
	}
	r := wkNSProgressAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TWkNSProgress) IsPaused() bool {
	if !m.IsValid() {
		return false
	}
	r := wkNSProgressAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TWkNSProgress) CancellationHandler() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := wkNSProgressAPI().SysCallN(12, m.Instance())
	return uintptr(r)
}

func (m *TWkNSProgress) PausingHandler() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := wkNSProgressAPI().SysCallN(13, m.Instance())
	return uintptr(r)
}

func (m *TWkNSProgress) IsIndeterminate() bool {
	if !m.IsValid() {
		return false
	}
	r := wkNSProgressAPI().SysCallN(14, m.Instance())
	return api.GoBool(r)
}

func (m *TWkNSProgress) FractionCompleted() (result float64) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkNSProgress) Release() {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(16, m.Instance())
}

func (m *TWkNSProgress) SetTotalUnitCount(value int64) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TWkNSProgress) SetCompletedUnitCount(value int64) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TWkNSProgress) SetLocalizedDescription(value string) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(19, m.Instance(), api.PasStr(value))
}

func (m *TWkNSProgress) SetLocalizedAdditionalDescription(value string) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(20, m.Instance(), api.PasStr(value))
}

func (m *TWkNSProgress) SetCancellable(value bool) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(21, m.Instance(), api.PasBool(value))
}

func (m *TWkNSProgress) SetPausable(value bool) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(22, m.Instance(), api.PasBool(value))
}

func (m *TWkNSProgress) SetCancellationHandler(value uintptr) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(23, m.Instance(), uintptr(value))
}

func (m *TWkNSProgress) SetPausingHandler(value uintptr) {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(24, m.Instance(), uintptr(value))
}

func (m *TWkNSProgress) Cancel() {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(25, m.Instance())
}

func (m *TWkNSProgress) Pause() {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(26, m.Instance())
}

func (m *TWkNSProgress) Publish() {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(27, m.Instance())
}

func (m *TWkNSProgress) Unpublish() {
	if !m.IsValid() {
		return
	}
	wkNSProgressAPI().SysCallN(28, m.Instance())
}

// Progress  is static instance
var Progress _ProgressClass

// _ProgressClass is class type defined by TWkNSProgress
type _ProgressClass uintptr

// CurrentProgress
//
//	Returns the progress instance, if any.
func (_ProgressClass) CurrentProgress() IWkNSProgress {
	r := wkNSProgressAPI().SysCallN(2)
	return AsWkNSProgress(r)
}

// ProgressWithTotalUnitCount
//
//	Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform.
func (_ProgressClass) ProgressWithTotalUnitCount(count int64) IWkNSProgress {
	r := wkNSProgressAPI().SysCallN(3, uintptr(base.UnsafePointer(&count)))
	return AsWkNSProgress(r)
}

// NewProgress class constructor
func NewProgress(data NSProgress) IWkNSProgress {
	r := wkNSProgressAPI().SysCallN(0, uintptr(data))
	return AsWkNSProgress(r)
}

var (
	wkNSProgressOnce   base.Once
	wkNSProgressImport *imports.Imports = nil
)

func wkNSProgressAPI() *imports.Imports {
	wkNSProgressOnce.Do(func() {
		wkNSProgressImport = api.NewDefaultImports()
		wkNSProgressImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNSProgress_Create", 0), // constructor NewProgress
			/* 1 */ imports.NewTable("TWkNSProgress_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNSProgress_CurrentProgress", 0), // static function CurrentProgress
			/* 3 */ imports.NewTable("TWkNSProgress_ProgressWithTotalUnitCount", 0), // static function ProgressWithTotalUnitCount
			/* 4 */ imports.NewTable("TWkNSProgress_TotalUnitCount", 0), // function TotalUnitCount
			/* 5 */ imports.NewTable("TWkNSProgress_CompletedUnitCount", 0), // function CompletedUnitCount
			/* 6 */ imports.NewTable("TWkNSProgress_LocalizedDescription", 0), // function LocalizedDescription
			/* 7 */ imports.NewTable("TWkNSProgress_LocalizedAdditionalDescription", 0), // function LocalizedAdditionalDescription
			/* 8 */ imports.NewTable("TWkNSProgress_IsCancellable", 0), // function IsCancellable
			/* 9 */ imports.NewTable("TWkNSProgress_IsPausable", 0), // function IsPausable
			/* 10 */ imports.NewTable("TWkNSProgress_IsCancelled", 0), // function IsCancelled
			/* 11 */ imports.NewTable("TWkNSProgress_IsPaused", 0), // function IsPaused
			/* 12 */ imports.NewTable("TWkNSProgress_CancellationHandler", 0), // function CancellationHandler
			/* 13 */ imports.NewTable("TWkNSProgress_PausingHandler", 0), // function PausingHandler
			/* 14 */ imports.NewTable("TWkNSProgress_IsIndeterminate", 0), // function IsIndeterminate
			/* 15 */ imports.NewTable("TWkNSProgress_FractionCompleted", 0), // function FractionCompleted
			/* 16 */ imports.NewTable("TWkNSProgress_Release", 0), // procedure Release
			/* 17 */ imports.NewTable("TWkNSProgress_SetTotalUnitCount", 0), // procedure SetTotalUnitCount
			/* 18 */ imports.NewTable("TWkNSProgress_SetCompletedUnitCount", 0), // procedure SetCompletedUnitCount
			/* 19 */ imports.NewTable("TWkNSProgress_SetLocalizedDescription", 0), // procedure SetLocalizedDescription
			/* 20 */ imports.NewTable("TWkNSProgress_SetLocalizedAdditionalDescription", 0), // procedure SetLocalizedAdditionalDescription
			/* 21 */ imports.NewTable("TWkNSProgress_SetCancellable", 0), // procedure SetCancellable
			/* 22 */ imports.NewTable("TWkNSProgress_SetPausable", 0), // procedure SetPausable
			/* 23 */ imports.NewTable("TWkNSProgress_SetCancellationHandler", 0), // procedure SetCancellationHandler
			/* 24 */ imports.NewTable("TWkNSProgress_SetPausingHandler", 0), // procedure SetPausingHandler
			/* 25 */ imports.NewTable("TWkNSProgress_Cancel", 0), // procedure Cancel
			/* 26 */ imports.NewTable("TWkNSProgress_Pause", 0), // procedure Pause
			/* 27 */ imports.NewTable("TWkNSProgress_Publish", 0), // procedure Publish
			/* 28 */ imports.NewTable("TWkNSProgress_Unpublish", 0), // procedure Unpublish
		}
	})
	return wkNSProgressImport
}
