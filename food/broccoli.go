package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Broccoli struct{}

// AlwaysConsumable ...
func (Broccoli) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Broccoli) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Broccoli) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Broccoli) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Broccoli) Edible() bool {
	return true
}

// EncodeItem ...
func (Broccoli) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:broccoli", 0
}

// Name ...
func (Broccoli) Name() string {
	return "Broccoli"
}

// Texture ...
func (Broccoli) Texture() image.Image {
	return textureFromName("broccoli")
}
