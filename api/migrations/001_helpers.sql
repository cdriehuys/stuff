CREATE FUNCTION _manage_updated_at(_tbl regclass) RETURNS VOID AS $$
BEGIN
    EXECUTE format('CREATE TRIGGER set_updated_at BEFORE UPDATE ON %s
                    FOR EACH ROW EXECUTE PROCEDURE _set_updated_at()', _tbl);
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION _set_updated_at() RETURNS trigger AS $$
BEGIN
    IF (
        NEW IS DISTINCT FROM OLD AND
        NEW.updated_at IS NOT DISTINCT FROM OLD.updated_at
    ) THEN
        NEW.updated_at := current_timestamp;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

---- create above / drop below ----

DROP FUNCTION _manage_updated_at;
DROP FUNCTION _set_updated_at;
