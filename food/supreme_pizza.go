package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SupremePizza struct{}

// AlwaysConsumable ...
func (SupremePizza) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SupremePizza) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SupremePizza) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SupremePizza) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SupremePizza) Edible() bool {
	return true
}

// EncodeItem ...
func (SupremePizza) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:supreme_pizza", 0
}

// Name ...
func (SupremePizza) Name() string {
	return "SupremePizza"
}

// Texture ...
func (SupremePizza) Texture() image.Image {
	return textureFromName("pizza2")
}
