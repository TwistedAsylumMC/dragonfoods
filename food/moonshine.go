package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Moonshine struct{}

// AlwaysConsumable ...
func (Moonshine) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Moonshine) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Moonshine) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Moonshine) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Moonshine) Edible() bool {
	return true
}

// EncodeItem ...
func (Moonshine) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:moonshine", 0
}

// Name ...
func (Moonshine) Name() string {
	return "Moonshine"
}

// Texture ...
func (Moonshine) Texture() image.Image {
	return textureFromName("moonshine")
}
