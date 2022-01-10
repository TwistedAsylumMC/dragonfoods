package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type YellowJello struct{}

// AlwaysConsumable ...
func (YellowJello) AlwaysConsumable() bool {
	return false
}

// Category ...
func (YellowJello) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (YellowJello) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (YellowJello) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (YellowJello) Edible() bool {
	return true
}

// EncodeItem ...
func (YellowJello) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellow_jello", 0
}

// Name ...
func (YellowJello) Name() string {
	return "Yellow Jello"
}

// Texture ...
func (YellowJello) Texture() image.Image {
	return textureFromName("60_jelly_dish")
}
