package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Curry struct{}

// AlwaysConsumable ...
func (Curry) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Curry) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Curry) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(9, 5.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Curry) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Curry) Edible() bool {
	return true
}

// EncodeItem ...
func (Curry) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:curry", 0
}

// Name ...
func (Curry) Name() string {
	return "Curry"
}

// Texture ...
func (Curry) Texture() image.Image {
	return textureFromName("curry")
}
