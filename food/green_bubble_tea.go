package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenBubbleTea struct{}

// AlwaysConsumable ...
func (GreenBubbleTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenBubbleTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenBubbleTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenBubbleTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenBubbleTea) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenBubbleTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_bubble_tea", 0
}

// Name ...
func (GreenBubbleTea) Name() string {
	return "GreenBubbleTea"
}

// Texture ...
func (GreenBubbleTea) Texture() image.Image {
	return textureFromName("btea3")
}
