package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SushiClamChowder struct{}

// AlwaysConsumable ...
func (SushiClamChowder) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SushiClamChowder) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SushiClamChowder) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (SushiClamChowder) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SushiClamChowder) Edible() bool {
	return true
}

// EncodeItem ...
func (SushiClamChowder) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sushi_clam_chowder", 0
}

// Name ...
func (SushiClamChowder) Name() string {
	return "SushiClamChowder"
}

// Texture ...
func (SushiClamChowder) Texture() image.Image {
	return textureFromName("sushi_03_nigiri_ika")
}
