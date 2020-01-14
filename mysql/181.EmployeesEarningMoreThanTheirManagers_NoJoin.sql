# Write your MySQL query statement below
SELECT a.Name AS Employee
FROM Employee AS a
WHERE a.Salary > (SELECT Salary FROM Employee WHERE Id = a.ManagerId)


