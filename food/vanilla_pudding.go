package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type VanillaPudding struct{}

// AlwaysConsumable ...
func (VanillaPudding) AlwaysConsumable() bool {
	return false
}

// Category ...
func (VanillaPudding) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (VanillaPudding) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (VanillaPudding) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (VanillaPudding) Edible() bool {
	return true
}

// EncodeItem ...
func (VanillaPudding) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:vanilla_pudding", 0
}

// Name ...
func (VanillaPudding) Name() string {
	return "Vanilla Pudding"
}

// Texture ...
func (VanillaPudding) Texture() image.Image {
	return textureFromName("pudding2")
}
