CREATE TABLE IF NOT EXISTS items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location_id INT NOT NULL,
    tag_id INT,
    amount INT NOT NULL DEFAULT 0,
    description TEXT,
    photo_path VARCHAR(255),
    FOREIGN KEY (location_id) REFERENCES locations(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);
