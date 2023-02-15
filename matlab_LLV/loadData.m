function import=loadData(fileName)
    filename = fileName+".xlsx";
    data = readtable(filename);
    importStr = table2array(data);
    
    import(length(importStr))=0.0;
    for i=1:length(importStr) 
        import(i)=str2double(importStr(i));
    end
    
end