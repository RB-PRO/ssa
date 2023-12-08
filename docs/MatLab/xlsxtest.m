clear;
exper="P1H1";
PathXls="endh/"+exper+"/"+exper+"_rPPG_output.csv";
disp(PathXls);
xls = xlsread(PathXls);

disp(1);

sgn = xls(:,3);
plot(sgn);
