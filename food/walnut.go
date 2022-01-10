package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Walnut struct{}

// AlwaysConsumable ...
func (Walnut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Walnut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Walnut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Walnut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Walnut) Edible() bool {
	return true
}

// EncodeItem ...
func (Walnut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:walnut", 0
}

// Name ...
func (Walnut) Name() string {
	return "Walnut"
}

// Texture ...
func (Walnut) Texture() image.Image {
	return textureFromName("walnut")
}
