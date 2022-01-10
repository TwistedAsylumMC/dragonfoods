package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Eclair struct{}

// AlwaysConsumable ...
func (Eclair) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Eclair) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Eclair) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(3, 1.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Eclair) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Eclair) Edible() bool {
	return true
}

// EncodeItem ...
func (Eclair) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:eclair", 0
}

// Name ...
func (Eclair) Name() string {
	return "Eclair"
}

// Texture ...
func (Eclair) Texture() image.Image {
	return textureFromName("eclair")
}
