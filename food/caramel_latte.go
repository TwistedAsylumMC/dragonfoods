package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CaramelLatte struct{}

// AlwaysConsumable ...
func (CaramelLatte) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CaramelLatte) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CaramelLatte) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CaramelLatte) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CaramelLatte) Edible() bool {
	return true
}

// EncodeItem ...
func (CaramelLatte) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:caramel_latte", 0
}

// Name ...
func (CaramelLatte) Name() string {
	return "CaramelLatte"
}

// Texture ...
func (CaramelLatte) Texture() image.Image {
	return textureFromName("caramellatte")
}
