You are given an integer array cars, where each element represents a car moving along a straight line.
The index of each element represents the car’s position on the line.
The absolute value of each element represents the car’s mass.
The sign represents the car’s direction:
Positive value → moving to the right
Negative value → moving to the left
All cars move at the same speed.
Collision Rules
A collision occurs only when a car moving right meets a car moving left.
When two cars collide:
The car with the smaller mass explodes (is removed).
If both have the same mass, both explode.
Cars moving in the same direction never collide.
Task
Return the final state of the cars after all possible collisions have occurred.
Examples
Example 1
Input:
cars = [5, 10, -5]
Output:
[5, 10]
Explanation:
The 10 (right) and -5 (left) collide → -5 explodes.
The 5 and 10 move in the same direction, so no collision.
Example 2
Input:
cars = [8, -8]
Output:
[]
Explanation:
Both have equal mass, so both explode.
Example 3
Input:
cars = [3, 5, -6, 2,-7, -1, 4]
Output:
[-6, 2, 4]
Explanation:
-6 collides with 5 → 5 explodes
-6 collides with 3 → 3 explodes
-6 continues moving left
2 collides with -1 → -1 explodes
2 and 4 move right and never collide


[-6, 2, 

[7, 5, -6, 2,-7, -1, 4]
	

[7,5,2,



entity
car 

behaviour
left, right => (direction)
spped


sudoku:

cars = [5, 10, -5]

iterate all the cars items values



var temp
isNegativeValue=false

if negative value
	temp=
	isNegativeValue=true
check highest value 5> 10 return 5 else return 10
	

main(){
	one
	
	
}


