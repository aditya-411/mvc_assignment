import os
import getpass
import bcrypt
from subprocess import run

# Function to run a shell command and capture its output
def out(command):
    result = run(command, shell=True, capture_output=True, text=True)
    return result.stdout, result.stderr

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

print(f"Creating database with name {database}")
result, err = out(f"mysql -h {host} -P {port} -u {user} -p{password} -e 'CREATE DATABASE {database}'")
if "Can't create database" in err:
    print("Database already exists with given name. Either drop it or try again with a different name.")
    exit(1)
if "Access denied" in err:
    print("Database access denied. Check the username and password.")
    exit(1)
print(f"Database {database} created, copying dump.sql to it now")
result, err = out(f"mysql -h {host} -P {port} -u {user} -p{password} {database} < dump.sql")
if err != "mysql: [Warning] Using a password on the command line interface can be insecure.\n":
    print("Error copying dump.sql to the database. Check the database name and credentials or maybe dump.sql doesn't exist.")
    exit(1)
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
result, err = out(f"mysql -h {host} -P {port} -u {user} -p{password} {database} -e 'DELETE FROM users'")
if err != "mysql: [Warning] Using a password on the command line interface can be insecure.\n":
    print("Error deleting existing users from the database. Check the database name and credentials.")
    exit(1)
result, err = out(f"mysql -h {host} -P {port} -u {user} -p{password} {database} -e 'INSERT INTO users (username, password, isadmin) VALUES (\"{admin_username}\", \"{hashed.decode()}\", 1)'")
if err != "mysql: [Warning] Using a password on the command line interface can be insecure.\n":
    print("Error adding admin user to the database. Check the database name and credentials.")
    exit(1)
print("done")


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

print("Building the binary")
os.system(f"go build -o mvc ./cmd/main.go")
print("done")

print("\n\nSetup complete! Run './mvc' to start the server from the built binary.")
print("You can also run 'go run ./cmd/main.go' to start the server without building the binary. Or run 'go build -o mvc ./cmd/main.go' to rebuild binary after any changes.")