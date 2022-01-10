package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Frittata struct{}

// AlwaysConsumable ...
func (Frittata) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Frittata) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Frittata) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(15, 9.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Frittata) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Frittata) Edible() bool {
	return true
}

// EncodeItem ...
func (Frittata) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:frittata", 0
}

// Name ...
func (Frittata) Name() string {
	return "Frittata"
}

// Texture ...
func (Frittata) Texture() image.Image {
	return textureFromName("frittata")
}
