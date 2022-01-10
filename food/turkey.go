package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Turkey struct{}

// AlwaysConsumable ...
func (Turkey) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Turkey) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Turkey) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(15, 9.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Turkey) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Turkey) Edible() bool {
	return true
}

// EncodeItem ...
func (Turkey) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:turkey", 0
}

// Name ...
func (Turkey) Name() string {
	return "Turkey"
}

// Texture ...
func (Turkey) Texture() image.Image {
	return textureFromName("85_roastedchicken")
}
