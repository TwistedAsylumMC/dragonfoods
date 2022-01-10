package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PeanutButterBananaandBaconTennessee struct{}

// AlwaysConsumable ...
func (PeanutButterBananaandBaconTennessee) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PeanutButterBananaandBaconTennessee) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PeanutButterBananaandBaconTennessee) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (PeanutButterBananaandBaconTennessee) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PeanutButterBananaandBaconTennessee) Edible() bool {
	return true
}

// EncodeItem ...
func (PeanutButterBananaandBaconTennessee) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:peanut_butter_bananaand_bacon_tennessee", 0
}

// Name ...
func (PeanutButterBananaandBaconTennessee) Name() string {
	return "PeanutButterBananaandBaconTennessee"
}

// Texture ...
func (PeanutButterBananaandBaconTennessee) Texture() image.Image {
	return textureFromName("sandwich6")
}
