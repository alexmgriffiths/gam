package tiles

type Tileset struct {
	Name    string
	Tilemap *Tilemap
}

func NewTileset(name string, tilemap *Tilemap) *Tileset {
	return &Tileset{
		Name:    name,
		Tilemap: tilemap,
	}
}
