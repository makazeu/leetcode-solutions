# Write your MySQL query statement below
SELECT Id
FROM Weather AS b
WHERE Temperature > (SELECT a.Temperature FROM Weather a WHERE DATEDIFF(b.RecordDate, a.RecordDate) = 1)

