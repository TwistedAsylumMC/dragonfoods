package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Tomato struct{}

// AlwaysConsumable ...
func (Tomato) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Tomato) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Tomato) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Tomato) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Tomato) Edible() bool {
	return true
}

// EncodeItem ...
func (Tomato) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tomato", 0
}

// Name ...
func (Tomato) Name() string {
	return "Tomato"
}

// Texture ...
func (Tomato) Texture() image.Image {
	return textureFromName("tomato")
}
