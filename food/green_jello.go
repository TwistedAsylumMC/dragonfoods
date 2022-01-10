package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenJello struct{}

// AlwaysConsumable ...
func (GreenJello) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenJello) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenJello) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenJello) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenJello) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenJello) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_jello", 0
}

// Name ...
func (GreenJello) Name() string {
	return "Green Jello"
}

// Texture ...
func (GreenJello) Texture() image.Image {
	return textureFromName("61jello")
}
