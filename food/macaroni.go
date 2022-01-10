package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Macaroni struct{}

// AlwaysConsumable ...
func (Macaroni) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Macaroni) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Macaroni) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Macaroni) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Macaroni) Edible() bool {
	return true
}

// EncodeItem ...
func (Macaroni) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:macaroni", 0
}

// Name ...
func (Macaroni) Name() string {
	return "Macaroni"
}

// Texture ...
func (Macaroni) Texture() image.Image {
	return textureFromName("noodle_macaroni")
}
