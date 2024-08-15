//go:build mage
// +build mage

package main

import (
	"github.com/zilinyo/gomake/mageutil"
	"os"
)

var Default = Build

func Build() {
	mageutil.Build()
}

func Start() {
	mageutil.InitForSSC()
	err := setMaxOpenFiles()
	if err != nil {
		mageutil.PrintRed("setMaxOpenFiles failed " + err.Error())
		os.Exit(1)
	}
	mageutil.StartToolsAndServices()
}

func Stop() {
	mageutil.StopAndCheckBinaries()
}

func Check() {
	mageutil.CheckAndReportBinariesStatus()
}

func Protocol() {
	mageutil.Protocol()
}
