function pw = rgb2pw2(pwc, VideoFile) 

cad = 30;         % 30 кадров/сек
lag = 100;        % наибольший лаг АКФ
pi2 = 2.0*pi;
fMi = 40.0/60.0;  % частота среза для 40 уд/мин (0.6667 Гц)
fMa = 240.0/60.0; % частота среза для 240 уд/мин (4.0 Гц)
lvl = 10;         % количество уровней разложения КМА 
wavname = 'db40'; % модель вейвлет-функций
%
dt   = 1.0/cad;          % интервал дискретизации времени, сек
len  = length(pwc(:,2)); % количество отсчетов G-сигнала
dfc  = cad/(len-1);      % интервал дискретизации частоты, Гц
lMax = 0.5*len;          % половина от количества отсчетов сигнала
tim(1) = 0.0; fa(1) = 0.0; fqc(1) = 0.0;
for i=2:len
        tim(i) = tim(i-1)+dt; % время в секундах
end
SfMi(1) = 0.0; CfMi(1) = 1.0; SfMa(1) = 0.0; CfMa(1) = 1.0;

alg = 'Cr';
pw = (112.0*pwc(:,1)-93.8*pwc(:,2)-18.2*pwc(:,3))./255.0;
SMO_med = floor(cad/fMi);
% % % Алгоритмы G, GR, Cr
DEV_med = medfilt1(pw.*pw,SMO_med); 
STD_med(:,1) = sqrt(DEV_med);
pw = pw./STD_med(:,1);



    file=fopen(NameVideoFile(VideoFile)+'_pw.txt','w'); 
    fprintf(file,'%f\n',pw);  
    fclose(file); 
end