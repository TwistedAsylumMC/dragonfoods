package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GlowBerryPie struct{}

// AlwaysConsumable ...
func (GlowBerryPie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GlowBerryPie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GlowBerryPie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(3, 1.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (GlowBerryPie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GlowBerryPie) Edible() bool {
	return true
}

// EncodeItem ...
func (GlowBerryPie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:glow_berry_pie", 0
}

// Name ...
func (GlowBerryPie) Name() string {
	return "Glow Berry Pie"
}

// Texture ...
func (GlowBerryPie) Texture() image.Image {
	return textureFromName("glowberrypie")
}
