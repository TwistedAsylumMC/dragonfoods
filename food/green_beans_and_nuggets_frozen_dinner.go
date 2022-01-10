package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenBeansAndNuggetsFrozenDinner struct{}

// AlwaysConsumable ...
func (GreenBeansAndNuggetsFrozenDinner) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenBeansAndNuggetsFrozenDinner) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenBeansAndNuggetsFrozenDinner) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(14, 8.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenBeansAndNuggetsFrozenDinner) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenBeansAndNuggetsFrozenDinner) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenBeansAndNuggetsFrozenDinner) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_beans_and_nuggets_frozen_dinner", 0
}

// Name ...
func (GreenBeansAndNuggetsFrozenDinner) Name() string {
	return "Green Beans And Nuggets Frozen Dinner"
}

// Texture ...
func (GreenBeansAndNuggetsFrozenDinner) Texture() image.Image {
	return textureFromName("tvdinner3")
}
