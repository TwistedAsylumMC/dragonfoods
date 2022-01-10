package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CheeseyBroccoli struct{}

// AlwaysConsumable ...
func (CheeseyBroccoli) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CheeseyBroccoli) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CheeseyBroccoli) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CheeseyBroccoli) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CheeseyBroccoli) Edible() bool {
	return true
}

// EncodeItem ...
func (CheeseyBroccoli) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cheesey_broccoli", 0
}

// Name ...
func (CheeseyBroccoli) Name() string {
	return "Cheesey Broccoli"
}

// Texture ...
func (CheeseyBroccoli) Texture() image.Image {
	return textureFromName("broccoli2")
}
