CREATE ROLE apollo_federation WITH PASSWORD 'password';
ALTER ROLE apollo_federation WITH LOGIN;
CREATE DATABASE apollo_federation_development;
GRANT ALL PRIVILEGES ON DATABASE apollo_federation_development TO apollo_federation;
