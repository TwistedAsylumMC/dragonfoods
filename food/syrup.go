package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Syrup struct{}

// AlwaysConsumable ...
func (Syrup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Syrup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Syrup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Syrup) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Syrup) Edible() bool {
	return true
}

// EncodeItem ...
func (Syrup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:syrup", 0
}

// Name ...
func (Syrup) Name() string {
	return "Syrup"
}

// Texture ...
func (Syrup) Texture() image.Image {
	return textureFromName("surope")
}
