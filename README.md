# Google I/O 2012 - Go Concurrency Patterns Code
The code in this repository implements and orders all the major code examples referenced in Rob Pike's talk found [here](https://www.youtube.com/watch?v=f6kdp27TYZs).

I have refactored it for clarity when possible as well as implemented enhancements.

# Enhancements

## Variadic and Slice Versions of `fanIn` Function

In the talk, Pike only covers a version of `fanIn` with a set number of two input channels. I implement two versions that can accept an arbitrary number of channels:
* `variadicFanIn` in `07_variadicFanIn.go` 
* `sliceFanIn` in `08_sliceFanIn.go`

In `11_sliceSelect.go`, I iterate on this to have one just call the other.

## Slice `select`

Pike only covers use of `select` for a set number of channels. I implement a `select` on an arbitrarily sized slice of input channels in `11_sliceSelect.go`.

## Google 3.5

Pike gets to Google 3.0 in his talk, which uses two replicas of each type of his simulated Google search. On my machine, this version still regularly failed the _80ms_ timeout. So I created a version in `17_googleSearch.go` called `Google3_5` that uses slices of 10 replicas per type and, not surprisingly, gets better performance than Pike's final version in the talk, sometimes sub-millisecond on my machine.