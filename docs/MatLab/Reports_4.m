clear; close all; clc;

    %% P1H1
    Name="P1H1";
%     Name="P1M1";
    [RGB,RGB_nc,RGB_med] = Rendes_RGB(Name);
    %% Исходный сигнал
    Rendes_chss(RGB_nc, Name);
%     %% Компенсация цвета
%     Rendes_chss(RGB, Name);
%     %% Компенсация цвета с медианной фильтрацией
%     Rendes_chss (RGB_med, Name);
    
    
%%
% close all;