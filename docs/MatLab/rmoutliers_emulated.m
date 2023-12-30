function B = rmoutliers_emulated(A, threshold)
    lower_prct = prctile(A, threshold(1));
    upper_prct = prctile(A, threshold(2));
    
    B = A;
    B(B < lower_prct | B > upper_prct) = 0; % Заменяем выбросы на NaN
end
