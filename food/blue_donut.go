package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueDonut struct{}

// AlwaysConsumable ...
func (BlueDonut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueDonut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueDonut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueDonut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueDonut) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueDonut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_donut", 0
}

// Name ...
func (BlueDonut) Name() string {
	return "Blue Donut"
}

// Texture ...
func (BlueDonut) Texture() image.Image {
	return textureFromName("blue_donut")
}
