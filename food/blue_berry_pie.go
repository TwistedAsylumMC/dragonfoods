package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueBerryPie struct{}

// AlwaysConsumable ...
func (BlueBerryPie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueBerryPie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueBerryPie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(3, 1.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueBerryPie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueBerryPie) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueBerryPie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_berry_pie", 0
}

// Name ...
func (BlueBerryPie) Name() string {
	return "BlueBerryPie"
}

// Texture ...
func (BlueBerryPie) Texture() image.Image {
	return textureFromName("blueberrypie")
}
