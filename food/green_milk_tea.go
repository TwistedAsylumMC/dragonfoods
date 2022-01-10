package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenMilkTea struct{}

// AlwaysConsumable ...
func (GreenMilkTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenMilkTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenMilkTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenMilkTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenMilkTea) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenMilkTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_milk_tea", 0
}

// Name ...
func (GreenMilkTea) Name() string {
	return "GreenMilkTea"
}

// Texture ...
func (GreenMilkTea) Texture() image.Image {
	return textureFromName("milktea2")
}
