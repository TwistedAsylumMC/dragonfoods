package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenCottonCandy struct{}

// AlwaysConsumable ...
func (GreenCottonCandy) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenCottonCandy) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenCottonCandy) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenCottonCandy) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenCottonCandy) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenCottonCandy) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_cotton_candy", 0
}

// Name ...
func (GreenCottonCandy) Name() string {
	return "Green Cotton Candy"
}

// Texture ...
func (GreenCottonCandy) Texture() image.Image {
	return textureFromName("cottoncandy3")
}
