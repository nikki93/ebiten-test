#!/bin/sh

pushd mobile
time go run github.com/nikki93/gomob -x -target ios -iosheader ios.h -o ../ios/Mobile.framework .
popd

pushd ios
xcodebuild
popd

ios-deploy -L -b ios/build/Release-iphoneos/*.app

