package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SalmonSushi struct{}

// AlwaysConsumable ...
func (SalmonSushi) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SalmonSushi) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SalmonSushi) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (SalmonSushi) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SalmonSushi) Edible() bool {
	return true
}

// EncodeItem ...
func (SalmonSushi) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:salmon_sushi", 0
}

// Name ...
func (SalmonSushi) Name() string {
	return "SalmonSushi"
}

// Texture ...
func (SalmonSushi) Texture() image.Image {
	return textureFromName("salmon_sushi")
}
