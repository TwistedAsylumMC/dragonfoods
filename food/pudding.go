package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pudding struct{}

// AlwaysConsumable ...
func (Pudding) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pudding) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pudding) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pudding) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pudding) Edible() bool {
	return true
}

// EncodeItem ...
func (Pudding) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pudding", 0
}

// Name ...
func (Pudding) Name() string {
	return "Pudding"
}

// Texture ...
func (Pudding) Texture() image.Image {
	return textureFromName("christmas_pudding_the_betweenlands")
}
