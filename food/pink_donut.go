package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PinkDonut struct{}

// AlwaysConsumable ...
func (PinkDonut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PinkDonut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PinkDonut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PinkDonut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PinkDonut) Edible() bool {
	return true
}

// EncodeItem ...
func (PinkDonut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pink_donut", 0
}

// Name ...
func (PinkDonut) Name() string {
	return "PinkDonut"
}

// Texture ...
func (PinkDonut) Texture() image.Image {
	return textureFromName("32donut")
}
