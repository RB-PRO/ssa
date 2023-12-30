function [out, lower_prct, upper_prct] = RaznFilter(signal, threshold)
    Nmed = 5;
    signal = signal-medfilt1(signal, Nmed);
    
    % Первая разность сигнала
    len = length(signal);
    FirstRaznSignal = [signal(1), signal];
    FirstRaznSignal = FirstRaznSignal-[signal, signal(len)];
    
    disp("Апертура фильтра для insFrc_AcfNrm для фильтра процентилей: " + Nmed);  
    prct = prctile(FirstRaznSignal, threshold);

    lower_prct=prct(1);
    upper_prct=prct(2);
    
    disp("lower_prct " + lower_prct);
    disp("upper_prct " + upper_prct);
    
%     signal=signal-medfilt1(signal, Nmed);
    out = FirstRaznSignal;
    
    figure();
    plot(FirstRaznSignal,'blue--'); hold on; grid on;
    line('XData', [0 200], 'YData', [upper_prct upper_prct], 'Color','black','LineStyle','--');
    line('XData', [0 200], 'YData', [lower_prct lower_prct], 'Color','black','LineStyle','--');
    
    index = 1;
    MemoryValue = FirstRaznSignal(1);
    
    % Цикл по всему сигналу
    for value = FirstRaznSignal
        if (value < lower_prct || value > upper_prct)
            out(index) = MemoryValue;
        end
       MemoryValue=out(index);
       index=index+1;
    end
    
     plot(out,'red'); hold on; grid on;
end