package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Sushi struct{}

// AlwaysConsumable ...
func (Sushi) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Sushi) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Sushi) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Sushi) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Sushi) Edible() bool {
	return true
}

// EncodeItem ...
func (Sushi) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sushi", 0
}

// Name ...
func (Sushi) Name() string {
	return "Sushi"
}

// Texture ...
func (Sushi) Texture() image.Image {
	return textureFromName("sushi2")
}
