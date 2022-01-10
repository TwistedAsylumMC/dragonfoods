package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PhillyCheeseSteak struct{}

// AlwaysConsumable ...
func (PhillyCheeseSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PhillyCheeseSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PhillyCheeseSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (PhillyCheeseSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PhillyCheeseSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (PhillyCheeseSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:philly_cheese_steak", 0
}

// Name ...
func (PhillyCheeseSteak) Name() string {
	return "Philly Cheese Steak"
}

// Texture ...
func (PhillyCheeseSteak) Texture() image.Image {
	return textureFromName("phillycheese")
}
