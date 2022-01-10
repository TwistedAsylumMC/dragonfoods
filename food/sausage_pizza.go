package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SausagePizza struct{}

// AlwaysConsumable ...
func (SausagePizza) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SausagePizza) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SausagePizza) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SausagePizza) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SausagePizza) Edible() bool {
	return true
}

// EncodeItem ...
func (SausagePizza) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sausage_pizza", 0
}

// Name ...
func (SausagePizza) Name() string {
	return "Sausage Pizza"
}

// Texture ...
func (SausagePizza) Texture() image.Image {
	return textureFromName("pizza6")
}
