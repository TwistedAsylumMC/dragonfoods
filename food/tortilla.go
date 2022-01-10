package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Tortilla struct{}

// AlwaysConsumable ...
func (Tortilla) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Tortilla) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Tortilla) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Tortilla) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Tortilla) Edible() bool {
	return true
}

// EncodeItem ...
func (Tortilla) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tortilla", 0
}

// Name ...
func (Tortilla) Name() string {
	return "Tortilla"
}

// Texture ...
func (Tortilla) Texture() image.Image {
	return textureFromName("tortilla")
}
