package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Eggplant struct{}

// AlwaysConsumable ...
func (Eggplant) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Eggplant) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Eggplant) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Eggplant) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Eggplant) Edible() bool {
	return true
}

// EncodeItem ...
func (Eggplant) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:eggplant", 0
}

// Name ...
func (Eggplant) Name() string {
	return "Eggplant"
}

// Texture ...
func (Eggplant) Texture() image.Image {
	return textureFromName("eggplant2")
}
