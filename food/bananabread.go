package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Bananabread struct{}

// AlwaysConsumable ...
func (Bananabread) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Bananabread) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Bananabread) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Bananabread) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Bananabread) Edible() bool {
	return true
}

// EncodeItem ...
func (Bananabread) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bananabread", 0
}

// Name ...
func (Bananabread) Name() string {
	return "Bananabread"
}

// Texture ...
func (Bananabread) Texture() image.Image {
	return textureFromName("loaf4")
}
