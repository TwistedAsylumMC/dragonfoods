package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GlazedDonut struct{}

// AlwaysConsumable ...
func (GlazedDonut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GlazedDonut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GlazedDonut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GlazedDonut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GlazedDonut) Edible() bool {
	return true
}

// EncodeItem ...
func (GlazedDonut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:glazed_donut", 0
}

// Name ...
func (GlazedDonut) Name() string {
	return "Glazed Donut"
}

// Texture ...
func (GlazedDonut) Texture() image.Image {
	return textureFromName("glazed_donut")
}
