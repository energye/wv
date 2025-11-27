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

// IWkNSURLAuthenticationChallenge Parent: lcl.IObject
type IWkNSURLAuthenticationChallenge interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSURLAuthenticationChallenge // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNSURLAuthenticationChallenge struct {
	lcl.TObject
}

func (m *TWkNSURLAuthenticationChallenge) Data() NSURLAuthenticationChallenge {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLAuthenticationChallengeAPI().SysCallN(1, m.Instance())
	return NSURLAuthenticationChallenge(r)
}

func (m *TWkNSURLAuthenticationChallenge) Release() {
	if !m.IsValid() {
		return
	}
	wkNSURLAuthenticationChallengeAPI().SysCallN(2, m.Instance())
}

// NewURLAuthenticationChallenge class constructor
func NewURLAuthenticationChallenge(data NSURLAuthenticationChallenge) IWkNSURLAuthenticationChallenge {
	r := wkNSURLAuthenticationChallengeAPI().SysCallN(0, uintptr(data))
	return AsWkNSURLAuthenticationChallenge(r)
}

var (
	wkNSURLAuthenticationChallengeOnce   base.Once
	wkNSURLAuthenticationChallengeImport *imports.Imports = nil
)

func wkNSURLAuthenticationChallengeAPI() *imports.Imports {
	wkNSURLAuthenticationChallengeOnce.Do(func() {
		wkNSURLAuthenticationChallengeImport = api.NewDefaultImports()
		wkNSURLAuthenticationChallengeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNSURLAuthenticationChallenge_Create", 0), // constructor NewURLAuthenticationChallenge
			/* 1 */ imports.NewTable("TWkNSURLAuthenticationChallenge_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNSURLAuthenticationChallenge_Release", 0), // procedure Release
		}
	})
	return wkNSURLAuthenticationChallengeImport
}
