package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SaltineCrackers struct{}

// AlwaysConsumable ...
func (SaltineCrackers) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SaltineCrackers) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SaltineCrackers) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (SaltineCrackers) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SaltineCrackers) Edible() bool {
	return true
}

// EncodeItem ...
func (SaltineCrackers) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:saltine_crackers", 0
}

// Name ...
func (SaltineCrackers) Name() string {
	return "SaltineCrackers"
}

// Texture ...
func (SaltineCrackers) Texture() image.Image {
	return textureFromName("saltinecracker")
}
