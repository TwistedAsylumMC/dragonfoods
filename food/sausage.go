package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Sausage struct{}

// AlwaysConsumable ...
func (Sausage) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Sausage) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Sausage) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(9, 5.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Sausage) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Sausage) Edible() bool {
	return true
}

// EncodeItem ...
func (Sausage) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sausage", 0
}

// Name ...
func (Sausage) Name() string {
	return "Sausage"
}

// Texture ...
func (Sausage) Texture() image.Image {
	return textureFromName("sausge")
}
