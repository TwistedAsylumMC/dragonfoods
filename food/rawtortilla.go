package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Rawtortilla struct{}

// AlwaysConsumable ...
func (Rawtortilla) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Rawtortilla) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Rawtortilla) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Rawtortilla) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Rawtortilla) Edible() bool {
	return true
}

// EncodeItem ...
func (Rawtortilla) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:rawtortilla", 0
}

// Name ...
func (Rawtortilla) Name() string {
	return "Rawtortilla"
}

// Texture ...
func (Rawtortilla) Texture() image.Image {
	return textureFromName("rawtortilla")
}
