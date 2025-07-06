# coding: utf-8

from array import array
from ctypes import cdll, c_double, POINTER
lib = cdll.LoadLibrary("./wave.so")

lib.WaveSize.argtypes = [c_double]
lib.SineWave.argtypes = [c_double, c_double, c_double, POINTER(c_double)]
lib.SquareWave.argtypes = [c_double, c_double, c_double, POINTER(c_double)]
lib.KarplusStrongWave.argtypes = [c_double, c_double, c_double, POINTER(c_double)]

def makewave(f,a,d, wavefct)-> array[float]:
    N = lib.WaveSize(d)
    out = array('d', [0 for i in range(N)]) 
    out_ptr = (c_double * len(out)).from_buffer(out)
    wavefct(f,a,d, out_ptr)
    return out

def SineWave(f,a,d) -> array[float]: return makewave(f,a,d, lib.SineWave)
def SquareWave(f,a,d) -> array[float]: return makewave(f,a,d, lib.SquareWave)
def KarplusStrongWave(f,a,d) -> array[float]: return makewave(f,a,d, lib.KarplusStrongWave)

import matplotlib.pyplot as plt
def plottimeseries(t,s):
    fig, ax = plt.subplots()
    ax.plot(t, s, 'y-',label="wave")
    ax.set_xlabel("time")
    ax.set_ylabel("value")
    plt.legend()
    plt.show()

samplerate = lib.WaveSize(1.) # number of samples in 1 second
print("samplerate = %d"%samplerate) 

f = 10. # Hz
a = 10.
d = 2.0 # seconds

def test01_sinewave():
    s = SineWave(f,a,d)
    t = [i/samplerate for i in range(len(s))]
    plottimeseries(t,s)

def test02_squarewave():
    s = SquareWave(f,a,d)
    t = [i/samplerate for i in range(len(s))]
    plottimeseries(t,s)

def test03_karplusStrongWave():
    s = KarplusStrongWave(f,a,d)
    t = [i/samplerate for i in range(len(s))]
    plottimeseries(t,s)

if __name__ == "__main__":
    test01_sinewave()
    test02_squarewave()
    test03_karplusStrongWave()