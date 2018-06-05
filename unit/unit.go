package unit

import (
	fmt "fmt"
)

//Level force level indicator
type Level int

/* force levels (in number of personnel)
squad     				   //8-14
platoon                    //15-45
company                    //80-150
battalion                  //300-800
regiment                   //1000-5500
division                   //10000-25000
corps                      //30000-50000
fieldArmy                  //100000-300000
armyGroup                  //2 or more FA
theater                    //4 or more AG
*/

//Force struct for establishing a unit of military force
type Force struct {
	level  int
	name   string
	count  int
	parent int
}

//CreateForce function for creating a unit of military force
func CreateForce(name string, count int) *Force {
	var lvl int
	if count >= 8 && count <= 14 {
		lvl = 0
	} else if count >= 15 && count <= 45 {
		lvl = 1
	} else if count >= 80 && count <= 150 {
		lvl = 2
	} else if count >= 300 && count <= 800 {
		lvl = 3
	} else if count >= 1000 && count <= 5500 {
		lvl = 4
	} else if count >= 10000 && count <= 25000 {
		lvl = 5
	} else if count >= 30000 && count <= 50000 {
		lvl = 6
	} else if count >= 100000 && count <= 300000 {
		lvl = 7
	}
	return &Force{
		level:  lvl,
		name:   name,
		count:  count,
		parent: 0,
	}
}

func (f *Force) String() string {
	enames := []string{
		"Squad",
		"Platoon",
		"Company",
		"Battalion",
		"Regiment",
		"Division",
		"Corps",
		"Field Army",
		"Army Group",
		"Theater",
	}
	return fmt.Sprintf(`
		Level: %s
		Name: %s
		Count: %d
		Parent: %d
	`, enames[f.level], f.name, f.count, f.parent)
}

//Pp pretty printer for force
func (f *Force) Pp() {
	fmt.Println(f.String())
}
