package main

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}

func (p *Player) DropItem(itemName string) {
	var i int
	for i = 0; i < len(p.Inventory); i++ {
		if p.Inventory[i].Name == itemName {
			break
		}
	}

	p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
}

func (p *Player) UseItem(item string) {

}

func main() {

}
