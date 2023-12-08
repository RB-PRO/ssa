function RGB = Median1RGB(RGB)
    fmin=0.15; % нижн€€ граница - частотный диапазон дыхательной волны
    dt=1/30;     % интервал временной дискретизации
    Nmed=1/(dt*fmin); % апертура фильтра
    nc=0:dt:length(RGB)*dt-dt;
    disp("ќкно медианной фильтрации: "+Nmed);
    
    figure('Name','RGB','Position', [0 0 600 300]); set(gcf,'name',"—равнение медианной фильтрации"); clf;
    subplot(1,2,1);
    plot(nc, RGB(:,1),"red", nc, RGB(:,2),"green", nc, RGB(:,3),"blue");
    title("RGB до медианной фильтрации"); xlabel("seconds"); grid on;
    
    % ћедианна€ фильтраци€
    RGB(:,1)=medfilt1(RGB(:,1),Nmed);
    RGB(:,2)=medfilt1(RGB(:,2),Nmed);
    RGB(:,3)=medfilt1(RGB(:,3),Nmed);
    
    subplot(1,2,2);
    plot(nc, RGB(:,1),"red", nc, RGB(:,2),"green", nc, RGB(:,3),"blue");
    title("RGB после медианной фильтрации"); xlabel("seconds"); grid on;
end