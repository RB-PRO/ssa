clear; close all; clc;

filename = 'pw.xlsx';
data = readtable(filename);
A = table2array(data);
plot(A); grid on;

% figure
% filename = 'xzx.xlsx';
% data = readtable(filename);
% A = table2array(data);
% plot(A); grid on;