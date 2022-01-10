package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PinkJello struct{}

// AlwaysConsumable ...
func (PinkJello) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PinkJello) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PinkJello) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (PinkJello) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PinkJello) Edible() bool {
	return true
}

// EncodeItem ...
func (PinkJello) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pink_jello", 0
}

// Name ...
func (PinkJello) Name() string {
	return "PinkJello"
}

// Texture ...
func (PinkJello) Texture() image.Image {
	return textureFromName("64jello")
}
