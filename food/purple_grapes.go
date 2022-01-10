package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PurpleGrapes struct{}

// AlwaysConsumable ...
func (PurpleGrapes) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PurpleGrapes) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PurpleGrapes) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (PurpleGrapes) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PurpleGrapes) Edible() bool {
	return true
}

// EncodeItem ...
func (PurpleGrapes) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:purple_grapes", 0
}

// Name ...
func (PurpleGrapes) Name() string {
	return "Purple Grapes"
}

// Texture ...
func (PurpleGrapes) Texture() image.Image {
	return textureFromName("red_grape")
}
