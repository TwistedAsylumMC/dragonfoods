package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueBerryFrezze struct{}

// AlwaysConsumable ...
func (BlueBerryFrezze) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueBerryFrezze) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueBerryFrezze) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueBerryFrezze) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueBerryFrezze) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueBerryFrezze) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_berry_frezze", 0
}

// Name ...
func (BlueBerryFrezze) Name() string {
	return "BlueBerryFrezze"
}

// Texture ...
func (BlueBerryFrezze) Texture() image.Image {
	return textureFromName("blueberryfrezze")
}
