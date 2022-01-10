package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pho struct{}

// AlwaysConsumable ...
func (Pho) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pho) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pho) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pho) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pho) Edible() bool {
	return true
}

// EncodeItem ...
func (Pho) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pho", 0
}

// Name ...
func (Pho) Name() string {
	return "Pho"
}

// Texture ...
func (Pho) Texture() image.Image {
	return textureFromName("pho")
}
