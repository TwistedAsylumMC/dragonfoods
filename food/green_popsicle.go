package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenPopsicle struct{}

// AlwaysConsumable ...
func (GreenPopsicle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenPopsicle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenPopsicle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenPopsicle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenPopsicle) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenPopsicle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_popsicle", 0
}

// Name ...
func (GreenPopsicle) Name() string {
	return "GreenPopsicle"
}

// Texture ...
func (GreenPopsicle) Texture() image.Image {
	return textureFromName("popsicle3")
}
