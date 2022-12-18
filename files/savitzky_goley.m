clear; close all; clc;

sgl=loadDataArr("sgl");
sgl2=loadDataArr("sgl2");
plot(sgl); grid on; hold on;
plot(sgl2);
legend("Original", "Savitzky goley");