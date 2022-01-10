package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GarlicBread struct{}

// AlwaysConsumable ...
func (GarlicBread) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GarlicBread) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GarlicBread) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GarlicBread) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GarlicBread) Edible() bool {
	return true
}

// EncodeItem ...
func (GarlicBread) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:garlic_bread", 0
}

// Name ...
func (GarlicBread) Name() string {
	return "Garlic Bread"
}

// Texture ...
func (GarlicBread) Texture() image.Image {
	return textureFromName("48_garlicbread")
}
