package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type IceWater struct{}

// AlwaysConsumable ...
func (IceWater) AlwaysConsumable() bool {
	return false
}

// Category ...
func (IceWater) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (IceWater) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (IceWater) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (IceWater) Edible() bool {
	return true
}

// EncodeItem ...
func (IceWater) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ice_water", 0
}

// Name ...
func (IceWater) Name() string {
	return "IceWater"
}

// Texture ...
func (IceWater) Texture() image.Image {
	return textureFromName("icewater")
}
