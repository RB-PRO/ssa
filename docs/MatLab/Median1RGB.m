function RGB = Median1RGB(RGB)
    fmin=0.15; % ������ ������� - ��������� �������� ����������� �����
    dt=1/30;     % �������� ��������� �������������
    Nmed=1/(dt*fmin); % �������� �������
    nc=0:dt:length(RGB)*dt-dt;
    disp("���� ��������� ����������: "+Nmed);
    
    figure('Name','RGB','Position', [0 0 600 300]); set(gcf,'name',"��������� ��������� ����������"); clf;
    subplot(1,2,1);
    plot(nc, RGB(:,1),"red", nc, RGB(:,2),"green", nc, RGB(:,3),"blue");
    title("RGB �� ��������� ����������"); xlabel("seconds"); grid on;
    
    % ��������� ����������
    RGB(:,1)=medfilt1(RGB(:,1),Nmed);
    RGB(:,2)=medfilt1(RGB(:,2),Nmed);
    RGB(:,3)=medfilt1(RGB(:,3),Nmed);
    
    subplot(1,2,2);
    plot(nc, RGB(:,1),"red", nc, RGB(:,2),"green", nc, RGB(:,3),"blue");
    title("RGB ����� ��������� ����������"); xlabel("seconds"); grid on;
end