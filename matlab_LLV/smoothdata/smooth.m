%#codegen
function [outputArr] = smooth()
    inputArr=[0 5 3 2 3 3 6 3 2 5 4];
    outputArr=smoothdata(inputArr,'rloess',0.3*200);
end