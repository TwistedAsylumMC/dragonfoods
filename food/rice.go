package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Rice struct{}

// AlwaysConsumable ...
func (Rice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Rice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Rice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Rice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Rice) Edible() bool {
	return true
}

// EncodeItem ...
func (Rice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:rice", 0
}

// Name ...
func (Rice) Name() string {
	return "Rice"
}

// Texture ...
func (Rice) Texture() image.Image {
	return textureFromName("rice")
}
