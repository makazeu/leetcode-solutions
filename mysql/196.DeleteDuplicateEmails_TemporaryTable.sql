# Write your MySQL query statement below
DELETE
FROM Person
WHERE Id IN (SELECT *
             FROM (SELECT Id
                   FROM Person AS a
                   WHERE a.Id > (SELECT MIN(b.Id) FROM Person b WHERE a.Email = b.Email)) c)


