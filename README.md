# Path Coverage by Exhaustion

Sometimes, trying for the hardest possible
case turns out to be easier than the "easy"
cases.

This is a demonstration of how easy it is
to get:
* full path coverage of code 
  * test _every single possible path_ through
  a body of code
* using exhaustion
  * you can guess from the name that it
might be hard (:-))

In fact, we get XX% coverage of a library
with a 74-line test of the main() function
of a caller.

The trick is to iterate across a 
deliberately-crafted set of values 
for each parameter instead of trying to 
hand-craft unit tests of everything.