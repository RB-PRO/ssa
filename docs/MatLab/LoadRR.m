function RR = LoadRR(FilePath)
    try
        xls = load(FilePath);
        RR = xls(:,2);
    catch
        disp("NO rr file");
        RR = [];
    end
end