INSERT INTO warehouse (id, title, available)
VALUES
    (1, 'main', true),
    (2, 'second', false),
    (3, 'emergency', true)
ON CONFLICT (id) DO NOTHING;

INSERT INTO item (id, title, code, reserved)
VALUES
    (1, 'sweater', 'cdbnth1', false),
    (2, 'hoodie', 'eelb2', true),
    (3, 'sneakers', 'rhjccjdrb3', true),
    (4, 'jeans', 'lbycs4', false),
    (5, 'shirt', 'hefirf5', true),
    (6, 'cap', 'rtgrf6', false)
ON CONFLICT (id) DO NOTHING;

DELETE FROM warehouse_item;

INSERT INTO warehouse_item (item_id, warehouse_id, quantity)
VALUES
    (1, 1, 15),
    (2, 1, 24),
    (3, 2, 11),
    (4, 2, 57),
    (5, 3, 33),
    (6, 3, 34);