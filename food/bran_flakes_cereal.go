package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BranFlakesCereal struct{}

// AlwaysConsumable ...
func (BranFlakesCereal) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BranFlakesCereal) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BranFlakesCereal) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BranFlakesCereal) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BranFlakesCereal) Edible() bool {
	return true
}

// EncodeItem ...
func (BranFlakesCereal) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bran_flakes_cereal", 0
}

// Name ...
func (BranFlakesCereal) Name() string {
	return "Bran Flakes Cereal"
}

// Texture ...
func (BranFlakesCereal) Texture() image.Image {
	return textureFromName("cereal2")
}
