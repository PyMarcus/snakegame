package snake

type Point struct{
	x, y int
}

func (p *Point) SetX(value int){
	p.x = value
}

func (p *Point) SetY(value int){
	p.y = value
}
