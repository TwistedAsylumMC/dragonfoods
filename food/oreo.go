package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Oreo struct{}

// AlwaysConsumable ...
func (Oreo) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Oreo) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Oreo) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Oreo) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Oreo) Edible() bool {
	return true
}

// EncodeItem ...
func (Oreo) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:oreo", 0
}

// Name ...
func (Oreo) Name() string {
	return "Oreo"
}

// Texture ...
func (Oreo) Texture() image.Image {
	return textureFromName("oreo_cookie")
}
