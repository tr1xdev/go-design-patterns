package main

import "fmt"

type State interface {
	ClickPlay()
	ClickLock()
	ClickNext()
	ClickPrevious()
}

type AudioPlayer struct {
	state     State
	isPlaying bool
}

func NewAudioPlayer() *AudioPlayer {
	player := &AudioPlayer{isPlaying: false}
	player.ChangeState(&ReadyState{player: player})
	return player
}

func (p *AudioPlayer) ChangeState(state State) {
	p.state = state
}

func (p *AudioPlayer) ClickPlay() {
	p.state.ClickPlay()
}

func (p *AudioPlayer) ClickLock() {
	p.state.ClickLock()
}

func (p *AudioPlayer) ClickNext() {
	p.state.ClickNext()
}

func (p *AudioPlayer) ClickPrevious() {
	p.state.ClickPrevious()
}

type ReadyState struct {
	player *AudioPlayer
}

func (s *ReadyState) ClickPlay() {
	s.player.isPlaying = true
	fmt.Println("Воспроизведение начато")
	s.player.ChangeState(&PlayingState{player: s.player})
}

func (s *ReadyState) ClickLock() {
	fmt.Println("Плеер заблокирован")
	s.player.ChangeState(&LockedState{player: s.player})
}

func (s *ReadyState) ClickNext() {
	fmt.Println("Переход к следующему треку")
}

func (s *ReadyState) ClickPrevious() {
	fmt.Println("Переход к предыдущему треку")
}

type PlayingState struct {
	player *AudioPlayer
}

func (s *PlayingState) ClickPlay() {
	s.player.isPlaying = false
	fmt.Println("Пауза")
	s.player.ChangeState(&ReadyState{player: s.player})
}

func (s *PlayingState) ClickLock() {
	fmt.Println("Плеер заблокирован во время воспроизведения")
	s.player.ChangeState(&LockedState{player: s.player})
}

func (s *PlayingState) ClickNext() {
	fmt.Println("Следующий трек (воспроизведение продолжается)")
}

func (s *PlayingState) ClickPrevious() {
	fmt.Println("Предыдущий трек (воспроизведение продолжается)")
}

type LockedState struct {
	player *AudioPlayer
}

func (s *LockedState) ClickPlay() {
	fmt.Println("Кнопки заблокированы, действие проигнорировано")
}

func (s *LockedState) ClickLock() {
	if s.player.isPlaying {
		fmt.Println("Плеер разблокирован (воспроизведение)")
		s.player.ChangeState(&PlayingState{player: s.player})
	} else {
		fmt.Println("Плеер разблокирован (готов)")
		s.player.ChangeState(&ReadyState{player: s.player})
	}
}

func (s *LockedState) ClickNext() {
	fmt.Println("Кнопки заблокированы, действие проигнорировано")
}

func (s *LockedState) ClickPrevious() {
	fmt.Println("Кнопки заблокированы, действие проигнорировано")
}

func main() {
	player := NewAudioPlayer()

	player.ClickPlay()
	player.ClickNext()
	player.ClickLock()
	player.ClickPlay()
	player.ClickLock()
	player.ClickPlay()
}
