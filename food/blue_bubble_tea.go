package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueBubbleTea struct{}

// AlwaysConsumable ...
func (BlueBubbleTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueBubbleTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueBubbleTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueBubbleTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueBubbleTea) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueBubbleTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_bubble_tea", 0
}

// Name ...
func (BlueBubbleTea) Name() string {
	return "BlueBubbleTea"
}

// Texture ...
func (BlueBubbleTea) Texture() image.Image {
	return textureFromName("btea2")
}
