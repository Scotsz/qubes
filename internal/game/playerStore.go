package game

import (
	"qubes/internal/model"
	"sync"
)

type PlayerStore struct {
	mu      sync.Mutex
	players map[model.ClientID]*Player
}

func NewPlayerStore() *PlayerStore {
	return &PlayerStore{
		players: make(map[model.ClientID]*Player),
	}
}
func (p *PlayerStore) Add(id model.ClientID, player *Player) {
	p.mu.Lock()
	p.players[id] = player
	p.mu.Unlock()

}
func (p *PlayerStore) Remove(id model.ClientID) {
	p.mu.Lock()
	delete(p.players, id)
	p.mu.Unlock()

}
func (p *PlayerStore) Get(id model.ClientID) (*Player, error) {
	p.mu.Lock()
	player := p.players[id]
	p.mu.Unlock()

	return player, nil
}

func (p *PlayerStore) All() map[model.ClientID]*Player {
	p.mu.Lock()
	ps := p.players
	p.mu.Unlock()
	return ps
}
