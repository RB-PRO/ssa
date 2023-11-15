clear; close all; clc;
VideoFile = "EUT_P1H1.txt"; Path="Files/"; % P1H1_edited

% Вычленение RGB из видео
% Face_tracking(Path+VideoFile);
% RGB=load(Path+NameVideoFile(VideoFile)+'_RGB.txt');

load(Path+'EUT_P1H1_RGB.mat');
RGB=EUT_P1H1_RGB; 
    file=fopen(strcat(Path+'EUT_P1H1_RGB.txt', '.txt'),'w'); 
    for i=1:length(RGB)
         fprintf(file,'%f;%f;%f\n',RGB(i, 1), RGB(i, 2), RGB(i, 3));
    end
    fclose(file);

% Фильтрация
%    ButterRGB(RGB, Path+VideoFile);
%    RGB2=load(Path+NameVideoFile(VideoFile)+'_but.txt'); 

% Получение pw
rgb2pw(RGB, Path+VideoFile);
pw=load(Path+NameVideoFile(VideoFile)+'_pw.txt');

load(Path+'EUT_P1H1_pwCr.mat');
pwOriginal=EUT_P1H1_pwCr;
% pw_smooth = smoothdata(pw,"movmean",5);  
pw_smooth = movmean(pw,5);
 
% figure(); plot(pwOriginal, '*k'); hold on; plot(pw, '--blue'); plot(pwOriginal, 'red');
% legend('Оптимальное pw','pw после CR','pw после с.к.'); 
% title('Сравнение сигнала фотоплетизмографии с оригиналом');
% (2400:2600)
figure(); plot(pwOriginal(2400:2600), '*k'); hold on; plot(pw(2400:2600), '--blue'); plot(pwOriginal(2400:2600), 'red');
legend('Оптимальное pw','pw после CR','pw после с.к.'); 
title('Сравнение сигнала фотоплетизмографии с оригиналом');

% Оценки средних частот основного тона для сегментов pw
chss2(pw_smooth);