% w[n] = 0.422323 – 0.49755 cos(2n/N) + 0.07922 cos(4n/N) 

% Шаг
N = 1024;
n = 0:N-1;
 
% Значение ненулевых коэффициентов семества окон Блэкмана-Хэрриса
Koef3_67db = [0.42323 0.49755 0.07922 0];
Koef3_61db = [0.44959 0.49364 0.05677 0];
Koef4_92db = [0.35875 0.48829 0.14128 0.01168]; % Using Matlab
Koef4_74db = [0.40217 0.49703 0.09392 0.00183];

a = Koef4_92db;

w = a(1) - a(2) * cos((2*pi/N)*1*n) + a(3) * cos((2*pi/N)*2*n) + a(4) * cos((2*pi/N)*3*n);
plot(w); hold on;

win=N;

% output=periodogram(spw_j,blackmanharris(win),win);

x = cos(pi/4*n)+rand(1);
plot(x);

figure;
per=periodogram(x,w,N);
plot(per)