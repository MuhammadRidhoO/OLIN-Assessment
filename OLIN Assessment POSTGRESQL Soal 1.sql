SELECT u.name, SUM(oi.price * oi.quantity) AS transaction
FROM users u
JOIN orders o ON u.id = o.user_id
JOIN order_item oi ON o.id = oi.orders_id
WHERE o.create_at >= '2022-01-01'
GROUP BY u.name
HAVING SUM(oi.price * oi.quantity) >= 100;
