# settlers-of-text
Game client for running a text based Settlers of Catan game

--------------------
### Sample Board
```
                      00 ---- 01
                      /   09   \
                     /          \
            02 ---- 03  GGGGGG  04 ---- 05
            /    12  \  GGGGGG  /   10   \
           /          \ GGGGGG /          \
  06 ---- 07   OOOOO   08 ---- 09  BBBBB  10 ---- 11
  /    A   \   OOOOO  /    N   \   BBBBB  /    E   \
 /     11   \  OOOOO /     5    \  BBBBB /     8    \
12          13 ---- 14          15 ---- 16           17
 \          /    M   \          /    O   \          /
  \    L   /     6    \    L   /     4    \    W   /
  18 ---- 19          20 ---- 21          22 ---- 23
  /    L   \          /    S   \          /    F   \
 /     4    \    G   /     11   \    O   /     3    \
24          25 ---- 26          27 ---- 28           29
 \          /    R   \          /    P   \          /
  \    B   /     3(R) \    O   /     9    \    G   /
  30 ---- 31          32 ---- 33          34 ---- 35
  /    K   \          /    Q   \          /    G   \
 /     8    \    W   /     10   \    G   /     6    \
36          37 ---- 38          39 ---- 40           41
 \          /    J   \          /    H   \          /
  \    L   /          \    W   /     2    \    L   /
  42 ---- 43          44 ---- 45          46 ---- 47
           \          /    I   \          /
            \        /     5    \    B   /
            48 ---- 49          50 ---- 51
                     \          /
                      \    W   /
                      52 ---- 53
```

The letters at the top are identifiers for each space. They are always the same for each game. The number below them is the value for the space - what dice roll activates that space.  `(R)` next to the number means the robber is currently on that space.

The letter on the bottom is the resource for that space. The letter could be color coded for easier visibility. Roads would also be colored.

B = Brick  
L = Lumber  
O = Ore  
G = Grain  
W = Wool  

Normally the letter identifier would only be visible when a player needs to use them to place a structure. Also allow player to filter a version of the board with only certain resources, or certain player structure/roads. 
