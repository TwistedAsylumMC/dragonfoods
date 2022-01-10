package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Spaghetti struct{}

// AlwaysConsumable ...
func (Spaghetti) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Spaghetti) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Spaghetti) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Spaghetti) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Spaghetti) Edible() bool {
	return true
}

// EncodeItem ...
func (Spaghetti) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spaghetti", 0
}

// Name ...
func (Spaghetti) Name() string {
	return "Spaghetti"
}

// Texture ...
func (Spaghetti) Texture() image.Image {
	return textureFromName("spaghetti")
}
