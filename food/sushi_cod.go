package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SushiCod struct{}

// AlwaysConsumable ...
func (SushiCod) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SushiCod) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SushiCod) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SushiCod) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SushiCod) Edible() bool {
	return true
}

// EncodeItem ...
func (SushiCod) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sushi_cod", 0
}

// Name ...
func (SushiCod) Name() string {
	return "Sushi Cod"
}

// Texture ...
func (SushiCod) Texture() image.Image {
	return textureFromName("sushi_03_nigiri_sake")
}
