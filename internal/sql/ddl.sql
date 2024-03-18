-- Create the 'users' table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    credentialType VARCHAR(10) NOT NULL,
    credentialValue VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

-- Create the 'profile' table
CREATE TABLE profile (
    id SERIAL PRIMARY KEY,
    userId INT NOT NULL,
    name VARCHAR(255),
    imageUrl VARCHAR(255),
    friendsId INT[],
    friendCount INT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users(id)
);

-- Create the 'posts' table
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    postInHTML VARCHAR(500) NOT NULL,
    tags VARCHAR(255)[] NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the 'comments' table
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    postsId INT NOT NULL,
    userId INT NOT NULL,
    comment VARCHAR(500) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (postsId) REFERENCES posts(id),
    FOREIGN KEY (userId) REFERENCES users(id)
);

-- Create the 'images' table
CREATE TABLE images (
    id SERIAL PRIMARY KEY,
    url VARCHAR(255) NOT NULL
);
