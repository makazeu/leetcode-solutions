# Write your MySQL query statement below
SELECT Email
FROM (SELECT Email, COUNT(Email) AS count FROM Person GROUP BY Email) AS a
WHERE count > 1

