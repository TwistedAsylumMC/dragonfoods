package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Can struct{}

// AlwaysConsumable ...
func (Can) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Can) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Can) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Can) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Can) Edible() bool {
	return true
}

// EncodeItem ...
func (Can) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:can", 0
}

// Name ...
func (Can) Name() string {
	return "Can"
}

// Texture ...
func (Can) Texture() image.Image {
	return textureFromName("foodcan")
}
