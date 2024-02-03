function out = first_diff_filter(signal, threshold)
    
    out = signal;
    diff_signal = diff(signal);
    
    lower_prct=0; upper_prct=0;
    if length(threshold) == 1
        lower_prct=-threshold; upper_prct=threshold;
    end
    if length(threshold) == 2 
        prct = prctile(diff_signal, threshold);
        lower_prct=prct(1); upper_prct=prct(2);
    end
    
    for i = 1:(length(out)-1)
        delta = signal(i+1)-signal(i);  
        if ((diff_signal(i)>upper_prct) || (diff_signal(i)<lower_prct))
            delta = 0;
        end
        out(i+1) = out(i) + delta;
    end
    
end