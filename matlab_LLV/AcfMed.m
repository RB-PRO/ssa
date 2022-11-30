function Acf = AcfMed(lagS,win,TS)
% lagS - �������� ���������� ���������� ���� (��) TS � ����������� ������������
% win  - ���������� �������� �� TS 
% TS   - ��, ���������� win ��������
   Y = zeros(win-lagS+1,lagS); % ����������� ������� �� TS
   for m=1:lagS
     Y(:,m) = TS(m:win-lagS+m); % m-� ������� ����������� ������� �� TS
   end
   Cor = Y'*Y; % lagS*lagS ������� ��������-� ������������
   lon = lagS;
   CorPro(1:lon) = diag(Cor); % �� ��������-�� ������������ ��� ���� 0 
   Acf(1) = median(CorPro(1:lon)); % ������� ������� ��������� CorPro 
   for m=2:lagS
      lon = lon-1;
      CorPro(1:lon) = diag(Cor,m-1); % �� ��������-�� ������������ ��� ���� m-1
      if m<=lagS
         Acf(m) = median(CorPro(1:lon))/Acf(1); % ������-� ������ ������������� ��� 
      end
   end
   Acf(1) = 1.0;    
end