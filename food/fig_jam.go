package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FigJam struct{}

// AlwaysConsumable ...
func (FigJam) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FigJam) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FigJam) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (FigJam) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FigJam) Edible() bool {
	return true
}

// EncodeItem ...
func (FigJam) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fig_jam", 0
}

// Name ...
func (FigJam) Name() string {
	return "Fig Jam"
}

// Texture ...
func (FigJam) Texture() image.Image {
	return textureFromName("figjam")
}
