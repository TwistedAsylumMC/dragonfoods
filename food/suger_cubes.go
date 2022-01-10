package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SugerCubes struct{}

// AlwaysConsumable ...
func (SugerCubes) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SugerCubes) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SugerCubes) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (SugerCubes) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SugerCubes) Edible() bool {
	return true
}

// EncodeItem ...
func (SugerCubes) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:suger_cubes", 0
}

// Name ...
func (SugerCubes) Name() string {
	return "SugerCubes"
}

// Texture ...
func (SugerCubes) Texture() image.Image {
	return textureFromName("sugercube")
}
