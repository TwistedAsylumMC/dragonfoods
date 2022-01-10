package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedCottonCandy struct{}

// AlwaysConsumable ...
func (RedCottonCandy) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedCottonCandy) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedCottonCandy) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedCottonCandy) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedCottonCandy) Edible() bool {
	return true
}

// EncodeItem ...
func (RedCottonCandy) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_cotton_candy", 0
}

// Name ...
func (RedCottonCandy) Name() string {
	return "Red Cotton Candy"
}

// Texture ...
func (RedCottonCandy) Texture() image.Image {
	return textureFromName("cottoncandy5")
}
