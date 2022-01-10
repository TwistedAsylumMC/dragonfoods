package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Carbonara struct{}

// AlwaysConsumable ...
func (Carbonara) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Carbonara) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Carbonara) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Carbonara) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Carbonara) Edible() bool {
	return true
}

// EncodeItem ...
func (Carbonara) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:carbonara", 0
}

// Name ...
func (Carbonara) Name() string {
	return "Carbonara"
}

// Texture ...
func (Carbonara) Texture() image.Image {
	return textureFromName("carbonara")
}
