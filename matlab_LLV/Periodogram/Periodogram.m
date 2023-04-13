function output = Periodogram(x) %#codegen 

    N = 1024;
    n = 0:N-1;

    % Значение ненулевых коэффициентов семества окон Блэкмана-Хэрриса 
    a = [0.35875 0.48829 0.14128 0.01168]; % Using Matlab 
    w = a(1) - a(2) * cos((2*pi/N)*1*n) + a(3) * cos((2*pi/N)*2*n) + a(4) * cos((2*pi/N)*3*n);
    
    output=periodogram(x,w,N);
end