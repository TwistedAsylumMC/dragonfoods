package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChickenSandwich struct{}

// AlwaysConsumable ...
func (ChickenSandwich) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChickenSandwich) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChickenSandwich) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChickenSandwich) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChickenSandwich) Edible() bool {
	return true
}

// EncodeItem ...
func (ChickenSandwich) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chicken_sandwich", 0
}

// Name ...
func (ChickenSandwich) Name() string {
	return "Chicken Sandwich"
}

// Texture ...
func (ChickenSandwich) Texture() image.Image {
	return textureFromName("burger33")
}
