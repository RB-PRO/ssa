clear; close all; clc;
VideoFile = "EUT_P1H1_RGB.txt"; Path="Files/"; % P1H1_edited

% Вычленение RGB из видео
% Face_tracking(Path+VideoFile);
RGB=load(Path+NameVideoFile(VideoFile)+'_RGB.txt');

load(Path+'EUT_P1H1_RGB.mat');
RGB=EUT_P1H1_RGB;

% Фильтрация
%   ButterRGB(RGB, Path+VideoFile);
%   RGB2=load(Path+NameVideoFile(VideoFile)+'_but.txt'); 

% Получение pw
rgb2pw(RGB, Path+VideoFile);
pw=load(Path+NameVideoFile(VideoFile)+'_pw.txt');
% pw=load(Path+'P1H1_edited_but_pw.txt');

% load(Path+'EUT_P1H1_pwCr.mat');
% pw=EUT_P1H1_pwCr;

% Оценки средних частот основного тона для сегментов pw
% chss(pw);
chss2(pw);
