package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Profiteroles struct{}

// AlwaysConsumable ...
func (Profiteroles) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Profiteroles) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Profiteroles) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Profiteroles) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Profiteroles) Edible() bool {
	return true
}

// EncodeItem ...
func (Profiteroles) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:profiteroles", 0
}

// Name ...
func (Profiteroles) Name() string {
	return "Profiteroles"
}

// Texture ...
func (Profiteroles) Texture() image.Image {
	return textureFromName("profiteroles")
}
