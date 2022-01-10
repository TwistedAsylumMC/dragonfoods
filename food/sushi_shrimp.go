package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SushiShrimp struct{}

// AlwaysConsumable ...
func (SushiShrimp) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SushiShrimp) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SushiShrimp) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (SushiShrimp) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SushiShrimp) Edible() bool {
	return true
}

// EncodeItem ...
func (SushiShrimp) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sushi_shrimp", 0
}

// Name ...
func (SushiShrimp) Name() string {
	return "Sushi Shrimp"
}

// Texture ...
func (SushiShrimp) Texture() image.Image {
	return textureFromName("sushi_03_nigiri_ebi")
}
