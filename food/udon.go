package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Udon struct{}

// AlwaysConsumable ...
func (Udon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Udon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Udon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Udon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Udon) Edible() bool {
	return true
}

// EncodeItem ...
func (Udon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:udon", 0
}

// Name ...
func (Udon) Name() string {
	return "Udon"
}

// Texture ...
func (Udon) Texture() image.Image {
	return textureFromName("udon")
}
