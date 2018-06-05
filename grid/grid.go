package grid

import (
	fmt "fmt"
	rand "math/rand"

	u "github.com/apottr/milsim/unit"
)

//Weather establishes weather object
type Weather struct {
	state int
	temp  int
	day   bool
}

func (w *Weather) String() string {
	eweather := []string{
		"Clear sky",
		"Few clouds",
		"Scattered clouds",
		"Broken clouds",
		"Rain shower",
		"Rain",
		"Thunderstorm",
		"Snow",
		"Mist",
	}
	return fmt.Sprintf(`
			State: %s
			Temperature: %dC
			Is Day: %v
	`, eweather[w.state], w.temp, w.day)
}

//Tile struct for handling geospatial tiles
type Tile struct {
	utmCoordinate string
	weatherState  *Weather
	units         []*u.Force
}

//CreateTile creates geospatial tile
func CreateTile(coords string, units []*u.Force) *Tile {
	return &Tile{
		coords,
		&Weather{
			rand.Intn(8),
			rand.Intn(100),
			true,
		},
		units,
	}
}

//Pp pretty prints tile obj
func (t *Tile) Pp() {
	fmt.Printf(`
		Coordinates: %s
		Weather: %s
		Unit count: %d 
	`, t.utmCoordinate, t.weatherState.String(), len(t.units))
	fmt.Println("Units in Tile:")
	for _, unit := range t.units {
		unit.Pp()
	}
}
