package config

var KingdomEmblemMap map[string]string
var DefaultInputFile string

//type InputSrcType map[string]map[string]string
var DefaultOptions InputOptions

type InputOptions struct {
	Source string
	Meta   metaStruct
}

type metaStruct struct {
	FilePath string
}

func init() { // Initialize default / config values
	DefaultInputFile = "input1.txt"
	KingdomEmblemMap = map[string]string{
		"LAND" : "PANDA",
		"WATER" : "OCTOPUS",
		"ICE" : "MAMMOTH",
		"AIR" : "OWL",
		"FIRE" : "DRAGON",
	}
	// InputSrcType enables simple addition for additional input options. Currently it's file, we can add http req also
	// and so on. (Note: Each property of map is self explanatory).
	DefaultOptions = InputOptions{
		Source: "File",
		Meta:   metaStruct{FilePath: DefaultInputFile,},
	}
}


//	ICE MAMMOTH THVAO
//	FIRE DRAGON JXGMUT
//	AIR ZORRO
//	WATER OCTO VJAVWBZ PUS
//  => SPACE FIRE AIR WATER

//	AIR ROZO
//	LAND FAIJWJSOOFAMAU
//	ICE STHSTSTVSASOS
//  => SPACE AIR LAND ICE