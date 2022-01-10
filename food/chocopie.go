package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Chocopie struct{}

// AlwaysConsumable ...
func (Chocopie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Chocopie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Chocopie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Chocopie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Chocopie) Edible() bool {
	return true
}

// EncodeItem ...
func (Chocopie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chocopie", 0
}

// Name ...
func (Chocopie) Name() string {
	return "Chocopie"
}

// Texture ...
func (Chocopie) Texture() image.Image {
	return textureFromName("chocopie")
}
