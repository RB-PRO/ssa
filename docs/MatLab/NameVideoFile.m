function Name = NameVideoFile(VideoFileName)
    Name = replace(VideoFileName, '.mp4', '');
    Name = replace(Name, '.avi', '');
    Name = replace(Name, '.mat', '');
    Name = replace(Name, '_RGB', '');
    Name = replace(Name, '.txt', '');
    Name = replace(Name, '_pw', '');
    Name = replace(Name, '_but', '');
end 