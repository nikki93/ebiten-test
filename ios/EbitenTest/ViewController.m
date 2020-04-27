//
//  ViewController.m
//  EbitenTest
//
//  Created by Nikhilesh Sigatapu on 4/22/20.
//  Copyright © 2020 Nikhilesh Sigatapu. All rights reserved.
//

#import "ViewController.h"

@import Mobile;


#import <UIKit/UIKit.h>

@interface MobileEbitenViewController : UIViewController

// onErrorOnGameUpdate is called on the main thread when an error happens when updating a game.
// You can define your own error handler, e.g., using Crashlytics, by overwriting this method.
- (void)onErrorOnGameUpdate:(NSError*)err;

// suspendGame suspends the game.
// It is recommended to call this when the application is being suspended e.g.,
// UIApplicationDelegate's applicationWillResignActive is called.
- (void)suspendGame;

// resumeGame resumes the game.
// It is recommended to call this when the application is being resumed e.g.,
// UIApplicationDelegate's applicationDidBecomeActive is called.
- (void)resumeGame;

@end


@interface ViewController ()

@property (nonatomic, strong) MobileEbitenViewController *ebitenVC;

@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    
    self.ebitenVC = [[MobileEbitenViewController alloc] init];
    [self addChildViewController:self.ebitenVC];
    self.ebitenVC.view.frame = self.view.frame;
    [self.view addSubview:self.ebitenVC.view];
    [self.ebitenVC didMoveToParentViewController:self];
}


@end
