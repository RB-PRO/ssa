function pw = rgb2pw2(pwc, VideoFile) 

cad = 30;         % 30 ������/���
lag = 100;        % ���������� ��� ���
pi2 = 2.0*pi;
fMi = 40.0/60.0;  % ������� ����� ��� 40 ��/��� (0.6667 ��)
fMa = 240.0/60.0; % ������� ����� ��� 240 ��/��� (4.0 ��)
lvl = 10;         % ���������� ������� ���������� ��� 
wavname = 'db40'; % ������ �������-�������
%
dt   = 1.0/cad;          % �������� ������������� �������, ���
len  = length(pwc(:,2)); % ���������� �������� G-�������
dfc  = cad/(len-1);      % �������� ������������� �������, ��
lMax = 0.5*len;          % �������� �� ���������� �������� �������
tim(1) = 0.0; fa(1) = 0.0; fqc(1) = 0.0;
for i=2:len
        tim(i) = tim(i-1)+dt; % ����� � ��������
end
SfMi(1) = 0.0; CfMi(1) = 1.0; SfMa(1) = 0.0; CfMa(1) = 1.0;

alg = 'Cr';
pw = (112.0*pwc(:,1)-93.8*pwc(:,2)-18.2*pwc(:,3))./255.0;
SMO_med = floor(cad/fMi);
% % % ��������� G, GR, Cr
DEV_med = medfilt1(pw.*pw,SMO_med); 
STD_med(:,1) = sqrt(DEV_med);
pw = pw./STD_med(:,1);



    file=fopen(NameVideoFile(VideoFile)+'_pw.txt','w'); 
    fprintf(file,'%f\n',pw);  
    fclose(file); 
end