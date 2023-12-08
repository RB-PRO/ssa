function Save(array, FilePathName)
    file=fopen(FilePathName,'w'); 
    frames=length(array);
    for i = 1:frames
        fprintf(file,'%f\n',array(1));
    end
    fclose(file); 
end
