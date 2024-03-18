-- Insert seed data into the 'users' table
INSERT INTO users (name, credentialType, credentialValue, password) VALUES
    ('John Doe', 'phone', '+6281234567890', 'password1'),
    ('Jane Smith', 'email', 'jane@example.com', 'password2'),
    ('Alice Johnson', 'phone', '+6289876543210', 'password3'),
    ('Bob Williams', 'email', 'bob@example.com', 'password4'),
    ('Eve Brown', 'phone', '+6285556667778', 'password5');

-- Insert seed data into the 'profile' table
INSERT INTO profile (userId, name, imageUrl, friendsId, friendCount) VALUES
    (1, 'John Doe', 'https://example.com/john.jpg', '{2,3}', 2),
    (2, 'Jane Smith', 'https://example.com/jane.jpg', '{1,4}', 2),
    (3, 'Alice Johnson', 'https://example.com/alice.jpg', '{1,4}', 2),
    (4, 'Bob Williams', 'https://example.com/bob.jpg', '{2,3}', 2),
    (5, 'Eve Brown', 'https://example.com/eve.jpg', '{2,4}', 2);

-- Insert seed data into the 'posts' table
INSERT INTO posts (postInHTML, tags) VALUES
    ('<p>This is a post about programming.</p>', '{"programming", "coding"}'),
    ('<p>Check out this new recipe!</p>', '{"food", "cooking"}'),
    ('<p>Traveling to beautiful destinations.</p>', '{"travel", "nature"}'),
    ('<p>Learning new skills and improving every day.</p>', '{"learning", "self-improvement"}'),
    ('<p>Sharing thoughts and ideas with the world.</p>', '{"communication", "ideas"}');

-- Insert seed data into the 'comments' table
INSERT INTO comments (postsId, userId, comment) VALUES
    (1, 2, 'Great post, I learned a lot!'),
    (1, 3, 'Interesting topic, looking forward to more.'),
    (2, 4, 'The recipe looks delicious, can''t wait to try it!'),
    (3, 5, 'Amazing photos, it makes me want to travel!'),
    (4, 1, 'Keep up the good work, always inspiring!');

-- Insert seed data into the 'images' table
INSERT INTO images (url) VALUES
    ('https://example.com/image1.jpg'),
    ('https://example.com/image2.jpg'),
    ('https://example.com/image3.jpg'),
    ('https://example.com/image4.jpg'),
    ('https://example.com/image5.jpg');
