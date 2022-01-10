package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Coffeebeans struct{}

// AlwaysConsumable ...
func (Coffeebeans) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Coffeebeans) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Coffeebeans) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Coffeebeans) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Coffeebeans) Edible() bool {
	return true
}

// EncodeItem ...
func (Coffeebeans) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:coffeebeans", 0
}

// Name ...
func (Coffeebeans) Name() string {
	return "Coffeebeans"
}

// Texture ...
func (Coffeebeans) Texture() image.Image {
	return textureFromName("coffee_bag")
}
