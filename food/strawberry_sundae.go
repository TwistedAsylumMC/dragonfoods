package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type StrawberrySundae struct{}

// AlwaysConsumable ...
func (StrawberrySundae) AlwaysConsumable() bool {
	return false
}

// Category ...
func (StrawberrySundae) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (StrawberrySundae) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (StrawberrySundae) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (StrawberrySundae) Edible() bool {
	return true
}

// EncodeItem ...
func (StrawberrySundae) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:strawberry_sundae", 0
}

// Name ...
func (StrawberrySundae) Name() string {
	return "Strawberry Sundae"
}

// Texture ...
func (StrawberrySundae) Texture() image.Image {
	return textureFromName("ice_cream_sundae_01")
}
