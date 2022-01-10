package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Macnchesseandpotatofrozendinner struct{}

// AlwaysConsumable ...
func (Macnchesseandpotatofrozendinner) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Macnchesseandpotatofrozendinner) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Macnchesseandpotatofrozendinner) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(14, 8.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Macnchesseandpotatofrozendinner) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Macnchesseandpotatofrozendinner) Edible() bool {
	return true
}

// EncodeItem ...
func (Macnchesseandpotatofrozendinner) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:macnchesseandpotatofrozendinner", 0
}

// Name ...
func (Macnchesseandpotatofrozendinner) Name() string {
	return "Macnchesseandpotatofrozendinner"
}

// Texture ...
func (Macnchesseandpotatofrozendinner) Texture() image.Image {
	return textureFromName("tvdinner2")
}
