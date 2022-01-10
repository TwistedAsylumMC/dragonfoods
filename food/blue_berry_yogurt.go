package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueBerryYogurt struct{}

// AlwaysConsumable ...
func (BlueBerryYogurt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueBerryYogurt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueBerryYogurt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueBerryYogurt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueBerryYogurt) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueBerryYogurt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_berry_yogurt", 0
}

// Name ...
func (BlueBerryYogurt) Name() string {
	return "Blue Berry Yogurt"
}

// Texture ...
func (BlueBerryYogurt) Texture() image.Image {
	return textureFromName("blueberryicecream")
}
