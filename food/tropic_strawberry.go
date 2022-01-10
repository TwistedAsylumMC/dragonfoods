package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type TropicStrawberry struct{}

// AlwaysConsumable ...
func (TropicStrawberry) AlwaysConsumable() bool {
	return false
}

// Category ...
func (TropicStrawberry) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (TropicStrawberry) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (TropicStrawberry) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (TropicStrawberry) Edible() bool {
	return true
}

// EncodeItem ...
func (TropicStrawberry) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tropic_strawberry", 0
}

// Name ...
func (TropicStrawberry) Name() string {
	return "Tropic Strawberry"
}

// Texture ...
func (TropicStrawberry) Texture() image.Image {
	return textureFromName("s_k_i_n")
}
