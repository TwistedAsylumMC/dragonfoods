package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BubbleTea struct{}

// AlwaysConsumable ...
func (BubbleTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BubbleTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BubbleTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BubbleTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BubbleTea) Edible() bool {
	return true
}

// EncodeItem ...
func (BubbleTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bubble_tea", 0
}

// Name ...
func (BubbleTea) Name() string {
	return "Bubble Tea"
}

// Texture ...
func (BubbleTea) Texture() image.Image {
	return textureFromName("btea")
}
