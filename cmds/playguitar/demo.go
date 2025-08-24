package main

import (
	"log"

	"galuma.net/synthetic/guitar"
	"galuma.net/synthetic/sound"
	"galuma.net/synthetic/wave"
	"github.com/gopxl/beep"
)

const sampleRate = wave.DefaultSampleRate

func init() {
	// Le speaker est initialisé avec un sample rate fixé. Tous les
	// signaux ([]float64) joués par ce speaker seront considérés comme
	// des sons avec ce sample rate. On doit donc générer des signaux
	// avec ce sample rate.
	err := sound.Init(sampleRate)
	if err != nil {
		log.Fatal(err)
	}
}

func chord(name string) guitar.Chord {
	return guitar.StandardChord(name)
}

var (
	Do  = chord("Do")
	Re  = chord("Re")
	Mi  = chord("Mi")
	Mim = chord("Mim")
	Fa  = chord("Fa")
	Sol = chord("Sol")
	La  = chord("La")
	Lam = chord("Lam")
)

// -----------------------------------------------------------
// Technical/Exercice examples

// Toutes les cordes à vide
func T01_play_open_strings() error {
	// We first create a guitar
	g := guitar.NewGuitar(sampleRate)

	// Then we strike some notes on this guitar. The sequence is
	// registered in a streamer
	s := beep.Seq(
		g.Silence(0.5),
		g.Pluck(guitar.Note{StringNum: 6, FretNum: 0}, 1), // 6 is the Lower frequency string (Low Mi)
		g.Pluck(guitar.Note{StringNum: 5, FretNum: 0}, 1), // 5 is the string named La
		g.Pluck(guitar.Note{StringNum: 4, FretNum: 0}, 1), // 4 is the string named Ré
		g.Pluck(guitar.Note{StringNum: 3, FretNum: 0}, 1), // 3 is the string named Sol
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, 1), // 2 is the string named Si
		g.Pluck(guitar.Note{StringNum: 1, FretNum: 0}, 1), // 1 is the Higher frequency string (High Mi)
	)

	// Then play the resulting streamer
	return sound.Play(s)
}

// Les accords de base principaux
func T02_main_chords() error {
	g := guitar.NewGuitar(sampleRate)

	duration := 0.8 // duration of the chord
	delay := 0.04   // delay between the string plucks
	labelledChord := func(label string) beep.Streamer {
		chord := guitar.StandardChord(label)
		stream := g.Chord(chord, duration, delay)
		return sound.LabelledStreamer(stream, label)
	}

	s := beep.Seq(
		g.Silence(0.5),
		labelledChord("Do"),
		labelledChord("Re"),
		labelledChord("Mi"),
		labelledChord("Mim"),
		labelledChord("Fa"),
		labelledChord("Sol"),
		labelledChord("La"),
		labelledChord("Lam"),
	)

	// Then play the resulting streamer
	return sound.Play(s)
}

// Gamme pentatonique en La
func T03_pentatonic_scale_La() error {
	g := guitar.NewGuitar(sampleRate)

	// Then we strike some notes on this guitar. The sequence is
	// registered in a streamer
	duration := 0.8
	gamme := func(fret guitar.FretNumber) beep.Streamer {
		return beep.Seq(
			g.Silence(0.5),
			g.Pluck(guitar.Note{StringNum: 6, FretNum: fret}, duration),
			g.Pluck(guitar.Note{StringNum: 6, FretNum: fret + 3}, duration),
			g.Pluck(guitar.Note{StringNum: 5, FretNum: fret}, duration),
			g.Pluck(guitar.Note{StringNum: 5, FretNum: fret + 2}, duration),
			g.Pluck(guitar.Note{StringNum: 4, FretNum: fret}, duration),
			g.Pluck(guitar.Note{StringNum: 4, FretNum: fret + 2}, duration),
			g.Pluck(guitar.Note{StringNum: 3, FretNum: fret}, duration),
			g.Pluck(guitar.Note{StringNum: 3, FretNum: fret + 2}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: fret}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: fret + 3}, duration),
			g.Pluck(guitar.Note{StringNum: 1, FretNum: fret}, duration),
			g.Pluck(guitar.Note{StringNum: 1, FretNum: fret + 3}, duration),
		)
	}
	s := gamme(5)

	// Then play the resulting streamer
	return sound.Play(s)
}

// -----------------------------------------------------------
// Songs examples

// Nocking on the heaven's door
func D01_Nocking_on_the_heavens_door() error {
	duration := 0.8 // duration of the chord
	delay := 0.04   // delay between the string plucks

	g := guitar.NewGuitar(sampleRate)
	s := beep.Seq(
		g.Silence(0.5),

		g.Chord(Sol, duration, delay),
		g.Chord(Sol, duration, delay),
		g.Chord(Re, duration, delay),
		g.Chord(Re, duration, delay),
		g.Chord(Lam, duration, delay),
		g.Chord(Lam, duration, delay),
		g.Chord(Lam, duration, delay),
		g.Chord(Lam, duration, delay),

		g.Chord(Sol, duration, delay),
		g.Chord(Sol, duration, delay),
		g.Chord(Re, duration, delay),
		g.Chord(Re, duration, delay),
		g.Chord(Do, duration, delay),
		g.Chord(Do, duration, delay),
		g.Chord(Do, duration, delay),
		g.Chord(Do, duration, delay),
	)

	return sound.Play(s)
}

func D02_U2_One() error {
	duration := 2. // duration of the chord
	delay := 0.1   // delay between the string plucks

	g := guitar.NewGuitar(sampleRate)
	s := beep.Seq(
		g.Silence(0.5),

		g.Chord(Do, duration, delay),
		g.Chord(Lam, duration, delay),
		g.Chord(Fa, duration, delay),
		g.Chord(Do, duration, delay),
	)

	return sound.Play(s)
}

// ACDC - Thunderstruck
func D03_ACDC_Thunderstruck() error {
	duration := 0.12 // duration of the chord
	g := guitar.NewGuitar(sampleRate)

	intro := func(fret guitar.FretNumber) beep.Streamer {
		return beep.Seq(
			g.Pluck(guitar.Note{StringNum: 2, FretNum: fret}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: fret + 3}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
		)
	}

	riff := func() beep.Streamer {
		return beep.Seq(
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 12}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 10}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),

			g.Pluck(guitar.Note{StringNum: 2, FretNum: 9}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 10}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),

			g.Pluck(guitar.Note{StringNum: 2, FretNum: 9}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 7}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),

			g.Pluck(guitar.Note{StringNum: 2, FretNum: 9}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 5}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),

			g.Pluck(guitar.Note{StringNum: 2, FretNum: 7}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 4}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),

			g.Pluck(guitar.Note{StringNum: 2, FretNum: 5}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 4}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 5}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 4}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 5}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 4}, duration),
			g.Pluck(guitar.Note{StringNum: 2, FretNum: 0}, duration),
		)
	}

	s := beep.Seq(
		g.Silence(0.5),
		intro(4), intro(4), intro(4), intro(4),
		intro(5), intro(5), intro(5), intro(5),
		intro(4), intro(4), intro(4), intro(4),
		intro(5), intro(5), intro(5), intro(5),
		riff(), riff(),
	)

	return sound.Play(s)
}

// Tostaky
func D04_NoirDesir_Tostaky() error {
	g := guitar.NewGuitar(sampleRate)

	chord := []guitar.Note{
		{StringNum: 5, FretNum: 7},
		{StringNum: 4, FretNum: 9},
		{StringNum: 3, FretNum: 9},
	}

	tempo := 0.25
	duration1 := 1.0 * tempo // duration of the chord
	duration2 := 1.5 * tempo
	duration3 := 0.5 * tempo
	delay := 0.04 // delay between the string plucks

	riff := func() beep.Streamer {
		return beep.Seq(
			g.Silence((duration3)),
			g.Pluck(guitar.Note{StringNum: 6, FretNum: 0}, duration1),
			g.Pluck(guitar.Note{StringNum: 6, FretNum: 0}, duration1),
			g.Chord(chord, duration2, delay),
			g.Silence(0.1),
			g.Pluck(guitar.Note{StringNum: 1, FretNum: 8}, duration3),
			g.Pluck(guitar.Note{StringNum: 1, FretNum: 7}, duration3),
			g.Pluck(guitar.Note{StringNum: 1, FretNum: 8}, duration3),
			g.Pluck(guitar.Note{StringNum: 1, FretNum: 9}, duration3),
			g.Pluck(guitar.Note{StringNum: 1, FretNum: 10}, duration3),
			g.Pluck(guitar.Note{StringNum: 1, FretNum: 10}, duration2),
		)
	}
	s := beep.Seq(
		g.Silence(0.5),
		riff(), riff(), riff(), riff(),
		g.Silence(0.5),
	)

	// Then we play the resulting streamer
	return sound.Play(s)
}

func D05_NoirDesir_Un_jour_en_France() error {
	g := guitar.NewGuitar(sampleRate)

	powerchord := func(fret guitar.FretNumber) []guitar.Note {
		return []guitar.Note{
			{StringNum: 6, FretNum: fret},
			{StringNum: 5, FretNum: fret + 2},
		}
	}

	tempo := 0.3
	duration := 1.0 * tempo // duration of the chord
	delay := 0.1 * tempo    // delay between the string plucks

	seq3 := func(fret guitar.FretNumber) beep.Streamer {
		return beep.Seq(
			g.Chord(powerchord(fret), duration, delay),
			g.Chord(powerchord(fret), duration, delay),
			g.Chord(powerchord(fret), duration, delay),
		)
	}

	seq4 := func(fret guitar.FretNumber) beep.Streamer {
		return beep.Seq(
			g.Chord(powerchord(fret), duration, delay),
			g.Chord(powerchord(fret), duration, delay),
			g.Chord(powerchord(fret), duration, delay),
			g.Chord(powerchord(fret), duration, delay),
		)
	}

	s := beep.Seq(
		g.Silence(2*tempo),
		g.Chord(powerchord(5), duration, delay),
		seq4(6), seq4(5), seq4(10), seq4(5), seq4(6), seq4(5),
		seq4(1), seq3(1), g.Chord(powerchord(5), duration, delay),
		seq4(6), seq4(5), seq4(10), seq4(5), seq4(6), seq4(5),
		seq4(1), seq4(10), seq3(10), g.Chord(powerchord(5), duration, delay),
		seq4(6), seq4(5), seq4(10), seq4(5), seq4(6), seq4(5),
		seq4(1), seq3(1), g.Chord(powerchord(5), duration, delay),
		g.Silence(4*tempo),
	)

	// Then we play the resulting streamer
	return sound.Play(s)
}

// Le vent l'emportera, pont
func D06_NoirDesir_Le_vent_l_emportera() error {
	g := guitar.NewGuitar(sampleRate)

	tempo := 1.
	duration := 1.0 * tempo // duration of the pluck

	sequence := func() beep.Streamer {
		return beep.Seq(
			g.Pluck(guitar.Note{StringNum: 6, FretNum: 0}, duration),
		)
	}

	s := beep.Seq(
		g.Silence(2*tempo),
		sequence(),
	)

	// Then we play the resulting streamer
	return sound.Play(s)
}

// Bloody Sunday, U2
func D07_U2_Bloody_Sunday() error {
	duration := 0.4 // duration of the pluck
	g := guitar.NewGuitar(sampleRate)
	s := beep.Seq(
		g.Silence(0.5),

		g.Pluck(guitar.Note{StringNum: 3, FretNum: 4}, duration),
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 3}, duration),
		g.Pluck(guitar.Note{StringNum: 1, FretNum: 2}, duration),
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 3}, duration),

		g.Pluck(guitar.Note{StringNum: 3, FretNum: 2}, duration),
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 3}, duration),
		g.Pluck(guitar.Note{StringNum: 1, FretNum: 2}, duration),
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 3}, duration),

		g.Pluck(guitar.Note{StringNum: 3, FretNum: 0}, duration),
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 3}, duration),
		g.Pluck(guitar.Note{StringNum: 1, FretNum: 0}, duration),
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 3}, duration),

		g.Pluck(guitar.Note{StringNum: 3, FretNum: 0}, duration),
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 3}, duration*0.5),
		g.Pluck(guitar.Note{StringNum: 1, FretNum: 0}, duration*1.5),
		g.Pluck(guitar.Note{StringNum: 2, FretNum: 3}, duration),
	)

	// Then we play the resulting streamer
	return sound.Play(s)
}

// Play a rythm Bas & Bas - Haut & Haut - Bas
func D08_Rythm_UpDown() error {
	duration := 0.3 // duration of the chord
	delay := 0.04   // delay between the string plucks
	g := guitar.NewGuitar(sampleRate)

	couplet := func(chord []guitar.Note) beep.Streamer {
		chordUp := guitar.Reverse(chord)
		return beep.Seq(
			g.Chord(chord, duration*2, delay),
			g.Chord(chord, duration*1, delay),
			g.Chord(chordUp, duration*2, delay),
			g.Chord(chordUp, duration*1, delay),
			g.Chord(chord, duration*2, delay),
		)
	}

	s := beep.Seq(
		g.Silence(0.5),
		couplet(Sol),
		couplet(Re),
		couplet(Lam),
		couplet(Lam),
	)

	// Then we play the resulting streamer
	return sound.Play(s)
}

func D09_Johnny_Cash_Hurt() error {
	g1 := guitar.NewGuitar(sampleRate)
	ks1 := guitar.NewKarplusStrongSynthesizer(0., 1.5, 0.1, sampleRate)
	g1.UseSynthesizer(ks1)

	g2 := guitar.NewGuitar(sampleRate)
	ks2 := guitar.NewKarplusStrongSynthesizer(0., 1., 0.05, sampleRate)
	g2.UseSynthesizer(ks2)

	s := beep.Mix(hurtLead(g1), hurtRhythm(g2))
	return sound.Play(s)
}
