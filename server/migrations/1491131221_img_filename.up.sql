ALTER TABLE image
  ADD COLUMN filename VARCHAR(255),
  ADD COLUMN nuid VARCHAR(32),
  ADD COLUMN ncp SMALLINT,
  ADD COLUMN ncreated TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'UTC') NOT NULL,
  ADD COLUMN nupdated TIMESTAMP WITH TIME ZONE;

UPDATE image SET
  filename = title || '.jpg',
  nuid = uid,
  ncp = cp,
  ncreated = created,
  nupdated = updated;

ALTER TABLE image
  DROP COLUMN uid,
  DROP COLUMN cp,
  DROP COLUMN created,
  DROP COLUMN updated;

ALTER TABLE image
  ALTER COLUMN filename SET NOT NULL,
  ALTER COLUMN ncp SET NOT NULL;

ALTER TABLE image RENAME COLUMN nuid TO uid;
ALTER TABLE image RENAME COLUMN ncp TO cp;
ALTER TABLE image RENAME COLUMN ncreated TO created;
ALTER TABLE image RENAME COLUMN nupdated TO updated;
