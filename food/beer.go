package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Beer struct{}

// AlwaysConsumable ...
func (Beer) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Beer) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Beer) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Beer) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Beer) Edible() bool {
	return true
}

// EncodeItem ...
func (Beer) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:beer", 0
}

// Name ...
func (Beer) Name() string {
	return "Beer"
}

// Texture ...
func (Beer) Texture() image.Image {
	return textureFromName("beer")
}
