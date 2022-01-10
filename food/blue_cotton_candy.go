package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueCottonCandy struct{}

// AlwaysConsumable ...
func (BlueCottonCandy) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueCottonCandy) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueCottonCandy) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueCottonCandy) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueCottonCandy) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueCottonCandy) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_cotton_candy", 0
}

// Name ...
func (BlueCottonCandy) Name() string {
	return "Blue Cotton Candy"
}

// Texture ...
func (BlueCottonCandy) Texture() image.Image {
	return textureFromName("cottoncandy2")
}
