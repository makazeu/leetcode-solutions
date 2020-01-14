# Write your MySQL query statement below
SELECT b.Name AS Department, a.Name AS Employee, Salary
FROM Employee a,
     Department b,
     (SELECT DepartmentId, MAX(Salary) AS MaxSalary FROM Employee GROUP BY DepartmentId) c
WHERE a.DepartmentId = c.DepartmentId
  AND a.DepartmentId = b.Id
  AND a.Salary = c.MaxSalary

