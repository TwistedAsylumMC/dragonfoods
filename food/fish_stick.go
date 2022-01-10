package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FishStick struct{}

// AlwaysConsumable ...
func (FishStick) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FishStick) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FishStick) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (FishStick) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FishStick) Edible() bool {
	return true
}

// EncodeItem ...
func (FishStick) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fish_stick", 0
}

// Name ...
func (FishStick) Name() string {
	return "Fish Stick"
}

// Texture ...
func (FishStick) Texture() image.Image {
	return textureFromName("fishstick")
}
