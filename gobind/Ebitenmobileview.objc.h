// Objective-C API for talking to github.com/hajimehoshi/ebiten/mobile/ebitenmobileview Go package.
//   gobind -lang=objc github.com/hajimehoshi/ebiten/mobile/ebitenmobileview
//
// File is generated by gobind. Do not edit.

#ifndef __Ebitenmobileview_H__
#define __Ebitenmobileview_H__

@import Foundation;
#include "ref.h"
#include "Universe.objc.h"


FOUNDATION_EXPORT void EbitenmobileviewLayout(double viewWidth, double viewHeight);

FOUNDATION_EXPORT void EbitenmobileviewResume(void);

// skipped function SetGame with unsupported parameter or return types


FOUNDATION_EXPORT void EbitenmobileviewSetUIView(int64_t uiview);

FOUNDATION_EXPORT void EbitenmobileviewSuspend(void);

FOUNDATION_EXPORT BOOL EbitenmobileviewUpdate(NSError* _Nullable* _Nullable error);

/**
 * UpdateTouchesOnAndroid is a dummy function for backward compatibility.
UpdateTouchesOnAndroid is called from ebiten/mobile package.
 */
FOUNDATION_EXPORT void EbitenmobileviewUpdateTouchesOnAndroid(long action, long id_, long x, long y);

FOUNDATION_EXPORT void EbitenmobileviewUpdateTouchesOnIOS(long phase, int64_t ptr, long x, long y);

#endif