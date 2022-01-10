package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type OystersRockefeller struct{}

// AlwaysConsumable ...
func (OystersRockefeller) AlwaysConsumable() bool {
	return false
}

// Category ...
func (OystersRockefeller) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (OystersRockefeller) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (OystersRockefeller) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (OystersRockefeller) Edible() bool {
	return true
}

// EncodeItem ...
func (OystersRockefeller) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:oysters_rockefeller", 0
}

// Name ...
func (OystersRockefeller) Name() string {
	return "OystersRockefeller"
}

// Texture ...
func (OystersRockefeller) Texture() image.Image {
	return textureFromName("oyster")
}
