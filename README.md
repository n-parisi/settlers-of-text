# settlers-of-text
Game client for running a text based Settlers of Catan game

--------------------
### Sample Board
```
                      00 ---- 01
                      /   09   \
                     /          \
            02 ---- 03  GGGGGG  04 ---- 05
            /   12   \  GGGGGG  /   10   \
           /          \ GGGGGG /          \
  06 ---- 07  OOOOOO   08 ---- 09 BBBBBB  10 ---- 11
  /    11  \  OOOOOO  /   05   \  BBBBBB  /   08   \
 /          \ OOOOOO /          \ BBBBBB /          \
12  LLLLLL  13 ---- 14  LLLLLL  15 ---- 16  WWWWWW  17
 \  LLLLLL  /   06   \  LLLLLL  /   04   \  WWWWWW  /
  \ LLLLLL /          \ LLLLLL /          \ WWWWWW /
  18 ---- 19  GGGGGG  20 ---- 21  OOOOOO  22 ---- 23
  /   04   \  GGGGGG  /   11   \  OOOOOO  /   03   \
 /          \ GGGGGG /          \ OOOOOO /          \
24  BBBBBB  25 ---- 26  OOOOOO  27 ---- 28  GGGGGG  29
 \  BBBBBB  /   03   \  OOOOOO  /   09   \  GGGGGG  /
  \ BBBBBB /   RRRR   \ OOOOOO /          \ GGGGGG /
  30 ---- 31  WWWWWW  32 ---- 33  GGGGGG  34 ---- 35
  /   08   \  WWWWWW  /   10   \  GGGGGG  /   06   \
 /          \ WWWWWW /          \ GGGGGG /          \
36  LLLLLL  37 ---- 38  WWWWWW  39 ---- 40  LLLLLL  41
 \  LLLLLL  /        \  WWWWWW  /   02   \  LLLLLL  /
  \ LLLLLL /          \ WWWWWW /          \ LLLLLL /
  42 ---- 43          44 ---- 45  BBBBBB  46 ---- 47
           \          /   05   \  BBBBBB  /
            \        /          \ BBBBBB /
            48 ---- 49  WWWWWW  50 ---- 51
                     \  WWWWWW  /
                      \ WWWWWW /
                      52 ---- 53
```

Edges are identified by a number. These numbers would not typically be shown to make the board more easily understandable, and would only need to be displayed when a player needs to designate where to build a structure.

The letter on the bottom is the resource for that space. The letter could be color coded for easier visibility. Roads would also be colored.

B = Brick  
L = Lumber  
O = Ore  
G = Grain  
W = Wool  
