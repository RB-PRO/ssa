function [NS,w_avr,w_med,w_iqr] = wav(N,S,W,res,sET)
   ET = zeros(N,S);
   for j=1:S % цикл по сегментам
      for i=1:W
         k = (j-1)*res;
         ET(i+k,j) = sET(i,j); % сдвинутый сегмент ET(:,j)
      end
   end
   TS = zeros(S);
   for i=1:N % цикл по отсчетам ET
      nSi = 0;
      for j=1:S % цикл по сегментам ET
         if ET(i,j)~=0
            nSi      = nSi+1; % текущий ненулевой отсчет ET(i,j)
            Smp(nSi) = ET(i,j); 
         end
      end
      NS(i)    = nSi; % кол-во сегментов для текущего i
      w_avr(i) = mean(Smp(1:nSi));   % выборочная средняя
      w_med(i) = median(Smp(1:nSi)); % медиана
      w_iqr(i) = (prctile(Smp(1:nSi),75)-prctile(Smp(1:nSi),25))/2.0;
   end
   NS    = NS';
   w_avr = w_avr'; % очищенная пульсовая волна
   w_med = w_med';
   w_iqr = w_iqr'; % половина интерквартильного диапазова
end