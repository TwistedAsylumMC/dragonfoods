package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Peppermint struct{}

// AlwaysConsumable ...
func (Peppermint) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Peppermint) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Peppermint) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Peppermint) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Peppermint) Edible() bool {
	return true
}

// EncodeItem ...
func (Peppermint) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:peppermint", 0
}

// Name ...
func (Peppermint) Name() string {
	return "Peppermint"
}

// Texture ...
func (Peppermint) Texture() image.Image {
	return textureFromName("peppermint")
}
