create table if not exists todos (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    body TEXT NOT NULL,
    due_date DATETIME NOT NULL,
    complete_at DATETIME,
    created_at DATETIME ,
    update_at DATETIME 
);