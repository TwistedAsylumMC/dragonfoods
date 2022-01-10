package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Nachos struct{}

// AlwaysConsumable ...
func (Nachos) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Nachos) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Nachos) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Nachos) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Nachos) Edible() bool {
	return true
}

// EncodeItem ...
func (Nachos) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:nachos", 0
}

// Name ...
func (Nachos) Name() string {
	return "Nachos"
}

// Texture ...
func (Nachos) Texture() image.Image {
	return textureFromName("72_nacho_dish")
}
