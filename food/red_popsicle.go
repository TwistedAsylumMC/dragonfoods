package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedPopsicle struct{}

// AlwaysConsumable ...
func (RedPopsicle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedPopsicle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedPopsicle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedPopsicle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedPopsicle) Edible() bool {
	return true
}

// EncodeItem ...
func (RedPopsicle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_popsicle", 0
}

// Name ...
func (RedPopsicle) Name() string {
	return "Red Popsicle"
}

// Texture ...
func (RedPopsicle) Texture() image.Image {
	return textureFromName("popsicle")
}
