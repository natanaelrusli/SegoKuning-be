-- Insert seed data into users table
INSERT INTO "users" ("email", "phone", "password", "name", "images_id", "created_at", "updated_at", "deleted_at")
VALUES
('user1@example.com', '1234567', 'password1', 'User One', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
('user2@example.com', '7654321', 'password2', 'User Two', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
('user3@example.com', '9876543', 'password3', 'User Three', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
('user4@example.com', '4567890', 'password4', 'User Four', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
('user5@example.com', '0987654', 'password5', 'User Five', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);

-- Insert seed data into images table
INSERT INTO "images" ("url", "created_at", "deleted_at")
VALUES
('https://example.com/image1.jpg', CURRENT_TIMESTAMP, NULL),
('https://example.com/image2.jpg', CURRENT_TIMESTAMP, NULL),
('https://example.com/image3.jpg', CURRENT_TIMESTAMP, NULL),
('https://example.com/image4.jpg', CURRENT_TIMESTAMP, NULL),
('https://example.com/image5.jpg', CURRENT_TIMESTAMP, NULL);

-- Insert seed data into friendships table
INSERT INTO "friendships" ("uid1", "uid2", "created_at", "deleted_at")
VALUES
(1, 2, CURRENT_TIMESTAMP, NULL),
(1, 3, CURRENT_TIMESTAMP, NULL),
(2, 3, CURRENT_TIMESTAMP, NULL),
(4, 5, CURRENT_TIMESTAMP, NULL),
(1, 5, CURRENT_TIMESTAMP, NULL);

-- Insert seed data into posts table
INSERT INTO "posts" ("user_id", "post_content", "tags", "created_at", "updated_at", "deleted_at")
VALUES
(1, 'Post by User One', 'tag1, tag2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(2, 'Post by User Two', 'tag3, tag4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(3, 'Post by User Three', 'tag5, tag6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(4, 'Post by User Four', 'tag7, tag8', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(5, 'Post by User Five', 'tag9, tag10', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);

-- Insert seed data into comments table
INSERT INTO "comments" ("post_id", "user_id", "created_at", "updated_at", "deleted_at")
VALUES
(1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(2, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(3, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(4, 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(5, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);
