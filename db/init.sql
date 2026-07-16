CREATE TABLE IF NOT EXISTS locations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE IF NOT EXISTS tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    location_id INT NOT NULL,
    amount INT,
    description TEXT,
    tag_id INT,
    FOREIGN KEY (location_id) REFERENCES locations(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);


INSERT INTO locations (name, description) VALUES
('Untere Lager', 'Lager neben dem Gebäudeeingang mit Schlagzeugen');

INSERT INTO tags (name) VALUES
('Reserviert');

INSERT INTO items (name, location_id, amount, description, tag_id) VALUES
('Papiertuch', '1', '18', 'Papiertuch also', '1');  