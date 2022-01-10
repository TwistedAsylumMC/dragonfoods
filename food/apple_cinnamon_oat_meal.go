package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type AppleCinnamonOatMeal struct{}

// AlwaysConsumable ...
func (AppleCinnamonOatMeal) AlwaysConsumable() bool {
	return false
}

// Category ...
func (AppleCinnamonOatMeal) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (AppleCinnamonOatMeal) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (AppleCinnamonOatMeal) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (AppleCinnamonOatMeal) Edible() bool {
	return true
}

// EncodeItem ...
func (AppleCinnamonOatMeal) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:apple_cinnamon_oat_meal", 0
}

// Name ...
func (AppleCinnamonOatMeal) Name() string {
	return "AppleCinnamonOatMeal"
}

// Texture ...
func (AppleCinnamonOatMeal) Texture() image.Image {
	return textureFromName("oatmeal3")
}
