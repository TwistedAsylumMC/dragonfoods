package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MeatTaco struct{}

// AlwaysConsumable ...
func (MeatTaco) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MeatTaco) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MeatTaco) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (MeatTaco) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MeatTaco) Edible() bool {
	return true
}

// EncodeItem ...
func (MeatTaco) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:meat_taco", 0
}

// Name ...
func (MeatTaco) Name() string {
	return "MeatTaco"
}

// Texture ...
func (MeatTaco) Texture() image.Image {
	return textureFromName("taco4")
}
