package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type EspressoCoffee struct{}

// AlwaysConsumable ...
func (EspressoCoffee) AlwaysConsumable() bool {
	return false
}

// Category ...
func (EspressoCoffee) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (EspressoCoffee) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (EspressoCoffee) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (EspressoCoffee) Edible() bool {
	return true
}

// EncodeItem ...
func (EspressoCoffee) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:espresso_coffee", 0
}

// Name ...
func (EspressoCoffee) Name() string {
	return "Espresso Coffee"
}

// Texture ...
func (EspressoCoffee) Texture() image.Image {
	return textureFromName("espressocoffe")
}
