function [C,LBD,RC] = SSA(N,M,X,nET)
%% Calculate covariance matrix C (Toeplitz approach)
%    covX  = xcorr(X,M-1,'unbiased');
%    Ctoep = toeplitz(covX(M:end));
%% Calculate covariance matrix (trajectory approach)
%  it ensures a positive semi-definite covariance matrix
   Y = zeros(N-M+1,M);
   for m=1:M
     Y(:,m) = X((1:N-M+1)+m-1);
   end
   Cemb = Y'*Y/(N-M+1);
%% Choose covariance estimation
%    C = Ctoep;
   C = Cemb;
%% Calculate eigenvalues LAMBDA and eigenvectors RHO
% Function eig returns two matrices,
% the matrix RHO with eigenvectors arranged in columns,
% the matrix LAMBDA with eigenvalues along the diagonal
   [RHO,LBD] = eig(C);
   LBD       = diag(LBD);           % extract the diagonal elements
   [LBD,ind] = sort(LBD,'descend'); % sort eigenvalues
   RHO       = RHO(:,ind);          % and eigenvectors
%% Calculate principal components PC
% The principal components are given as the scalar product
% between Y, the time-delayed embedding of X, and the eigenvectors RHO
   PC = Y*RHO;
%% Calculate reconstructed components RC
% In order to determine the reconstructed components RC,
% we have to invert the projecting PC = Y*RHO;i.e. RC = Y*RHO*RHO'=PC*RHO'
% Averaging along anti-diagonals gives the RCs for the original input X
   RC = zeros(N,nET);
   for m=1:nET
     buf = PC(:,m)*RHO(:,m)'; % invert projection
     buf = buf(end:-1:1,:);
% Anti-diagonal averaging
     for n=1:N
       RC(n,m) = mean(diag(buf,-(N-M+1)+n));
     end
end
