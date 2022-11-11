

xnn = table2array(readtable("xn.xlsx"));
for i = 1:300
    xn(i)=str2double(cell2mat(xnn(i)));
end
xn=xn';



