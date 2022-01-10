package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cabbage struct{}

// AlwaysConsumable ...
func (Cabbage) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cabbage) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cabbage) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cabbage) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cabbage) Edible() bool {
	return true
}

// EncodeItem ...
func (Cabbage) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cabbage", 0
}

// Name ...
func (Cabbage) Name() string {
	return "Cabbage"
}

// Texture ...
func (Cabbage) Texture() image.Image {
	return textureFromName("cabbage")
}
