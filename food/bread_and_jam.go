package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BreadAndJam struct{}

// AlwaysConsumable ...
func (BreadAndJam) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BreadAndJam) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BreadAndJam) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BreadAndJam) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BreadAndJam) Edible() bool {
	return true
}

// EncodeItem ...
func (BreadAndJam) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bread_and_jam", 0
}

// Name ...
func (BreadAndJam) Name() string {
	return "Bread And Jam"
}

// Texture ...
func (BreadAndJam) Texture() image.Image {
	return textureFromName("62_jam_dish")
}
