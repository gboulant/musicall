# Synthetic - Create synthetic sounds with go

The core package is [wave](wave) that defines the concept of Synthesizer
to generates tha samples of a wave form. Then comes the package
[sound](sound) that can creates a sound from a wave samples. And the
package [music](music) that introduces the concept of notes (Do, Ré, Mi,
etc). The package [guitar](guitar) is a special package for playing
notes with a synthesizer that emulates the a guitar timbre (Karplus
Strong algotithm).

The guitar package is inspired from the [timiskhakov
music](https://github.com/timiskhakov/music) project, with an improvment
that consists in calculating the whole set of frequencies from the La3
reference instead of tabulate all the possible frequencies of all
couples (string number, fret number).
