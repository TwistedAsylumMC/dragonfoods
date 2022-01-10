package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MangoLassi struct{}

// AlwaysConsumable ...
func (MangoLassi) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MangoLassi) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MangoLassi) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (MangoLassi) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MangoLassi) Edible() bool {
	return true
}

// EncodeItem ...
func (MangoLassi) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mango_lassi", 0
}

// Name ...
func (MangoLassi) Name() string {
	return "MangoLassi"
}

// Texture ...
func (MangoLassi) Texture() image.Image {
	return textureFromName("mangolassi")
}
