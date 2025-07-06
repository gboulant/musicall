# coding: utf-8

from array import array
from ctypes import cdll, c_double, POINTER
lib = cdll.LoadLibrary("./wave.so")

lib.WaveSize.argtypes = [c_double]
lib.SineWave.argtypes = [c_double, c_double, c_double, POINTER(c_double)]

samplerate = lib.WaveSize(1.) # number of samples in 1 second
print("samplerate = %d"%samplerate) 

def SineWave(f,a,d) -> array[float]:
    N = lib.WaveSize(d)
    out = array('d', [0 for i in range(N)]) 
    out_ptr = (c_double * len(out)).from_buffer(out)
    lib.SineWave(f,a,d, out_ptr)
    return out

import matplotlib.pyplot as plt
def plottimeseries(t,s):
    fig, ax = plt.subplots()
    ax.plot(t, s, 'y-',label="wave")
    ax.set_xlabel("time")
    ax.set_ylabel("value")
    plt.legend()
    plt.show()

f = 10. # Hz
a = 10.
d = 2.0 # seconds
s = SineWave(f,a,d)
t = [i/samplerate for i in range(len(s))]
plottimeseries(t,s)
