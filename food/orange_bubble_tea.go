package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type OrangeBubbleTea struct{}

// AlwaysConsumable ...
func (OrangeBubbleTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (OrangeBubbleTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (OrangeBubbleTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (OrangeBubbleTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (OrangeBubbleTea) Edible() bool {
	return true
}

// EncodeItem ...
func (OrangeBubbleTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:orange_bubble_tea", 0
}

// Name ...
func (OrangeBubbleTea) Name() string {
	return "OrangeBubbleTea"
}

// Texture ...
func (OrangeBubbleTea) Texture() image.Image {
	return textureFromName("btea5")
}
