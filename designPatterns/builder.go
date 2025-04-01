package designpatterns

type Destroyer struct {
	MainGun string
	Torpedo string
	AA      string
}

type DestroyerBuilder interface {
	BuildMainGun(v string)
	BuildTorpedo(v string)
	BuildAA(v string)
	Reset()
}

type DruidBuilder struct {
	Ship Destroyer
}

func (b *DruidBuilder) Reset() {
	b.Ship = Destroyer{}
}

func (b *DruidBuilder) BuildMainGun(v string) {
	b.Ship.MainGun = v
}

func (b *DruidBuilder) BuildTorpedo(v string) {
	b.Ship.Torpedo = v
}

func (b *DruidBuilder) BuildAA(v string) {
	b.Ship.AA = v
}

func (b *DruidBuilder) GetDestroyer() Destroyer {
	destroyer := b.Ship
	b.Reset()
	return destroyer
}

type Director struct {
	Builder DestroyerBuilder
}

func NewDirector(b DestroyerBuilder) *Director {
	return &Director{Builder: b}
}

func (d *Director) Build(class string) {
	switch class {
	case "Druid":
		d.Builder.Reset()
		d.Builder.BuildMainGun("127mm, 2*2")
		d.Builder.BuildAA("low")
		// 德鲁伊没有鱼雷
	}
}
