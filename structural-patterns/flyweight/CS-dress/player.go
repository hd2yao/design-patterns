package main

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
