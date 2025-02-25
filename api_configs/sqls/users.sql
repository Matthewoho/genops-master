CREATE TABLE USERS (
    ID BIGINT AUTO_INCREMENT,                  -- 用户ID
    USERNAME VARCHAR(50) NOT NULL UNIQUE DEFAULT '',  -- 用户名
    REALNAME VARCHAR(50),                       -- 真实姓名
    PASSWORD VARCHAR(255) NOT NULL DEFAULT '',  -- 密码
    SALT VARCHAR(50),                           -- 密码盐
    EMAIL VARCHAR(100),                         -- 邮箱
    PHONE VARCHAR(20),                          -- 电话
    ROLE_CODE VARCHAR(20),                      -- 角色编码
    ROLE_NAME VARCHAR(50),                      -- 角色名称
    STATUS TINYINT,                            -- 状态
    DEL_FLAG TINYINT DEFAULT 0,                -- 删除标志，默认0表示未删除
    REMARK TEXT,                                -- 备注
    CREATE_BY VARCHAR(50),                      -- 创建人
    CREATE_AT BIGINT,                          -- 创建时间 (Unix 时间戳)
    UPDATE_BY VARCHAR(50),                      -- 更新人
    UPDATE_AT BIGINT,                          -- 更新时间 (Unix 时间戳)
    CLIENT_IP VARCHAR(45),                      -- 客户端IP
    FAILED_CNT BIGINT DEFAULT 0,               -- 登录失败次数，默认0
    NICKNAME VARCHAR(50),                       -- 昵称
    DEPARTMENT VARCHAR(50),                     -- 部门
    EN_NAME VARCHAR(50),                        -- 英文名
    AVATAR VARCHAR(255),                        -- 头像
    IS_ADMIN TINYINT DEFAULT 0,                -- 是否为管理员，默认0表示否
    LAST_LOGIN_AT BIGINT,                      -- 最后登录时间 (Unix 时间戳)
    CREATED_IP VARCHAR(45),                     -- 创建时IP
    LAST_PASSWORD_CHANGE_AT BIGINT,            -- 最后修改密码时间 (Unix 时间戳)
    TWO_FACTOR_ENABLED TINYINT DEFAULT 0,      -- 是否启用双重认证，默认0表示否
    PRIMARY KEY (ID)                           -- 主键(goctl 必须使用这个方式定义主键)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
