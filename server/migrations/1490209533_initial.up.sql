-- Image Table
CREATE TABLE image
(
  id      VARCHAR(32) PRIMARY KEY             NOT NULL,
  title   VARCHAR(64),
  uid     VARCHAR(32),
  cp      SMALLINT                            NOT NULL,
  created TIMESTAMP DEFAULT current_timestamp NOT NULL,
  updated TIMESTAMP
);
CREATE INDEX IDX_image__uid
  ON image (uid);
COMMENT ON TABLE image IS 'Image metadata';

-- Gallery Table
CREATE TABLE gallery
(
  id      VARCHAR(32) PRIMARY KEY             NOT NULL,
  title   VARCHAR(64),
  uid     VARCHAR(32),
  cp      SMALLINT                            NOT NULL,
  created TIMESTAMP DEFAULT current_timestamp NOT NULL,
  updated TIMESTAMP
);
CREATE INDEX IDX_gallery__uid
  ON gallery (uid);
COMMENT ON TABLE gallery IS 'Gallery metadata';

-- Image identity table
CREATE TABLE image_ids
(
    iid VARCHAR(32) NOT NULL,
    sid VARCHAR(32) NOT NULL,
    gid VARCHAR(32),
    CONSTRAINT PK_image_ids PRIMARY KEY (iid, sid),
    CONSTRAINT FK_image_ids__iid FOREIGN KEY (iid) REFERENCES image (id) ON DELETE CASCADE,
    CONSTRAINT FK_image_ids__gid FOREIGN KEY (gid) REFERENCES gallery (id) ON DELETE CASCADE
);
CREATE INDEX IDX_image_ids__gid ON image_ids (gid);
COMMENT ON TABLE image_ids IS 'Image identity';
