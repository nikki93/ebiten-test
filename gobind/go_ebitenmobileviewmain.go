// Code generated by gobind. DO NOT EDIT.

// Package main is an autogenerated binder stub for package ebitenmobileview.
//
//   autogenerated by gobind -lang=go github.com/hajimehoshi/ebiten/mobile/ebitenmobileview
package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "seq.h"
#include "ebitenmobileview.h"

*/
import "C"

import (
	"github.com/hajimehoshi/ebiten/mobile/ebitenmobileview"
	_seq "golang.org/x/mobile/bind/seq"
)

// suppress the error if seq ends up unused
var _ = _seq.FromRefNum

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
	res_0 := ebitenmobileview.Update()
	var _res_0 C.int32_t = _seq.NullRefNum
	if res_0 != nil {
		_res_0 = C.int32_t(_seq.ToRefNum(res_0))
	}
	return _res_0
}

//export proxyebitenmobileview__UpdateTouchesOnAndroid
func proxyebitenmobileview__UpdateTouchesOnAndroid(param_action C.nint, param_id C.nint, param_x C.nint, param_y C.nint) {
	_param_action := int(param_action)
	_param_id := int(param_id)
	_param_x := int(param_x)
	_param_y := int(param_y)
	ebitenmobileview.UpdateTouchesOnAndroid(_param_action, _param_id, _param_x, _param_y)
}

//export proxyebitenmobileview__UpdateTouchesOnIOS
func proxyebitenmobileview__UpdateTouchesOnIOS(param_phase C.nint, param_ptr C.int64_t, param_x C.nint, param_y C.nint) {
	_param_phase := int(param_phase)
	_param_ptr := int64(param_ptr)
	_param_x := int(param_x)
	_param_y := int(param_y)
	ebitenmobileview.UpdateTouchesOnIOS(_param_phase, _param_ptr, _param_x, _param_y)
}
