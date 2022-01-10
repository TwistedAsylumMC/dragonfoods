package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Corn struct{}

// AlwaysConsumable ...
func (Corn) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Corn) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Corn) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Corn) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Corn) Edible() bool {
	return true
}

// EncodeItem ...
func (Corn) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:corn", 0
}

// Name ...
func (Corn) Name() string {
	return "Corn"
}

// Texture ...
func (Corn) Texture() image.Image {
	return textureFromName("corn")
}
