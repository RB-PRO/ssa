clear; close all; clc;

dt=1/30;     % �������� ��������� �������������
fmin=0.15; % ������ ������� - ��������� �������� ����������� �����
dt=1/30;     % �������� ��������� �������������
Nmed=1/(dt*fmin); % �������� �������

% Names
names = [ "P1H1" "P1LC1" "P1LC2" "P1LC3" "P1LC4" "P1LC5" "P1LC6" "P1LC7" "P1M1" "P1M2" "P1M3" "P2LC1" "P2LC2" "P2LC3" "P2LC4" "P2LC5" "P3LC1" "P3LC2" "P3LC3" "P3LC4" "P3LC5" ];
% names = [ "P1H1" "P1LC1" ];
  names = ["P1H1"];
zrp="endh/"; % Zero Folder

%%% ����� �� ���������:
for iName = 1:length(names)
    Name = names(iName);
    Path=zrp+Name+'/';
    
    try
        RGB=load(Path+Name+"_RGB.txt");
        RGB_nc=load(Path+Name+"_nc_RGB.txt");
        nc=0:dt:length(RGB)*dt-dt;
    catch
        disp("ERROR: ������ ��� ������������ ������ �� "+Name);
        continue
    end

        %% ����� RGB
        disp(Name);
        
        figure('Position', [0 0 400 100]); title(Name,'FontSize',36);
    
        figure('Name','RGB','Position', [0 0 1800 900]); set(gcf,'name',"��������� ��������� ���������� ��� "+Name); clf;
        subplot(2,2,1); plot(nc, RGB(:,1),"red", nc, RGB(:,2),"green", nc, RGB(:,3),"blue");
        title("��������� ��� RGB �� ����������� �����"); grid on;
        xlabel("�������"); ylabel("������������� �������� �������"); % ylim([40;105]);

        subplot(2,2,2); plot(nc, RGB_nc(:,1),"red", nc, RGB_nc(:,2),"green", nc, RGB_nc(:,3),"blue");
        title("��������� ��� RGB ����� ����������� �����"); grid on;
        xlabel("�������"); ylabel("������������� �������� �������"); % ylim([40;105]);

        subplot(2,2,[3,4]);
        RGB_med=RGB_nc;
        RGB_med(:,1)=RGB_med(:,1)-medfilt1(RGB_med(:,1),Nmed);
        RGB_med(:,2)=RGB_med(:,2)-medfilt1(RGB_med(:,2),Nmed);
        RGB_med(:,3)=RGB_med(:,3)-medfilt1(RGB_med(:,3),Nmed);
        plot(nc, RGB_med(:,1),"red", nc, RGB_med(:,2),"green", nc, RGB_med(:,3),"blue");
        title("��������� ��� RGB ����� ����������� ����� � ���������� ������� � ����� N_m_e_d="+Nmed);
        xlabel("�������"); ylabel("������������� �������� �������"); grid on; ylim([-4;4]);
        
        %% �������� ������
        try
            chss2(rgb2pw(RGB, "Cr"), Path, Name);
        catch
            disp("ERROR: ������ ��� ������������ ������ �� �������� ������");
        end
        CloseFigure
        
        %% ����������� �������
        try
            chss2(rgb2pw(RGB_nc, "Cr"), Path, Name);
        catch
            disp("ERROR: ������ ��� ������������ ������ �� ������ � ������������ �����");
        end
        CloseFigure
        
        %% ����������� ������� + ��������� ����������
        try
            chss2(rgb2pw(RGB_med, "Cr"), Path, Name);
        catch
            disp("ERROR: ������ ��� ������������ ������ �� ������ � ������������ ����� � ��������� ����������");
        end
        CloseFigure
end
close all;
