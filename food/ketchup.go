package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Ketchup struct{}

// AlwaysConsumable ...
func (Ketchup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Ketchup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Ketchup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Ketchup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Ketchup) Edible() bool {
	return true
}

// EncodeItem ...
func (Ketchup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ketchup", 0
}

// Name ...
func (Ketchup) Name() string {
	return "Ketchup"
}

// Texture ...
func (Ketchup) Texture() image.Image {
	return textureFromName("ketchup")
}
