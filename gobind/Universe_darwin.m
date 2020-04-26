// Objective-C API for talking to  Go package.
//   gobind -lang=objc 
//
// File is generated by gobind. Do not edit.

#include <Foundation/Foundation.h>
#include "seq.h"
#include "_cgo_export.h"
#include "Universe.objc.h"

@implementation Universeerror {
}

- (nonnull instancetype)initWithRef:(id)ref {
	if (self) {
		__ref = ref;
		self = [super initWithDomain:@"go" code:1 userInfo:@{NSLocalizedDescriptionKey: [self error]}];
	}
	return self;
}

- (NSString* _Nonnull)error {
	int32_t refnum = go_seq_go_to_refnum(self._ref);
	nstring r0 = proxy_error_Error(refnum);
	NSString *_ret0_ = go_seq_to_objc_string(r0);
	return _ret0_;
}

@end



nstring cproxy_error_Error(int32_t refnum) {
	@autoreleasepool {
		Universeerror* o = go_seq_objc_from_refnum(refnum);
		NSString* _Nonnull ret0_;
		ret0_ = [o localizedDescription];
		nstring _ret0_ = go_seq_from_objc_string(ret0_);
		return _ret0_;
	}
}

__attribute__((constructor)) static void init() {
	init_seq();
}
