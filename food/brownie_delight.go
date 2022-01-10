package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BrownieDelight struct{}

// AlwaysConsumable ...
func (BrownieDelight) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BrownieDelight) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BrownieDelight) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BrownieDelight) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BrownieDelight) Edible() bool {
	return true
}

// EncodeItem ...
func (BrownieDelight) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:brownie_delight", 0
}

// Name ...
func (BrownieDelight) Name() string {
	return "BrownieDelight"
}

// Texture ...
func (BrownieDelight) Texture() image.Image {
	return textureFromName("brownie_delight")
}
