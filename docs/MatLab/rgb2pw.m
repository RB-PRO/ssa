function pw = rgb2pw(pwc, alg) 
    fMi  = 40.0/60.0;   % ������� ����� ��� 40 ��/��� (0.6667 ��)
    cad=30;
    len=length(pwc(:,1));

    % ������������� ������� ��� ������������ ��������� ���� �� �������� �������
    % �������� ������: 1->R; 2->G; 3->B

    if alg=="G"
        pw = pwc(:,2);
    end

    if alg=="G"
        pw = pwc(:,2)-pwc(:,1);
    end

    if alg=="Cr"
        pw = (112.0*pwc(:,1)-93.8*pwc(:,2)-18.2*pwc(:,3))./255.0;
        pw=pw-movmean(pw,32); % �������� �����
    end

    if alg=="CHROM"
        ws(:,1) = (3.0*pwc(:,1)-2.0*pwc(:,2))./sqrt(13.0);
        ws(:,2) = (-1.5*pwc(:,1)-pwc(:,2)+1.5*pwc(:,3))./sqrt(11.0/2.0); % �������� �� (-1) ��� pw=ws(:,1)+ws(:,2)
    end
    
    if alg=="POS"
        ws(:,1) = (pwc(:,2)-pwc(:,3))./sqrt(2.0);
        ws(:,2) = (pwc(:,2)+pwc(:,3)-2.0*pwc(:,1))./sqrt(6.0);
    end

    %% �������������� � ������������� ��������� ����� pw
    SMO_med = floor(cad/fMi);

    if alg=="G" || alg=="GR" || alg=="Cr"
        pw2 = pw.*pw;
        DEV_med = medfilt1(pw2,SMO_med); 
        STD_med(:,1) = sqrt(DEV_med);
        pw = pw./STD_med(:,1);
    end

    if alg=="CHROM" || alg=="POS"
        DEV_med = medfilt1(ws(:,1).*ws(:,1),SMO_med); STD_med(:,1) = sqrt(DEV_med);
        DEV_med = medfilt1(ws(:,2).*ws(:,2),SMO_med); STD_med(:,2) = sqrt(DEV_med);
        pw = (ws(:,1)./STD_med(:,1))+(ws(:,2)./STD_med(:,2));
        [iPer, D_time] = iPer(video, pw, alg);
        HRV (video, iPer, D_time, alg);
    end

    % ������ ���������� ��������
    prcMi = prctile(pw,0.1); 
    prcMa = prctile(pw,99.9); % ���������� �� ������� 0.1% � 99.9%
    for i=1:len
        if pw(i)<prcMi
            pw(i) = prcMi;    
        end
        if pw(i)>prcMa
            pw(i) = prcMa;    
        end
    end
    STD = std(pw);
    pw = pw./STD; % ������������� pw
    pw = movmean(pw,5);
    
end