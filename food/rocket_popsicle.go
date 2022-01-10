package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RocketPopsicle struct{}

// AlwaysConsumable ...
func (RocketPopsicle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RocketPopsicle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RocketPopsicle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (RocketPopsicle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RocketPopsicle) Edible() bool {
	return true
}

// EncodeItem ...
func (RocketPopsicle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:rocket_popsicle", 0
}

// Name ...
func (RocketPopsicle) Name() string {
	return "RocketPopsicle"
}

// Texture ...
func (RocketPopsicle) Texture() image.Image {
	return textureFromName("popsicle_rocket_pop")
}
