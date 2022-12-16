function import = loadVar(direction, fileName)
    filename = string(direction)+"\"+fileName+".txt";
    import=importdata(filename);
end