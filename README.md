# lem-in

This project is a digital version of an ant farm.

program lem-in that will read from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

Standard output in the following format:
```
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

* x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.
* A room is defined by "name coord_x coord_y", and will usually look like "Room 1 2", "nameoftheroom 1 6", "4 6 7"
* The links are defined by "name1-name2" and will usually look like "1-2", "2-5".

# Usage
```
student$ ./lem-in test0.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
student$
```
