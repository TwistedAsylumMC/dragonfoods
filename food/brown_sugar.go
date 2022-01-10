package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BrownSugar struct{}

// AlwaysConsumable ...
func (BrownSugar) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BrownSugar) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BrownSugar) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BrownSugar) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BrownSugar) Edible() bool {
	return true
}

// EncodeItem ...
func (BrownSugar) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:brown_sugar", 0
}

// Name ...
func (BrownSugar) Name() string {
	return "BrownSugar"
}

// Texture ...
func (BrownSugar) Texture() image.Image {
	return textureFromName("sugarbrown")
}
