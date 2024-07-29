package levels

type LevelMap struct {
	Heightmap [][]int
	Objectmap [][]int
}

func NewLevelMap(heightmap, objectmap [][]int) *LevelMap {
	return &LevelMap{
		Heightmap: heightmap,
		Objectmap: objectmap,
	}
}
