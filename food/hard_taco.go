package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HardTaco struct{}

// AlwaysConsumable ...
func (HardTaco) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HardTaco) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HardTaco) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (HardTaco) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HardTaco) Edible() bool {
	return true
}

// EncodeItem ...
func (HardTaco) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:hard_taco", 0
}

// Name ...
func (HardTaco) Name() string {
	return "Hard Taco"
}

// Texture ...
func (HardTaco) Texture() image.Image {
	return textureFromName("hardtaco1")
}
