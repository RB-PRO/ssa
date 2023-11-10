function chss(pw)
    len   = length(pw); % количество отсчетов pw
    cad = 30;      % 30 кадров/сек
%     cad=1427/60;
    dt  = 1.0/cad; % интервал дискретизации времени, сек
    tim=zeros(len,1);
    tim(1) = 0.0;
    for i=2:len
       tim(i) = tim(i-1)+dt; % время в секундах
    end

    nPart = 20; % количество долей res
    win = 1024;
    fMi  = 40.0/60.0;   % частота среза для 40 уд/мин (0.6667 Гц)
    fMa  = 240.0/60.0;  % частота среза для 240 уд/мин (4.0 Гц)
%     Nf   = 1+win/2;     % кол-во отсчетов частоты
%     df   = cad/(win-1); % интервал дискретизации частоты, Гц 



    %%%%%%%%%%%%%%%%%%%%%%%%


    % Матрица перекрывающихся сегментов pw
    col = 1; Imin = 1; Imax = win;
    res = len-win*floor(len/win); res = floor(res/nPart);
    LenArray = round((len-Imax)/res,-1);
    ns = zeros(LenArray, 1);
    tseg = zeros(LenArray, 1);
    while Imax<=len
       ns(col) = col; % номер текущего сегмента pw
       Imin    = Imin+res;
       Imid  = floor(Imin+res/2);
       Imax    = Imax+res;
       tseg(col) = tim(Imid);
       col     = col+1;
    end
    col = col-1; % кол-во перекрывающихся сегментов в пределах len
%     NSF = win+res*(col-1); % номер финального отсчета финального сегмента <=len
    
    spw=zeros(win, col);
    for j=1:col
       for i=1:win
          k        = (j-1)*res;
          spw(i,j) = pw(k+i); % текущий сегмент pw длинною win 
       end
    end

    % Оценки СПМ перекрывающихся сегменов pw
    f(1) = 0.0; df = cad/(win-1); % интервал дискретизации частоты, Гц
    row  = 1+win/2;
    Fmin = fMi-10*df; Fmax = fMa+10*df; % частота в Гц
    for i=2:row
       f(i) = f(i-1)+df; % частота в герцах
       if abs(f(i)-Fmin)<=df
          iGmin = i;
       end
       if abs(f(i)-Fmax)<=df
          iGmax = i;
       end
    end
    for i=1:iGmax
       fG(i) = f(i); % сетка частот 3D-графика
    end
    f = f'; smopto = 3; % параметр сглаживания периодограммы Томсона
    figure();
    zpg_spw=zeros(iGmax, col);
    zto_spw=zeros(iGmax, col);
    for j=1:col
       pg_spw(:,j) = periodogram(spw(:,j),blackmanharris(win),win);
       stem(f(iGmin:iGmax),pg_spw(iGmin:iGmax,j),'LineStyle','-','Marker','none'); hold on;
       pto_spw(:,j) = pmtm(spw(:,j),smopto,win);   % периодограмма Томсона
       for i=1:iGmax
          zpg_spw(i,j) = pg_spw(i,j);
          zto_spw(i,j) = pto_spw(i,j);
       end
    end
    grid on; title('СПМ spw');
    xlabel("f,Hz",'interp','none'); ylabel("Psd",'interp','none');
    
    
ppto=load('Files/pto.txt')';
figure(); 
subplot(1,2, 1); imagesc(ppto);
subplot(1,2, 2); imagesc(pto_spw);
% pto_spw=ppto;


    % Оценки средних частот основного тона для сегментов pw
    for j=1:col
        [~,I] = sort(pg_spw(:,j),'descend');
        pg_fMAX(j) = f(I(1)); % I(1) - индекс частоты(Гц) максимума pg_spw(:,j)   
        [~,I] = sort(pto_spw(:,j),'descend');
        pto_fMAX(j) = f(I(1));
     end
     pto_fMAX = pto_fMAX';
     smo_pto_fMAX = smooth(pto_fMAX,0.3*col,'rloess'); 
%       EUT_P1LC3_Fmaxcr = smo_pto_fMAX; 
%       save('EUT_P1LC3_Fmaxcr','EUT_P1LC3_Fmaxcr');

     figure();
    % plot(ns,pg_fMAX,'b','Marker','o'); hold on;
     plot(pto_fMAX,'g'); hold on; % tseg
     plot(smo_pto_fMAX,'k','LineWidth',0.8);
     grid off; title('Частоты основного тона периодограммы Томсона spw ');
     xlabel("t, s",'interp','none'); ylabel("fMAX,Гц",'interp','none');
end