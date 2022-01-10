package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueBarrys struct{}

// AlwaysConsumable ...
func (BlueBarrys) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueBarrys) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueBarrys) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueBarrys) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueBarrys) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueBarrys) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_barrys", 0
}

// Name ...
func (BlueBarrys) Name() string {
	return "Blue Barrys"
}

// Texture ...
func (BlueBarrys) Texture() image.Image {
	return textureFromName("bluebarrys")
}
