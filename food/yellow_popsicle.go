package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type YellowPopsicle struct{}

// AlwaysConsumable ...
func (YellowPopsicle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (YellowPopsicle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (YellowPopsicle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (YellowPopsicle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (YellowPopsicle) Edible() bool {
	return true
}

// EncodeItem ...
func (YellowPopsicle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellow_popsicle", 0
}

// Name ...
func (YellowPopsicle) Name() string {
	return "Yellow Popsicle"
}

// Texture ...
func (YellowPopsicle) Texture() image.Image {
	return textureFromName("popsicle4")
}
