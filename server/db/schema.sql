CREATE TABLE users (
    uid INTEGER PRIMARY KEY AUTOINCREMENT,
    alexaId TEXT NOT NULL,
    alexaDevice TEXT NOT NULL,
    email TEXT,
    createTime TEXT DEFAULT (strftime('%Y-%m-%d %H:%M:%S:%s', 'now', 'localtime') ),
    updateTime TEXT DEFAULT (strftime('%Y-%m-%d %H:%M:%S:%s', 'now', 'localtime') )
);

CREATE TRIGGER update_users_updateTime
        BEFORE UPDATE
            ON users
BEGIN
    UPDATE users
       SET updateTime = strftime('%Y-%m-%d %H:%M:%S:%s', 'now', 'localtime') 
     WHERE uid = old.uid;
END;

-- qid is for question IDs, but i'm just making it a string to make things easier to
-- extend in the future when i'm not here
CREATE TABLE answers (
    uid INTEGER PRIMARY KEY AUTOINCREMENT,
    qid TEXT NOT NULL,
    answer TEXT NOT NULL,
    createTime TEXT DEFAULT (strftime('%Y-%m-%d %H:%M:%S:%s', 'now', 'localtime') ),
    updateTime TEXT DEFAULT (strftime('%Y-%m-%d %H:%M:%S:%s', 'now', 'localtime') )
);

CREATE TRIGGER update_answers_updateTime
        BEFORE UPDATE
            ON answers
BEGIN
    UPDATE answers
       SET updateTime = strftime('%Y-%m-%d %H:%M:%S:%s', 'now', 'localtime') 
     WHERE uid = old.uid;
END;