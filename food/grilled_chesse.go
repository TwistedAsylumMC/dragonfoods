package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GrilledChesse struct{}

// AlwaysConsumable ...
func (GrilledChesse) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GrilledChesse) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GrilledChesse) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (GrilledChesse) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GrilledChesse) Edible() bool {
	return true
}

// EncodeItem ...
func (GrilledChesse) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:grilled_chesse", 0
}

// Name ...
func (GrilledChesse) Name() string {
	return "Grilled Chesse"
}

// Texture ...
func (GrilledChesse) Texture() image.Image {
	return textureFromName("grilled_chesse")
}
