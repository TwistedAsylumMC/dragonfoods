package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Taco struct{}

// AlwaysConsumable ...
func (Taco) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Taco) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Taco) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Taco) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Taco) Edible() bool {
	return true
}

// EncodeItem ...
func (Taco) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:taco", 0
}

// Name ...
func (Taco) Name() string {
	return "Taco"
}

// Texture ...
func (Taco) Texture() image.Image {
	return textureFromName("taco")
}
