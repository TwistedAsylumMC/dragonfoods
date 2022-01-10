package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Donut struct{}

// AlwaysConsumable ...
func (Donut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Donut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Donut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Donut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Donut) Edible() bool {
	return true
}

// EncodeItem ...
func (Donut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:donut", 0
}

// Name ...
func (Donut) Name() string {
	return "Donut"
}

// Texture ...
func (Donut) Texture() image.Image {
	return textureFromName("34_donut")
}
