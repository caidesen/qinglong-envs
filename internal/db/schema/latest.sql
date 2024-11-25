CREATE TABLE users
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    username    TEXT NOT NULL UNIQUE,
    password    TEXT,
    wx_pusher_uid TEXT,
    admin       BOOLEAN NOT NULL DEFAULT 0

);