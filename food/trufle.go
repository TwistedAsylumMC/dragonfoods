package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Trufle struct{}

// AlwaysConsumable ...
func (Trufle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Trufle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Trufle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Trufle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Trufle) Edible() bool {
	return true
}

// EncodeItem ...
func (Trufle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:trufle", 0
}

// Name ...
func (Trufle) Name() string {
	return "Trufle"
}

// Texture ...
func (Trufle) Texture() image.Image {
	return textureFromName("truffles")
}
