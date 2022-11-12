function plotwave(Nmin,Nmax,t,w_avr,w_med,w_rng)
   sp1 = subplot(3,1,1); plot(t(Nmin:Nmax),w_avr(Nmin:Nmax),'b-');
   title(sp1,'Mean wave '); grid on;
   xlabel("t,s",'interp','none'); ylabel("wave",'interp','none');
   sp2 = subplot(3,1,2); plot(t(Nmin:Nmax),w_med(Nmin:Nmax),'g-'); grid on;
   title(sp2,'Median wave '); grid on;
   xlabel("t,s",'interp','none'); ylabel("wave",'interp','none');
   sp3 = subplot(3,1,3); plot(t(Nmin:Nmax),w_rng(Nmin:Nmax),'r-'); grid on;
   title(sp3,'Range wave '); grid on;
   xlabel("t,s",'interp','none'); ylabel("Range",'interp','none');
end