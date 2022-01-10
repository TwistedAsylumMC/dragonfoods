package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SwedishMeatballs struct{}

// AlwaysConsumable ...
func (SwedishMeatballs) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SwedishMeatballs) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SwedishMeatballs) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (SwedishMeatballs) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SwedishMeatballs) Edible() bool {
	return true
}

// EncodeItem ...
func (SwedishMeatballs) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:swedish_meatballs", 0
}

// Name ...
func (SwedishMeatballs) Name() string {
	return "SwedishMeatballs"
}

// Texture ...
func (SwedishMeatballs) Texture() image.Image {
	return textureFromName("70_meatball_dish")
}
