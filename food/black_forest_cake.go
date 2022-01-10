package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlackForestCake struct{}

// AlwaysConsumable ...
func (BlackForestCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlackForestCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlackForestCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlackForestCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlackForestCake) Edible() bool {
	return true
}

// EncodeItem ...
func (BlackForestCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:black_forest_cake", 0
}

// Name ...
func (BlackForestCake) Name() string {
	return "BlackForestCake"
}

// Texture ...
func (BlackForestCake) Texture() image.Image {
	return textureFromName("91_strawberrycake_dish")
}
