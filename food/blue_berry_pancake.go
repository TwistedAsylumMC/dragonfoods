package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueBerryPancake struct{}

// AlwaysConsumable ...
func (BlueBerryPancake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueBerryPancake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueBerryPancake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueBerryPancake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueBerryPancake) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueBerryPancake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_berry_pancake", 0
}

// Name ...
func (BlueBerryPancake) Name() string {
	return "Blue Berry Pancake"
}

// Texture ...
func (BlueBerryPancake) Texture() image.Image {
	return textureFromName("pancake2")
}
