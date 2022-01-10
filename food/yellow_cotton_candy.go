package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type YellowCottonCandy struct{}

// AlwaysConsumable ...
func (YellowCottonCandy) AlwaysConsumable() bool {
	return false
}

// Category ...
func (YellowCottonCandy) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (YellowCottonCandy) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (YellowCottonCandy) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (YellowCottonCandy) Edible() bool {
	return true
}

// EncodeItem ...
func (YellowCottonCandy) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellow_cotton_candy", 0
}

// Name ...
func (YellowCottonCandy) Name() string {
	return "Yellow Cotton Candy"
}

// Texture ...
func (YellowCottonCandy) Texture() image.Image {
	return textureFromName("cottoncraft3")
}
