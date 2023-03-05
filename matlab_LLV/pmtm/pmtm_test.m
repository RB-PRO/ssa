clear; close;
filename = "pmtm" + ".xlsx"; 
data = readtable(filename);
x = table2array(data);

%

pm_x = pmtm(x,3,1024);
plot(pm_x);