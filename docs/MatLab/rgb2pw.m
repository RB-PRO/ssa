function pw = rgb2pw(pwc, VideoFile) 
    fMi  = 40.0/60.0;   % частота среза для 40 уд/мин (0.6667 Гц)
    cad = 30;      % 30 кадров/сек 
    dt  = 1.0/cad; % интервал дискретизации времени, сек
    len=length(pwc(:,1));
    tim(1) = 0.0;
    for i=2:len
        tim(i) = tim(i-1)+dt; % время в секундах
    end
    
    % Информативные сигналы для формирования пульсовых волн по цветовым каналам
    % Цветовые каналы: 1->R; 2->G; 3->B
    %  Алгоритм G
    %  alg = 'G';
% pw = pwc(:,2);
    % Алгоритм GR
    % alg = 'RG';
    % pw = pwc(:,2)-pwc(:,1);
    % Алгоритм Cr
% alg = 'Cr';
    pw = (112.0*pwc(:,1)-93.8*pwc(:,2)-18.2*pwc(:,3))./255.0;

%     Вычитаем тренд
%     pw_smooth = smoothdata(pw,"movmean",32);  
    pw_smooth = movmean(pw,32);
%     smoot_pw=smooth(pw); 
    pw=pw-pw_smooth;
    
    % Алгоритм CHROM
%     alg = 'CHROM';
%     ws(:,1) = (3.0*pwc(:,1)-2.0*pwc(:,2))./sqrt(13.0);
%     ws(:,2) = (-1.5*pwc(:,1)-pwc(:,2)+1.5*pwc(:,3))./sqrt(11.0/2.0); % умножено на (-1) для pw=ws(:,1)+ws(:,2)
    % Алгоритм POS
%      alg = 'POS';
% ws(:,1) = (pwc(:,2)-pwc(:,3))./sqrt(2.0);
% ws(:,2) = (pwc(:,2)+pwc(:,3)-2.0*pwc(:,1))./sqrt(6.0);
    %% Центрированная и нормированная пульсовая волна pw
    SMO_med = floor(cad/fMi);
    % Алгоритмы G, GR, Cr
   pw2 = pw.*pw;
    DEV_med = medfilt1(pw2,SMO_med); 
    figure();plot(pw2);
    figure();plot(DEV_med);
    STD_med(:,1) = sqrt(DEV_med);
    pw = pw./STD_med(:,1);
    figure();plot(pw);
    % Алгоритмы CHOM, POS
% DEV_med = medfilt1(ws(:,1).*ws(:,1),SMO_med); STD_med(:,1) = sqrt(DEV_med);
% DEV_med = medfilt1(ws(:,2).*ws(:,2),SMO_med); STD_med(:,2) = sqrt(DEV_med);
% pw = (ws(:,1)./STD_med(:,1))+(ws(:,2)./STD_med(:,2));
% [iPer, D_time] = iPer(video, pw, alg);
% HRV (video, iPer, D_time, alg);

    % Анализ аномальных значений
    prcMi = prctile(pw,0.1); 
    prcMa = prctile(pw,99.9); % процентили по уровням 0.1% и 99.9%
    for i=1:len
       if pw(i)<prcMi
          pw(i) = prcMi;    
       end
       if pw(i)>prcMa
          pw(i) = prcMa;    
       end
    end
    STD = std(pw); 
    figure();plot(pw);
    pw = pw./STD; % нормированная pw
    
    pw = movmean(pw,5);
    
    figure();
    plot(tim,pw); grid on;
    title('Центрированная и нормированная pw');
    xlabel("t,s",'interp','none'); ylabel("pw",'interp','none');
    
    file=fopen(NameVideoFile(VideoFile)+'_pw.txt','w'); 
    fprintf(file,'%f\n',pw);  
    fclose(file);
end