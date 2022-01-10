package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pistacchio struct{}

// AlwaysConsumable ...
func (Pistacchio) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pistacchio) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pistacchio) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pistacchio) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pistacchio) Edible() bool {
	return true
}

// EncodeItem ...
func (Pistacchio) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pistacchio", 0
}

// Name ...
func (Pistacchio) Name() string {
	return "Pistacchio"
}

// Texture ...
func (Pistacchio) Texture() image.Image {
	return textureFromName("pistacchio")
}
