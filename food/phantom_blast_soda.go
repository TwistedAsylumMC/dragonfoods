package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PhantomBlastSoda struct{}

// AlwaysConsumable ...
func (PhantomBlastSoda) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PhantomBlastSoda) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PhantomBlastSoda) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (PhantomBlastSoda) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PhantomBlastSoda) Edible() bool {
	return true
}

// EncodeItem ...
func (PhantomBlastSoda) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:phantom_blast_soda", 0
}

// Name ...
func (PhantomBlastSoda) Name() string {
	return "Phantom Blast Soda"
}

// Texture ...
func (PhantomBlastSoda) Texture() image.Image {
	return textureFromName("soft_drink_yellow")
}
