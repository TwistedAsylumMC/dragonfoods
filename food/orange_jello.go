package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type OrangeJello struct{}

// AlwaysConsumable ...
func (OrangeJello) AlwaysConsumable() bool {
	return false
}

// Category ...
func (OrangeJello) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (OrangeJello) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (OrangeJello) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (OrangeJello) Edible() bool {
	return true
}

// EncodeItem ...
func (OrangeJello) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:orange_jello", 0
}

// Name ...
func (OrangeJello) Name() string {
	return "OrangeJello"
}

// Texture ...
func (OrangeJello) Texture() image.Image {
	return textureFromName("jello67")
}
