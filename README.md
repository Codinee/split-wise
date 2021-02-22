# split-wise

A split-wise calculator that given a spend position and participants, outputs the most optimized funds-due matrix that reduces the total number of individual payments. 

Input file format
----------------
|Items|Person|Spent|Weightage|
|----------------------------|
|Itm1|P1|10|1|
||P2|0|2

Items - item for which share is calculated
Person - friends name who share the expense
Spent - amount spent by person for the item 
Weightage - weightage of the persons share for the item ie. if 6 persons go to hotel and 3 of them want to split the cost in non-even ratio for eg. 4:1:1. So weightages will be 4, 1 and 1 

1. Download package
2. Execute the main program with sharing item details in expense.csv file in the shary folder.
3. The split details will be updated in the same expense.csv file. 

