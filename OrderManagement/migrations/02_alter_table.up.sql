alter table orders add column meals int[];

alter table orders alter column closed_at drop default;

alter table orders alter column closed_at set default current_timestamp;