package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedPinkCake struct{}

// AlwaysConsumable ...
func (BakedPinkCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedPinkCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedPinkCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedPinkCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedPinkCake) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedPinkCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_pink_cake", 0
}

// Name ...
func (BakedPinkCake) Name() string {
	return "BakedPinkCake"
}

// Texture ...
func (BakedPinkCake) Texture() image.Image {
	return textureFromName("pink_pink")
}
