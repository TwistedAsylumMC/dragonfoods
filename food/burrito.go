package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Burrito struct{}

// AlwaysConsumable ...
func (Burrito) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Burrito) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Burrito) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Burrito) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Burrito) Edible() bool {
	return true
}

// EncodeItem ...
func (Burrito) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:burrito", 0
}

// Name ...
func (Burrito) Name() string {
	return "Burrito"
}

// Texture ...
func (Burrito) Texture() image.Image {
	return textureFromName("burrio")
}
