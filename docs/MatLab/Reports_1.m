clear; close all; clc;

dt=1/30;     % интервал временной дискретизации
fmin=0.15; % нижняя граница - частотный диапазон дыхательной волны
dt=1/30;     % интервал временной дискретизации
Nmed=1/(dt*fmin); % апертура фильтра

% Names
names = [ "P1H1" "P1LC1" "P1LC2" "P1LC3" "P1LC4" "P1LC5" "P1LC6" "P1LC7" "P1M1" "P1M2" "P1M3" "P2LC1" "P2LC2" "P2LC3" "P2LC4" "P2LC5" "P3LC1" "P3LC2" "P3LC3" "P3LC4" "P3LC5" ];
% names = [ "P1H1" "P1LC1" ];
  names = ["P1H1"];
zrp="endh/"; % Zero Folder

%%% Отчёт по видеоряду:
for iName = 1:length(names)
    Name = names(iName);
    Path=zrp+Name+'/';
    
    try
        RGB=load(Path+Name+"_RGB.txt");
        RGB_nc=load(Path+Name+"_nc_RGB.txt");
        nc=0:dt:length(RGB)*dt-dt;
    catch
        disp("ERROR: Ошибка при формировании отчёта по "+Name);
        continue
    end

        %% Вывод RGB
        disp(Name);
        
        figure('Position', [0 0 400 100]); title(Name,'FontSize',36);
    
        figure('Name','RGB','Position', [0 0 1800 900]); set(gcf,'name',"Сравнение медианной фильтрации для "+Name); clf;
        subplot(2,2,1); plot(nc, RGB(:,1),"red", nc, RGB(:,2),"green", nc, RGB(:,3),"blue");
        title("Временной ряд RGB до компенсации цвета"); grid on;
        xlabel("Секунды"); ylabel("Интенсивность цветовых каналов"); % ylim([40;105]);

        subplot(2,2,2); plot(nc, RGB_nc(:,1),"red", nc, RGB_nc(:,2),"green", nc, RGB_nc(:,3),"blue");
        title("Временной ряд RGB после компенсации цвета"); grid on;
        xlabel("Секунды"); ylabel("Интенсивность цветовых каналов"); % ylim([40;105]);

        subplot(2,2,[3,4]);
        RGB_med=RGB_nc;
        RGB_med(:,1)=RGB_med(:,1)-medfilt1(RGB_med(:,1),Nmed);
        RGB_med(:,2)=RGB_med(:,2)-medfilt1(RGB_med(:,2),Nmed);
        RGB_med(:,3)=RGB_med(:,3)-medfilt1(RGB_med(:,3),Nmed);
        plot(nc, RGB_med(:,1),"red", nc, RGB_med(:,2),"green", nc, RGB_med(:,3),"blue");
        title("Временной ряд RGB после компенсации цвета и медианного фильтра с окном N_m_e_d="+Nmed);
        xlabel("Секунды"); ylabel("Интенсивность цветовых каналов"); grid on; ylim([-4;4]);
        
        %% Исходные данные
        try
            chss2(rgb2pw(RGB, "Cr"), Path, Name);
        catch
            disp("ERROR: Ошибка при формировании отчёта из исходных данных");
        end
        CloseFigure
        
        %% Компенсация яркости
        try
            chss2(rgb2pw(RGB_nc, "Cr"), Path, Name);
        catch
            disp("ERROR: Ошибка при формировании отчёта из данных с компенсацией цвета");
        end
        CloseFigure
        
        %% Компенсация яркости + медианная фильтрация
        try
            chss2(rgb2pw(RGB_med, "Cr"), Path, Name);
        catch
            disp("ERROR: Ошибка при формировании отчёта из данных с компенсацией цвета и медианной фильтрации");
        end
        CloseFigure
end
close all;
