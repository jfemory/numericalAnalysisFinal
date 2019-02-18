syms a t
s = a*(50*t - 2500 + exp(-0.02*t)*2500)+exp(-0.02*t)*150-50;
fs = fsurf(s, [0.0 0.02 0 180]);
zlim([-20,0]);
%rotate to where you see the projection onto the a,t-plane.
title('Zeros of $\frac{ds}{dt}$','Interpreter','latex')
xlabel('Acceleration $[ms^{-2}]$','Interpreter','latex')
ylabel('Time $[s]$','Interpreter','latex')
zlabel('Distance $[m]$','Interpreter','latex')