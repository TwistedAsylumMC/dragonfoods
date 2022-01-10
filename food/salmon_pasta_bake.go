package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SalmonPastaBake struct{}

// AlwaysConsumable ...
func (SalmonPastaBake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SalmonPastaBake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SalmonPastaBake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (SalmonPastaBake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SalmonPastaBake) Edible() bool {
	return true
}

// EncodeItem ...
func (SalmonPastaBake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:salmon_pasta_bake", 0
}

// Name ...
func (SalmonPastaBake) Name() string {
	return "SalmonPastaBake"
}

// Texture ...
func (SalmonPastaBake) Texture() image.Image {
	return textureFromName("salmonpastabake")
}
