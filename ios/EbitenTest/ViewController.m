//
//  ViewController.m
//  EbitenTest
//
//  Created by Nikhilesh Sigatapu on 4/22/20.
//  Copyright © 2020 Nikhilesh Sigatapu. All rights reserved.
//

#import "ViewController.h"
#import <Mobile/MobileEbitenViewController.h>


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
