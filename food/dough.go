package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Dough struct{}

// AlwaysConsumable ...
func (Dough) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Dough) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Dough) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Dough) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Dough) Edible() bool {
	return true
}

// EncodeItem ...
func (Dough) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:dough", 0
}

// Name ...
func (Dough) Name() string {
	return "Dough"
}

// Texture ...
func (Dough) Texture() image.Image {
	return textureFromName("dough")
}
