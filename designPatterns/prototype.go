package designpatterns

type Prototype interface {
	clone() Prototype
}

type BattleShip struct {
	MainGun string
}

func (b *BattleShip) clone() Prototype {
	return &BattleShip{
		MainGun: b.MainGun,
	}
}
