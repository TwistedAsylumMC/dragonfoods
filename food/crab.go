package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Crab struct{}

// AlwaysConsumable ...
func (Crab) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Crab) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Crab) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Crab) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Crab) Edible() bool {
	return true
}

// EncodeItem ...
func (Crab) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:crab", 0
}

// Name ...
func (Crab) Name() string {
	return "Crab"
}

// Texture ...
func (Crab) Texture() image.Image {
	return textureFromName("crab")
}
