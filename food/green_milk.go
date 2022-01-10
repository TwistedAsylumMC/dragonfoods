package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenMilk struct{}

// AlwaysConsumable ...
func (GreenMilk) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenMilk) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenMilk) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenMilk) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenMilk) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenMilk) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_milk", 0
}

// Name ...
func (GreenMilk) Name() string {
	return "GreenMilk"
}

// Texture ...
func (GreenMilk) Texture() image.Image {
	return textureFromName("kawiimilk")
}
