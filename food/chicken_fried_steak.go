package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChickenFriedSteak struct{}

// AlwaysConsumable ...
func (ChickenFriedSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChickenFriedSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChickenFriedSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChickenFriedSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChickenFriedSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (ChickenFriedSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chicken_fried_steak", 0
}

// Name ...
func (ChickenFriedSteak) Name() string {
	return "Chicken Fried Steak"
}

// Texture ...
func (ChickenFriedSteak) Texture() image.Image {
	return textureFromName("chickenfriedsteak")
}
