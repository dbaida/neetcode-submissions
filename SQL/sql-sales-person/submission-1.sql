-- Write your query below
SELECT sp.name
FROM sales_person AS sp
LEFT JOIN orders AS o ON o.sales_id = sp.sales_id
LEFT JOIN company AS c ON c.com_id = o.com_id
GROUP BY sp.sales_id
HAVING COUNT(o.order_id) FILTER (WHERE c.name = 'CRIMSON') = 0
