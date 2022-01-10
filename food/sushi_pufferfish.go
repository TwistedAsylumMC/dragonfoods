package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SushiPufferfish struct{}

// AlwaysConsumable ...
func (SushiPufferfish) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SushiPufferfish) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SushiPufferfish) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SushiPufferfish) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SushiPufferfish) Edible() bool {
	return true
}

// EncodeItem ...
func (SushiPufferfish) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sushi_pufferfish", 0
}

// Name ...
func (SushiPufferfish) Name() string {
	return "Sushi Pufferfish"
}

// Texture ...
func (SushiPufferfish) Texture() image.Image {
	return textureFromName("sushi_03_nigiri_tai")
}
