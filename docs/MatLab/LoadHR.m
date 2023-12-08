function HR = LoadHR(FilePath)
    try
        xls = xlsread(FilePath);
        HR = xls(:,3);
    catch
        disp("NO hr file");
        HR = [];
    end
end