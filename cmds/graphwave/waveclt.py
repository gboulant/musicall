# coding: utf-8

import os
from ctypes import cdll, c_double, POINTER

libpath = os.path.join(os.path.dirname(__file__), "wavelib.so")
lib = cdll.LoadLibrary(libpath)

lib.WaveSize.argtypes = [c_double]
lib.SineWave.argtypes = [c_double, c_double, c_double, POINTER(c_double)]
lib.SquareWave.argtypes = [c_double, c_double, c_double, POINTER(c_double)]
lib.KarplusStrongWave.argtypes = [c_double, c_double, c_double, POINTER(c_double)]

from typing import List
from array import array

def makewave(f,a,d, wavefct) -> List[float]:
    N = lib.WaveSize(d)
    out = array('d', [0 for i in range(N)]) 
    out_ptr = (c_double * len(out)).from_buffer(out)
    wavefct(f,a,d, out_ptr)
    return out

def SineWave(f,a,d) -> List[float]: return makewave(f,a,d, lib.SineWave)
def SquareWave(f,a,d) -> List[float]: return makewave(f,a,d, lib.SquareWave)
def KarplusStrongWave(f,a,d) -> List[float]: return makewave(f,a,d, lib.KarplusStrongWave)

samplerate = lib.WaveSize(1.) # number of samples in 1 second
