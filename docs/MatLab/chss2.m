function chss2(pw, Path, Name)
%     dt=1/30;     % �������� ��������� �������������
    fmin=0.15; % ������ ������� - ��������� �������� ����������� �����
    dt=1/30;     % �������� ��������� �������������
    Nmed=1/(dt*fmin); % �������� �������
    PathHR=Path+replace(Name,"_nc","")+"_rPPG_output.csv";
    hr = LoadHR(PathHR);
    % PathHR=Path+replace(Name,"_nc","")+"_Mobi_RR-intervals.rr";
    % hr = LoadRR(PathHR);

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
            k = (j-1)*res;
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
    % fmp=zeros(S,1);
    for j=1:S % ���� �� ��������� pw
        %    L(j) = floor(cad/fmp(j)); % ���-�� �������� ��������� ���� pw
        L(j) = floor(cad/1.5); % ���-�� �������� ��������� ���� pw
    end
    L = L';
    K = 5; % ���-�� �������� ��� ��������� ��������
    M = K*max(L); % �������� �������� � ����������� ������������ 
    %% SSA- ������ ��������� pw
    nET = 4;   % ���-�� ����������� ����� ��� ��������� pw
    for j=1:S  % ���� �� ���������
        %% SSA time series
        %    M = K*L(j); % �������� �������� � ����������� ������������
        [C,LBD,RC] = SSA(win,M,spw(:,j),nET);
        %% Estimation of the spw(:,j) reconstructed with the pair of RC   
        sET12(:,j) = sum(RC(:,1:2),2);   
        %    sET34(:,j) = sum(RC(:,3:4),2);   
        %% Compare reconstruction and original time series
        if j==100 % ����� �������� pw ��� ������������
            figure('name','Covariance matrix'); clf;
            imagesc(C); axis square; set(gca,'clim',[-1 1]); colorbar;
            figure('name','Eigenvalues'); clf; plot(LBD,'o-');
            figure('name','Original time series and reconstruction'); clf;
            plot(tim(1:win),spw(:,j),'b-',tim(1:win),sET12(:,j),'r-');
            legend('Original','sET12'); xlabel("t,s",'interp','none'); ylabel("sET",'interp','none');
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
    Time=0:dt:lag*dt-dt;
    figure('name','��� ����������� ����� sET12 ��������� pw'); clf;
    % mesh(ns,lgl,Acf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    % xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
    mesh(ns,Time,Acf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
    zlabel("Acf",'interp','none'); grid on;
    %% ��������� �� �������� ��������� ���������� abs(acf_sET12)
    for j=1:S % ���� �� ��������� ���
        absTS = abs(Acf_sET12(:,j));
        AT1   = absTS(1);
        AT2   = absTS(2);
        maxTS = zeros(lag,1); maxTS(1) = AT1;
        maxN  = zeros(lag,1); maxN(1)  = 1;
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
        EnvAcf_sET12(:,j) = interp1(NumMax,maxTS(1:Nmax),lgl,'pchip');
        AcfNrm_sET12(:,j) = Acf_sET12(1:lag,j)./EnvAcf_sET12(:,j); % ������������� ���
    end
    figure('name','������������� ��� ����������� ����� sET12 ��������� pw'); clf;
    % mesh(ns,lgl,AcfNrm_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    % xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
    mesh(ns,Time,AcfNrm_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
    zlabel("Acf_Nrm",'interp','none'); grid on;
    %% ���������� ������� ������������� ��� ����������� ����� sET12 ��� ��������� pw
    pi2 = 2.0*pi;
    for j=1:S % ���� �� ��������� ���
        PhaAcfNrm = abs(acos(AcfNrm_sET12(:,j))); % ���������� ����
        pAcf = pchip(lgl,PhaAcfNrm); % ������������ ������������ pchip
        for m=2:lag
            FrcAcfNrm(m) = abs(pAcf.coefs(m-1,3))/pi2/dt; % ���������� ������� ��������-�� ���, ��        
        end
        FrcAcfNrm(1) = FrcAcfNrm(2); 
        %    FrcAcfNrm = abs(diff(PhaAcfNrm))/pi2/dt; % ���������� ������� ��������-�� ���, ��
        insFrc_AcfNrm(j) = median(FrcAcfNrm); % ������� ���������� �������� j-�� �������� pw 
    end
    
    disp("�������� ������� ��� insFrc_AcfNrm: " + Nmed);  
%     insFrc_AcfNrm=medfilt1(insFrc_AcfNrm, 5);
      
    smo_insFrc_AcfNrm = smoothdata(insFrc_AcfNrm, 'rloess', 0.25*S); % smo_insFrc_AcfNrm = smooth(insFrc_AcfNrm,0.25*S,'rlowess');
    figure('name','������� ������-�� ��� ��������-� ����� ��������� pw','Position', [0 0 1400 800]); clf;
%             insFrc_AcfNrm=medfilt1(insFrc_AcfNrm,Nmed);

    p1 = plot(ns,insFrc_AcfNrm,'b','LineWidth',0.8); hold on;
    plot(ns,smo_insFrc_AcfNrm,'r','LineWidth',0.8); grid on;
    xlabel("ns",'interp','none'); ylabel("insFrc_AcfNrm,Hz",'interp','none');
    title("������� ������-�� ��� ��������-� ����� ��������� pw");
%     legend(p1,'sET12');
    if length(hr)>100
        ns_hr = (length(ns)/length(hr) : length(ns)/length(hr) : length(ns))';
        % yyaxis right; 
        plot(ns_hr,hr./60,'black'); ylabel("HR[bpm]",'interp','none');
        % legend(p1,'insFrc_AcfNrm','rloess','HR[bpm]');
        hr_med=medfilt1(hr,Nmed*5);
        hr_diff_med=hr-hr_med;
        plot(ns_hr,hr_med./60,'cyan--');
        plot(ns_hr,hr_diff_med./60,'magenta'); %ylabel("HR[bpm]",'interp','none');
        legend('insFrc AcfNrm','rloess','HR','HR[medfilt]','HR[HR-medfilt1]')
    end
    %!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
    subplot(1,3,[1 2]);
    plot(ns,insFrc_AcfNrm,'b','LineWidth',0.8); hold on;
    plot(ns,smo_insFrc_AcfNrm,'r','LineWidth',0.8); grid on;
    xlabel("ns",'interp','none'); ylabel("insFrc_AcfNrm,Hz",'interp','none');
    title("������� ������-�� ��� ��������-� ����� ��������� pw");
%     legend(p1,'sET12');
    if length(hr)>100
        ns_hr = (length(ns)/length(hr) : length(ns)/length(hr) : length(ns))';
        % yyaxis right; 
        plot(ns_hr,hr./60,'black'); ylabel("HR[bpm]",'interp','none');
        % legend(p1,'insFrc_AcfNrm','rloess','HR[bpm]');
        hr_med=medfilt1(hr,Nmed*5);
        hr_diff_med=hr-hr_med;
        plot(ns_hr,hr_med./60,'cyan--');
        
        legend('insFrc AcfNrm','rloess','HR','HR[medfilt]'); 

        
        
    subplot(1,3,3);
        [outs, lower_prct, upper_prct] =  RaznFilter(insFrc_AcfNrm, [1 99]); 
        plot((insFrc_AcfNrm-medfilt1(insFrc_AcfNrm, 5)),'black--'); hold on;
%         plot(rmoutliers_emulated(insFrc_AcfNrm, [10 90])./60,'red'); 
%         plot(filloutliers(insFrc_AcfNrm,"next")./60,'red'); 
        
%         plot(outs, 'blue');
        line('XData', [0 200], 'YData', [lower_prct lower_prct], 'Color','red','LineStyle','--');
        line('XData', [0 200], 'YData', [upper_prct upper_prct],'Color','red','LineStyle','--');


        
%         �������� �� �������� ������ ����������� ������� �����������
%         plot(ns_hr,hr_diff_med./60,'magenta'); %ylabel("HR[bpm]",'interp','none');
        legend('HR[HR-medfilt1]');  ylabel("HR[bpm]",'interp','none'); xlabel("ns",'interp','none'); grid on;
    end

    %% ������ ��� ����������� ����� ��� �������� pw
    smopto = 3; % �������� ����������� ������������� �������
    for j=1:S
        %    disp(spw(:,j));
        pto_sET12(:,j) = periodogram(spw(:,j),blackmanharris(win),win); % ��������-�������
        pto_sET12(:,j) = pmtm(sET12(:,j),smopto,win); % ������������� �������
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
    figure('name','������������� ������� sET12 ��������� pw'); clf;
    mesh(ns,fG(iGmin:iGmax),pto_sET12(iGmin:iGmax,:),'FaceAlpha',0.5,'FaceColor','flat');
    colorbar; grid on;
    xlabel("ns",'interp','none'); ylabel("f,Hz",'interp','none'); zlabel("P(f)",'interp','none');
    %% ������ ������� ������ ��������� ���� ����������� ����� ��������� pw
    for j=1:S
        [B,I] = sort(pto_sET12(:,j),'descend');
        pto_fMAX12(j) = f(I(1)); % I(1) - ������ �������(��) ��������� pto_sET12(:,j)
    end
    pto_fMAX12 = pto_fMAX12';
    smo_pto_fMAX12 = smoothdata(pto_fMAX12,'rloess',0.3*S); 
    % smo_pto_fMAX12 = smooth(pto_fMAX12,0.3*S,'rloess');
    figure('name','������� ��������� ���� sET ��������� pw','Position', [800 0 800 600]); clf;
%             pto_fMAX12=medfilt1(pto_fMAX12,Nmed);
    p=plot(ns,pto_fMAX12,'b'); hold on;
    plot(ns,smo_pto_fMAX12,'r','LineWidth',0.8); grid on;
    xlabel("ns",'interp','none'); ylabel("fMAX,Hz",'interp','none');
    title("������� ��������� ���� sET ��������� pw");
    if length(hr)>100
        ns_hr = (length(ns)/length(hr) : length(ns)/length(hr) : length(ns))';
        % yyaxis right;
        plot(ns_hr,hr./60,'black'); ylabel("HR[bpm]",'interp','none');
        % legend(p,'pto_fMAX12','rloess','HR[bpm]');
        hr_med=medfilt1(hr,Nmed*5);
        hr_diff_med=hr-hr_med;
        plot(ns_hr,hr_med./60,'cyan--');
        plot(ns_hr,hr_diff_med./60,'magenta'); %ylabel("HR[bpm]",'interp','none');
        legend('pto sET12','smoothdata','HR[bpm]','medfilt1','hr-medfilt1')
    end
    saveas(p,Path+Name+'_���_sET.png');

    %% ������������� ��������� ��������� ��������� ����� cpw
    [NumS,cpw_avr,cpw_med,cpw_iqr] = wav(NSF,S,win,res,sET12);
    % figure('name','Pulse wave'); clf;
    % plotwave(1,NSF,tim,cpw_avr,cpw_med,cpw_iqr);
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
    figure('name','Unwrape phase pulse wave'); clf;
    sp1 = subplot(2,1,1); plot(tim(1:NSF),unwPc_cpw); grid on;
    xlabel("t,s",'interp','none'); ylabel("Phase cont",'interp','none');
    title(sp1,'����������� ����������� ���� pw');
    sp2 = subplot(2,1,2); plot(tim(1:NSF),unwPd_cpw); grid on; 
    xlabel("t,s",'interp','none'); ylabel("Phase disc",'interp','none');
    title(sp2,'��������� ����������� ���� pw');
    %% ���������� ������� � ������� ��������� ��������� �����
    % ������ ������ ����������� ����������� ���������� ����������� ���������� ����
    % � ������� ����������� ����� ������� ���������� ������������� ���������� ������
    t = 1:NSF;
    p = pchip(t,unwPc_cpw);
    insF_cpw(1) = 0.0;
    for i=2:NSF
        insF_cpw(i) = p.coefs(i-1,3)/pi2/dt; % ���������� ������� ����������� ���������� cpw, ��
    end
    insF_cpw(1) = insF_cpw(2);
    insF_cpw = insF_cpw';
    smo_insF_cpw = smoothdata(insF_cpw, 'rloess', 0.03*NSF); 
    % smo_insF_cpw = smooth(insF_cpw,0.03*NSF,'rloess');
    res_insF_cpw = insF_cpw-smo_insF_cpw; % ������� ������� 
    dev_insF_cpw = smoothdata(res_insF_cpw.^2, 'rloess', 0.03*NSF); 
    % dev_insF_cpw = smooth(res_insF_cpw.^2,0.03*NSF,'rloess');
    std_insF_cpw = abs(sqrt(dev_insF_cpw));
    figure('name','Frequencie and energy pulse wave'); clf;
    sp1 = subplot(2,1,1); plot(tim(1:NSF),insF_cpw); hold on;
    plot(tim(1:NSF),smo_insF_cpw,'Color','r','LineWidth',0.8); hold on;
    % plot(tim(1:NSF),std_insF_cpw);
    xlabel("t,s",'interp','none'); ylabel("insF,Hz",'interp','none');
    grid on; title(sp1,'���������� ������� cpw'); ylim([1.0 3.0]);
    sp2 = subplot(2,1,2); plot(tim(1:NSF),insE_cpw.^2); 
    xlabel("t,s",'interp','none'); ylabel("insE^2",'interp','none');
    grid on; title(sp2,'���������� ������� cpw');
    save(Path+Name+"_nc"+".mat");
end