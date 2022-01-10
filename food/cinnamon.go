package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cinnamon struct{}

// AlwaysConsumable ...
func (Cinnamon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cinnamon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cinnamon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cinnamon) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Cinnamon) Edible() bool {
	return true
}

// EncodeItem ...
func (Cinnamon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cinnamon", 0
}

// Name ...
func (Cinnamon) Name() string {
	return "Cinnamon"
}

// Texture ...
func (Cinnamon) Texture() image.Image {
	return textureFromName("cinnamon")
}
