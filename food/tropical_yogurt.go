package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type TropicalYogurt struct{}

// AlwaysConsumable ...
func (TropicalYogurt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (TropicalYogurt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (TropicalYogurt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (TropicalYogurt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (TropicalYogurt) Edible() bool {
	return true
}

// EncodeItem ...
func (TropicalYogurt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tropical_yogurt", 0
}

// Name ...
func (TropicalYogurt) Name() string {
	return "Tropical Yogurt"
}

// Texture ...
func (TropicalYogurt) Texture() image.Image {
	return textureFromName("tropicicecream")
}
