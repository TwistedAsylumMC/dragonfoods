package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Strawberrys struct{}

// AlwaysConsumable ...
func (Strawberrys) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Strawberrys) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Strawberrys) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Strawberrys) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Strawberrys) Edible() bool {
	return true
}

// EncodeItem ...
func (Strawberrys) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:strawberrys", 0
}

// Name ...
func (Strawberrys) Name() string {
	return "Strawberrys"
}

// Texture ...
func (Strawberrys) Texture() image.Image {
	return textureFromName("strawberry")
}
