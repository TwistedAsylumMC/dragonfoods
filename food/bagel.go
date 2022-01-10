package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Bagel struct{}

// AlwaysConsumable ...
func (Bagel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Bagel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Bagel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Bagel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Bagel) Edible() bool {
	return true
}

// EncodeItem ...
func (Bagel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bagel", 0
}

// Name ...
func (Bagel) Name() string {
	return "Bagel"
}

// Texture ...
func (Bagel) Texture() image.Image {
	return textureFromName("20_bagel")
}
