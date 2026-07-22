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
    photo_path VARCHAR(255) NUll,
    FOREIGN KEY (location_id) REFERENCES locations(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);


INSERT INTO locations (name, description) VALUES
('Untere Lager', 'Lager neben dem Gebaudeeingang mit Schlagzeugen'),
('Shoplager', 'Lager neben christian stein'),
('Getrankelager', 'Lager mit Getranke in Nirvana'),
('Icinga Schrank', 'Icinga Dings'),
('Netways Schrank', 'Netways Dings'),
('Putzraum', 'Putzraum mit putz dings??');


INSERT INTO tags (name) VALUES
('Reserviert'),
('Kann weg'),
('Gebraucht');

INSERT INTO items (name, location_id, amount, description, tag_id) VALUES
('Papiertuch', 1, 18, 'Papiertuch also', NULL),
('Skeleton Head', 1, 2, 'Skeleton Head fur Deko', NULL),
('Club-Mate', 3, 32, 'club matte mag ich nicht', NULL),
('Rasberry Pi 4', 2, 4, 'Rasberry PI 4 4 gb ram', NULL),
('Lenovo Laptop', 2, 1, 'ThinkPad Lenovo', NULL),
('Briefumschlag', 6, 300, 'umschlag fur briefe', NULL),
('Carbon Dioxide', 6, 1, 'dont smoke gas', NULL),
('Christmass stuff', 5, 1, 'hohoho', NULL),
('cool tshirt', 5, 5, 'cool tshirt', NULL),
('Einlassband', 5, 50, 'Einglassband Icinga und Netways', NULL),
('Feuchttuch', 6, 8, 'Tucher fur saubere Hande??', NULL),
('Flussigseife', 6, 5, 'Seife fur saubere Hande??', NULL),
('Klo Papier', 6, 60, 'obere regal(falls du schnell brauchst)', NULL),
('Kulis', 4, 20, 'nice kulis netways kulis', NULL),
('Markers', 6, 10, 'markers fur wande taggen', NULL),
('Notizblock', 4, 60, 'block fur notizen und malen und skizzen und zeichnen', NULL),
('Ordner', 6, 20, 'ordner fur organisation und dings', NULL),
('Stickers', 5, 20, 'ubertreib nicht mit stickers pls', NULL),
('Spulmaschine Kapseln', 6, 6, 'Kapseln fur Spulmaschine', NULL);

