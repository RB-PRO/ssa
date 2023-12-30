function [out, lower_prct, upper_prct] = ProcentileFilter(signal, threshold)


    Nmed = 5;
    signal = signal-medfilt1(signal, Nmed);
    
    disp("Апертура фильтра для insFrc_AcfNrm для фильтра процентилей: " + Nmed);  
    prct = prctile(signal, threshold);

    lower_prct=prct(1);
    upper_prct=prct(2);
    
%     lower_prct = prctile(signal-medfilt1(signal, Nmed), 30);
%     upper_prct = prctile(signal-medfilt1(signal, Nmed), 30);

    disp("lower_prct " + lower_prct);
    disp("upper_prct " + upper_prct);
    
%     line('XData', [0 200], 'YData', [lower_prct lower_prct], 'Color','red','LineStyle','--');
%     line('XData', [0 200], 'YData', [upper_prct upper_prct],'Color','red','LineStyle','--');
    
    
    signal=signal-medfilt1(signal, Nmed);
    out = signal;
    
    index = 1;
    MemoryValue = signal(1);
    
    % Цикл по всему сигналу
    for value = signal
       
        if (value < lower_prct || value > upper_prct)
            out(index) = MemoryValue;
        end
       
       MemoryValue=out(index);
       index=index+1;
    end
     
end
