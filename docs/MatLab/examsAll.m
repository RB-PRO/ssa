clear; close all; clc;
% Names
names = ["P1H1"    "P1LC1" "P1LC2" "P1LC3" "P1LC4" "P1LC5" "P1LC6" "P1LC7"     "P1M1" "P1M2" "P1M3"     "P2LC1" "P2LC2" "P2LC3" "P2LC4" "P2LC5"     "P3LC1" "P3LC2" "P3LC3" "P3LC4" "P3LC5"];
% names = ["P1H1" "P1LC1"];
zrp="endh/"; % Zero Folder

%%% Отчёт по видеоряду:
for iName = 1:length(names)
    Name = names(iName);
    Path=zrp+Name+'/';
    
    %% С цветовой компенсацией
    try
        disp(Name+" с цветовой компенсацией"); ncPrefix="";
        RGB=load(Path+Name+ncPrefix+'_RGB.txt');
        RGBmed=Median1RGB(RGB);
        pw=rgb2pw(RGBmed, "Cr");
        Save(pw,Path+Name+ncPrefix+'_pw.txt');
        chss2(pw, Path, Name+ncPrefix);
        CloseFigure
    catch
        disp("ERROR: Ошибка при формировании отчёта по "+Name+" с цветовой компенсацией");
    end
    
    %% Без цветовой компенсацией
    try
        disp(Name+" без цветовой компенсацией"); ncPrefix="_nc";
        RGB=load(Path+Name+ncPrefix+'_RGB.txt');
        RGBmed=Median1RGB(RGB);
        pw=rgb2pw(RGBmed, "Cr");
        Save(pw,Path+Name+ncPrefix+'_pw.txt');
        chss2(pw, Path, Name+ncPrefix);
        CloseFigure
    catch
        disp("ERROR: Ошибка при формировании отчёта по "+Name+" без цветовой компенсацией");
    end
end
close all;