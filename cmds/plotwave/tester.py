# coding: utf8

import waveclt

import numpy as np
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

def test01_sinewave():
    s = waveclt.SineWave(f,a,d)
    t = [i/waveclt.samplerate for i in range(len(s))]
    plottimeseries(t,s)

def test02_squarewave():
    s = waveclt.SquareWave(f,a,d)
    t = [i/waveclt.samplerate for i in range(len(s))]
    plottimeseries(t,s)

def test03_karplusStrongWave():
    s = waveclt.KarplusStrongWave(f,a,d)
    t = [i/waveclt.samplerate for i in range(len(s))]
    plottimeseries(t,s)

def test04_dissonance():
    d = 4.
    s1 = waveclt.SineWave(f,a,d)
    s2 = waveclt.SineWave(3.*f/2.,a,d)
    s3 = waveclt.SineWave(3.1*f/2.,a,d)
    t = [i/waveclt.samplerate for i in range(len(s1))]

    a12 = np.add(s1,s2)
    plottimeseries(t,a12)

    a13 = np.add(s1,s3)
    plottimeseries(t,a13)

def test05_vibrato():
    d = 4.
    deltaf = 6.
    s1 = waveclt.SineWave(f,a,d)
    s2 = waveclt.SineWave(f+deltaf,0.2*a,d)
    t = [i/waveclt.samplerate for i in range(len(s1))]

    a12 = np.add(s1,s2)
    plottimeseries(t,a12)

if __name__ == "__main__":
    #test01_sinewave()
    #test02_squarewave()
    #test03_karplusStrongWave()
    #test04_dissonance()
    test05_vibrato()