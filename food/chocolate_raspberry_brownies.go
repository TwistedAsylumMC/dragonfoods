package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChocolateRaspberryBrownies struct{}

// AlwaysConsumable ...
func (ChocolateRaspberryBrownies) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChocolateRaspberryBrownies) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChocolateRaspberryBrownies) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChocolateRaspberryBrownies) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChocolateRaspberryBrownies) Edible() bool {
	return true
}

// EncodeItem ...
func (ChocolateRaspberryBrownies) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chocolate_raspberry_brownies", 0
}

// Name ...
func (ChocolateRaspberryBrownies) Name() string {
	return "Chocolate Raspberry Brownies"
}

// Texture ...
func (ChocolateRaspberryBrownies) Texture() image.Image {
	return textureFromName("pie4")
}
