package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Brownie struct{}

// AlwaysConsumable ...
func (Brownie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Brownie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Brownie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Brownie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Brownie) Edible() bool {
	return true
}

// EncodeItem ...
func (Brownie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:brownie", 0
}

// Name ...
func (Brownie) Name() string {
	return "Brownie"
}

// Texture ...
func (Brownie) Texture() image.Image {
	return textureFromName("brownie")
}
