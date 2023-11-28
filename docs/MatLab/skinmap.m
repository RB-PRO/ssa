function [out, bin] = skinmap(img_orig)

%   The function reads an image file given by the input parameter string
%   filename, read by the MATLAB function 'imread'.
%   out - contains the skinmap overlayed onto the image with skin pixels
%   marked in blue color.
%   bin - contains the binary skinmap, with skin pixels as '1'.
%    
    if nargin > 1 || nargin < 1
        error('usage: generate_skinmap(image)');
    end
    
    %Read the image, and capture the dimensions
    height = size(img_orig,1);
    width = size(img_orig,2);
    
    %Initialize the output images
    out = img_orig;
    bin = zeros(height,width);
%      img = img_orig;
    
    %Apply Grayworld Algorithm for illumination compensation
    A_lin = rgb2lin(img_orig);
    percentiles = 10;
    illuminant = illumgray(A_lin,percentiles);
    B_lin = chromadapt(A_lin,illuminant,'ColorSpace','linear-rgb');
    img = lin2rgb(B_lin);
    
    %Convert the image from RGB to YCbCr
    img_ycbcr = rgb2ycbcr(img);
    Cb = img_ycbcr(:,:,2);
    Cr = img_ycbcr(:,:,3);
    
    %Convert to HUE
    [hue,~,~] = rgb2hsv(img);
    
    %Detect Skin
    [r,c,~] = find(Cb>=98 & Cb<=142 & Cr>=135 & Cr<=177 & hue<=0.1 & 0.01<=hue);
    numind = size(r,1);
    
    %Mark Skin Pixels
    for i=1:numind
        out(r(i),c(i),:) = [0 0 255];
        bin(r(i),c(i)) = 1;
    end
end














