package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Kabob struct{}

// AlwaysConsumable ...
func (Kabob) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Kabob) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Kabob) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Kabob) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Kabob) Edible() bool {
	return true
}

// EncodeItem ...
func (Kabob) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:kabob", 0
}

// Name ...
func (Kabob) Name() string {
	return "Kabob"
}

// Texture ...
func (Kabob) Texture() image.Image {
	return textureFromName("shish_kabob")
}
