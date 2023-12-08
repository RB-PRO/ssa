clear; close all; clc;
% P1H1
% P1H1_noComp
% P2LC1
% P2LC1_noComp
VideoFile = "P2LC1.avi"; Path="Files/"; % P1H1_edited

% Extracting a time series of RGB segments
% Face_tracking(Path+VideoFile);
RGB=load(Path+NameVideoFile(VideoFile)+'_RGB.txt');

% Calculation of the photoplethysmography signal
rgb2pw(RGB, Path+VideoFile);
pw=load(Path+NameVideoFile(VideoFile)+'_pw.txt');


% Estimates of the average pitch frequencies for PW segments
chss2(pw);