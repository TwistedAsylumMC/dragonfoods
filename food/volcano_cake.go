package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type VolcanoCake struct{}

// AlwaysConsumable ...
func (VolcanoCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (VolcanoCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (VolcanoCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (VolcanoCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (VolcanoCake) Edible() bool {
	return true
}

// EncodeItem ...
func (VolcanoCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:volcano_cake", 0
}

// Name ...
func (VolcanoCake) Name() string {
	return "VolcanoCake"
}

// Texture ...
func (VolcanoCake) Texture() image.Image {
	return textureFromName("donut_1")
}
