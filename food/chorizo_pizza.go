package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChorizoPizza struct{}

// AlwaysConsumable ...
func (ChorizoPizza) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChorizoPizza) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChorizoPizza) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChorizoPizza) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChorizoPizza) Edible() bool {
	return true
}

// EncodeItem ...
func (ChorizoPizza) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chorizo_pizza", 0
}

// Name ...
func (ChorizoPizza) Name() string {
	return "ChorizoPizza"
}

// Texture ...
func (ChorizoPizza) Texture() image.Image {
	return textureFromName("pizza7")
}
