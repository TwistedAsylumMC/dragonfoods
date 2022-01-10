package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedBlueBerryPie struct{}

// AlwaysConsumable ...
func (BakedBlueBerryPie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedBlueBerryPie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedBlueBerryPie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedBlueBerryPie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedBlueBerryPie) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedBlueBerryPie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_blue_berry_pie", 0
}

// Name ...
func (BakedBlueBerryPie) Name() string {
	return "BakedBlueBerryPie"
}

// Texture ...
func (BakedBlueBerryPie) Texture() image.Image {
	return textureFromName("bakedblueberrypie")
}
