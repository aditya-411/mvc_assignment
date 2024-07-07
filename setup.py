import os
import getpass
import bcrypt

# Run a shell command
print("running: go mod vendor")
os.system("go mod vendor")
print("done")

print("running: go mod tidy")
os.system("go mod tidy")
print("done")


user = input("What is the database username (root used in example): ")
password = getpass.getpass("What is the database password: ")
confirm_password = getpass.getpass("Confirm the database password: ")
while password != confirm_password:
    print("Passwords do not match. Try again.")
    password = getpass.getpass("What is the database password: ")
    confirm_password = getpass.getpass("Confirm the database password: ")

print("If the provided username or password is incorrect, the database creation will fail and application won't connect to db.")


database = input("What is the database name: ")
host = input("What is the database host (localhost used in example): ")
port = input("What is the database port (3306 used in example): ")

jwt_secret = getpass.getpass("Select a secret for JWT cookies: ")
confirm_jwt_secret = getpass.getpass("Confirm the JWT secret: ")
while jwt_secret != confirm_jwt_secret:
    print("Secrets do not match. Try again.")
    jwt_secret = getpass.getpass("Select a secret for JWT cookies: ")
    confirm_jwt_secret = getpass.getpass("Confirm the JWT secret: ")

salt_rounds = input("Select a number of salt rounds for password hashing (10 default): ")

lines = []
lines.append(f'DB_USERNAME="{user}"')
lines.append(f'DB_PASSWORD="{password}"')
lines.append(f'DB_HOST="{host}"')
lines.append(f'DB_PORT="{port}"')
lines.append(f'DB_NAME="{database}"')
lines.append(f'salt_rounds={salt_rounds}')
lines.append(f'jwt_secret="{jwt_secret}"')

print("Writing to .env file")
with open(".env", "w") as file:
    file.write("\n".join(lines))
print("done")

print(f"Creating database with name {database}")
os.system(f"mysql -h {host} -P {port} -u {user} -p{password} -e 'CREATE DATABASE {database}'")
print(f"Database {database} created, copying dump.sql to it now")
os.system(f"mysql -h {host} -P {port} -u {user} -p{password} {database} < dump.sql")
print("done")

admin_username = input("Enter the admin username: ")
admin_password = getpass.getpass("Enter the admin password: ")
confirm_admin_password = getpass.getpass("Confirm the admin password: ")
while admin_password != confirm_admin_password:
    print("Passwords do not match. Try again.")
    admin_password = getpass.getpass("Enter the admin password: ")
    confirm_admin_password = getpass.getpass("Confirm the admin password: ")

print("Adding admin user to the database")
salt = bcrypt.gensalt(int(salt_rounds))
hashed = bcrypt.hashpw(admin_password.encode(), salt)
os.system(f"mysql -h {host} -P {port} -u {user} -p{password} {database} -e 'DELETE FROM users'")
os.system(f"mysql -h {host} -P {port} -u {user} -p{password} {database} -e 'INSERT INTO users (username, password, isadmin) VALUES (\"{admin_username}\", \"{hashed.decode()}\", 1)'")
print("done")

print("Building the binary")
os.system(f"go build -o mvc ./cmd/main.go")
print("done")

print("\n\nSetup complete! Run './mvc' to start the server from the built binary.")
print("You can also run 'go run ./cmd/main.go' to start the server without building the binary. Or run 'go build -o mvc ./cmd/main.go' to rebuild binary after any changes.")