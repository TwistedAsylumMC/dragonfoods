package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenGrapes struct{}

// AlwaysConsumable ...
func (GreenGrapes) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenGrapes) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenGrapes) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenGrapes) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenGrapes) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenGrapes) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_grapes", 0
}

// Name ...
func (GreenGrapes) Name() string {
	return "Green Grapes"
}

// Texture ...
func (GreenGrapes) Texture() image.Image {
	return textureFromName("green_grape")
}
