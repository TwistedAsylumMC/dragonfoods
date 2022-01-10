package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Blackjello struct{}

// AlwaysConsumable ...
func (Blackjello) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Blackjello) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Blackjello) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Blackjello) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Blackjello) Edible() bool {
	return true
}

// EncodeItem ...
func (Blackjello) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blackjello", 0
}

// Name ...
func (Blackjello) Name() string {
	return "Blackjello"
}

// Texture ...
func (Blackjello) Texture() image.Image {
	return textureFromName("66jello")
}
