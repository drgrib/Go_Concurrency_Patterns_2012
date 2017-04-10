# Google I/O 2012 - Go Concurrency Patterns
https://www.youtube.com/watch?v=f6kdp27TYZs

The code in this repository implements and orders all the major code examples referenced in Rob Pike's talk, linked to above. I have refactored it for clarity when possible as well as implemented enhancements.

# Enhancements

## Variadic and Slice Versions of the `fanIn` Function

In the talk, Pike only covers a version of `fanIn` with a set number (2) of input channels. I implement two versions that can accept an arbitrary number of channels:
* `variadicFanIn` in `07_variadicFanIn.go` 
* `sliceFanIn` in `08_sliceFanIn.go`

In `11_sliceSelect.go`, I further iterate on this to have one just call the other.

## Slice `select`

Pike only covers use of `select` for a set number of channels. I implement a `select` on an arbitrarily sized slice of input channels in `11_sliceSelect.go`.

## Google 4.0

Pike gets to Google 3.0 in his talk, which uses two replicas of his simulated Google search. On my machine, this version still regularly failed the 80ms timeout. So I created a version in `17_googleSearch.go` called `Google4` that uses 10 replicas and, not surprisingly, gets better performance than Pike's final version in the talk, sometimes sub-millisecond.