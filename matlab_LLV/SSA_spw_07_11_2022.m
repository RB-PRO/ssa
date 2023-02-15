clear; close all; clc
load('EUT_P1H1_pwCr');
pw = EUT_P1H1_pwCr; % ����������� ������ ��������� �����
load('EUT_P1H1_Fmaxcr');
fmp = EUT_P1H1_FmaxCr; % ������� ��������� ���� ��������� ��������� �����

%% C������� pw
N   = length(pw); % ���������� �������� pw
win = 1024;
res = N-win*floor(N/win);
nPart = 20; % ���������� ����� res
res = floor(res/nPart); overlap = (win-res)/win;
S = 1; Imin = 1; Imax = win;
while Imax<=N
   ns(S) = S; % ����� �������� �������� pw
   Imin  = Imin+res;
   Imax  = Imax+res;
   S     = S+1;
end
S = S-1; % ���-�� ��������������� ��������� pw � �������� N
NSF = win+res*(S-1); % ����� ���������� ������� ���������� �������� <= N
for j=1:S
   for i=1:win
      k        = (j-1)*res;
      spw(i,j) = pw(k+i); % ������� ������� pw ������� win 
   end
end
%% Set general parameters
cad = 30;      % 30 ������/���
dt  = 1.0/cad; % �������� ������������� �������, ���
tim(1) = 0.0;
for i=2:N
   tim(i) = tim(i-1)+dt; % ����� � ��������
end
% tim = tim';
ns  = (1:S)'; % ������ ��������� pw
%
for j=1:S % ���� �� ��������� pw
   L(j) = floor(cad/fmp(j)); % ���-�� �������� ��������� ���� pw
end
L = L';
K = 5; % ���-�� �������� ��� ��������� ��������
M = K*max(L); % �������� �������� � ����������� ������������
%
%% SSA- ������ ��������� pw
seg = 100; % ����� �������� pw ��� ������������
nET = 4;   % ���-�� ����������� ����� ��� ��������� pw
for j=1:S  % ���� �� ���������
%% SSA time series
%    M = K*L(j); % �������� �������� � ����������� ������������
   [C,LBD,RC] = SSA(win,M,spw(:,j),nET);
%% Estimation of the spw(:,j) reconstructed with the pair of RC   
   sET12(:,j) = sum(RC(:,1:2),2);   
   sET34(:,j) = sum(RC(:,3:4),2);   
%% Compare reconstruction and original time series
   if j==seg
      figure();
      set(gcf,'name','Covariance matrix');
      clf;
      imagesc(C);
      axis square;
      set(gca,'clim',[-1 1]);
      colorbar;
%
      figure();
      set(gcf,'name','Eigenvalues')
      clf;
      plot(LBD,'o-');
%      
      figure();
      set(gcf,'name','Original time series and reconstruction')
      clf;
      plot(tim(1:win),spw(:,j),'b-',tim(1:win),sET12(:,j),'r-');
      legend('Original','sET12');
      xlabel("t,s",'interp','none'); ylabel("sET",'interp','none');
%      
      figure();
      set(gcf,'name','Original time series and reconstruction')
      clf;
      plot(tim(1:win),spw(:,j),'b-',tim(1:win),sET34(:,j),'r-');
      legend('Original','sET34');
      xlabel("t,s",'interp','none'); ylabel("sET",'interp','none');
   end
end
%% ������ ��� ����������� ����� ��� ��������� pw
lag  = floor(win/10); % ���������� ��� ��� <= win/10
lagS = 2*lag;
for j=1:S
   Acf_sET12(:,j) = AcfMed(lagS,win,sET12(:,j)); % ������������� ��� j-�� ��������
%    Acf_sET12(:,j) = autocorr(sET12(:,j),'NumLags',lag); % ������������� ��� j-�� ��������
end
%% ������������ ��� ����������� ����� ��� ��������� pw
lgl = 1:lag; % ����� 3D-������� ���
Time(1) = 0.0;
for m=2:lag
   Time(m) = Time(m-1)+dt; % ����� � ��������
end
figure();
set(gcf,'name','��� ����������� ����� sET12 ��������� pw');
clf;
% mesh(ns,lgl,Acf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
% xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
mesh(ns,Time,Acf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
zlabel("Acf",'interp','none'); grid on;
%% ��������� �� �������� ��������� ���������� abs(acf_sET12)
power = 0.75; % �������� ������������ ��������������
for j=1:S % ���� �� ��������� ���
   absTS = abs(Acf_sET12(:,j));
   AT1   = absTS(1);
   AT2   = absTS(2);
   maxTS = zeros(lag,1); 
   maxTS(1) = AT1;
   maxN  = zeros(lag,1); 
   maxN(1)  = 1;
   Nmax = 1;
   for m=3:lag
      AT3 = absTS(m);
      if (AT1<=AT2)&&(AT2>=AT3)
         Nmax = Nmax+1; % ����� ���������� ���� ������������ (������� ����������)
         maxN(Nmax) = m-1; % ����� ���������� ��������� ��� ���� absTS
         maxTS(Nmax) = AT2; % ������ ���������� ���� ������������
      end
      AT1 = AT2;
      AT2 = AT3;
   end
   Nmax = Nmax+1; % ���������� ����� ������������
   maxN(Nmax)  = lag; % ����� ������� absTS ���������� ���� ������������
   maxTS(Nmax) = absTS(lag); % ������ absTS ���������� ���� ������������
   NumMax = maxN(1:Nmax); % ������ ���������� �� absTS
    % ������������ ��������� ���
    % 'pchip','cubic','v5cubic','makima','spline'
   EnvAcf_sET12(:,j) = interp1(NumMax,maxTS(1:Nmax),lgl,'pchip'); % pchip(NumMax,maxTS(1:Nmax),lgl);
   AcfNrm_sET12(:,j) = Acf_sET12(1:lag,j)./EnvAcf_sET12(:,j); % ������������� ���
end
figure();
set(gcf,'name','��������� ��� ����������� ����� sET12 ��������� pw');
clf;
% mesh(ns,lgl,EnvAcf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
% xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
mesh(ns,Time,EnvAcf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
zlabel("Env_Acf",'interp','none'); grid on;
figure();
set(gcf,'name','������������� ��� ����������� ����� sET12 ��������� pw');
clf;
% mesh(ns,lgl,AcfNrm_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
% xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
mesh(ns,Time,AcfNrm_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
zlabel("Acf_Nrm",'interp','none'); grid on;
%% ����������� ��������� � ����������� ������������  
for j=1:S % ���� �� ��������� ���
   [BC_EnvAcf_sET12(:,j),lambda(j)] = boxcox(EnvAcf_sET12(:,j)); % �������������� �����-�����
   lin_Env1 = -((-BC_EnvAcf_sET12(1,j)).^power); % ����������� ��������������
   der(1,j) = lin_Env1;
   for m=2:lag
      lin_Env2 = -((-BC_EnvAcf_sET12(m,j)).^power);
      der(m,j) = lin_Env2-lin_Env1;
      lin_Env1 = lin_Env2;
   end
% ������ ��������� � �������� ������������   
   med_der(j) = median(der(:,j)); % ������� ����������� ��������� ���
   linear1 = der(1,j); % ��������� ��������
   BC_Env1 = -(-linear1)^(1.0/power);
% �������� �������������� �����-�����
   if lambda(j)==0
      AprEnvAcf_sET12(1,j) = exp(BC_Env1);
   else
      AprEnvAcf_sET12(1,j) = (abs(1.0+lambda(j)*BC_Env1))^(1.0/lambda(j));
   end
   for m=2:lag
      linear2 = linear1+med_der(j); % �������� ������ � ����������� ������������
      BC_Env2 = -(-linear2)^(1.0/power);
      if lambda(j)==0
         AprEnvAcf_sET12(m,j) = exp(BC_Env2);
      else
         AprEnvAcf_sET12(m,j) = (abs(1.0+lambda(j)*BC_Env2))^(1.0/lambda(j));
      end
      linear1 = linear2;
   end
end
figure();
set(gcf,'name','����������� ��������� ��� ����������� ����� sET12 ��������� pw');
clf;
mesh(ns,lgl,der(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
figure();
set(gcf,'name','������������� ��������� ��� � �������� ������������');
clf;
mesh(ns,lgl,AprEnvAcf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
zlabel("AprEnv_Acf",'interp','none'); grid on;
figure();
set(gcf,'name','������ ������������� ��������� ��� � �������� ������������');
clf;
for m=1:lag
   EroEnvAcf_sET12(m,:) = EnvAcf_sET12(m,:)-AprEnvAcf_sET12(m,:);
end
mesh(ns,lgl,EroEnvAcf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
zlabel("EroEnv_Acf",'interp','none'); grid on;
%% ���������� ������� ������������� ��� ����������� ����� sET12 ��� ��������� pw
pi2 = 2.0*pi;
for j=1:S % ���� �� ��������� ���
   PhaAcfNrm = acos(AcfNrm_sET12(:,j)); % ���������� ����
   FrcAcfNrm = abs(diff(PhaAcfNrm))/pi2/dt; % ���������� ������� ��������-�� ���, ��
   insFrc_AcfNrm(j) = median(FrcAcfNrm); % ������� ���������� �������� j-�� �������� pw 
end
smo_insFrc_AcfNrm = smooth(insFrc_AcfNrm,0.2*S,'rloess');
figure();
set(gcf,'name','������� ������-�� ��� ��������-� ����� ��������� pw');
clf;
p1 = plot(ns,insFrc_AcfNrm,'b','LineWidth',0.8); hold on;
plot(ns,smo_insFrc_AcfNrm,'r','LineWidth',0.8); grid on;
xlabel("ns",'interp','none'); ylabel("insFrc_AcfNrm,Hz",'interp','none');
legend(p1,'sET12');
%% ��������������� ������ ������������� ��� ����������� ����� sET12 ��� ��������� pw 
for j=1:S
   for m=1:lag
      AprAcf_sET12(m,j) = AprEnvAcf_sET12(m,j)*cos(pi2*smo_insFrc_AcfNrm(j)*m*dt); % �������������
   end
end
figure();
set(gcf,'name','������������� ��� ����������� ����� sET12 ��������� pw');
clf;
mesh(ns,lgl,AprAcf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
mesh(ns,Time,AprAcf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
zlabel("Apr_Acf",'interp','none'); grid on;
%% ������ ��� ����������� ����� ��� �������� pw
smopto = 3; % �������� ����������� ������������� �������
for j=1:S
   pto_sET12(:,j) = pmtm(sET12(:,j),smopto,win); % ������������� �������
   pto_sET34(:,j) = pmtm(sET34(:,j),smopto,win);
end
%% ������������ ��� ����������� ����� �������� pw
fmi  = 40.0/60.0;   % ������� ����� ��� 40 ��/��� (0.6667 ��)
fma  = 240.0/60.0;  % ������� ����� ��� 240 ��/��� (4.0 ��)
Nf   = 1+win/2;     % ���-�� �������� �������
df   = cad/(win-1); % �������� ������������� �������, ��
Fmin = fmi-10*df; Fmax = fma+10*df; % ������� � ��
f(1) = 0.0;
for i=2:Nf
   f(i) = f(i-1)+df; % ������� � ������
   if abs(f(i)-Fmin)<=df
      iGmin = i;
   end
   if abs(f(i)-Fmax)<=df
      iGmax = i;
   end
end
for i=1:iGmax
   fG(i) = f(i); % ����� ������ 3D-�������
end
f = f';
figure();
set(gcf,'name','������������� ������� sET12 ��������� pw');
clf;
mesh(ns,fG(iGmin:iGmax),pto_sET12(iGmin:iGmax,:),'FaceAlpha',0.5,'FaceColor','flat');
colorbar; grid on;
xlabel("ns",'interp','none'); ylabel("f,Hz",'interp','none');
zlabel("P(f)",'interp','none');
%
figure();
set(gcf,'name','������������� ������� sET34 ��������� pw');
clf;
mesh(ns,fG(iGmin:iGmax),pto_sET34(iGmin:iGmax,:),'FaceAlpha',0.5,'FaceColor','flat');
colorbar; grid on;
xlabel("ns",'interp','none'); ylabel("f,Hz",'interp','none');
zlabel("P(f)",'interp','none');
%% ������ ������� ������ ��������� ���� ����������� ����� ��������� pw
for j=1:S
   [B,I] = sort(pto_sET12(:,j),'descend');
   pto_fMAX12(j) = f(I(1)); % I(1) - ������ �������(��) ��������� pto_sET12(:,j)
   [B,I] = sort(pto_sET34(:,j),'descend');
   pto_fMAX34(j) = f(I(1));
end
pto_fMAX12 = pto_fMAX12';
smo_pto_fMAX12 = smooth(pto_fMAX12,0.3*S,'rloess');
pto_fMAX34 = pto_fMAX34'; 
smo_pto_fMAX34 = smooth(pto_fMAX34,0.3*S,'rloess');
%
figure();
set(gcf,'name','������� ��������� ���� sET ��������� pw');
clf;
p1 = plot(ns,pto_fMAX12,'b'); hold on;
plot(ns,smo_pto_fMAX12,'r','LineWidth',0.8); grid on;
p2 = plot(ns,pto_fMAX34,'g');
plot(ns,smo_pto_fMAX34,'m','LineWidth',0.8);
xlabel("ns",'interp','none'); ylabel("fMAX,Hz",'interp','none');
legend([p1,p2],'sET12','sET34');
%% ������������� ��������� ��������� ��������� ����� cpw
[NumS,cpw_avr,cpw_med,cpw_iqr] = wav(NSF,S,win,res,sET12);
%
figure();
set(gcf,'name','Pulse wave')
clf;
% plotwave(NSF,tim,cpw_avr,cpw_med,cpw_iqr);
plotwave(1,NSF,tim,cpw_avr,cpw_med,cpw_iqr);
%% ������������� ��������� ����� ET34
[NumS,ET34_avr,ET34_med,ET34_iqr] = wav(NSF,S,win,res,sET34);
%
figure();
set(gcf,'name','Wave ET34')
clf;
% plotwave(NSF,tim,ET34_avr,ET34_med,ET34_iqr);
plotwave(1,NSF,tim,ET34_avr,ET34_med,ET34_iqr);
%% ����������� ���������� ���� cpw
% cpw = cpw_avr; % ������ ��������� ��������� �����
cpw = cpw_med; % ������ ��������� ��������� �����
cutoff = pi; pi2 = 2.0*pi;
H_cpw = hilbert(cpw);
insE_cpw = abs(H_cpw); % ���������� ���������
unwPha = unwrap(angle(H_cpw),cutoff); % ����������� ���������� ����
% �����������-(�) � ���������-(d) ���������� ����������� ���������� ����
unwPc_cpw(1) = unwPha(1); unwPd_cpw(1) = 0.0;
for i=2:NSF
   dif = unwPha(i)-unwPha(i-1);
   unwPc_cpw(i) = unwPc_cpw(i-1); % �����������
   unwPd_cpw(i) = unwPd_cpw(i-1); % ��������� 
   if dif>=0.0
      unwPc_cpw(i) = unwPc_cpw(i)+dif;
   else
      unwPd_cpw(i) = unwPd_cpw(i)+dif+pi2;
   end
end
unwPc_cpw = unwPc_cpw'; unwPd_cpw = unwPd_cpw';
figure();
set(gcf,'name','Unwrape phase pulse wave')
clf;
sp1 = subplot(2,1,1); plot(tim(1:NSF),unwPc_cpw); grid on;
xlabel("t,s",'interp','none'); ylabel("Phase cont",'interp','none');
title(sp1,'����������� ����������� ���� pw');
sp2 = subplot(2,1,2); plot(tim(1:NSF),unwPd_cpw); grid on; 
xlabel("t,s",'interp','none'); ylabel("Phase disc",'interp','none');
title(sp2,'��������� ����������� ���� pw');
%% ���������� ������� � ������� ��������� ��������� �����
% ������ ������ ����������� ����������� ���������� ����������� ���������� ����
% � ������� ����������� ����� ������� ���������� ������������� ���������� ������
t = 1:NSF; cof = 1:4;
p = pchip(t,unwPc_cpw);
insF_cpw(1) = 0.0;
for i=2:NSF
   insF_cpw(i) = p.coefs(i-1,3)/pi2/dt; % ���������� ������� ����������� ���������� cpw, ��
end
insF_cpw = insF_cpw';
smo_insF_cpw = smooth(insF_cpw,0.03*NSF,'rloess');
res_insF_cpw = insF_cpw-smo_insF_cpw; % ������� ������� 
dev_insF_cpw = smooth(res_insF_cpw.^2,0.03*NSF,'rloess');
std_insF_cpw = abs(sqrt(dev_insF_cpw));
figure();
set(gcf,'name','Frequencie and energy pulse wave')
clf;
sp1 = subplot(2,1,1); plot(tim(1:NSF),insF_cpw); hold on; 
plot(tim(1:NSF),smo_insF_cpw,'Color','r','LineWidth',0.8); hold on;
% plot(tim(1:NSF),std_insF_cpw);
xlabel("t,s",'interp','none'); ylabel("insF,Hz",'interp','none');
grid on; title(sp1,'���������� ������� cpw');
ylim([1.0 3.0]);
sp2 = subplot(2,1,2); plot(tim(1:NSF),insE_cpw.^2); 
xlabel("t,s",'interp','none'); ylabel("insE",'interp','none');
grid on; title(sp2,'���������� ������� cpw');
%%