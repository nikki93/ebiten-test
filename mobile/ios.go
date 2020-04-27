// +build darwin,ios

package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/mobile/ebitenmobileview"
)

/*
#cgo CFLAGS: -DGLES_SILENCE_DEPRECATION -Werror -Wno-deprecated-declarations -fmodules -fobjc-arc -x objective-c

#include "stdint.h"
#include "ios.h"

*/
import "C"

//export test_app_go
func test_app_go() {
	fmt.Println("ios.go: test_app_go")
}

//export proxyebitenmobileview__Layout
func proxyebitenmobileview__Layout(param_viewWidth C.double, param_viewHeight C.double) {
	_param_viewWidth := float64(param_viewWidth)
	_param_viewHeight := float64(param_viewHeight)
	ebitenmobileview.Layout(_param_viewWidth, _param_viewHeight)
}

//export proxyebitenmobileview__Resume
func proxyebitenmobileview__Resume() {
	ebitenmobileview.Resume()
}

// skipped function SetGame with unsupported parameter or result types
//export proxyebitenmobileview__SetUIView
func proxyebitenmobileview__SetUIView(param_uiview C.int64_t) {
	_param_uiview := int64(param_uiview)
	ebitenmobileview.SetUIView(_param_uiview)
}

//export proxyebitenmobileview__Suspend
func proxyebitenmobileview__Suspend() {
	ebitenmobileview.Suspend()
}

//export proxyebitenmobileview__Update
func proxyebitenmobileview__Update() C.int32_t {
	_ = ebitenmobileview.Update()
	//var _res_0 C.int32_t = _seq.NullRefNum
	//if res_0 != nil {
	//    _res_0 = C.int32_t(_seq.ToRefNum(res_0))
	//}
	return 0
}

//export proxyebitenmobileview__UpdateTouchesOnAndroid
func proxyebitenmobileview__UpdateTouchesOnAndroid(param_action C.int, param_id C.int, param_x C.int, param_y C.int) {
	_param_action := int(param_action)
	_param_id := int(param_id)
	_param_x := int(param_x)
	_param_y := int(param_y)
	ebitenmobileview.UpdateTouchesOnAndroid(_param_action, _param_id, _param_x, _param_y)
}

//export proxyebitenmobileview__UpdateTouchesOnIOS
func proxyebitenmobileview__UpdateTouchesOnIOS(param_phase C.int, param_ptr C.int64_t, param_x C.int, param_y C.int) {
	_param_phase := int(param_phase)
	_param_ptr := int64(param_ptr)
	_param_x := int(param_x)
	_param_y := int(param_y)
	ebitenmobileview.UpdateTouchesOnIOS(_param_phase, _param_ptr, _param_x, _param_y)
}
