package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Orange struct{}

// AlwaysConsumable ...
func (Orange) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Orange) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Orange) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Orange) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Orange) Edible() bool {
	return true
}

// EncodeItem ...
func (Orange) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:orange", 0
}

// Name ...
func (Orange) Name() string {
	return "Orange"
}

// Texture ...
func (Orange) Texture() image.Image {
	return textureFromName("orange")
}
