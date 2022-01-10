package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Twinkie struct{}

// AlwaysConsumable ...
func (Twinkie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Twinkie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Twinkie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Twinkie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Twinkie) Edible() bool {
	return true
}

// EncodeItem ...
func (Twinkie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:twinkie", 0
}

// Name ...
func (Twinkie) Name() string {
	return "Twinkie"
}

// Texture ...
func (Twinkie) Texture() image.Image {
	return textureFromName("twinkie2")
}
