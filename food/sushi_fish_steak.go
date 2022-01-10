package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SushiFishSteak struct{}

// AlwaysConsumable ...
func (SushiFishSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SushiFishSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SushiFishSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (SushiFishSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SushiFishSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (SushiFishSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sushi_fish_steak", 0
}

// Name ...
func (SushiFishSteak) Name() string {
	return "SushiFishSteak"
}

// Texture ...
func (SushiFishSteak) Texture() image.Image {
	return textureFromName("sushi_03_nigiri_maguro")
}
