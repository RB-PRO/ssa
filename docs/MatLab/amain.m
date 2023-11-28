clear; close all; clc;
VideoFile = "P2LC1_edited.avi"; Path="Files/"; % P1H1_edited

% Extracting a time series of RGB segments
 Face_tracking(Path+VideoFile);
RGB=load(Path+NameVideoFile(VideoFile)+'_RGB.txt');

%   Filter
   ButterRGB(RGB, Path+VideoFile);
   RGB2=load(Path+NameVideoFile(VideoFile)+'_but.txt'); 

% Calculation of the photoplethysmography signal
rgb2pw(RGB, Path+VideoFile);
pw=load(Path+NameVideoFile(VideoFile)+'_pw.txt');


% Estimates of the average pitch frequencies for PW segments
chss2(pw);