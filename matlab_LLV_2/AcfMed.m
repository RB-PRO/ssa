function Acf = AcfMed(lagS,win,TS)
% lagS - параметр погружения временного ряда (ВР) TS в траекторное пространство
% win  - количество отсчетов ВР TS 
% TS   - ВР, содержащий win отсчетов
   Y = zeros(win-lagS+1,lagS); % траекторная матрица ВР TS
   for m=1:lagS
     Y(:,m) = TS(m:win-lagS+m); % m-й столбец траекторной матрица ВР TS
   end
   Cor = Y'*Y; % lagS*lagS матрица корреляц-х произведений
   lon = lagS;
   CorPro(1:lon) = diag(Cor); % ВР корреляц-го произведения для лага 0 
   Acf(1) = median(CorPro(1:lon)); % медиана главной диагонали CorPro 
   for m=2:lagS
      lon = lon-1;
      CorPro(1:lon) = diag(Cor,m-1); % ВР корреляц-го произведения для лага m-1
      if m<=lagS
         Acf(m) = median(CorPro(1:lon))/Acf(1); % медиан-я оценка нормированной АКФ 
      end
   end
   Acf(1) = 1.0;    
end