package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RamenNoodles struct{}

// AlwaysConsumable ...
func (RamenNoodles) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RamenNoodles) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RamenNoodles) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (RamenNoodles) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RamenNoodles) Edible() bool {
	return true
}

// EncodeItem ...
func (RamenNoodles) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ramen_noodles", 0
}

// Name ...
func (RamenNoodles) Name() string {
	return "Ramen Noodles"
}

// Texture ...
func (RamenNoodles) Texture() image.Image {
	return textureFromName("87_ramen")
}
