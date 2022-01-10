package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MreSteak struct{}

// AlwaysConsumable ...
func (MreSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MreSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MreSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (MreSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MreSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (MreSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mre_steak", 0
}

// Name ...
func (MreSteak) Name() string {
	return "MreSteak"
}

// Texture ...
func (MreSteak) Texture() image.Image {
	return textureFromName("mre2")
}
