package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PinkMilk struct{}

// AlwaysConsumable ...
func (PinkMilk) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PinkMilk) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PinkMilk) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (PinkMilk) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PinkMilk) Edible() bool {
	return true
}

// EncodeItem ...
func (PinkMilk) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pink_milk", 0
}

// Name ...
func (PinkMilk) Name() string {
	return "PinkMilk"
}

// Texture ...
func (PinkMilk) Texture() image.Image {
	return textureFromName("cute_strawberry_milk")
}
