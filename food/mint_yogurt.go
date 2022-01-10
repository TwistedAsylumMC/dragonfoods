package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MintYogurt struct{}

// AlwaysConsumable ...
func (MintYogurt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MintYogurt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MintYogurt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (MintYogurt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MintYogurt) Edible() bool {
	return true
}

// EncodeItem ...
func (MintYogurt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mint_yogurt", 0
}

// Name ...
func (MintYogurt) Name() string {
	return "Mint Yogurt"
}

// Texture ...
func (MintYogurt) Texture() image.Image {
	return textureFromName("mintjam")
}
