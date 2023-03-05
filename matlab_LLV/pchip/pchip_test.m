% Тестирование pchip
lgl=1:102;
PhaAcfNrm=importdata("test.txt");
pAcf = pchip(lgl,PhaAcfNrm);

% Сохранить матрицу в файл
writematrix(pAcf.coefs, 'coefs.xlsx');