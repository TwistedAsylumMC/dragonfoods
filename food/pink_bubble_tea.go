package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PinkBubbleTea struct{}

// AlwaysConsumable ...
func (PinkBubbleTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PinkBubbleTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PinkBubbleTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (PinkBubbleTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PinkBubbleTea) Edible() bool {
	return true
}

// EncodeItem ...
func (PinkBubbleTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pink_bubble_tea", 0
}

// Name ...
func (PinkBubbleTea) Name() string {
	return "Pink Bubble Tea"
}

// Texture ...
func (PinkBubbleTea) Texture() image.Image {
	return textureFromName("btea6")
}
