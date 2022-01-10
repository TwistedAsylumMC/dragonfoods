package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CornOnTheCob struct{}

// AlwaysConsumable ...
func (CornOnTheCob) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CornOnTheCob) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CornOnTheCob) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CornOnTheCob) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CornOnTheCob) Edible() bool {
	return true
}

// EncodeItem ...
func (CornOnTheCob) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:corn_on_the_cob", 0
}

// Name ...
func (CornOnTheCob) Name() string {
	return "CornOnTheCob"
}

// Texture ...
func (CornOnTheCob) Texture() image.Image {
	return textureFromName("corn2")
}
