package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Vindaloo struct{}

// AlwaysConsumable ...
func (Vindaloo) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Vindaloo) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Vindaloo) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Vindaloo) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Vindaloo) Edible() bool {
	return true
}

// EncodeItem ...
func (Vindaloo) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:vindaloo", 0
}

// Name ...
func (Vindaloo) Name() string {
	return "Vindaloo"
}

// Texture ...
func (Vindaloo) Texture() image.Image {
	return textureFromName("vindaloo")
}
