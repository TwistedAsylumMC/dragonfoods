package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Beans struct{}

// AlwaysConsumable ...
func (Beans) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Beans) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Beans) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Beans) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Beans) Edible() bool {
	return true
}

// EncodeItem ...
func (Beans) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:beans", 0
}

// Name ...
func (Beans) Name() string {
	return "Beans"
}

// Texture ...
func (Beans) Texture() image.Image {
	return textureFromName("beans")
}
