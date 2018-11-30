# Tic-Tac-Toe

This project provides a simple set of tic-tac-toe playing bots. It was developed
as a playground to explore the ideas presented in the first chapter of [Reinforcement
Learning: An Introduction](http://incompleteideas.net/book/the-book.html) by Sutton
and Barto.

## Project Structure

At the project root is the tic-tac-toe gamestate. This provides a simple board
representation and function to determine if there is a winner.

The directory bots contains several bots that respond with a move when given a
gamestate.

In cmd are several small programs for experimenting with the behavior of the bots.
Arena competes selected bots head to head. Training trains the Reinforcement
Learning bot and checks its performance against other bots periodically.
Rigged provides a contrived example of how Reinforcement Learning can outperform
MiniMax, even though MiniMax supposedly finds the optimal solution.

## Gamestate Representation

Game state is represented by treating the lower 18 bits of a unit32 as 9 "cells".
If the cell is 0b00, it is considered empty. 0b01 is X, 0b10 is 0, and 0x11 is
invalid, of course. Several masks are provided for pulling out specific wins, as
well as conversion functions to and from strings.

## Bots

### RandomX

There are a variety of random bots. They attempt to make a random move, or follow
simple opportunistic strategies, but never look more than one move ahead.

### MiniMax

Minimax exhaustively searches the space of possible games from the current move.
Accordingly, it will play the known optimal strategy.

Additionally a variant of Minimax that makes a "mistake" at a specified random
frequency is implemented.

### Reinforcement Learning

The Reinforcement Learning bot is very simple, using a lookup table for its value
function. Because gamestate is represented as bits in an uint, the table is simply
sized to the maximum possible game state, and lookups can occur by treating the
gamestate as a number. The bot can be used in two modes, either a learning mode
(which performs exploratory moves and alters its knowledge) or only make what it
believes is the optimal move, avoiding updating its data.

The value signal it receives is weather it one (applies 1) or lost (applies 0)
a game. As draws are ignored, it treats them as valuable as the initial state, 0.5.

The value function it implements is very simple, it simply looks at the set of
possible moves, and picks the one that in its experience has lead to the most
success.

## Results

*Arena*
```
 win/loss/draw  Players
 431  405  164  random vs random
 417  435  148  random vs random-spoiler
 254  659   87  random vs random-opportunistic
 285  650   65  random vs random-opportunistic-spoiler
 491  457   52  random-opportunistic vs random-opportunistic-spoiler
 485  465   50  random-opportunistic vs random-opportunistic
 672  258   70  random-opportunistic vs random-spoiler
 472  430   98  random-spoiler vs random-spoiler
 212  656  132  random vs minimax-sometimes-random-0.500
 369  540   91  random-opportunistic-spoiler vs minimax-sometimes-random-0.500
   0  900  100  random vs minimax
   0  890  110  random-opportunistic-spoiler vs minimax
 422  377  201  minimax-sometimes-random-0.500 vs minimax-sometimes-random-0.500
 519  218  263  minimax-sometimes-random-0.250 vs minimax-sometimes-random-0.500
 621   51  328  minimax-sometimes-random-0.050 vs minimax-sometimes-random-0.500
 293  300  407  minimax-sometimes-random-0.250 vs minimax-sometimes-random-0.250
 100   96  804  minimax-sometimes-random-0.050 vs minimax-sometimes-random-0.050
 688    0  312  minimax vs minimax-sometimes-random-0.500
 428    0  572  minimax vs minimax-sometimes-random-0.250
  89    0  911  minimax vs minimax-sometimes-random-0.050
   0    0 1000  minimax vs minimax
```
Above is an abbreviated list of the results from the arena.

*Trainer*
```
     random            self         0.1 minimax        minimax
 win/loss/draw    win/loss/draw    win/loss/draw    win/loss/draw
 453  534   13    500  500    0     55  944    1      0 1000    0
 599  340   61      0    0 1000    105  473  422      0  500  500
 791  127   82      0    0 1000    177    5  818      0    0 1000
 880   34   86      0    0 1000    182    0  818      0    0 1000
```
Trainer pits the RL bot against several opponents. Starting with no training,
then increasing the amount of training each round. This has my favourite result:
In the second round, the RL bot has learned to beat minimax given a specific
starting position (first or second), but has not figured out the other position,
as evidenced by its perfect 50% win loss ratio. Also interesting is that against
an opponent that 10% of the time makes a random move, it take much longer to learn
to hold that perfect no-loss record. The ideal opponent will only expose RL to
some 1300 game states, but an opponent that sometimes makes a random move (based
on the 1% win rate, sometimes several per game) the RL bot has to learn more paths.
Being patient and adding even more rounds of training (or more rounds against a
purely random player) will get the RL bot to get near the 9-1 win-draw ration that
minimax has against random.

*Rigged*
```
 rl v corners     rl v minimax    corners v minimax
 win/loss/draw    win/loss/draw    win/loss/draw
 500  500    0      0 1000    0      0  500  500
1000    0    0      0 1000    0      0  500  500
1000    0    0      0 1000    0      0  500  500
1000    0    0      0    0 1000      0  500  500
```

Here we see that RL fairly quickly learns to defeat corners, the contrived bot meant
to allow a non-optimal player to defeat it. Whereas minimax never learns to win more
than 50% of games (which bot starts first determines the outcome). What is impressive
is that in the last round, after the rl bot has had some thousands of rounds against
minimax, it has learned to both defeat corners, and always play minimax to a draw.

## Design Mistakes

This project was quickly hacked together for fun in a couple bars. Naturally, it
does not reflect the professional practice of the author. In hindsight, the
following mistakes were made:

* Making moves -
The player is expected to return exactly the bit it wants OR'd against the existing
gamestate. Naturally, a truly malicious bot would use this to win every game. This
is also awkward to use as a developer. As well, the frequent use of AllX and AllO
to mask against possible moves, having to be sensitive to which player is current,
causes a lot of repeated code.

* Game representation -
Parts of the game representation would likely have been easier to make use of if
instead of representing the gamestate as 9 cells of X, O, or empty, it had been
9 bits of X followed by 9 bits of O. Win state evaluation would have been "Check
these nine bits", followed by shifting 9 bits to the right and repeating with a
different possible winner. Similarly, when speculating moves we could have limited
checks to the player making the move. Which moves were free could have been checked
with (gamestate | (gamestate >> 9)) & 0x1FF. Logic passed to players could have
ignored which was the current player and simply ensured the lower 9 bits represented
the current player.

* MiniMax calculation -
Assuming one has a deterministic strategy, they only need to remember 490 moves
to play all possible games of tic-tac-toe against an opponent. This includes 1 for
the desired initial move. As move calculation for Minimax initially took 2ms on
most hardware, simulations could have occurred much faster with a faster bot.
Even though this would not have worked optimally for the SometimesRandom variant,
recalculation would still have been reduced.

* RL value function representation -
The naive representation currently allows for 262144 game states, even though far
fewer states are actually reachable in tic tac toe. For one, this allows for illegal
boards such as all X's, and for two it allows games that contain multiple wins.
Using a map, or a more clever  would use less space. That said, calculating a move
for the RL bot takes less time than a random player, so performance is not a problem,
and the data behind the value function is likely highly compressible (as most states
will not have been encountered, and long sequences of 0.5 should exist).
This design is somewhat a consequence of the model-free nature of the bot.
