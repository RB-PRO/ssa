function SaveRGB(RGB, FilePathName)
    file=fopen(FilePathName,'w'); 
    frames=length(RGB);
    for i = 1:frames
        fprintf(file,'%f;%f;%f\n',RGB(i, 1), RGB(i, 2), RGB(i, 3));
    end
    fclose(file); 
end
