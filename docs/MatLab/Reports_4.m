clear; close all; clc;

    %% P1H1
    Name="P1H1";
%     Name="P1M1";
    [RGB,RGB_nc,RGB_med] = Rendes_RGB(Name);
    %% �������� ������
    Rendes_chss(RGB_nc, Name);
%     %% ����������� �����
%     Rendes_chss(RGB, Name);
%     %% ����������� ����� � ��������� �����������
%     Rendes_chss (RGB_med, Name);
    
    
%%
% close all;