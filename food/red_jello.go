package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedJello struct{}

// AlwaysConsumable ...
func (RedJello) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedJello) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedJello) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedJello) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedJello) Edible() bool {
	return true
}

// EncodeItem ...
func (RedJello) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_jello", 0
}

// Name ...
func (RedJello) Name() string {
	return "Red Jello"
}

// Texture ...
func (RedJello) Texture() image.Image {
	return textureFromName("62jello")
}
