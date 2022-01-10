package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Fennel struct{}

// AlwaysConsumable ...
func (Fennel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Fennel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Fennel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Fennel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Fennel) Edible() bool {
	return true
}

// EncodeItem ...
func (Fennel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fennel", 0
}

// Name ...
func (Fennel) Name() string {
	return "Fennel"
}

// Texture ...
func (Fennel) Texture() image.Image {
	return textureFromName("fennel_bulb")
}
