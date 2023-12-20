function [RGB,RGB_nc,RGB_med] = Rendes_RGB(Name)
    fmin=0.15; % ������ ������� - ��������� �������� ����������� �����
    dt=1/30;     % �������� ��������� �������������
    Nmed=1/(dt*fmin); % �������� �������
    zrp="endh/"; % Zero Folder
    Path=zrp+Name+'/';
%   P1H1_����������������������_RGB
%   P1H1_��������������������_RGB
    try
        RGB=load(Path+Name+"_��������������������_RGB.txt");
        RGB_nc=load(Path+Name+"_����������������������_RGB.txt");
        nc=0:dt:length(RGB)*dt-dt;
    catch
        disp("ERROR: ������ ��� ������������ ������ �� "+Name);
        RGB=[];RGB_nc=[];RGB_med=[];
        return
    end
    
    figure('Name','RGB','Position', [0 0 1400 900]); set(gcf,'name',"��������� ��������� ���������� ��� "+Name); clf;
    subplot(2,2,1); plot(nc, RGB_nc(:,1),"red", nc, RGB_nc(:,2),"green", nc, RGB_nc(:,3),"blue");
    title("��������� ��� RGB �� ����������� �����"); grid on;
    xlabel("�������"); ylabel("������������� �������� �������"); % ylim([40;105]);

    subplot(2,2,2); plot(nc, RGB(:,1),"red", nc, RGB(:,2),"green", nc, RGB(:,3),"blue");
    title("��������� ��� RGB ����� ����������� �����"); grid on;
    xlabel("�������"); ylabel("������������� �������� �������"); % ylim([40;105]);

    subplot(2,2,[3,4]);
    RGB_med=RGB;
    RGB_med(:,1)=RGB_med(:,1)-medfilt1(RGB_med(:,1),Nmed);
    RGB_med(:,2)=RGB_med(:,2)-medfilt1(RGB_med(:,2),Nmed);
    RGB_med(:,3)=RGB_med(:,3)-medfilt1(RGB_med(:,3),Nmed);
    plot(nc, RGB_med(:,1),"red", nc, RGB_med(:,2),"green", nc, RGB_med(:,3),"blue");
    title("��������� ��� RGB ����� ����������� ����� � ���������� ������� � ����� N_m_e_d="+Nmed);
    xlabel("�������"); ylabel("������������� �������� �������"); grid on; ylim([-4;4]);
end
