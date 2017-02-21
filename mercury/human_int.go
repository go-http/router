package router

import (
	"fmt"
)

const KiB = 1024
const MiB = 1024 * KiB
const GiB = 1024 * MiB

type HumanInt int

func (v HumanInt) String() string {
	if v > GiB {
		return fmt.Sprintf("%.1fG", float32(v)/float32(GiB))
	} else if v > MiB {
		return fmt.Sprintf("%.1fM", float32(v)/float32(MiB))
	} else if v > KiB {
		return fmt.Sprintf("%.1fK", float32(v)/float32(KiB))
	} else {
		return fmt.Sprintf("%.1fB", float32(v))
	}
}
