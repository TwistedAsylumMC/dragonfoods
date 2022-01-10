package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type StrawberryJam struct{}

// AlwaysConsumable ...
func (StrawberryJam) AlwaysConsumable() bool {
	return false
}

// Category ...
func (StrawberryJam) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (StrawberryJam) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (StrawberryJam) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (StrawberryJam) Edible() bool {
	return true
}

// EncodeItem ...
func (StrawberryJam) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:strawberry_jam", 0
}

// Name ...
func (StrawberryJam) Name() string {
	return "Strawberry Jam"
}

// Texture ...
func (StrawberryJam) Texture() image.Image {
	return textureFromName("strawberry_jam")
}
