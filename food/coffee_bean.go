package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CoffeeBean struct{}

// AlwaysConsumable ...
func (CoffeeBean) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CoffeeBean) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CoffeeBean) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CoffeeBean) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CoffeeBean) Edible() bool {
	return true
}

// EncodeItem ...
func (CoffeeBean) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:coffee_bean", 0
}

// Name ...
func (CoffeeBean) Name() string {
	return "CoffeeBean"
}

// Texture ...
func (CoffeeBean) Texture() image.Image {
	return textureFromName("coffeebean")
}
