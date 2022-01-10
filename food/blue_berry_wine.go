package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueBerryWine struct{}

// AlwaysConsumable ...
func (BlueBerryWine) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueBerryWine) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueBerryWine) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueBerryWine) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueBerryWine) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueBerryWine) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_berry_wine", 0
}

// Name ...
func (BlueBerryWine) Name() string {
	return "BlueBerryWine"
}

// Texture ...
func (BlueBerryWine) Texture() image.Image {
	return textureFromName("wine2")
}
