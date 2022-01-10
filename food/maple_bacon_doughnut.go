package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MapleBaconDoughnut struct{}

// AlwaysConsumable ...
func (MapleBaconDoughnut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MapleBaconDoughnut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MapleBaconDoughnut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (MapleBaconDoughnut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MapleBaconDoughnut) Edible() bool {
	return true
}

// EncodeItem ...
func (MapleBaconDoughnut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:maple_bacon_doughnut", 0
}

// Name ...
func (MapleBaconDoughnut) Name() string {
	return "Maple Bacon Doughnut"
}

// Texture ...
func (MapleBaconDoughnut) Texture() image.Image {
	return textureFromName("maplebacondoughnut")
}
