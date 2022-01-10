package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueberryPineappleSmoothie struct{}

// AlwaysConsumable ...
func (BlueberryPineappleSmoothie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueberryPineappleSmoothie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueberryPineappleSmoothie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueberryPineappleSmoothie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueberryPineappleSmoothie) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueberryPineappleSmoothie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blueberry_pineapple_smoothie", 0
}

// Name ...
func (BlueberryPineappleSmoothie) Name() string {
	return "BlueberryPineappleSmoothie"
}

// Texture ...
func (BlueberryPineappleSmoothie) Texture() image.Image {
	return textureFromName("smoothie")
}
