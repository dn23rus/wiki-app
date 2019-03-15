CREATE TABLE IF NOT EXISTS pages
(
  id         INT(11) UNSIGNED AUTO_INCREMENT,
  slug       VARCHAR(255)                         NOT NULL,
  title     VARCHAR(255)                         NOT NULL,
  content    TEXT                                 NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP() NOT NULL,
  PRIMARY KEY (id)
)
  DEFAULT CHARACTER SET utf8
  DEFAULT COLLATE utf8_general_ci;

CREATE UNIQUE INDEX page_slug ON pages (slug);
