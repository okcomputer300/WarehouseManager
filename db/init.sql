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
('Untere Lager', 'Lager neben dem Gebäudeeingang mit Schlagzeugen'),
('Shoplager', 'Lager neben christian stein'),
('Getränkelager', 'Lager mit Getränke in Nirvana');

INSERT INTO tags (name) VALUES
('Reserviert'),
('Kann weg'),
('Wird gebraucht');

INSERT INTO items (name, location_id, amount, description, tag_id) VALUES
('Papiertuch', 1, 18, 'Papiertuch also', NULL),
('Skeleton Head', 1, 2, 'Skeleton Head für Deko', NULL),
('Club-Mate', 3, 32, 'club matte mag ich nicht', NULL),
('Rasberry Pi 4', 2, 4, 'Rasberry PI 4 4 gb ram', NULL),
('Lenovo Laptop', 2, 1, 'ThinkPad Lenovo', NULL);

