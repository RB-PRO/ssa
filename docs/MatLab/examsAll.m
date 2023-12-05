clear; close all; clc;

% Names
names = ["P1H1"    "P1LC1" "P1LC2" "P1LC3" "P1LC4" "P1LC5" "P1LC6" "P1LC7"     "P1M1" "P1M2" "P1M3"     "P2LC1" "P2LC2" "P2LC3" "P2LC4" "P2LC5"     "P3LC1" "P3LC2" "P3LC3" "P3LC4" "P3LC5"];


zrp="endh/"; % Zero Folder

for iName = 1:length(names)
    Name = names(iName);
    disp(Name);
    Path=zrp+Name+'/';
    
    try
        ncPrefix=""; %Prefix for modework
        RGB=load(Path+Name+ncPrefix+'_RGB.txt');
        rgb2pw(RGB, Path+Name+ncPrefix);
        pw=load(Path+Name+ncPrefix+'_pw.txt');
        chss2(pw, Path, Name+ncPrefix);
    catch
        disp("ERROR: "+ncPrefix);
    end
    
    try
        ncPrefix="_nc"; %Prefix for modework
        RGB=load(Path+Name+ncPrefix+'_RGB.txt');
        rgb2pw(RGB, Path+Name+ncPrefix);
        pw=load(Path+Name+ncPrefix+'_pw.txt');
        chss2(pw, Path, Name+ncPrefix);
    catch
        disp("ERROR: "+ncPrefix);
    end

    close all;
end