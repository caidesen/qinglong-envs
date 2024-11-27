CREATE TABLE users -- 用户
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    username      TEXT     NOT NULL UNIQUE,
    password      TEXT,
    wx_pusher_uid TEXT,                       -- wxpuhser 用户ID
    admin         BOOLEAN  NOT NULL DEFAULT 0 -- 是否为管理员
);

CREATE TABLE panels -- 青龙面板
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name          TEXT     NOT NULL,
    url           TEXT     NOT NULL,
    client_id     TEXT     NOT NULL,
    client_secret TEXT     NOT NULL
);

CREATE TABLE project -- 项目
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name        TEXT     NOT NULL,
    description TEXT     NOT NULL DEFAULT '',
    env_name    TEXT     NOT NULL,            -- 环境变量名
    env_limit   INTEGER  NOT NULL DEFAULT -1, -- 环境变量个数限制
    enabled     BOOLEAN  NOT NULL DEFAULT 1   -- 是否启用
);

CREATE TABLE panel_project -- 青龙面板<=>项目
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    panel_id   INTEGER  NOT NULL,
    project_id INTEGER  NOT NULL,
    UNIQUE (panel_id, project_id),
    FOREIGN KEY (panel_id) REFERENCES panels (id),
    FOREIGN KEY (project_id) REFERENCES project (id)
);

CREATE TABLE project_user -- 项目<=>用户
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    project_id INTEGER  NOT NULL,            -- 加入项目
    user_id    INTEGER  NOT NULL,
    env_value  TEXT     NOT NULL,            -- 环境变量值（ck）
    enabled    BOOLEAN  NOT NULL DEFAULT 1,  -- 是否启用
    UNIQUE (project_id, user_id, env_value), -- 一个用户在一个项目中环境变量不能重复
    FOREIGN KEY (project_id) REFERENCES project (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);