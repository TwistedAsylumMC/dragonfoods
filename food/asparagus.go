package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Asparagus struct{}

// AlwaysConsumable ...
func (Asparagus) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Asparagus) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Asparagus) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Asparagus) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Asparagus) Edible() bool {
	return true
}

// EncodeItem ...
func (Asparagus) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:asparagus", 0
}

// Name ...
func (Asparagus) Name() string {
	return "Asparagus"
}

// Texture ...
func (Asparagus) Texture() image.Image {
	return textureFromName("asparagus_02")
}
