function newpw = TrendSubtraction(pw)
    smoot_pw = smoothdata(pw,"movmean",32); 
%     smooth_pw = smooth(pw,128); 
    newpw=pw-smoot_pw;
    
    figure();
    subplot(2,1,1); 
    plot(pw); hold on; plot(smoot_pw); grid on; hold off; 
    legend('pw','smoothdata');
%     plot(pw); hold on; plot(smoot_pw); grid on; plot(smooth_pw); hold off; 
%     legend('pw','smoothdata','smooth');
    
    subplot(2,1,2);
    plot(newpw); grid on;
    
end