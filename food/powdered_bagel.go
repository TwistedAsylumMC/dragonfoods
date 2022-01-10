package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PowderedBagel struct{}

// AlwaysConsumable ...
func (PowderedBagel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PowderedBagel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PowderedBagel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (PowderedBagel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PowderedBagel) Edible() bool {
	return true
}

// EncodeItem ...
func (PowderedBagel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:powdered_bagel", 0
}

// Name ...
func (PowderedBagel) Name() string {
	return "PowderedBagel"
}

// Texture ...
func (PowderedBagel) Texture() image.Image {
	return textureFromName("bagel4")
}
