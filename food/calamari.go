package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Calamari struct{}

// AlwaysConsumable ...
func (Calamari) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Calamari) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Calamari) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Calamari) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Calamari) Edible() bool {
	return true
}

// EncodeItem ...
func (Calamari) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:calamari", 0
}

// Name ...
func (Calamari) Name() string {
	return "Calamari"
}

// Texture ...
func (Calamari) Texture() image.Image {
	return textureFromName("calamari")
}
