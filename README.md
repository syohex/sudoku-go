# sudoku solver in golang

## Usage

```bash
% cat data/text.txt
0,2,0,0,0,0,7,6,0
0,0,0,0,0,2,0,0,4
8,0,0,0,7,3,2,5,0
0,8,0,9,0,0,0,0,0
9,0,3,0,0,0,1,0,5
0,0,0,0,0,7,0,9,0
0,3,7,6,4,0,0,0,2
4,0,0,3,0,0,0,0,0
0,9,1,0,0,0,0,3,0
% go run main.go < data/text.txt
## Input
0 2 0  | 0 0 0  | 7 6 0
0 0 0  | 0 0 2  | 0 0 4
8 0 0  | 0 7 3  | 2 5 0
- - - - - - - - - - - -
0 8 0  | 9 0 0  | 0 0 0
9 0 3  | 0 0 0  | 1 0 5
0 0 0  | 0 0 7  | 0 9 0
- - - - - - - - - - - -
0 3 7  | 6 4 0  | 0 0 2
4 0 0  | 3 0 0  | 0 0 0
0 9 1  | 0 0 0  | 0 3 0

## Answer
3 2 4  | 1 9 5  | 7 6 8
7 5 9  | 8 6 2  | 3 1 4
8 1 6  | 4 7 3  | 2 5 9
- - - - - - - - - - - -
1 8 5  | 9 3 4  | 6 2 7
9 7 3  | 2 8 6  | 1 4 5
6 4 2  | 5 1 7  | 8 9 3
- - - - - - - - - - - -
5 3 7  | 6 4 1  | 9 8 2
4 6 8  | 3 2 9  | 5 7 1
2 9 1  | 7 5 8  | 4 3 6
```