clear; close all; clc;

filename = 'xx.xlsx';
data = readtable(filename);
A = table2array(data);
plot(A); grid on;

% figure
% filename = 'xzx.xlsx';
% data = readtable(filename);
% A = table2array(data);
% plot(A); grid on;