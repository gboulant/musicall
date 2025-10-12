# MusicAll - Create synthetic sounds with go

This project is a tutorial framework that illustrates some basic
elements of the music theory. Basic knowledge of what is a sound wave
can be required (amplitude, frequency, timbre).

**contact**: [Guillaume Boulant](mailto:gboulant@gmail.com?subject=musicall)

## Description of the project

The motivation for this project came after reading the post [Programming
Guitar
Music](https://timiskhakov.github.io/posts/programming-guitar-music) and
studying the associated [timiskhakov
music](https://github.com/timiskhakov/music) project from Timor
Iskhakov.

I am also a software programmer an try to learn playing guitar. I
started this project (strongly inspired from the [timiskhakov
music](https://github.com/timiskhakov/music) project, many thanks to
him) to test some elements of the music theory, in particular how
the musical scales work, how the music notes are designed and related
to a physical sound wave, how to calculate the frequency of a guitar
pluck, given the string number and the fret number.

With this tutorial project, you can play with a set of executable
programs (in the folder [cmds](cmds) as traditional in the go programing
convention) that illustrates some basic elements of the music theory.
Basic knowledge of what is a sound wave is required (at least its
mathematical representation in terms of amplitude, frequency and
timbre).

Even if it is a tutorial project, the packages may be used in other
applications for other purposes (at least drawing wave forms, making
sounds, or playing music, with or without a guitar). Let me know.

## Quick start guide

A good way to start with this project is to play with the executable
programs in the [cmds](cmds) folder, in particular with:

* [cmds/examples](cmds/examples): play different sound examples based on
  the usage of the packages wave, sound and music.
* [cmds/playguitar](cmds/playguitar): play differents songs or musical
  scales with a sound timbre emulating the guitar sound (use the
  KarplusStrong synthesizer). These examples illustrate the usage of the
  guitar package.

For the examples:

```shell
cd cmds/examples
go build
./examples -l     # to display all possible examples
./examples -n D02 # to play the demo example D02 (vibrato)
```

For the guitar player:

```shell
cd cmds/playguitar
go build
./playguitar -l     # to display all possible examples
./playguitar -n D03 # to play the demo example D03 (ACDC, Thunderstruck)
```

Please, listen to the example D09 (Hurt from Johnny Cash), whose
tablature is taken from [timiskhakov
music](https://github.com/timiskhakov/music) project and played using
this package with very light adaptation (just the frequencies are
calculated instead of read from a table).

## Description of the packages

The core package [wave](wave) defines the concept of Synthesizer to
generates the samples of a wave form at a given sampling rate. Different
kind of synthesizer are implemented (SineWave, SquareWave, Triangle, Saw
Tooth, etc., and even a simple Karplus Strong synthesizer that emulates
the wave form of the sound of a guitar pluck). This package does not
make sound, only samples, i.e. a dataset of float values that represent
a timeseries at a given sample rate. The package provides a simple
plotter (based on the external package
[go-echarts](https://github.com/go-echarts/go-echarts)) for visual
control of the timeseries.

Then comes the package [sound](sound) that can create a sound from a
wave samples. This package relies on the external package
[beep](https://github.com/gopxl/beep). This package implements the low
level features for playing a sound from a sample dataset, as created for
example with the package [wave](wave).

The package [music](music) introduces the concept of notes (Do, Ré, Mi,
etc.). We show in this package how to calculate the frequency of a note,
characterized by an octave number and an index of the note in this
octave. This calculation is based on the tempered musical scale, in
which the interval between two notes is 1/12 of an octave.

The package [guitar](guitar) is a special package for playing notes with
a synthesizer that emulates the guitar timbre (Karplus Strong
algotithm). It also defines a special definition of a note,
characterized by the guitar string being plucked and the fret being
pressed. From this guitar note, we show how to determine a
[music](music) note in terms of octave and index as defined above, and
then calculate its frequency.

The [guitar](guitar) package is inspired from the [timiskhakov
music](https://github.com/timiskhakov/music) project, with a
modification that consists in calculating the frequency of a note by
derivation from the La3 reference, instead of picking the frequency
value in a table of all possible frequencies.

## Technical features

The executable program [cmds/plotwave](cmds/plotwave) shows how to plot
the wave form using the matplotlib python library. It is in fact a
technical demonstration of how to call go functions from a python context
using C go.
