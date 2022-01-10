package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Turnip struct{}

// AlwaysConsumable ...
func (Turnip) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Turnip) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Turnip) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Turnip) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Turnip) Edible() bool {
	return true
}

// EncodeItem ...
func (Turnip) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:turnip", 0
}

// Name ...
func (Turnip) Name() string {
	return "Turnip"
}

// Texture ...
func (Turnip) Texture() image.Image {
	return textureFromName("turnip")
}
