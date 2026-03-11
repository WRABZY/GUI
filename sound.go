package gui

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type InterruptableSound struct {
	player *audio.Player
}

func NewInterruptableSound(sound *wav.Stream) *InterruptableSound {
	audioPlayer, err := audio.NewContext(48000).NewPlayerF32(sound)
	if err != nil {
		log.Fatal(err)
	}
	return &InterruptableSound{audioPlayer}
}

func (s *InterruptableSound) Volume(volume float64) *InterruptableSound {
	s.player.SetVolume(volume)
	return s
}

func (s *InterruptableSound) Play() {
	if s.player == nil {
		return
	}

	if s.player.IsPlaying() {
		s.player.Pause()
	}

	if err := s.player.Rewind(); err != nil {
		log.Println(err)
		return
	}

	s.player.Play()
}
