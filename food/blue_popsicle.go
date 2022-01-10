package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BluePopsicle struct{}

// AlwaysConsumable ...
func (BluePopsicle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BluePopsicle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BluePopsicle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BluePopsicle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BluePopsicle) Edible() bool {
	return true
}

// EncodeItem ...
func (BluePopsicle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_popsicle", 0
}

// Name ...
func (BluePopsicle) Name() string {
	return "BluePopsicle"
}

// Texture ...
func (BluePopsicle) Texture() image.Image {
	return textureFromName("popsicle2")
}
