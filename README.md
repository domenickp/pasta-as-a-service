# pasta-as-a-service
The family pasta recipe, implemented as a REST API.

## History
Back in the 90's, my father taught me how to make pasta from scratch, using the recipe
my grandfather learned when he was a boy in Italy, passed to him by his parents.  When
I looked at the ingredient amounts, they were all ratios based on the number of people
you intended to feed.  It seemed like the perfect thing to codify, so I wrote some vbs
(the shame!) to make the Petrella Pasta Calculator, which provided a simple Windows UI
to calculate the ingredient amounts, based on the number of people that needed pasta.

Rewriting the code to scale the recipe up and down has been on my mind for a while, so
here it is.  This first version takes care of the calculations, and converting the results
to standard cup measurements, and figuring out the number of eggs required.  A few more
tweaks are needed there, then I will convert to a REST API, and finally make a web client.

## Scaling the recipe
The first version of this I wrote way back when listed the ingredient amounts in decimal
form.  I didn't mind measuring 1.5938 cups of flour, but knew it was a little weird.  This
time I wrote a binary search to figure out the closest decimal equivalent of a standard
fractional measurement, and convert that to words.  The plan is to make this a general
service to scale any recipe.

## TODO
A lot.
* Convert to a REST API
* Add a function to output ingredient amounts in fractions, not just words
* Re-organize the code, rename some variables for clarity
* Come up with a data storage scheme for recipes.


