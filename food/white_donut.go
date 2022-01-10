package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type WhiteDonut struct{}

// AlwaysConsumable ...
func (WhiteDonut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (WhiteDonut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (WhiteDonut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (WhiteDonut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (WhiteDonut) Edible() bool {
	return true
}

// EncodeItem ...
func (WhiteDonut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:white_donut", 0
}

// Name ...
func (WhiteDonut) Name() string {
	return "White Donut"
}

// Texture ...
func (WhiteDonut) Texture() image.Image {
	return textureFromName("36_donut")
}
