package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PinkPopsicle struct{}

// AlwaysConsumable ...
func (PinkPopsicle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PinkPopsicle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PinkPopsicle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PinkPopsicle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PinkPopsicle) Edible() bool {
	return true
}

// EncodeItem ...
func (PinkPopsicle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pink_popsicle", 0
}

// Name ...
func (PinkPopsicle) Name() string {
	return "Pink Popsicle"
}

// Texture ...
func (PinkPopsicle) Texture() image.Image {
	return textureFromName("popsicle5")
}
