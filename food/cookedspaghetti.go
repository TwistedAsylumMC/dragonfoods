package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cookedspaghetti struct{}

// AlwaysConsumable ...
func (Cookedspaghetti) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cookedspaghetti) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cookedspaghetti) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cookedspaghetti) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cookedspaghetti) Edible() bool {
	return true
}

// EncodeItem ...
func (Cookedspaghetti) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cookedspaghetti", 0
}

// Name ...
func (Cookedspaghetti) Name() string {
	return "Cookedspaghetti"
}

// Texture ...
func (Cookedspaghetti) Texture() image.Image {
	return textureFromName("94_spaghetti")
}
