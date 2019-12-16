package kingdoms

type kingdom struct {
	Name   string
	Emblem string
}

var Kingdoms map[string]kingdom

func init() {
	// Possible allies only i.e. List of all nations of Southeros minus the Message sender Kingdom.
	Kingdoms = make(map[string]kingdom)
	Kingdoms["LAND"] = kingdom{Name: "LAND", Emblem: "PANDA",}
	Kingdoms["WATER"] = kingdom{"WATER", "OCTOPUS"}
	Kingdoms["ICE"] = kingdom{"ICE", "MAMMOTH"}
	Kingdoms["AIR"] = kingdom{"AIR", "OWL"}
	Kingdoms["FIRE"] = kingdom{"FIRE", "DRAGON"}
}

//NewKingdom Returns a pointer to the zero initialized struct
func NewKingdom(name string, emblem string) *kingdom {
	return &kingdom{Name: name, Emblem: emblem}
}

//GetKingdomsList Returns a List of possible ally kingdoms.
func GetKindomData(key string) kingdom {
	return Kingdoms[key]
}
