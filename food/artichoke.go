package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Artichoke struct{}

// AlwaysConsumable ...
func (Artichoke) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Artichoke) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Artichoke) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Artichoke) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Artichoke) Edible() bool {
	return true
}

// EncodeItem ...
func (Artichoke) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:artichoke", 0
}

// Name ...
func (Artichoke) Name() string {
	return "Artichoke"
}

// Texture ...
func (Artichoke) Texture() image.Image {
	return textureFromName("artichoke")
}
