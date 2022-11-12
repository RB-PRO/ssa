clear; close all; clc;

% Загружаю исходный сигнал
X_xlsx = table2array(readtable("xn.xlsx"));
X(length(X_xlsx))=0.0;
for i = 1:length(X_xlsx)
    X(i)=str2double(cell2mat(X_xlsx(i)));
end
X=X';

N=300;
M=40;
nET=2;

[C,LBD,RC] = SSA(N,M,X,nET);