package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedRoll struct{}

// AlwaysConsumable ...
func (BakedRoll) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedRoll) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedRoll) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedRoll) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedRoll) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedRoll) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_roll", 0
}

// Name ...
func (BakedRoll) Name() string {
	return "Baked Roll"
}

// Texture ...
func (BakedRoll) Texture() image.Image {
	return textureFromName("roll2")
}
