package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GrapeJam struct{}

// AlwaysConsumable ...
func (GrapeJam) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GrapeJam) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GrapeJam) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (GrapeJam) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GrapeJam) Edible() bool {
	return true
}

// EncodeItem ...
func (GrapeJam) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:grape_jam", 0
}

// Name ...
func (GrapeJam) Name() string {
	return "GrapeJam"
}

// Texture ...
func (GrapeJam) Texture() image.Image {
	return textureFromName("72_jam")
}
