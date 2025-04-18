-- accountテーブルのテストデータ
INSERT INTO accounts (account_name, path_hash) VALUES
('John Doe', 'abc123'),
('Jane Smith', 'def456'),
('Alice Johnson', 'ghi789'),
('Bob Brown', 'jkl012');

-- taskテーブルのテストデータ
INSERT INTO tasks (account_id, title, due_date, progress) VALUES
(1, 'Complete project report', '2025-04-20 10:00:00', 50),
(2, 'Prepare presentation slides', '2025-04-22 15:00:00', 20),
(3, 'Fix critical bugs', '2025-04-18 09:00:00', 80),
(4, 'Plan team meeting', '2025-04-25 14:00:00', 10);