package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cucumber struct{}

// AlwaysConsumable ...
func (Cucumber) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cucumber) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cucumber) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cucumber) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cucumber) Edible() bool {
	return true
}

// EncodeItem ...
func (Cucumber) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cucumber", 0
}

// Name ...
func (Cucumber) Name() string {
	return "Cucumber"
}

// Texture ...
func (Cucumber) Texture() image.Image {
	return textureFromName("cucumber")
}
