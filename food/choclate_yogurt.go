package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateYogurt struct{}

// AlwaysConsumable ...
func (ChoclateYogurt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateYogurt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateYogurt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateYogurt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateYogurt) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateYogurt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_yogurt", 0
}

// Name ...
func (ChoclateYogurt) Name() string {
	return "Choclate Yogurt"
}

// Texture ...
func (ChoclateYogurt) Texture() image.Image {
	return textureFromName("chocicecream")
}
