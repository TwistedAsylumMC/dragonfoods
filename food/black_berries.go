package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlackBerries struct{}

// AlwaysConsumable ...
func (BlackBerries) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlackBerries) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlackBerries) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlackBerries) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlackBerries) Edible() bool {
	return true
}

// EncodeItem ...
func (BlackBerries) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:black_berries", 0
}

// Name ...
func (BlackBerries) Name() string {
	return "Black Berries"
}

// Texture ...
func (BlackBerries) Texture() image.Image {
	return textureFromName("blackberries")
}
