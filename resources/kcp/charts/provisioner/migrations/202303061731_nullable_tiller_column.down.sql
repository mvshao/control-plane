BEGIN;
ALTER TABLE kyma_release ALTER COLUMN tiller_yaml SET NOT NULL;
COMMIT;
