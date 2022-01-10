package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChickenAlfredo struct{}

// AlwaysConsumable ...
func (ChickenAlfredo) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChickenAlfredo) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChickenAlfredo) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChickenAlfredo) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChickenAlfredo) Edible() bool {
	return true
}

// EncodeItem ...
func (ChickenAlfredo) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chicken_alfredo", 0
}

// Name ...
func (ChickenAlfredo) Name() string {
	return "Chicken Alfredo"
}

// Texture ...
func (ChickenAlfredo) Texture() image.Image {
	return textureFromName("chickenalfredo")
}
