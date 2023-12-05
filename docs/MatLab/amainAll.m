clear; close all; clc;

% Names
names = ["P1H1"    "P1LC1" "P1LC2" "P1LC3" "P1LC4" "P1LC5" "P1LC6" "P1LC7"     "P1M1" "P1M2" "P1M3"     "P2LC1" "P2LC2" "P2LC3" "P2LC4" "P2LC5"     "P3LC1" "P3LC2" "P3LC3" "P3LC4" "P3LC5"];


frt = ".avi"; zrp="endh/";
ncPrefix="";

for iName = 1:length(names)
    Name = names(iName);
    disp(Name);
    
    Path=zrp+Name+'/';
    
    % Extracting a time series of RGB segments
    try
        Face_tracking(Path+Name+ncPrefix+frt);
        RGB=load(Path+Name+ncPrefix+'_RGB.txt');
    catch
        disp("Œÿ»¡ ¿: Face_tracking")
    end

    % Calculation of the photoplethysmography signal
    try
        rgb2pw(RGB, Path+Name+ncPrefix);
        pw=load(Path+Name+ncPrefix+'_pw.txt');
    catch
        disp("ERROR: rgb2pw")
    end

    % Estimates of the average pitch frequencies for PW segments
    try
        chss2(pw, Path, Name+ncPrefix);
    catch
        disp("ERROR: chss2")
    end
    
    close all;
end
