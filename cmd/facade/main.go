package main

import "fmt"

type AudioPlayer struct {
}

type VideoPlayer struct {
	name string
}

type ScreenManager struct {
}

func (v *VideoPlayer) playVideo() {
	fmt.Println("Playing video " + v.name)

}

func (a *AudioPlayer) playAudio() {
	fmt.Println("Playing audio")
}

func (s *ScreenManager) showScreen() {
	fmt.Println("Showing screen")
}

type MultimediaFacade struct {
	a *AudioPlayer
	v *VideoPlayer
	s *ScreenManager
}

func NewMultimediaFacade() *MultimediaFacade {
	return &MultimediaFacade{
		v: &VideoPlayer{name: "name"},
		a: &AudioPlayer{},
		s: &ScreenManager{},
	}
}

func (m *MultimediaFacade) playMovie() {
	m.a.playAudio()
	m.v.playVideo()
	m.s.showScreen()
}

func main() {
	mf := NewMultimediaFacade()
	mf.playMovie()
}
