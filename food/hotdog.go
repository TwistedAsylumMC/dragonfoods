package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Hotdog struct{}

// AlwaysConsumable ...
func (Hotdog) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Hotdog) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Hotdog) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Hotdog) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Hotdog) Edible() bool {
	return true
}

// EncodeItem ...
func (Hotdog) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:hotdog", 0
}

// Name ...
func (Hotdog) Name() string {
	return "Hotdog"
}

// Texture ...
func (Hotdog) Texture() image.Image {
	return textureFromName("54_hotdog")
}
