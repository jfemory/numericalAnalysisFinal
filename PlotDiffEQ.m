syms a t
s = a*(50*t - 2500 + exp(-0.02*t)*2500)+exp(-0.02*t)*150-50;
fs = fsurf(s, [0.0 0.02 0 180]);
zlim([0,110]);
title('Family of solutions to $\frac{ds}{dt}$','Interpreter','latex')
xlabel('Acceleration $[ms^{-2}]$','Interpreter','latex')
ylabel('Time $[s]$','Interpreter','latex')
zlabel('Distance $[m]$','Interpreter','latex')