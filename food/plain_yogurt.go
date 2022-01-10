package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PlainYogurt struct{}

// AlwaysConsumable ...
func (PlainYogurt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PlainYogurt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PlainYogurt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (PlainYogurt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PlainYogurt) Edible() bool {
	return true
}

// EncodeItem ...
func (PlainYogurt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:plain_yogurt", 0
}

// Name ...
func (PlainYogurt) Name() string {
	return "Plain Yogurt"
}

// Texture ...
func (PlainYogurt) Texture() image.Image {
	return textureFromName("plain_yogurt")
}
