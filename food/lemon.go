package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Lemon struct{}

// AlwaysConsumable ...
func (Lemon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Lemon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Lemon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Lemon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Lemon) Edible() bool {
	return true
}

// EncodeItem ...
func (Lemon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:lemon", 0
}

// Name ...
func (Lemon) Name() string {
	return "Lemon"
}

// Texture ...
func (Lemon) Texture() image.Image {
	return textureFromName("lemon")
}
