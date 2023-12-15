clear; close all; clc;

exp="P3LC5";
folder="endh/"+exp+"/";
RGB=load(folder+exp+"_RGB.txt");
RGB_nc=load(folder+exp+"_nc_RGB.txt");

RGB(1:30,:)=[];
RGB_nc(1:30,:)=[];

dt=1/30;     % �������� ��������� �������������
nc=0:dt:length(RGB)*dt-dt;
fmin=0.15; % ������ ������� - ��������� �������� ����������� �����
dt=1/30;     % �������� ��������� �������������
Nmed=1/(dt*fmin); % �������� �������

figure('Name','RGB','Position', [0 0 600 300]); set(gcf,'name',"��������� ��������� ����������"); clf;
subplot(2,2,1);
plot(nc, RGB(:,1),"red", nc, RGB(:,2),"green", nc, RGB(:,3),"blue");
title("��������� ��� RGB �� ����������� �����"); grid on;
ylim([40;105]); xlabel("�������"); ylabel("������������� �������� �������");

subplot(2,2,2);
plot(nc, RGB_nc(:,1),"red", nc, RGB_nc(:,2),"green", nc, RGB_nc(:,3),"blue");
title("��������� ��� RGB ����� ����������� �����"); grid on;
ylim([40;105]); xlabel("�������"); ylabel("������������� �������� �������");


subplot(2,2,[3,4]);
RGB_med=RGB_nc;
RGB_med(:,1)=RGB_med(:,1)-medfilt1(RGB_med(:,1),Nmed);
RGB_med(:,2)=RGB_med(:,2)-medfilt1(RGB_med(:,2),Nmed);
RGB_med(:,3)=RGB_med(:,3)-medfilt1(RGB_med(:,3),Nmed);
plot(nc, RGB_med(:,1),"red", nc, RGB_med(:,2),"green", nc, RGB_med(:,3),"blue");
title("��������� ��� RGB ����� ����������� ����� � ���������� ������� � ����� N_m_e_d="+Nmed); grid on;
ylim([-4;4]); xlabel("�������"); ylabel("������������� �������� �������");


