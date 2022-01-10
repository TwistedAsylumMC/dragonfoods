package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Saki struct{}

// AlwaysConsumable ...
func (Saki) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Saki) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Saki) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Saki) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Saki) Edible() bool {
	return true
}

// EncodeItem ...
func (Saki) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:saki", 0
}

// Name ...
func (Saki) Name() string {
	return "Saki"
}

// Texture ...
func (Saki) Texture() image.Image {
	return textureFromName("saki")
}
