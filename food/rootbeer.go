package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Rootbeer struct{}

// AlwaysConsumable ...
func (Rootbeer) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Rootbeer) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Rootbeer) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Rootbeer) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Rootbeer) Edible() bool {
	return true
}

// EncodeItem ...
func (Rootbeer) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:rootbeer", 0
}

// Name ...
func (Rootbeer) Name() string {
	return "Rootbeer"
}

// Texture ...
func (Rootbeer) Texture() image.Image {
	return textureFromName("soda_bottle")
}
