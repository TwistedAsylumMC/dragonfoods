package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Whisk struct{}

// AlwaysConsumable ...
func (Whisk) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Whisk) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Whisk) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Whisk) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Whisk) Edible() bool {
	return true
}

// EncodeItem ...
func (Whisk) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:whisk", 0
}

// Name ...
func (Whisk) Name() string {
	return "Whisk"
}

// Texture ...
func (Whisk) Texture() image.Image {
	return textureFromName("whisk")
}
