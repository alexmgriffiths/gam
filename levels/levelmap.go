package levels

type LevelMap struct {
	Heightmap [][]int
	Objectmap [][]int
	Lightmap  [][]int
}

func NewLevelMap(heightmap, objectmap, lightmap [][]int) *LevelMap {
	return &LevelMap{
		Heightmap: heightmap,
		Objectmap: objectmap,
		Lightmap:  lightmap,
	}
}
