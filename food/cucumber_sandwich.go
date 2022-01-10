package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CucumberSandwich struct{}

// AlwaysConsumable ...
func (CucumberSandwich) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CucumberSandwich) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CucumberSandwich) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CucumberSandwich) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CucumberSandwich) Edible() bool {
	return true
}

// EncodeItem ...
func (CucumberSandwich) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cucumber_sandwich", 0
}

// Name ...
func (CucumberSandwich) Name() string {
	return "Cucumber Sandwich"
}

// Texture ...
func (CucumberSandwich) Texture() image.Image {
	return textureFromName("sandwich5")
}
