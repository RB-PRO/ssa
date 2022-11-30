function import=loadData(direction, fileName)
    filename = string(direction)+"\"+fileName+".xlsx";
    data = readtable(filename);
    import = table2array(data);
end