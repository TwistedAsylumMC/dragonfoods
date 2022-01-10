package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Butter struct{}

// AlwaysConsumable ...
func (Butter) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Butter) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Butter) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Butter) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Butter) Edible() bool {
	return true
}

// EncodeItem ...
func (Butter) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:butter", 0
}

// Name ...
func (Butter) Name() string {
	return "Butter"
}

// Texture ...
func (Butter) Texture() image.Image {
	return textureFromName("butter")
}
