#!/bin/sh

set -e

pushd mobile
go run github.com/nikki93/gomob -target ios -iosheader ios.h -o ../ios/Mobile.framework .
popd

pushd ios
xcodebuild | xcpretty
popd

ios-deploy -L -b ios/build/Release-iphoneos/*.app

