package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedGlowBerryPie struct{}

// AlwaysConsumable ...
func (BakedGlowBerryPie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedGlowBerryPie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedGlowBerryPie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedGlowBerryPie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedGlowBerryPie) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedGlowBerryPie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_glow_berry_pie", 0
}

// Name ...
func (BakedGlowBerryPie) Name() string {
	return "BakedGlowBerryPie"
}

// Texture ...
func (BakedGlowBerryPie) Texture() image.Image {
	return textureFromName("bakedglowberrypie")
}
