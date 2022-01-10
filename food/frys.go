package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Frys struct{}

// AlwaysConsumable ...
func (Frys) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Frys) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Frys) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Frys) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Frys) Edible() bool {
	return true
}

// EncodeItem ...
func (Frys) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:frys", 0
}

// Name ...
func (Frys) Name() string {
	return "Frys"
}

// Texture ...
func (Frys) Texture() image.Image {
	return textureFromName("45_frenchfries_dish")
}
