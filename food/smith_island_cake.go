package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SmithIslandCake struct{}

// AlwaysConsumable ...
func (SmithIslandCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SmithIslandCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SmithIslandCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (SmithIslandCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SmithIslandCake) Edible() bool {
	return true
}

// EncodeItem ...
func (SmithIslandCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:smith_island_cake", 0
}

// Name ...
func (SmithIslandCake) Name() string {
	return "Smith Island Cake"
}

// Texture ...
func (SmithIslandCake) Texture() image.Image {
	return textureFromName("pie5")
}
