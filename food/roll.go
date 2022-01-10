package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Roll struct{}

// AlwaysConsumable ...
func (Roll) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Roll) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Roll) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Roll) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Roll) Edible() bool {
	return true
}

// EncodeItem ...
func (Roll) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:roll", 0
}

// Name ...
func (Roll) Name() string {
	return "Roll"
}

// Texture ...
func (Roll) Texture() image.Image {
	return textureFromName("roll")
}
