package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedBubbleTea struct{}

// AlwaysConsumable ...
func (RedBubbleTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedBubbleTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedBubbleTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedBubbleTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedBubbleTea) Edible() bool {
	return true
}

// EncodeItem ...
func (RedBubbleTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_bubble_tea", 0
}

// Name ...
func (RedBubbleTea) Name() string {
	return "Red Bubble Tea"
}

// Texture ...
func (RedBubbleTea) Texture() image.Image {
	return textureFromName("btea4")
}
