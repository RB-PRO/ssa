function import=loadDataArr(fileName)
    filename = fileName+".xlsx";
    data = readtable(filename);
    import = table2array(data);
end