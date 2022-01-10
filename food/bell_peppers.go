package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BellPeppers struct{}

// AlwaysConsumable ...
func (BellPeppers) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BellPeppers) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BellPeppers) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BellPeppers) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BellPeppers) Edible() bool {
	return true
}

// EncodeItem ...
func (BellPeppers) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bell_peppers", 0
}

// Name ...
func (BellPeppers) Name() string {
	return "BellPeppers"
}

// Texture ...
func (BellPeppers) Texture() image.Image {
	return textureFromName("bell_pepper")
}
