package director

import (
	t "github.com/apottr/milsim/grid"
)

//Director handles worldstate
type Director struct {
	tiles [][][]*t.Tile
	time  int
}

//InitDirector initializes the director struct
func InitDirector(w int, h int) *Director {
	return &Director{}
}

//Pp pretty prints director object
func (d *Director) Pp() {

}

//TimeStep step world one time unit
func (d *Director) TimeStep() {
	ty := make([][]*t.Tile, len(d.tiles[d.time][0]))
	for y := range d.tiles[d.time] {
		tx := make([]*t.Tile, len(d.tiles[d.time]))
		for x := range d.tiles[d.time][y] {
			tile := d.tiles[d.time][y][x]
			tileNew := tile.Step()
			tx[x] = tileNew
		}
		ty[y] = tx
	}
	d.time++
	d.tiles = append(d.tiles, ty)
}
