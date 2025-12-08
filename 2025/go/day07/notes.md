# AoC 2025 day 7

## Part 2

counting number of possible paths the beam can go

count distinct beams?

Problems with test data working extremly fast and getting right answer
But my real input completely explodes and runs forever (I force quit after a while)

IMPORTANT
Need to memoize my recursive calculations before the number of loops, doing the exact same calculations, spin out of control.
I've tried several easy approaches to this problem but they time out due to underestimating the number of splits and beams going into the same path.

The memoize I did will not work if using goroutines to paralellize.
Would need some kind of lock, "sync.Mutex" perhaps

---

Check Johans part2, really clean. No memoize.
I did something similar with weights, but copied the board in every loop - very costly
