# Write your MySQL query statement below
SELECT id,
       CASE
           WHEN id % 2 = 0 THEN (SELECT student FROM seat WHERE id = a.id - 1)
           WHEN id = (SELECT MAX(id) FROM seat) THEN (SELECT student FROM seat WHERE id = a.id)
           ELSE (SELECT student FROM seat WHERE id = a.id + 1) END AS student
FROM seat a

