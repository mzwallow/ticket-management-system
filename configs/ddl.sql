create type status as enum ('PENDING', 'ACCEPTED', 'RESOLVED', 'REJECTED');

alter type status owner to "ticket-admin";

create table tickets
(
    id                  serial
        constraint tickets_pk
            primary key,
    title               varchar(255),
    status              status       default 'PENDING'::status,
    description         text,
    contact_information text,
    created_at          timestamp    default CURRENT_TIMESTAMP,
    updated_at          timestamp    default CURRENT_TIMESTAMP
);

alter table tickets
    owner to "ticket-admin";