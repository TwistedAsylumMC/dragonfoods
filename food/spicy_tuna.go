package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SpicyTuna struct{}

// AlwaysConsumable ...
func (SpicyTuna) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SpicyTuna) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SpicyTuna) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SpicyTuna) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SpicyTuna) Edible() bool {
	return true
}

// EncodeItem ...
func (SpicyTuna) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spicy_tuna", 0
}

// Name ...
func (SpicyTuna) Name() string {
	return "Spicy Tuna"
}

// Texture ...
func (SpicyTuna) Texture() image.Image {
	return textureFromName("tuna_can2")
}
