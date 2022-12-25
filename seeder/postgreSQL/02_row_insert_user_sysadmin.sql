-- +migrate Up
-- +migrate StatementBegin

-- password hash sha256 dari @Sysadmin37
INSERT INTO users (user_id, email, password, client_id, status)
VALUES (1, 'bg.drive01@gmail.com',
        '6fe88193540bbe2b9113b349b0eacbc50938ed19943696c9d568d68aa4ee55d5',
        'af3483c1-8bd5-4502-9ba9-7098c06c623b', 2);

-- password hash sha256 dari @Sysadmin37
INSERT INTO users (user_id, email, password, client_id, status)
VALUES (2, 'bg.drive02@gmail.com',
        '6fe88193540bbe2b9113b349b0eacbc50938ed19943696c9d568d68aa4ee55d5',
        'ff2fd762-95ee-4e20-9150-de70e0c73f8d', 2);

-- +migrate StatementEnd
-- +migrate Down
