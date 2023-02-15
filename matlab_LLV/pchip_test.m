 
% Тестирование pchip
S=102;
lgl=1:S;
PhaAcfNrm=loadData("PhaAcfNrm");
pAcf = pchip(lgl,PhaAcfNrm);
disp(pAcf.coefs);