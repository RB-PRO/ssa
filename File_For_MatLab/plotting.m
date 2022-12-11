clear; close all; clc;

%% Compare reconstruction and original time series
    %% Covariance matrix - Picture 1
    C=loadData(1,"C");
    figure();
    set(gcf,'name','Covariance matrix');
    clf;
    imagesc(C);
    axis square;
    set(gca,'clim',[-1 1]);
    colorbar;

    %% Eigenvalues - Picture 2
    LBD=loadData(2,"LBD");
    figure();
    set(gcf,'name','Eigenvalues')
    clf;
    plot(LBD,'o-');

    %% Original time series and reconstruction - Picture 3
    seg = 100; % номер сегмента pw для визуализации
    win = 1024;
    spw=loadData(3,"spw");
    tim=loadData(3,"tim")';
    sET12=loadData(3,"sET12");
    figure();
    set(gcf,'name','Original time series and reconstruction'); clf;
    plot(tim(1:win),spw(:,seg),'b-',tim(1:win),sET12(:,seg),'r-');
    legend('Original','sET12'); xlabel("t,s",'interp','none'); ylabel("sET",'interp','none');
    
    %% Original time series and reconstruction - Picture 4
    seg = 100; % номер сегмента pw для визуализации
    win = 1024;
    spw=loadData(4,"spw");
    tim=loadData(4,"tim")';
    sET34=loadData(4,"sET34");
    figure();
    set(gcf,'name','Original time series and reconstruction'); clf;
    plot(tim(1:win),spw(:,seg),'b-',tim(1:win),sET34(:,seg),'r-');
    legend('Original','sET34'); xlabel("t,s",'interp','none'); ylabel("sET",'interp','none');
    
    %% Визуализация АКФ сингулярных троек для сегментов pw - Picture 6
    lag  = floor(win/10); % наибольший лаг АКФ <= win/10
    lagS = 2*lag;
    ns=loadData(5,"ns");
    Time=loadData(5,"Time");
    Acf_sET12=loadData(5,"Acf_sET12");
    figure();
    set(gcf,'name','АКФ сингулярных троек sET12 сегментов pw'); clf;
    % mesh(ns,lgl,Acf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    % xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
    mesh(ns,Time,Acf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
    zlabel("Acf",'interp','none'); grid on;
    
    %% - Picture 3
    ns=loadData(6,"ns");
    Time=loadData(6,"Time")';
    EnvAcf_sET12=loadData(6,"EnvAcf_sET12");
    figure();
    set(gcf,'name','Огибающие АКФ сингулярных троек sET12 сегментов pw');
    clf;
    % mesh(ns,lgl,EnvAcf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    % xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
    mesh(ns,Time,EnvAcf_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
    zlabel("Env_Acf",'interp','none'); grid on;

    %% - Picture 7
    ns=loadData(7,"ns");
    Time=loadData(7,"Time")';
    AcfNrm_sET12=loadData(7,"AcfNrm_sET12");
    figure();
    set(gcf,'name','Нормированные АКФ сингулярных троек sET12 сегментов pw');
    clf;
    % mesh(ns,lgl,AcfNrm_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    % xlabel("ns",'interp','none'); ylabel("lag",'interp','none');
    mesh(ns,Time,AcfNrm_sET12(1:lag,:),'FaceAlpha',0.5,'FaceColor','flat'); colorbar;
    xlabel("ns",'interp','none'); ylabel("lag,s",'interp','none');
    zlabel("Acf_Nrm",'interp','none'); grid on;
    
    %% - Picture 8
    ns=loadData(8,"ns");
    insFrc_AcfNrm=loadData(8,"insFrc_AcfNrm");
    smo_insFrc_AcfNrm=loadData(8,"smo_insFrc_AcfNrm")';
    figure();
    set(gcf,'name','Частоты нормир-ой АКФ сингуляр-х троек сегментов pw');
    clf;
    p1 = plot(ns,insFrc_AcfNrm,'b','LineWidth',0.8); hold on;
    plot(ns,smo_insFrc_AcfNrm,'r','LineWidth',0.8); grid on; % smo_insFrc_AcfNrm
    xlabel("ns",'interp','none'); ylabel("insFrc_AcfNrm,Hz",'interp','none');
    legend(p1,'sET12');
    
%%
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    