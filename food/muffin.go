package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Muffin struct{}

// AlwaysConsumable ...
func (Muffin) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Muffin) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Muffin) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Muffin) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Muffin) Edible() bool {
	return true
}

// EncodeItem ...
func (Muffin) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:muffin", 0
}

// Name ...
func (Muffin) Name() string {
	return "Muffin"
}

// Texture ...
func (Muffin) Texture() image.Image {
	return textureFromName("muffin")
}
