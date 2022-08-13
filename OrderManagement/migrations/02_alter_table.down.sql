alter table orders drop column meals;

alter table orders alter column closed_at set default null;

alter table orders alter column closed_at drop default;