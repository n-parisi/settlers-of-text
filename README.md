# settlers-of-text
Game client for running a text based Settlers of Catan game

--------------------
### Sample Board
```
                       + ---- +
                      /    C   \
                     /     9    \
             + ---- +            + ---- +
            /    B   \          /    D   \
           /     12   \    G   /     10   \
   + ---- +            + ---- +            + ---- +
  /    A   \          /    N   \          /    E   \
 /     11   \    O   /     5    \    B   /     8    \
+            + ---- +            + ---- +            +
 \          /    M   \          /    O   \          /
  \    L   /     6    \    L   /     4    \    W   /
   + ---- +            + ---- +            + ---- +
  /    L   \          /    S   \          /    F   \
 /     4    \    G   /     11   \    O   /     3    \
+            + ---- +            + ---- +            +
 \          /    R   \          /    P   \          /
  \    B   /     3(R) \    O   /     9    \    G   /
   + ---- +            + ---- +            + ---- +
  /    K   \          /    Q   \          /    G   \
 /     8    \    W   /     10   \    G   /     6    \
+            + ---- +            + ---- +            +
 \          /    J   \          /    H   \          /
  \    L   /          \    W   /     2    \    L   /
   + ---- +      X     + ---- +            + ---- +
           \          /    I   \          /
            \        /     5    \    B   /
             + ---- +            + ---- +
                     \          /
                      \    W   /
                       + ---- +
```

The letters at the top are identifiers for each space. They are always the same for each game. The number below them is the value for the space - what dice roll activates that space.  `(R)` next to the number means the robber is currently on that space.

The letter on the bottom is the resource for that space. The letter could be color coded for easier visibility. Roads would also be colored.

B = Brick  
L = Lumber  
O = Ore  
G = Grain  
W = Wool  

Normally the letter identifier would only be visible when a player needs to use them to place a structure. Also allow player to filter a version of the board with only certain resources, or certain player structure/roads. 
