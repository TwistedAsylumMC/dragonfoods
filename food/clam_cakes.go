package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ClamCakes struct{}

// AlwaysConsumable ...
func (ClamCakes) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ClamCakes) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ClamCakes) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ClamCakes) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ClamCakes) Edible() bool {
	return true
}

// EncodeItem ...
func (ClamCakes) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:clam_cakes", 0
}

// Name ...
func (ClamCakes) Name() string {
	return "ClamCakes"
}

// Texture ...
func (ClamCakes) Texture() image.Image {
	return textureFromName("clamcakes")
}
